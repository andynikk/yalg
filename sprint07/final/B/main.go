package main

import (
	"strconv"
	"strings"
)

func distance(dp [][]int, s, t []string) int {

	for j := 1; j < len(t)+1; j++ {
		for i := 1; i < len(s)+1; i++ {
			coefficient := 0
			if s[i-1] != t[j-1] {
				coefficient = 1
			}

			dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+coefficient)
		}
	}

	return dp[len(s)][len(t)]
}

// <template>
func distanceLowenstein(res *[]byte) string {

	lines := strings.Split(string(*res), "\n")

	s := strings.Split(lines[0], "")
	t := strings.Split(lines[1], "")

	dp := make([][]int, len(s)+1)
	for i := 0; i < len(s)+1; i++ {
		dp[i] = make([]int, len(t)+1)
		dp[i][0] = i
	}

	for j := 0; j < len(t)+1; j++ {
		dp[0][j] = j
	}

	return strconv.Itoa(distance(dp, s, t)) + "\n"
}
