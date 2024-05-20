package sprint3

import (
	"strconv"
	"strings"
)

func Task7(res *[]byte) string {

	lines := strings.Split(string(*res), "\n")
	k, _ := strconv.Atoi(lines[0])
	if k == 0 {
		return ""
	}
	d := strings.Split(strings.Replace(lines[1], "\r", "", -1), " ")

	idxs := make([]int, 3)

	for i := 0; i < len(d); i++ {
		v, _ := strconv.Atoi(d[i])
		idxs[v]++
	}

	idx := 0
	for i := 0; i < len(idxs); i++ {
		for j := 0; j < idxs[i]; j++ {
			d[idx] = strconv.Itoa(i)
			idx++
		}
	}

	return strings.Join(d, " ")
}
