package main

import (
	"strconv"
	"strings"
)

func catalanNumber(n int) int {
	if n == 0 {
		return 1
	} else {
		return (4*n - 2) * catalanNumber(n-1) / (n + 1)
	}
}

func Solution(res *[]byte) string {
	lines := strings.Split(string(*res), "\n")
	n, _ := strconv.Atoi(lines[0])

	return strconv.Itoa(catalanNumber(n))
}
