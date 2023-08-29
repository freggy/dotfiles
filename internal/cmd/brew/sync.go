package brew

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/freggy/dotfiles/internal/cmd"
	"github.com/freggy/dotfiles/internal/packages"
	"github.com/freggy/dotfiles/internal/packages/brew"
	"github.com/freggy/dotfiles/internal/sh"
	"github.com/spf13/cobra"
)

func syncc(state *packages.State) *cobra.Command {
	return &cobra.Command{
		Use:   "sync",
		Short: "syncs brew casks and formulas",
		RunE:  brewSync(state),
	}
}

func brewSync(state *packages.State) cmd.RunEFunc {
	return func(cmd *cobra.Command, args []string) (reterr error) {
		// In order to persist local package changes we need to do the following
		// * git pull to retrieve the newest package state
		//   if we do not do this we risk having inconsistencies
		// * apply the new state

		pkgs, err := installedBrewPackages()
		if err != nil {
			return fmt.Errorf("get installed brew packages: %w", err)
		}

		var remove []string
		// find packages that are installed on the system,
		// but not contained in the desired state. we want
		// to uninstall these.
		for _, actual := range pkgs {
			found := false
			for _, desired := range state.Brew.Packages {
				if desired.Name == actual {
					found = true
					break
				}
			}
			if !found {
				// TODO: remove deleted entries from pkgs (use map?)
				remove = append(remove, actual)
			}
		}

		s := strings.Join(remove, " ")
		log.Printf("removing brew packages %s\n", s)

		if _, err := sh.ExecArgs("brew uninstall", s); err != nil {
			return fmt.Errorf("uninstall brew packages: %w", err)
		}

		out, err := sh.Cmd("brew info --json").Append(pkgs...).Exec(nil)
		if err != nil {
			return err
		}

		var list []brew.Package
		if err := json.Unmarshal(out, &list); err != nil {
			return err
		}

		state.Brew.Update(brew.State{
			Packages: list,
		})

		if err := state.Flush(); err != nil {
			return fmt.Errorf("flush state: %w", err)
		}

		// TODO: git add . && git commit && git push

		return nil
	}
}

func installedBrewPackages() ([]string, error) {
	out, err := sh.Exec("brew leaves -r")
	if err != nil {
		return nil, err
	}
	return list(out), nil
}

func list(in []byte) []string {
	var l []string
	for _, i := range strings.Split(string(in), "\n") {
		if i == "" {
			continue
		}
		l = append(l, i)
	}
	return l
}
