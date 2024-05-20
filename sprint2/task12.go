package sprint2

import (
	"math"
	"strconv"
	"strings"
)

func Task12(res *[]byte) string {

	lines := strings.Split(string(*res), "\n")
	arrLine := strings.Split(lines[0], " ")

	n, _ := strconv.Atoi(arrLine[0])
	k, _ := strconv.Atoi(arrLine[1])

	prev1 := 1
	prev2 := 1
	a := math.Pow(10, float64(k))

	for n > 1 {
		p := prev1 + prev2
		v := p % int(a)

		prev1 = prev2
		prev2 = v

		n--
	}

	return strconv.Itoa(prev2) + "\n"
}
