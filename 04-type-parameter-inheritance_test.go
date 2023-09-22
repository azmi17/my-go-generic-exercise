package gogeneric

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Employee interface {
	GetName() string
}

// Generic functions
func GetName[T Employee](param T) string {
	return param.GetName()
}

type Manager interface {
	GetName() string
	GetManagerName() string
}

type MyManager struct {
	Name string
}

func (m *MyManager) GetName() string {
	return m.Name
}

func (m *MyManager) GetManagerName() string {
	return m.Name
}

type VicePresident interface {
	GetName() string
	GetVicePresidentName() string
}

type MyVicePresident struct {
	Name string
}

func (v *MyVicePresident) GetName() string {
	return v.Name
}

func (v *MyVicePresident) GetVicePresidentName() string {
	return v.Name
}

func TestTypeParamInheritance(t *testing.T) {
	assert.Equal(t, "Azmi", GetName[Manager](&MyManager{Name: "Azmi"}))
	assert.Equal(t, "Farhan", GetName[VicePresident](&MyVicePresident{Name: "Farhan"}))
}
