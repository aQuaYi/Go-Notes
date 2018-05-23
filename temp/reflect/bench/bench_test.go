package bench

import "testing"

func Benchmark_set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		set(100)
	}
}


func Benchmark_rset(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rset(100)
	}
}

func Benchmark_call(b *testing.B) {
	for i := 0; i < b.N; i++ {
	call()
	}
}


func Benchmark_rcall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rcall()
	}
}
