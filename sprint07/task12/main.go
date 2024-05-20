package main

import (
	"strconv"
	"strings"
)

func maxWeight(dp [][]int, weight []int, items, W int) int {

	i := 0
	for {
		for j := 1; j <= W; j++ {

			w := 0
			if j-weight[i] >= 0 {
				w = weight[i] + dp[0][j-weight[i]]
			}

			dp[1][j] = max(dp[0][j], w)
		}

		for j := 0; j <= W; j++ {
			dp[0][j] = dp[1][j]
			dp[1][j] = 0
		}
		i++
		if i == items {
			break
		}
	}

	return dp[0][W]
}

// <template>
func goldLeprechauns(res *[]byte) string {

	lines := strings.Split(strings.Replace(string(*res), "\r", "", -1), "\n")

	items, _ := strconv.Atoi(strings.Split(lines[0], " ")[0])
	W, _ := strconv.Atoi(strings.Split(lines[0], " ")[1])

	//dp := make([][]int, items+1)
	dp := make([][]int, 2)

	for i := 0; i <= 1; i++ {
		dp[i] = make([]int, W+1)
	}

	weightItems := strings.Split(lines[1], " ")
	weight := make([]int, items)

	for i := 0; i < len(weightItems); i++ {
		weight[i], _ = strconv.Atoi(weightItems[i])
	}

	maximum := maxWeight(dp, weight, items, W)

	return strconv.Itoa(maximum) + "\n"
}
