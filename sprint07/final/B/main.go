package main

import (
	"strconv"
	"strings"
)

func dividedTwo(dp []bool, inputData []int, n, m int) bool {

	for i := 0; i < n; i++ {
		for j := m; j >= inputData[i]; j-- {
			dp[j] = dp[j] || dp[j-inputData[i]]
		}
	}

	return dp[m]

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

	dp := make([]bool, sum/2+1)
	dp[0] = true

	result := strconv.FormatBool(dividedTwo(dp, intInputs, n, sum/2))
	result = strings.ToUpper(result[:1]) + result[1:]

	return result

}
