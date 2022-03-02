package algors

import "fmt"

var memRegCommon map[string]bool

func isMatchCommon(s string, p string) bool {
	memRegCommon = make(map[string]bool)

	return dpRegCommon(s, 0, p, 0)
}

func dpRegCommon(s string, i int, p string, j int) bool {
	m := len(s)
	n := len(p)

	if j == n {
		return i == m
	}
	if i == m {
		if j == n {
			return true
		}

		// 剩余的字符是 ? 或者 *
		for ; j < n; j++ {
			if p[j] != '*' {
				return false
			}
		}
		return true
	}
	var res bool
	var ok bool
	key := fmt.Sprintf("%d, %d", i, j)
	if res, ok = memRegCommon[key]; ok {
		return res
	}

	if s[i] == p[j] || p[j] == '?' || p[j] == '*' {
		if p[j] != '*' {
			res = dpRegCommon(s, i+1, p, j+1)
			memRegCommon[key] = res
			return res
		}
		res = dpRegCommon(s, i, p, j+1) || dpRegCommon(s, i+1, p, j)
		memRegCommon[key] = res
		return res
	}

	res = false
	memRegCommon[key] = res
	return res
}
