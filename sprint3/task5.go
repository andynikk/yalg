package sprint3

import (
	"strconv"
	"strings"
)

func Task5(res *[]byte) string {

	lines := strings.Split(string(*res), "\n")

	t := strings.Split(strings.Replace(lines[0], "\r", "", -1), " ")
	h := strings.Split(strings.Replace(lines[1], "\r", "", -1), " ")

	//n, _ := strconv.Atoi(t[0])
	b, _ := strconv.Atoi(t[1])

	Sort(h)

	houses := 0

	for i := 0; i < len(h); i++ {
		p, _ := strconv.Atoi(h[i])
		b -= p

		switch b < 0 {
		case true:
			break
		default:
			houses++
		}
	}

	return strconv.Itoa(houses)
}
