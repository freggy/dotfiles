package packages

import "slices"

type Brew struct {
	Packages []string `json:"packages"`
	Casks    []string `json:"casks"`
}

// Update checks whether there are changes between
// the new brew state and the current one saved on disc.
// If there are, the corresponding fields will be overridden
// with the new ones. Members cannot be nil.
func (b *Brew) Update(newState Brew) {
	if newState.Packages != nil &&
		!slices.Equal(b.Packages, newState.Packages) {
		b.Packages = newState.Packages
	}
	if newState.Casks != nil &&
		!slices.Equal(b.Casks, newState.Casks) {
		b.Casks = newState.Casks
	}
}
