package sprint1

import (
	"strconv"
	"strings"
)

func Task7(res *[]byte) ([]string, string) {

	lines := strings.Split(string(*res), "\n")
	k, err := strconv.ParseInt(strings.Replace(lines[0], "\r", "", -1), 0, 64)
	if err != nil {
		panic(err)
	}
	if k == 0 {
		return []string{"0"}, ""
	}

	o := 0

	total := ""
	for {
		if k == 0 {
			if o != 0 {
				total = strconv.Itoa(o) + total
			}
			break
		}
		o := k % 2
		k = k / 2

		total = strconv.Itoa(int(o)) + total
	}

	return []string{total}, ""
}
