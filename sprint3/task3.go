package sprint3

import (
	"strings"
)

func Task3(res *[]byte) string {

	lines := strings.Split(string(*res), "\n")
	s := lines[0]
	t := lines[1]

	sov := 0

	idx := 0
	for i := 0; i < len(s); i++ {
		symbol1 := string(s[i])
		for j := idx; j < len(t); j++ {
			idx++
			if symbol1 == string(t[j]) {
				sov++
				break
			}
		}
	}

	if sov == len(s) {
		return "True\n"
	}
	return "False\n"
}
