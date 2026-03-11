package table_driven

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertTemperature(t *testing.T) {

	temperatures := []struct {
		temperature float64
		scale       string
		expected    float64
	}{
		{0, "C", 32},
		{100, "C", 212},
		{-40, "C", -40},
	}

	for _, s := range temperatures {
		result, err := ConvertTemperature(s.temperature, s.scale)

		if err != nil {
			t.Errorf("Error converting temperature %v %v: %v", s.temperature, s.scale, err)
		}
		assert.Equal(t, s.expected, result)
	}

}
