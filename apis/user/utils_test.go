package user

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultUsername(t *testing.T) {
	name := defaultUsername()
	assert.LessOrEqual(t, len(name), 32)
}
