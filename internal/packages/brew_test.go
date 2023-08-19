package packages

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBrew(t *testing.T) {
	tests := []struct {
		name     string
		current  Brew
		new      Brew
		expected Brew
	}{
		{
			name: "do not update nil values",
			current: Brew{
				Packages: []string{"abc"},
				Casks:    []string{"def"},
			},
			new: Brew{},
			expected: Brew{
				Packages: []string{"abc"},
				Casks:    []string{"def"},
			},
		},
		{
			name: "update both values",
			current: Brew{
				Packages: []string{"abc"},
				Casks:    []string{"def"},
			},
			new: Brew{
				Packages: []string{"abc", "123"},
				Casks:    []string{"def", "456"},
			},
			expected: Brew{
				Packages: []string{"abc", "123"},
				Casks:    []string{"def", "456"},
			},
		},
		{
			name: "update empty values",
			current: Brew{
				Packages: []string{"abc"},
				Casks:    []string{"def"},
			},
			new: Brew{
				Packages: []string{"abc"},
				Casks:    []string{""},
			},
			expected: Brew{
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
