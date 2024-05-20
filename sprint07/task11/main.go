package main

import (
	"slices"
	"strconv"
	"strings"
)

func maxSequence(sequenceA, sequenceB []int) (int, []string, []string) {
	dp := make([][]int, len(sequenceA)+1)
	for i := 0; i <= len(sequenceA); i++ {
		row := make([]int, len(sequenceB)+1)
		dp[i] = row
	}

	for i := 1; i <= len(sequenceA); i++ {
		for j := 1; j <= len(sequenceB); j++ {
			if sequenceA[i-1] == sequenceB[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	answerA := make([]string, 0)
	answerB := make([]string, 0)

	i, j := len(sequenceA), len(sequenceB)

	for i != 0 && j != 0 {
		if sequenceA[i-1] == sequenceB[j-1] {
			answerA = append(answerA, strconv.Itoa(j))
			answerB = append(answerB, strconv.Itoa(i))
			i--
			j--
		} else {
			if dp[i][j] == dp[i-1][j] {
				i--
			} else {
				j--
			}
		}
	}
	slices.Reverse(answerA)
	slices.Reverse(answerB)

	return dp[len(sequenceA)][len(sequenceB)], answerA, answerB
}

func horoscopes(res *[]byte) string {

	lines := strings.Split(strings.Replace(string(*res), "\r", "", -1), "\n")

	lenSequenceA, _ := strconv.Atoi(lines[0])
	sequenceA := make([]int, lenSequenceA)
	for i := range strings.Split(lines[1], " ") {
		sequenceA[i], _ = strconv.Atoi(strings.Split(lines[1], " ")[i])
	}

	lenSequenceB, _ := strconv.Atoi(lines[2])
	sequenceB := make([]int, lenSequenceB)
	for i := range strings.Split(lines[3], " ") {
		sequenceB[i], _ = strconv.Atoi(strings.Split(lines[3], " ")[i])
	}

	maximum, answerA, answerB := maxSequence(sequenceA, sequenceB)

	strA := ""
	if len(answerA) != 0 {
		strA = strings.Join(answerA, " ")
		strA = strings.TrimSuffix(strA, " ") + "\n"
	}

	strB := ""
	if len(answerB) != 0 {
		strB = strings.Join(answerB, " ")
		strB = strings.TrimSuffix(strB, " ") + "\n"
	}

	return strconv.Itoa(maximum) + "\n" + strB + strA
}
