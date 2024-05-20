package sprint1

import (
	"strconv"
	"strings"
)

func Task1(res *[]byte) ([]string, string) {

	lines := strings.Split(string(*res), "\n")
	arr := strings.Split(strings.Replace(lines[0], "\r", "", -1), " ")

	a, err := strconv.ParseInt(strings.Replace(arr[0], "\r", "", -1), 0, 0)
	if err != nil {
		panic(err)
	}

	x, err := strconv.ParseInt(strings.Replace(arr[1], "\r", "", -1), 0, 0)
	if err != nil {
		panic(err)
	}

	b, err := strconv.ParseInt(strings.Replace(arr[2], "\r", "", -1), 0, 0)
	if err != nil {
		panic(err)
	}

	c, err := strconv.ParseInt(strings.Replace(arr[3], "\r", "", -1), 0, 0)
	if err != nil {
		panic(err)
	}

	val := a*(x*x) + b*x + c
	return []string{strconv.Itoa(int(val))}, " "
}
