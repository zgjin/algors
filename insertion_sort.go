package algors

func InsertionSort(a []int) {
	for i := 1; i < len(a); i++ {
		for j := i; j > 0 && a[j] < a[j-1]; j-- {
			swap(a, j, j-1)
		}
	}
}
