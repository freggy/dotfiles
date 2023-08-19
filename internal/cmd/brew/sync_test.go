package brew

import (
	"testing"

	"github.com/freggy/dotfiles/internal/packages"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	state := &packages.State{}
	err := syncc(state).RunE(nil, nil)
	assert.NoError(t, err)
}
