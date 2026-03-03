package go_generics

import (
	"testing"
)

func BenchmarkMinMaxInt(b *testing.B) {

	data := []int{5, 10, 8, 22, 14}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		(data)
	}
}
