package sprint3

import (
	"sort"
	"strconv"
	"strings"
)

func mod(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

func splitB(arr []int) int {

	b := 0
	min := 0
	step := 1

	for i := 0; i < len(arr); i++ {

		if i+step >= len(arr) {
			b++
			return b
		}

		tmpArr := arr[i : i+step]
		if len(tmpArr) == 1 {
			if tmpArr[0] == min {
				min++
				b++
			} else {
				i--
				step++
			}
			continue
		}

		sort.Ints(tmpArr)
		allPlusOne := true
		findMin := false
		maxInt := 0
		for j := 1; j < len(tmpArr); j++ {
			if mod(tmpArr[j-1]-tmpArr[j]) != 1 {
				allPlusOne = false
			}
			if tmpArr[j-1] == min || tmpArr[j] == min {
				findMin = true
			}
			if tmpArr[j-1] > maxInt || tmpArr[j] > maxInt {
				maxInt = max(tmpArr[j-1], tmpArr[j])
			}
		}
		if allPlusOne && findMin {
			min = maxInt + 1
			i += len(tmpArr) - 1
			step = 1
			b++
		} else {
			i--
			step++
		}
	}

	return b
}

func Task16(res *[]byte) string {

	lines := strings.Split(string(*res), "\n")
	arr := strings.Split(strings.Replace(lines[1], "\r", "", -1), " ")

	if len(arr) == 1 {
		return strconv.Itoa(1)
	}

	arrInt := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		arrInt[i], _ = strconv.Atoi(arr[i])
	}

	b := splitB(arrInt)

	return strconv.Itoa(b)
}
