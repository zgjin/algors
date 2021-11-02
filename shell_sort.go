package algors

func ShellSort(a []int) {
	h := 1
	for h < len(a)/3 {
		h = 3*h + 1
	}
	for h >= 1 {
		for i := h; i < len(a); i++ {
			for j := i; j >= h && a[j] > a[j-h]; j -= h {
			}
		}

		h = h / 3
	}
}
