package sprint1

import (
	"strconv"
	"strings"
)

func Task8(res *[]byte) ([]string, string) {

	lines := strings.Split(string(*res), "\n")
	d1 := lines[0]
	d2 := lines[1]

	maxLen := max(len(d1), len(d2))
	var cel int64 = 0

	v1 := 0
	v2 := 0
	ost := 0

	total := ""
	for i := 0; i < maxLen; i++ {

		idx1 := len(d1) - 1 - i
		idx2 := len(d2) - 1 - i

		switch idx1 >= 0 {
		case true:
			vPr, _ := strconv.ParseInt(string(d1[idx1]), 0, 0)
			v1 = int(vPr)
		default:
			v1 = 0
		}

		switch idx2 >= 0 {
		case true:
			vPr, _ := strconv.ParseInt(string(d2[idx2]), 0, 0)
			v2 = int(vPr)
		default:
			v2 = 0
		}

		v := v1 + v2 + ost

		switch v {
		case 0:
			cel = 0
			ost = 0
		case 1:
			cel = 1
			ost = 0
		case 2:
			cel = 0
			ost = 1
		default:
			cel = 1
			ost = 1
		}

		total = strconv.Itoa(int(cel)) + total
	}

	if ost == 1 {
		total = "1" + total
	}

	val := []string{total}
	return val, ""
}
