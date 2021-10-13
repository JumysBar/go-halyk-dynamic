package perftest

import "testing"

const size = 10

func BenchmarkSecondSimpleSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice := make([]int, 0)
		for n := 0; n < size; n++ {
			slice = append(slice, n)
		}
	}
}

func BenchmarkSecondPreAllocSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice := make([]int, size)
		for n := 0; n < size; n++ {
			slice = append(slice, n)
		}
	}
}
