package main

import (
	"strconv"
	"strings"
)

const mod = 1000000007

// <template>
func jumpingStairs(res *[]byte) string {

	lines := strings.Split(string(*res), "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	lines = strings.Split(lines[0], " ")

	n, _ := strconv.Atoi(lines[0])
	k, _ := strconv.Atoi(lines[1])

	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		for j := 1; j <= k && i-j >= 0; j++ {
			dp[i] = (dp[i] + dp[i-j]) % mod
		}
	}

	return strconv.Itoa(dp[len(dp)-1]) + "\n"
}
