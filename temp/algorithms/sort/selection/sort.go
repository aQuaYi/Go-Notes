package selection

//Sort 选择排序实现的算法
func Sort(a []int) {
	n := len(a)

	for i := 0; i < n-1; i++ {
		index := i
		min := a[i]
		for j := i + 1; j < n; j++ {
			if a[j] < min {
				index = j
				min = a[j]
			}
		}
		a[i], a[index] = a[index], a[i]
	}
}
