package table_driven

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseString(t *testing.T) {

	cases := []struct {
		name     string
		input    string
		expected string
	}{
		{"An empty string", "  ", "  "},
		{"A palindrome", "civic", "civic"},
		{"A string with special characters.", "§test", "tset§"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := ReverseString(c.input)
			assert.Equal(t, c.expected, result)
		})
	}

}
