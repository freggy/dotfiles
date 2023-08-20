package packages

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/freggy/dotfiles/internal/packages/brew"
	"github.com/freggy/dotfiles/internal/sh"
)

// TODO: rename(?)

// State holds all information about packages
// that should be installed across different
// package managers
type State struct {
	Brew brew.State `json:"brew"`

	stateFile string
}

func StateFromFile(path string) (*State, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var state *State
	if err := json.Unmarshal(data, &state); err != nil {
		return nil, err
	}
	state.stateFile = path
	return state, nil
}

// Flush writes the source to disc
func (s *State) Flush() error {
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(s.stateFile, data, 0777); err != nil {
		return err
	}
	return nil
}

// Apply installs all configured packages
func (s *State) Apply() error {
	if installed("brew") {
		if err := s.Brew.Apply(); err != nil {
			return fmt.Errorf("apply brew: %w", err)
		}
	}
	return nil
}

func installed(name string) bool {
	if _, err := sh.ExecArgs("command", "-v", name); err != nil {
		if !errors.Is(err, &exec.ExitError{}) {
			log.Printf("WARN failed to execute command -v %s reason: %v ", name, err)
			return false
		}
		return false
	}
	return true
}
