package packages

import (
	"slices"

	"github.com/freggy/dotfiles/internal/sh"
)

type Brew struct {
	Packages []BrewPackage `json:"packages"`
}

type BrewPackage struct {
	Name string `json:"name"`
	Tap  string `json:"tap"`
}

// Update checks whether there are changes between
// the new brew state and the current one saved on disc.
// If there are, the corresponding fields will be overridden
// with the new ones. Members cannot be nil.
func (b *Brew) Update(new Brew) {
	if new.Packages != nil &&
		!slices.Equal(b.Packages, new.Packages) {
		b.Packages = new.Packages
	}
}

func (b *Brew) Apply() error {
	for _, p := range b.Packages {
		if _, err := sh.ExecArgs("brew", "tap", p.Tap); err != nil {
			return err
		}
		if _, err := sh.ExecArgs("brew", "install", p.Name); err != nil {
			return err
		}
	}
	return nil
}
