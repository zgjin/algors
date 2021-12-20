package algors

import (
	"log"
	"strconv"
	"strings"
)

func printNumbers(n int) {
	s := make([]string, n)
	pn(s, 0, n)
}

func pn(s []string, index int, n int) {
	if index >= n {
		log.Println(strings.Join(s, ""))
		return
	}

	for i := 0; i < 10; i++ {
		s[index] = strconv.Itoa(i)
		pn(s, index+1, n)
	}
}
