package cmd

import "github.com/spf13/cobra"

func Up() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "up",
		RunE: up,
	}
	return cmd
}

func up(cmd *cobra.Command, args []string) error {
	return nil
}
