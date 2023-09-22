package gogeneric

import (
	"fmt"
	"testing"
)

type Bag[T any] []T

func PrintBag[T any](bag Bag[T]) {
	for _, val := range bag {
		fmt.Println(val)
	}
}

func TestBagString(t *testing.T) {
	names := Bag[string]{"Azmi", "Farhan", "Al-Khawarizmi"}
	PrintBag(names)
}

func TestBagInt(t *testing.T) {
	numbers := Bag[int]{1, 2, 3, 4, 5}
	PrintBag[int](numbers)
}
