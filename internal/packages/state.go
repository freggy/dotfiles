package packages

import (
	"encoding/json"
	"os"
)

type State struct {
	Brew Brew `json:"brew"`

	stateFile string
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
