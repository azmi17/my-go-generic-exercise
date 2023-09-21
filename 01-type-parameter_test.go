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

func TestTypeParameter(t *testing.T) {
	var resultStr string = Length[string]("Azmi")
	assert.Equal(t, "Azmi", resultStr)

	var resultNum int = Length[int](99)
	assert.Equal(t, 99, resultNum)

	var resultBool bool = Length[bool](false)
	assert.Equal(t, false, resultBool)
}
