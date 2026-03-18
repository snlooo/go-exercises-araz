package unit_testing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {

	a := 2
	b := 3

	result, _ := Calculate(a, b, "+")

	assert.Equal(t, result, 5)

}

func TestMinus(t *testing.T) {

	a := 6
	b := 3

	result, _ := Calculate(a, b, "-")

	assert.Equal(t, result, 3)

}

func TestMultiply(t *testing.T) {
	a := 6
	b := 2

	result, _ := Calculate(a, b, "*")

	assert.Equal(t, result, 12)
}

func TestDivide(t *testing.T) {

	a := 6
	b := 3

	result, err := Calculate(a, b, "/")
	if err != nil {
		assert.Error(t, err)
	} else {
		assert.Equal(t, result, 2)
	}
}

func TestUnSupportedOperation(t *testing.T) {

	a := 6
	b := 2

	_, err := Calculate(a, b, "§")
	if err != nil {
		assert.Error(t, err)
	}
}
