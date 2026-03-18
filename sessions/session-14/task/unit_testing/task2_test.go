package unit_testing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAge(t *testing.T) {

	name := "Araz"
	age := 24

	profile := NewUserProfile(name, age)
	if age < 18 {
		assert.Equal(t, profile.IsMinor, true)
	} else {
		assert.Equal(t, name, profile.Name, false)
	}
}
