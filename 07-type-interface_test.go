package gogeneric

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypeInterface(t *testing.T) {
	assert.Equal(t, Age(150), Min2(Age(150), Age(700)))
}
