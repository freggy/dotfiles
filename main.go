package main

import (
	"log"

	"github.com/freggy/dotfiles/internal/cmd"
	"github.com/freggy/dotfiles/internal/cmd/brew"
	"github.com/freggy/dotfiles/internal/packages"
	"github.com/spf13/cobra"
)

func main() {
	root := &cobra.Command{
		Use:              "dof",
		TraverseChildren: true,
	}
	state := &packages.State{}
	root.AddCommand(
		cmd.Apply(),
		brew.Cmd(state),
	)
	if err := root.Execute(); err != nil {
		log.Fatalf("cmd: %v", err)
	}
}
