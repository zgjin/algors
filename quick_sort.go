package algors

func partition(a []int, lo, hi int) (p int) {
	v := a[lo]
	i := lo
	j := hi + 1
	for {
		for i = i + 1; i <= hi; i++ {
			if a[i] > v {
				break
			}
		}

		for j = j - 1; j > lo; j-- {
			if a[j] < v {
				break
			}
		}
		if i >= j {
			break
		}

		swap(a, i, j)
	}

	swap(a, i, j)

	return
}

func quickSort(a []int, lo, hi int) {
	if lo == hi {
		return
	}

	p := partition(a, lo, hi)
	quickSort(a, lo, p-1)
	quickSort(a, p+1, hi)
}
