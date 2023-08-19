package cmd

import "github.com/spf13/cobra"

func Apply() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "applies the current configuration",
		RunE: apply,
	}
	return cmd
}

func apply(cmd *cobra.Command, args []string) error {
	return nil
}
