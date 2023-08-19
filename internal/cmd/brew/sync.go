package brew

import (
	"log"
	"strings"
	"sync"

	"github.com/freggy/dotfiles/internal/cmd"
	"github.com/freggy/dotfiles/internal/packages"
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
		brew := packages.Brew{}
		// these commands take a while to complete,
		// so run them in parallel.
		wg := sync.WaitGroup{}
		wg.Add(2)
		go func() {
			log.Println("syncing installed brew packages")
			out, err := sh.Exec("brew leaves -r")
			if err != nil {
				reterr = err
				return
			}
			for _, p := range list(out) {
				brew.Packages = append(brew.Packages, p)
			}
			wg.Done()
		}()
		go func() {
			log.Println("syncing installed casks")
			out, err := sh.Exec("brew list -1 --casks")
			if err != nil {
				reterr = err
				return
			}
			for _, c := range list(out) {
				brew.Casks = append(brew.Casks, c)
			}
			wg.Done()
		}()
		wg.Wait()
		if reterr != nil { // ???
			return reterr
		}
		state.Brew.Update(brew)
		log.Println(state)
		return nil
		//return state.Flush()
	}
}

func list(in []byte) []string {
	return strings.Split(string(in), "\n")
}
