package orderprint

import "testing"

func Benchmark_channel(b *testing.B) {
	for i := 1; i < b.N; i++ {
		channel()
	}
}
