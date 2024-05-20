package main

import (
	"strconv"
	"strings"
)

func distance(dp [][2]int, s, t []string) int {

	for j := 1; j < len(t)+1; j++ {
		for i := 1; i < len(s)+1; i++ {
			dp[0][1] = j
			coefficient := 0
			if s[i-1] != t[j-1] {
				coefficient = 1
			}

			dp[i][1] = min(dp[i-1][1]+1, dp[i][0]+1, dp[i-1][0]+coefficient)
			dp[i-1][0] = dp[i-1][1]
		}
		dp[len(s)][0] = dp[len(s)][1]
	}

	return dp[len(s)][0]
}

// <template>
func Levenshtein(res *[]byte) string {

	lines := strings.Split(string(*res), "\n")

	s := strings.Split(lines[0], "")
	t := strings.Split(lines[1], "")

	if len(s) > len(t) {
		s, t = t, s
	}

	dp := make([][2]int, len(s)+1)
	for i := 0; i < len(s)+1; i++ {
		dp[i][0] = i
	}

	return strconv.Itoa(distance(dp, s, t)) + "\n"
}
