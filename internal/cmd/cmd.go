package cmd

import (
	"github.com/freggy/dotfiles/internal/packages"
	"github.com/spf13/cobra"
)

type RunEFunc func(cmd *cobra.Command, args []string) error

// TODO: populate fields

type Config struct {
	PackageState *packages.State
	InstallDir   string
	HomeDir      string
}
