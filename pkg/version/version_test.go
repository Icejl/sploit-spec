package version

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVer_String(t *testing.T) {
	assert.Equal(t, "sploit public version library-import, build library-import at library-import", DefaultVer().String())
}
