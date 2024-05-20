package sprint3

import (
	"strconv"
	"strings"
)

func Task12(res *[]byte) string {

	lines := strings.Split(string(*res), "\n")
	n, _ := strconv.Atoi(lines[2])

	arr := strings.Split(strings.Replace(lines[1], "\r", "", -1), " ")
	keep := false

	val := []string{}
	for i := 0; i < len(arr); i++ {
		if keep {
			break
		}

		switch len(val) {
		case 0:
			o, _ := strconv.Atoi(arr[i])
			if o >= n {
				val = append(val, strconv.Itoa(i+1))
			}
		default:
			o, _ := strconv.Atoi(arr[i])
			o = o - n*len(val)

			if o >= n {
				val = append(val, strconv.Itoa(i+1))
				keep = true
			}
		}
	}

	lenVal := len(val)
	for i := 0; i < 2-lenVal; i++ {
		val = append(val, "-1")
	}

	return strings.Join(val, " ")
}
