package sprint3

import (
	"strconv"
	"strings"
)

func SortStrMax(array []string, k int) {
	countedValues := make([]int, k)
	for _, value := range array {
		intV, _ := strconv.Atoi(value)
		countedValues[intV]++
	}

	index := 0
	for i := len(countedValues) - 1; i >= 0; i-- {
		for j := 0; j < countedValues[i]; j++ {
			array[index] = strconv.Itoa(i)
			index++
		}
	}
}

func Task6(res *[]byte) string {

	lines := strings.Split(string(*res), "\n")
	o := strings.Split(strings.Replace(lines[1], "\r", "", -1), " ")
	SortStrMax(o, MaxInArr(o)+1)

	sumO := 0

	for i := 0; i < len(o); i++ {
		if sumO != 0 {
			break
		}
		for j := 1; j < len(o); j++ {
			if i == j {
				continue
			}
			if sumO != 0 {
				break
			}
			for g := 2; g < len(o); g++ {
				if j == g {
					continue
				}

				a, _ := strconv.Atoi(o[i])
				b, _ := strconv.Atoi(o[j])
				c, _ := strconv.Atoi(o[g])

				if (c < a+b && a <= b && b <= c) ||
					(a < c+b && c <= b && b <= a) ||
					(b < a+c && a <= c && c <= b) {

					sumO = a + b + c
					break
				}
			}
		}
	}

	return strconv.Itoa(sumO) + "\n"
}
