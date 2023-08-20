package brew

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBrew(t *testing.T) {
	tests := []struct {
		name     string
		current  State
		new      State
		expected State
	}{
		{
			name: "do not update nil values",
			current: State{
				Packages: []string{"abc"},
				Casks:    []string{"def"},
			},
			new: State{},
			expected: State{
				Packages: []string{"abc"},
				Casks:    []string{"def"},
			},
		},
		{
			name: "update both values",
			current: State{
				Packages: []string{"abc"},
				Casks:    []string{"def"},
			},
			new: State{
				Packages: []string{"abc", "123"},
				Casks:    []string{"def", "456"},
			},
			expected: State{
				Packages: []string{"abc", "123"},
				Casks:    []string{"def", "456"},
			},
		},
		{
			name: "update empty values",
			current: State{
				Packages: []string{"abc"},
				Casks:    []string{"def"},
			},
			new: State{
				Packages: []string{"abc"},
				Casks:    []string{""},
			},
			expected: State{
				Packages: []string{"abc"},
				Casks:    []string{""},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			tt.current.Update(tt.new)
			assert.Equal(t, tt.current, tt.expected)
		})
	}
}
