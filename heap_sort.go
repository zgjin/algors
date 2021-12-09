package algors

func sink(a []int, k, n int) {
	for k*2 <= n {
		j := k * 2
		if j < n && a[j] < a[j+1] {
			j += 1
		}
		if !(a[k] < a[j]) {
			break
		}

		a[k], a[j] = a[j], a[k]
		k = j
	}
}

// sort a[1:]
func heap_sort(a []int) {
	n := len(a[1:])
	for k := n / 2; k >= 1; k-- {
		sink(a, k, n)
	}

	k := n
	for k > 1 {
		swap(a, 1, k)
		k -= 1
		sink(a, 1, k)
	}
}
