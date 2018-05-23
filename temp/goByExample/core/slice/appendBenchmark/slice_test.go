package slice

import "testing"

const (
	N = 10000 //切片长度
)

func Benchmark_Slice_Assignment(b *testing.B) {
	b.Log("设置好len和cap为n，利用s[j]=j的方式赋值。")
	n := N
	s := make([]int, n, n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			s[j] = j
		}
	}
}

func Benchmark_Slice_Append(b *testing.B) {
	b.Log("设置len为0，cap为n，利用s=append(s, j)的方式来赋值")
	n := N
	var s []int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		s = make([]int, 0, n)
		b.StartTimer()
		for j := 0; j < n; j++ {
			s = append(s, j)
		}
	}
}
func Benchmark_Slice_Append_startFrom0cap(b *testing.B) {
	b.Log("设置len和cap皆为0，利用s=append(s, j)的方式来赋值")
	n := N
	var s []int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		s = make([]int, 0, 0)
		b.StartTimer()
		for j := 0; j < n; j++ {
			s = append(s, j)
		}
	}
}
func Benchmark_Slice_Append_All(b *testing.B) {
	b.Log("设置len=0，cap=n，利用s=append(s, a...)来一次性赋值")
	n := N
	var s []int
	a := make([]int, n, n)
	for i := 0; i < n; i++ {
		a[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		s = make([]int, 0, n)
		b.StartTimer()
		s = append(s, a...)
	}
}
func Benchmark_Slice_Append_All_0Cap(b *testing.B) {
	b.Log("设置len=0，cap=0，利用s=append(s, a...)来一次性赋值")
	n := N
	var s []int
	a := make([]int, n, n)
	for i := 0; i < n; i++ {
		a[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		s = make([]int, 0, 0)
		b.StartTimer()
		s = append(s, a...)
	}
}
