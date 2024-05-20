package sprint4

import (
	"sort"
	"strconv"
	"strings"
)

func SortString(s string) string {
	arrS := strings.Split(s, "")
	sort.Strings(arrS)
	return strings.Join(arrS, "")
}

func Task9(res *[]byte) string {

	lines := strings.Split(string(*res), "\n")
	arr := strings.Split(strings.Replace(lines[1], "\r", "", -1), " ")

	valArr := make([][]int, len(arr))

	m := map[string][]int{}
	for i := 0; i < len(arr); i++ {

		arr[i] = SortString(arr[i])
		if v, ok := m[arr[i]]; ok {
			m[arr[i]] = append(v, i)
		} else {
			m[arr[i]] = []int{i}
		}
	}

	for _, v := range m {
		first := v[0]
		valArr[first] = v
	}

	val := []string{}
	for _, v := range valArr {
		if len(v) == 0 {
			continue
		}

		podVal := []string{}
		for i := 0; i < len(v); i++ {
			podVal = append(podVal, strconv.Itoa(v[i]))
		}
		val = append(val, strings.Join(podVal, " "))
	}

	return strings.Join(val, "\n")
}
