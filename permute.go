package algors

import "fmt"

var res [][]int
var root []int

func contains(a []int, b int, i int) bool {
	for j := 0; j <= i; j++ {
		if a[j] == b {
			return true
		}
	}

	return false
}

func permute(nums []int) [][]int {
	res = [][]int{}
	root = make([]int, len(nums))
	i := -1

	backtrace(nums, root, i)
	return res
}

func backtrace(nums []int, roots []int, i int) {
	if i == len(nums)-1 {
		cp := []int{}
		cp = append(cp, roots...)
		res = append(res, cp)
		return
	}

	for j, num := range nums {
		ok := contains(roots, num, i)
		if ok {
			fmt.Println("i = ", i, " num = ", num, "j = ", j)
			continue
		}

		i++
		roots[i] = num
		backtrace(nums, roots, i)
		i--
	}
}
