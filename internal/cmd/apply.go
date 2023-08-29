package cmd

import (
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/freggy/dotfiles/internal/sh"
	"github.com/spf13/cobra"
)

type applyOpts struct {
	Pkgs bool
	Pull bool
}

func Apply(config Config) *cobra.Command {
	opts := &applyOpts{}
	cmd := &cobra.Command{
		Use:   "apply",
		Short: "applies the configuration",
		RunE:  apply(config, opts),
	}
	cmd.Flags().BoolVarP(&opts.Pkgs, "install", "i", false, "whether packages should be reconciled")
	cmd.Flags().BoolVarP(&opts.Pull, "pull", "p", true, "whether new state should be pulled from git")
	return cmd
}

func apply(config Config, opts *applyOpts) RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		if opts.Pull {
			log.Println("pulling new updates from remote")
			if _, err := sh.ExecArgs("git", "-C", config.InstallDir+"/dotfiles", "pull"); err != nil {
				return fmt.Errorf("git pull: %w", err)
			}
		}
		if opts.Pkgs {
			log.Println("installing packages")
			if err := config.PackageState.Apply(); err != nil {
				return fmt.Errorf("apply state: %w", err)
			}
		}
		return copyFiles(config.InstallDir+"/dotfiles/rootfs", config.HomeDir)
	}
}

func copyFiles(root string, homedir string) error {
	return filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		// we are only interested in files
		if d.IsDir() {
			return nil
		}
		loc := strings.Replace(path, "__homedir", homedir, 1)
		loc = filepath.Join("/", loc)
		log.Printf("copy %s\n", loc)
		info, err := d.Info()
		if err != nil {
			return fmt.Errorf("file info: %w", err)
		}
		if err := os.MkdirAll(filepath.Dir(loc), 0777); err != nil {
			return fmt.Errorf("mkdir all: %w", err)
		}
		f, err := os.OpenFile(loc, os.O_TRUNC|os.O_CREATE|os.O_RDWR, info.Mode())
		if err != nil {
			return fmt.Errorf("open: %w", err)
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("fs read: %w", err)
		}
		if _, err := io.Copy(f, bytes.NewReader(data)); err != nil {
			return fmt.Errorf("copy: %w", err)
		}
		return nil
	})
}
