package benchmarking

import "testing"

func BenchmarkConcatenateWithJoin(b *testing.B) {
	var str = []string{"Araz", "is", "a", "Gopher "}

	for i := 0; i < b.N; i++ {
		ConcatenateWithJoin(str)
	}
}

func BenchmarkConcatenateWithPlus(b *testing.B) {
	var str = []string{"Araz", "is", "a", "Developer "}

	for i := 0; i < b.N; i++ {
		ConcatenateWithPlus(str)
	}
}
