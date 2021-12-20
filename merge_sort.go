package algors

var aux []int

func merge(a []int, lo, mid, hi int) {
	for k := lo; k <= hi; k++ {
		aux[k] = a[k]
	}

	i := lo
	j := mid + 1
	for k := lo; k <= hi; k++ {
		if i > mid {
			a[k] = aux[j]
			j++
		} else if j > hi {
			a[k] = aux[i]
			i++
		} else if aux[j] < aux[i] {
			a[k] = aux[j]
			j++
		} else {
			a[k] = aux[i]
			i++
		}
	}
}

func MergeSort(a []int, lo, hi int) {
	if hi <= lo {
		return
	}

	mid := lo + (hi-lo)/2
	MergeSort(a, lo, mid)
	MergeSort(a, mid+1, hi)

	merge(a, lo, mid, hi)
}

func MergeBU(a []int) {
	aux = make([]int, len(a))

	for sz := 1; sz < len(a); sz += sz {
		for lo := 0; lo < len(a)-sz; lo = lo + sz + sz {
			merge(a, lo, lo+sz-1, min(lo+sz+sz-1, len(a)-1))
		}
	}
}
