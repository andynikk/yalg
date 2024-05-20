package sprint2

import (
	"strconv"
	"strings"
)

func Task1(res *[]byte) string {

	lines := strings.Split(string(*res), "\n")
	n, err := strconv.ParseInt(strings.Replace(lines[0], "\r", "", -1), 0, 0)
	if err != nil {
		panic(err)
	}

	m, err := strconv.ParseInt(strings.Replace(lines[1], "\r", "", -1), 0, 0)
	if err != nil {
		panic(err)
	}
	if n == 0 && m == 0 {
		return ""
	}

	arr := make([]string, m)
	for i := 2; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}
		marr := strings.Split(strings.Replace(lines[i], " ", "\n", -1), "\n")

		for j := 0; j < len(marr); j++ {
			arr[j] = arr[j] + marr[j] + " "
			if i+2 == len(lines) {
				arr[j] = arr[j][:len(arr[j])-1]
			}
		}
	}

	return strings.Join(arr, "\n") + "\n"

}
