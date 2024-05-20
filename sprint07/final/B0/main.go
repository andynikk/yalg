package main

import (
	"strconv"
	"strings"
)

func dividedTwo(dp [][2]int, inputData []int, n, m int) int {

	for i := 0; i < n; i++ {

		for j := 1; j <= m; j++ {
			p := 0
			if j-inputData[i] >= 0 {
				p = dp[j-inputData[i]][1] + inputData[i]
			}

			dp[j][0], dp[j][1] = max(dp[j][0], p), dp[j][0]
		}
	}

	return dp[m][0]
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

	dp := make([][2]int, sum/2+1)

	if dividedTwo(dp, intInputs, n, sum/2) == sum/2 {
		return "True"
	}
	return "False"
}
