package sprint1

import (
	"strconv"
	"strings"
)

func Task11(res *[]byte) ([]string, string) {

	lines := strings.Split(string(*res), "\n")

	strX := strings.Replace(lines[1], " ", "", -1)
	strK := strings.Replace(lines[2], " ", "", -1)

	maxLen := max(len(strX), len(strK))
	var cel int64 = 0

	v1 := 0
	v2 := 0
	ost := 0

	val := make([]string, maxLen+1)
	for i := 0; i < maxLen; i++ {

		idx1 := len(strX) - 1 - i
		idx2 := len(strK) - 1 - i

		switch idx1 >= 0 {
		case true:
			vPr, _ := strconv.ParseInt(string(strX[idx1]), 0, 0)
			v1 = int(vPr)
		default:
			v1 = 0
		}

		switch idx2 >= 0 {
		case true:
			vPr, _ := strconv.ParseInt(string(strK[idx2]), 0, 0)
			v2 = int(vPr)
		default:
			v2 = 0
		}

		v := v1 + v2 + ost

		switch v >= 10 {
		case true:
			cel = int64(v - 10)
			ost = int(int64(v)-cel) / 10
		default:
			cel = int64(v)
			ost = 0
		}

		val[maxLen-i] = strconv.Itoa(int(cel))
	}

	if ost == 1 {
		val[0] = strconv.Itoa(ost)
	} else {
		val = val[1:]
	}

	return val, " "
}
