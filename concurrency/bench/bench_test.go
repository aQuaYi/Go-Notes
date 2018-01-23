package bench

import "testing"

func Benchmark_test(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test()
	}
}

func Benchmark_testBlock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testBlock()
	}
}
