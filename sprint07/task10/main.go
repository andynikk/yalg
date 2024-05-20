package main

import (
	"strconv"
	"strings"
)

func maxFlowers(nI, mJ int, field [][]int, dp [][]int) (int, string) {

	for j := 1; j < mJ; j++ {
		dp[nI-1][j] += dp[nI-1][j-1]
		field[nI-1][j] = dp[nI-1][j]
	}

	for i := nI - 2; i >= 0; i-- {
		dp[i][0] += dp[i+1][0]
		field[i][0] = dp[i][0]
	}

	for i := nI - 2; i >= 0; i-- {
		for j := 1; j < mJ; j++ {
			dp[i][j] = field[i][j] + max(field[i+1][j], field[i][j-1])
			field[i][j] = dp[i][j]
		}
	}

	kI := 0
	kJ := mJ - 1

	way := make([]string, 0)
	for {

		if kI == nI-1 && kJ == 0 {
			break
		}

		if kI == nI-1 {
			way = append(way, "R")
			kJ--
			if kJ <= 0 {
				break
			}
			continue
		}

		if kJ == 0 {
			way = append(way, "U")
			kI++
			if kI >= nI-1 {
				break
			}
			continue
		}

		valueDown := field[kI+1][kJ]
		valueLeft := field[kI][kJ-1]

		if valueDown == max(valueDown, valueLeft) {
			way = append(way, "U")
			kI++
		} else {
			way = append(way, "R")
			kJ--
		}

	}

	for i := 0; i < len(way)/2; i++ {
		way[i], way[len(way)-i-1] = way[len(way)-i-1], way[i]
	}

	return dp[0][mJ-1], strings.Join(way, "")
}

// <template>
func fieldFlowers(res *[]byte) string {

	lines := strings.Split(strings.Replace(string(*res), "\r", "", -1), "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	n, _ := strconv.Atoi(strings.Split(lines[0], " ")[0])
	m, _ := strconv.Atoi(strings.Split(lines[0], " ")[1])

	dp := make([][]int, n)
	field := make([][]int, n)
	for i := 1; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}

		a := make([]int, m)
		b := make([]int, m)
		for index, l := range []rune(lines[i]) {
			v, _ := strconv.Atoi(string(l))

			a[index] = v
			b[index] = v
		}
		field[i-1] = a
		dp[i-1] = b
	}

	maximum, way := maxFlowers(n, m, field, dp)

	return strconv.Itoa(maximum) + "\n" + way + "\n"
}
