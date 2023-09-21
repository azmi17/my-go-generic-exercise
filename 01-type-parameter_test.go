package gogeneric

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Length[T any](param T) T {
	fmt.Println(param)
	return param
}

func TestSample(t *testing.T) {
	var result string = Length[string]("Azmi")
	assert.Equal(t, "Azmi", result)

	var resultNum int = Length[int](99)
	assert.Equal(t, 99, resultNum)
}
