package gogeneric

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Number interface {
	int | int8 | int16 | int32 | int64 |
		float32 | float64
}

func Min[T Number](first T, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestTypeSetInerface(t *testing.T) {
	// var res int = Min[int](600, 200)
	// fmt.Println(res)
	assert.Equal(t, int32(150), Min[int32](int32(150), int32(200)))
	assert.Equal(t, float32(150), Min[float32](float32(150), float32(400)))
	assert.Equal(t, int(150), Min[int](int(150), int(700)))
}
