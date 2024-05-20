package sprint1

import (
	"strconv"
	"strings"
)

func Final1(res *[]byte) ([]string, string) {

	lines := strings.Split(string(*res), "\n")

	arr := strings.Split(strings.Replace(lines[1], "\r", "", -1), " ")
	lines = nil

	zeroWas := false

	v := 0
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == "0" {
			v = 0
			zeroWas = true
			continue
		}

		switch zeroWas {
		case false:
			arr[i] = strconv.Itoa(i + 1)
			continue
		default:
			v++
			arr[i] = strconv.Itoa(v)
		}
	}

	zeroWas = false
	for i := 0; i < len(arr); i++ {
		if arr[i] == "0" {
			v = 0
			zeroWas = true
			continue
		}

		switch zeroWas {
		case true:
			v++
			vv, _ := strconv.Atoi(arr[i])
			if v < vv {
				arr[i] = strconv.Itoa(v)
			}
		default:
			continue
		}

	}

	return arr, " "
}
