package deferBenchmark

import "testing"

func Benchmark_call(b *testing.B) {
	for i := 0; i < b.N; i++ {
		call()
	}
}
func Benchmark_deferCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deferCall()
	}
}
