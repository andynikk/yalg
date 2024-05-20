package main

import (
	"strconv"
	"strings"
)

func dividedTwo(dp [][]int, inputData []int, n, m int) int {

	for i := 1; i < n+1; i++ {

		for j := 1; j <= m; j++ {
			p := 0
			if j-inputData[i-1] >= 0 {
				p = dp[i-1][j-inputData[i]] + inputData[i]
			}
			dp[i+1][j] = max(dp[i][j], p)
		}

		//for j := 0; j < m; j++ {
		//	dp[0][j] = dp[1][j]
		//	dp[1][j] = 0
		//}

		i++
		if i >= n {
			break
		}

	}

	return dp[n][m]
}

func sameAmounts(res *[]byte) string {

	lines := strings.Split(string(*res), "\n")

	n, _ := strconv.Atoi(lines[0])
	inputs := strings.Split(lines[1], " ")
	intInputs := make([]int, len(inputs))

	sum := 0
	for i := range inputs {
		intInputs[i], _ = strconv.Atoi(inputs[i])
		sum += intInputs[i]
	}

	if sum%2 != 0 {
		return "False"
	}

	dp := make([][]int, len(inputs)+1)
	for i := 0; i < len(inputs)+1; i++ {
		dp[i] = make([]int, sum/2+1)
	}

	if dividedTwo(dp, intInputs, n, sum/2) == sum/2 {
		return "True"
	}
	return "False"
}
