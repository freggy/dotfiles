package brew

import (
	"slices"

	"github.com/freggy/dotfiles/internal/sh"
)

type State struct {
	Packages []Package `json:"packages"`
}

type Package struct {
	Name string `json:"name"`
	Tap  string `json:"tap"`
}

// Update checks whether there are changes between
// the new brew state and the current one saved on disc.
// If there are, the corresponding fields will be overridden
// with the new ones. Members cannot be nil.
func (s *State) Update(new State) {
	if new.Packages != nil &&
		!slices.Equal(s.Packages, new.Packages) {
		s.Packages = new.Packages
	}
}

func (s *State) Apply() error {
	for _, p := range s.Packages {
		if _, err := sh.ExecArgs("brew", "tap", p.Tap); err != nil {
			return err
		}
		if _, err := sh.ExecArgs("brew", "install", p.Name); err != nil {
			return err
		}
	}
	return nil
}
