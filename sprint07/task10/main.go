package main

import (
	"slices"
	"strconv"
	"strings"
)

func findLIS(n int, ratings []int) ([]string, int) {
	dp := make([]int, n)
	parent := make([]int, n)
	maxLength := 0
	endIndex := 0

	for i := 0; i < n; i++ {
		dp[i] = 1
		parent[i] = -1
		for j := 0; j < i; j++ {
			if ratings[i] > ratings[j] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
				parent[i] = j
			}
		}
		if dp[i] > maxLength {
			maxLength = dp[i]
			endIndex = i
		}
	}

	result := make([]string, 0)
	for endIndex != -1 {
		result = append(result, strconv.Itoa(endIndex+1))
		endIndex = parent[endIndex]
	}
	slices.Reverse(result)

	return result, maxLength
}

// <template>
func journey(res *[]byte) string {

	lines := strings.Split(strings.Replace(string(*res), "\r", "", -1), "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	n, _ := strconv.Atoi(strings.Split(lines[0], " ")[0])
	line := strings.Split(lines[1], " ")

	ratings := make([]int, n)
	for i := range line {
		ratings[i], _ = strconv.Atoi(line[i])
	}

	subsequence, length := findLIS(n, ratings)

	return strconv.Itoa(length) + "\n" + strings.Join(subsequence, " ") + "\n"
}
