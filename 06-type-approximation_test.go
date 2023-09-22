package gogeneric

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Age int

type Number2 interface {
	~int | int8 | int16 | int32 | int64 |
		float32 | float64
}

func Min2[T Number2](first T, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestTypeSetInerface2(t *testing.T) {
	assert.Equal(t, Age(150), Min2[Age](Age(150), Age(700)))
}
