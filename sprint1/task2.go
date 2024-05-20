package sprint1

import (
	"strconv"
	"strings"
)

func Task2(res *[]byte) ([]string, string) {

	lines := strings.Split(string(*res), "\n")
	arr := strings.Split(strings.Replace(lines[0], "\r", "", -1), " ")

	a1, err := strconv.ParseInt(strings.Replace(arr[0], "\r", "", -1), 0, 0)
	if err != nil {
		panic(err)
	}

	a2, err := strconv.ParseInt(strings.Replace(arr[1], "\r", "", -1), 0, 0)
	if err != nil {
		panic(err)
	}

	a3, err := strconv.ParseInt(strings.Replace(arr[2], "\r", "", -1), 0, 0)
	if err != nil {
		panic(err)
	}

	val := "FAIL"
	if (a1%2 == 0 && a2%2 == 0 && a3%2 == 0) ||
		(a1%2 != 0 && a2%2 != 0 && a3%2 != 0) {
		val = "WIN"
	}

	return []string{val}, " "
}
