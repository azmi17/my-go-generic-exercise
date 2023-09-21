package gogeneric

import (
	"fmt"
	"testing"
)

func MultipleTypeParams[T1 any, T2 any](param1 T1, param2 T2) {
	fmt.Println(param1)
	fmt.Println(param2)
}

func TestMultipleTypeParams(t *testing.T) {
	MultipleTypeParams[int, string](99, "Azmi")
	MultipleTypeParams[string, int]("Azmi", 99)
}
