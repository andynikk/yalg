package sprint3

import (
	"strconv"
	"strings"
)

func Sort(s []string) {
	SortStrMin(s, MaxInArr(s)+1)
}

func MaxInArr(s []string) int {
	maxInArr := 0
	for _, v := range s {
		intV, _ := strconv.Atoi(v)
		if intV > maxInArr {
			maxInArr = intV
		}
	}
	return maxInArr
}

func SortStrMin(array []string, k int) {
	countedValues := make([]int, k)
	for _, value := range array {
		intV, _ := strconv.Atoi(value)
		countedValues[intV]++
	}

	index := 0
	for value := 0; value < k; value++ {
		for amount := 0; amount < countedValues[value]; amount++ {
			array[index] = strconv.Itoa(value)
			index++
		}
	}
}

func Task4(res *[]byte) string {

	lines := strings.Split(string(*res), "\n")
	//n, _ := strconv.Atoi(lines[0])
	//m, _ := strconv.Atoi(lines[2])

	greed := strings.Split(strings.Replace(lines[1], "\r", "", -1), " ")
	size := strings.Split(strings.Replace(lines[3], "\r", "", -1), " ")

	Sort(greed)
	Sort(size)

	happening := 0

	idx := 0
	for i := 0; i < len(greed); i++ {
		g, _ := strconv.Atoi(greed[i])
		for j := idx; j < len(size); j++ {
			idx++
			s, _ := strconv.Atoi(size[j])

			if g <= s {
				happening++
				break
			}
		}
	}

	return strconv.Itoa(happening) + "\n"
}
