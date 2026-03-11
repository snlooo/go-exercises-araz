package benchmarking

import "testing"

func BenchmarkBubbleSort(t *testing.B) {

	//var arr = []int{33, 3, 25, 7, 9}
	var arr2 = []int{33, 3, 25, 7, 9, 76, 89}

	for i := 0; i < t.N; i++ {

		BubbleSort(arr2)
	}
}
