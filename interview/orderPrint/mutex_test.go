package orderprint

import "testing"

func Benchmark_mutex(b *testing.B) {
	for i := 1; i < b.N; i++ {
		mutex()
	}
}
