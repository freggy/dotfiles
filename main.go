package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/freggy/dotfiles/internal/cmd"
	"github.com/freggy/dotfiles/internal/cmd/brew"
	"github.com/freggy/dotfiles/internal/packages"
	"github.com/spf13/cobra"
)

func main() {
	root := &cobra.Command{
		Use:              "dof",
		TraverseChildren: true,
		SilenceUsage:     true,
	}
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("could not get home dir: %v", err)
	}

	dir := filepath.Join(home, "/.config/dof")
	state, err := packages.StateFromFile(dir + "/packages.state.json")
	if err != nil {
		log.Printf("WARN could not read state from file: %v\n", err)
	}

	config := cmd.Config{
		PackageState: state,
		InstallDir:   dir,
		HomeDir:      home,
	}

	root.AddCommand(
		cmd.Apply(config),
		brew.Cmd(state),
	)
	if err := root.Execute(); err != nil {
		log.Fatalf("%v", err)
	}
}
