package sprint3

import (
	"strconv"
	"strings"
)

func Task3(res *[]byte) string {

	lines := strings.Split(string(*res), "\n")
	n, _ := strconv.Atoi(lines[0])

	arr := strings.Split(strings.Replace(lines[1], "\r", "", -1), " ")

	val := []string{}
	for i := 1; i < n; i++ {
		keep := true
		for j := 0; j < len(arr)-1; j++ {
			v1, _ := strconv.Atoi(arr[j])
			v2, _ := strconv.Atoi(arr[j+1])

			if v1 > v2 {
				keep = false
				arr[j], arr[j+1] = strconv.Itoa(v2), strconv.Itoa(v1)
			}
		}
		if keep {
			if len(val) == 0 {
				val = append(val, strings.Join(arr, " "))
			}
			break
		}
		val = append(val, strings.Join(arr, " "))
	}

	return strings.Join(val, "\n") + "\n"
}
