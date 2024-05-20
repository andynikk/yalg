package main

import (
	"strconv"
	"strings"
)

func maxFlowers(nI, mJ int, field [][]int, dp [][]int) int {

	for j := 1; j < mJ; j++ {
		dp[nI-1][j] += dp[nI-1][j-1]
		field[nI-1][j] = dp[nI-1][j]
	}

	for i := nI - 2; i >= 0; i-- {
		dp[i][0] += dp[i+1][0]
		field[i][0] = dp[i][0]
	}

	for i := nI - 2; i >= 0; i-- {
		for j := 1; j < mJ; j++ {
			dp[i][j] = field[i][j] + max(field[i+1][j], field[i][j-1])
			field[i][j] = dp[i][j]
		}
	}

	return dp[0][mJ-1]
}

// <template>
func fieldFlowers(res *[]byte) string {

	lines := strings.Split(strings.Replace(string(*res), "\r", "", -1), "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	n, _ := strconv.Atoi(strings.Split(lines[0], " ")[0])
	m, _ := strconv.Atoi(strings.Split(lines[0], " ")[1])

	dp := make([][]int, n)
	field := make([][]int, n)
	for i := 1; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}

		a := make([]int, m)
		b := make([]int, m)
		for index, l := range []rune(lines[i]) {
			v, _ := strconv.Atoi(string(l))

			a[index] = v
			b[index] = v
		}
		field[i-1] = a
		dp[i-1] = b
	}

	maximum := maxFlowers(n, m, field, dp)

	return strconv.Itoa(maximum) + "\n"
}
