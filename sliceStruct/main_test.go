package main

import (
	"testing"
)

const LEN = 1024

// 通过以下测试可知
// 
//

func Benchmark_createArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = [LEN]int{}
	}
}

func Benchmark_createArrayAndReturn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		createArray()
	}
}

func Benchmark_arrayAssign(b *testing.B) {
	a := [LEN]int{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < len(a); i++ {
			a[i] = i
		}
	}
}

func Benchmark_createArrayAndAssignAndReturn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		createArrayAndAssign()
	}
}

func Benchmark_makeSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = make([]int, LEN)
	}
}

func Benchmark_makeSliceAndReturn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		makeSlice()
	}
}

func Benchmark_sliceAssign(b *testing.B) {
	s := make([]int, LEN)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < len(s); i++ {
			s[i] = i
		}
	}
}

func Benchmark_makeSliceAndAssignAndReturn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		makeSliceAndAssign()
	}
}
