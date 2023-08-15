package main

import (
	"log"

	"github.com/freggy/dotfiles/cmd"
	"github.com/spf13/cobra"
)

func main() {
	root := &cobra.Command{
		Use:              "env",
		TraverseChildren: true,
	}
	root.AddCommand(
		cmd.Up(),
	)
	if err := root.Execute(); err != nil {
		log.Fatalf("cmd: %v", err)
	}
}
