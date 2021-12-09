package algors

func binary_search(a []int, val int) int {
	lo := 0
	hi := len(a) - 1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if a[mid] > val {
			hi = mid - 1
		} else if a[mid] < val {
			lo = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
