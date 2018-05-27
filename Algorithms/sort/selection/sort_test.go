package selection

import (
	"fmt"
	"testing"
)

func isSorted(a []int) bool {
	for i := 1; i < len(a); i++ {
		if a[i-1] > a[i] {
			return false
		}
	}
	return true
}

func Test_Sort(t *testing.T) {
	h := make([]int, 11)
	for i := 0; i < len(h); i++ {
		h[i] = len(h) - i - 1
	}
	fmt.Println(h)
	Sort(h)
	fmt.Println(h)
	if !isSorted(h) {
		t.Error("没能排序好")
	}
}

func Test_Sort_Sorted(t *testing.T) {
	h := make([]int, 11)
	for i := 0; i < len(h); i++ {
		h[i] = i
	}
	fmt.Println(h)
	Sort(h)
	fmt.Println(h)
	if !isSorted(h) {
		t.Error("没能排序好")
	}
}
func Test_Sort_with_Repeating(t *testing.T) {
	n := 10
	h := make([]int, n*3)
	for i := 0; i < n; i++ {
		h[i] = n - i - 1
		h[i+n], h[i+n+n] = h[i], h[i]
	}
	fmt.Println(h)
	Sort(h)
	fmt.Println(h)
	if !isSorted(h) {
		t.Error("没能排序好")
	}
}

func Test_Sort_100(t *testing.T) {
	h := make([]int, 100)
	for i := 0; i < len(h); i++ {
		h[i] = len(h) - i - 1
	}
	Sort(h)
	if !isSorted(h) {
		t.Error("没能排序好")
	}
}

func Test_Sort_1000(t *testing.T) {
	h := make([]int, 1000)
	for i := 0; i < len(h); i++ {
		h[i] = len(h) - i - 1
	}
	Sort(h)
	if !isSorted(h) {
		t.Error("没能排序好")
	}

}

func Test_Sort_10000(t *testing.T) {
	h := make([]int, 10000)
	for i := 0; i < len(h); i++ {
		h[i] = len(h) - i - 1
	}
	Sort(h)
	if !isSorted(h) {
		t.Error("没能排序好")
	}
}

func Test_Sort_100000(t *testing.T) {
	h := make([]int, 100000)
	for i := 0; i < len(h); i++ {
		h[i] = len(h) - i - 1
	}
	Sort(h)
	if !isSorted(h) {
		t.Error("没能排序好")
	}
}
