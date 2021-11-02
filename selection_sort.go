package algors

// SelectionSort 选择排序
// 时间复杂度：o(n^2/2)
func SelectionSort(a []int) {
	for i := 0; i < len(a); i++ {
		min := i
		for j := i + 1; j < len(a); j++ {
			if a[min] > a[j] {
				min = j
			}
		}

		swap(a, i, min)
	}
}

func swap(a []int, i, j int) {
	a[j], a[i] = a[i], a[j]
}
