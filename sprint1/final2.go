package sprint1

import (
	"strconv"
	"strings"
)

func Task2(res *[]byte) []string {

	const sch int = 4
	lines := strings.Split(string(*res), "\n")

	k, err := strconv.ParseInt(strings.Replace(lines[0], "\r", "", -1), 0, 0)
	if err != nil {
		panic(err)
	}
	maxPress := k * 2

	arr := lines[1:]
	lines = nil

	hash := map[string]int{}
	for i := 0; i < sch; i++ {
		for j := 0; j < sch; j++ {
			symbol := string(arr[i][j])
			if symbol == "." {
				continue
			}

			val, ok := hash[symbol]
			if ok {
				hash[symbol] = val + 1
			} else {
				hash[symbol] = 1
			}
		}
	}

	winningPoint := 0
	for _, val := range hash {
		if val <= int(maxPress) {
			winningPoint++
		}
	}

	return []string{strconv.Itoa(winningPoint)}
}
