package main

import (
	"fmt"
	"sort"
)

//IntSlice 定义interface{},并实现sort.Interface接口的三个方法
type IntSlice []int

func (is IntSlice) Len() int {
	return len(is)
}

func (is IntSlice) Less(i, j int) bool {
	return is[i] < is[j]
}

func (is IntSlice) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

//Ascending 让is是升序排列
func (is IntSlice) Ascending() {
	sort.Sort(is)
}

//Descending 让is是降序排列
func (is IntSlice) Descending() {
	//NOTICE: 通过sort.Reverse()的转换，可以直接实现另一种降序排序
	//NOTICE: 通过sort.Reverse()的转换，可以直接实现另一种降序排序
	//NOTICE: 通过sort.Reverse()的转换，可以直接实现另一种降序排序
	sort.Sort(sort.Reverse(is))
}

func main() {
	a := IntSlice{1, 3, 5, 7, 2, 8, 4, 6, 9}

	fmt.Println("a的原始排序")
	fmt.Println(a)

	a.Ascending()
	fmt.Println("a经过升序排序后")
	fmt.Println(a)

	a.Descending()
	fmt.Println("a经过降序排序后")
	fmt.Println(a)

	fmt.Println("//NOTICE:sort.Sort是原地排序，也就是说，改变了a切片本身。")
}
