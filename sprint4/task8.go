package sprint4

import (
	"strconv"
	"strings"
)

func Task8(res *[]byte) string {

	lines := strings.Split(string(*res), "\n")
	str := lines[0]

	if len(str) == 0 {
		return "0\n"
	}

	m := map[string]int{}
	symbol := str[0:1]

	//a b c a b c b b
	//0 1 2 3 4 5 6 7

	m[symbol] = 0
	maximum := 1
	finStr := symbol

	for j := 1; j < len(str); j++ {
		symbol = str[j : j+1]

		find := strings.Contains(finStr, symbol)

		switch find {
		case true:
			j, m[symbol] = m[symbol], j
			finStr = ""
		default:
			m[symbol] = j
			finStr += symbol

			if maximum < len(finStr) {
				maximum = len(finStr)
			}
		}
	}

	val := strconv.Itoa(maximum)
	return val + "\n"
}
