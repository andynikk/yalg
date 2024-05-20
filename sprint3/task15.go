package sprint3

import (
	"sort"
	"strconv"
	"strings"
)

func valueK(arr []int, min int, max int, k int) int {
	if min == max {
		return min
	}

	ras := 0
	if ras = max - min; ras < 0 {
		ras *= -1
	}
	mid := min + (ras / 2)

	kol := 0
	j := 0

	for i := 1; i < len(arr); i++ {
		for g := j; g < len(arr); g++ {
			if arr[i]-arr[g] <= mid {
				break
			}
			j++
		}
		kol += i - j
	}

	switch kol >= k {
	case true:
		min = valueK(arr, min, mid, k)
	default:
		min = valueK(arr, mid+1, max, k)
	}

	return min
}

func Task15(res *[]byte) string {

	lines := strings.Split(string(*res), "\n")
	//n, _ := strconv.Atoi(lines[0])
	k, _ := strconv.Atoi(lines[2])

	arr := strings.Split(strings.Replace(lines[1], "\r", "", -1), " ")
	//n := len(arr) * (len(arr) - 1) / 2

	arrInt := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		arrInt[i], _ = strconv.Atoi(arr[i])
	}
	sort.Ints(arrInt)

	min := 0
	max := arrInt[len(arrInt)-1] - arrInt[0]

	valK := valueK(arrInt, min, max, k)

	val := strconv.Itoa(valK)
	return val
}
