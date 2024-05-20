package sprint1

import (
	"strconv"
	"strings"
)

func Task9(res *[]byte) ([]string, string) {

	const d int64 = 4

	lines := strings.Split(string(*res), "\n")
	k, err := strconv.ParseInt(strings.Replace(lines[0], "\r", "", -1), 0, 64)
	if err != nil {
		panic(err)
	}

	if k == 1 {
		return []string{"True"}, ""
	}

	for {
		o := k % d
		if o != 0 {
			break
		}

		k = k / d

		switch k {
		case 0:
			return []string{"False"}, ""
		case 1:
			return []string{"True"}, ""
		default:
			continue
		}
	}

	return []string{"False"}, ""
}
