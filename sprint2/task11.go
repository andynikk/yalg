package sprint2

import (
	"strconv"
	"strings"
)

func NumberFibonachi(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	return NumberFibonachi(n-1) + NumberFibonachi(n-2)
}

func Task11(res *[]byte) string {

	lines := strings.Split(string(*res), "\n")
	n, err := strconv.ParseInt(strings.Replace(lines[0], "\r", "", -1), 0, 0)
	if err != nil {
		panic(err)
	}

	val := []string{strconv.Itoa(NumberFibonachi(int(n)))}

	return strings.Join(val, "\n") + "\n"
}
