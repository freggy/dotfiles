package brew

import (
	"github.com/freggy/dotfiles/internal/packages"
	"github.com/spf13/cobra"
)

func Cmd(state *packages.State) *cobra.Command {
	brew := &cobra.Command{
		Use:   "brew",
		Short: "provides commands to manage brew packages",
	}
	brew.AddCommand(
		syncc(state),
	)
	return brew
}
