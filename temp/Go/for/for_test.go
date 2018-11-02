package fortest

import "testing"

const size = 1024

func Benchmark_for_i_j(b *testing.B) {
	array2d := [size][size]int{}
	res := 0
	for k := 1; k < b.N; k++ {
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				res += array2d[i][j]
			}
		}
	}
}

func Benchmark_for_j_i(b *testing.B) {
	array2d := [size][size]int{}
	res := 0
	for k := 1; k < b.N; k++ {
		for j := 0; j < size; j++ {
			for i := 0; i < size; i++ {
				res += array2d[i][j]
			}
		}
	}
}
