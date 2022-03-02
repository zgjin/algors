package algors

import "fmt"

var mem map[string]bool

// s 字符串 p 正则表达式
func isMatch(s, p string) (ok bool) {

	mem = make(map[string]bool)

	return dpReg(s, 0, p, 0)
}

func dpReg(s string, i int, p string, j int) bool {
	m := len(s)
	n := len(p)

	if j == n {
		return i == m
	}
	if i == m {
		// 检查 p[j..] 是否可以匹配空字符串
		if (n-j)%2 == 1 {
			return false
		}
		// 是否为 x*y*z*
		for ; j+1 < n; j += 2 {
			if p[j+1] != '*' {
				return false
			}
		}

		return true
	}

	var res bool
	key := fmt.Sprintf("%d,%d", i, j)
	var ok bool
	if res, ok = mem[key]; ok {
		return res
	}

	if s[i] == p[j] || p[j] == '.' {
		if j < n-1 && p[j+1] == '*' {
			res = dpReg(s, i, p, j+2) || dpReg(s, i+1, p, j)
		} else {
			res = dpReg(s, i+1, p, j+1)
		}
	} else {
		if j < n-1 && p[j+1] == '*' {
			res = dpReg(s, i, p, j+2)
		} else {
			res = false
		}
	}

	mem[key] = res
	return res
}
