package main

import (
	"strconv"
	"strings"
)

func distance(dp [][2]int, s, t []string) int {

	for j := 1; j < len(t)+1; j++ {
		for i := 1; i < len(s)+1; i++ {
			dp[1][0] = i
			coefficient := 0
			if s[i-1] != t[j-1] {
				coefficient = 1
			}

			//dp[i0][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+coefficient)
			dp[j][0] = min(dp[j][1]+1, dp[j-1][0]+1, dp[j-1][1]+coefficient)
		}
	}

	return dp[len(s)][1]
}

// <template>
func distanceLowenstein(res *[]byte) string {

	lines := strings.Split(string(*res), "\n")

	s := strings.Split(lines[0], "")
	t := strings.Split(lines[1], "")

	dp := make([][2]int, len(t)+1)
	for i := 0; i < len(t)+1; i++ {
		dp[i][1] = i
	}

	//for j := 0; j < len(t)+1; j++ {
	//	dp[0][j] = j
	//}

	return strconv.Itoa(distance(dp, s, t)) + "\n"
}
