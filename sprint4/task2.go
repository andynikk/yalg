package sprint4

import (
	"strconv"
	"strings"
)

func Task2(res *[]byte) string {

	lines := strings.Split(string(*res), "\n")
	arr := strings.Split(strings.Replace(lines[1], " ", "\n", -1), "\n")

	m := map[int][]int{}
	o := 0
	max := 0
	m[0] = []int{0, 0}

	for i := 0; i < len(arr); i++ {
		pp := i + 1
		if arr[i] == "0" {
			o--
		} else {
			o++
		}

		_, ok := m[o]
		if !ok && o != 0 {
			m[o] = []int{pp, pp}
		} else {
			m[o][1] = pp
		}

		if m[o][1]-m[o][0] > max {
			max = m[o][1] - m[o][0]
		}
	}

	return strconv.Itoa(max) + "\n"
}
