package sprint3

import (
	"strconv"
	"strings"
)

func median(arr1 []string, arr2 []string) float64 {
	firstArrayLen := len(arr1)
	secondArrayLen := len(arr2)
	allLen := firstArrayLen + secondArrayLen

	var mid int
	i := 0
	j := 0
	var k int
	mid = (allLen)/2 + 1

	if (allLen)%2 == 1 {
		var median float64

		for k < mid {
			n1 := 0
			if firstArrayLen != i {
				n1, _ = strconv.Atoi(arr1[i])
			}

			n2 := 0
			if secondArrayLen != j {
				n2, _ = strconv.Atoi(arr2[j])
			}

			if i < firstArrayLen && j < secondArrayLen {
				if n1 <= n2 {
					median = float64(n1)
					i++
				} else {
					median = float64(n2)
					j++
				}
			} else if i < firstArrayLen {
				median = float64(n1)
				i++
			} else {
				median = float64(n2)
				j++
			}
			k++

		}
		return median
	} else {
		var median1 float64
		var median2 float64

		for k < mid {
			n1 := 0
			if firstArrayLen != i {
				n1, _ = strconv.Atoi(arr1[i])
			}

			n2 := 0
			if secondArrayLen != j {
				n2, _ = strconv.Atoi(arr2[j])
			}

			median1 = median2
			if i < firstArrayLen && j < secondArrayLen {
				if n1 <= n2 {
					median2 = float64(n1)
					i++
				} else {
					median2 = float64(n2)
					j++
				}
			} else if i < firstArrayLen {
				median2 = float64(n1)
				i++
			} else {
				median2 = float64(n2)
				j++
			}
			k++
		}
		return (median1 + median2) / 2
	}
}

func Task13(res *[]byte) string {

	lines := strings.Split(string(*res), "\n")
	//n, _ := strconv.Atoi(lines[0])
	//m, _ := strconv.Atoi(lines[1])

	arr1 := strings.Split(strings.Replace(lines[2], "\r", "", -1), " ")
	arr2 := strings.Split(strings.Replace(lines[3], "\r", "", -1), " ")

	return strconv.FormatFloat(median(arr1, arr2), 'f', -1, 64) + "\n"
}
