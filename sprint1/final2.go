package sprint1

import (
	"strconv"
	"strings"
)

func Final2(res *[]byte) ([]string, string) {

	const sch int = 4
	lines := strings.Split(string(*res), "\n")

	k, err := strconv.Atoi(strings.Replace(lines[0], "\r", "", -1))
	if err != nil {
		panic(err)
	}
	maxPress := k * 2

	arr := lines[1:]

	win := 0

	arrNum := make([]int, 9)
	for i := 0; i < sch; i++ {
		for j := 0; j < sch; j++ {
			symbol := string(arr[i][j])
			if symbol == "." {
				continue
			}

			v, _ := strconv.Atoi(symbol)
			idx := v - 1

			if arrNum[idx] == 0 {
				win++
			}

			arrNum[idx] = arrNum[idx] + 1

			if arrNum[idx]-maxPress == 1 {
				win--
			}
		}
	}

	return []string{strconv.Itoa(win)}, " "
}
