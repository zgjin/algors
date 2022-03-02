package algors

import (
	"fmt"
	"math"
)

var memEgg map[string]int

func superEggDrop(k int, n int) int {
	memEgg = make(map[string]int)
	return dpEgg(k, n)
}

func dpEgg(k, n int) int {
	if k == 1 {
		return n
	}
	if n == 0 {
		return 0
	}

	key := fmt.Sprintf("%v,%v", k, n)
	if res, ok := memEgg[key]; ok {
		return res
	}

	res := int(math.MaxInt64)

	for i := 1; i <= n; i++ {
		res = min(res, max(dpEgg(k-1, i-1), dpEgg(k, n-i))+1)
	}

	memEgg[key] = res

	return res
}
