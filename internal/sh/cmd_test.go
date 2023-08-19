package sh_test

import (
	"testing"

	"github.com/freggy/dotfiles/internal/sh"
	"github.com/stretchr/testify/assert"
)

func TestCommand(t *testing.T) {
	tests := []struct {
		name     string
		cmd      sh.Cmd
		expected string
		err      bool
	}{
		{
			name:     "simple command",
			cmd:      "echo hello",
			expected: "hello\n",
		},
		{
			name: "fail",
			cmd:  "exit 3",
			err:  true,
		},
		// with the test below we can also ensure that
		// our string splitting logic works with entries
		// below 2
		{
			name: "fail single command",
			cmd:  "test",
			err:  true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			out, err := tt.cmd.Exec(nil)
			if tt.err {
				assert.Error(t, err)
				return
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.expected, string(out))
		})
	}
}
