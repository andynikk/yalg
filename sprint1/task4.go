package sprint1

import (
	"strconv"
	"strings"
)

func Task4(res *[]byte) ([]string, string) {

	lines := strings.Split(string(*res), "\n")

	k, err := strconv.ParseInt(strings.Replace(lines[0], "\r", "", -1), 0, 0)
	if err != nil {
		panic(err)
	}

	arr := strings.Split(strings.Replace(lines[1], "\r", "", -1), " ")
	total := 0
	for i := 0; i < int(k); i++ {
		currentTemp, err := strconv.ParseInt(arr[i], 0, 0)
		if err != nil {
			panic(err)
		}

		prev := false
		if i-1 >= 0 {
			prevTemp, err := strconv.ParseInt(arr[i-1], 0, 0)
			if err != nil {
				panic(err)
			}
			if prevTemp < currentTemp {
				prev = true
			}
		} else {
			prev = true
		}

		after := false
		if i+1 < len(arr) {
			afterTemp, err := strconv.ParseInt(arr[i+1], 0, 0)
			if err != nil {
				panic(err)
			}
			if afterTemp < currentTemp {
				after = true
			}

		} else {
			after = true
		}

		if prev && after {
			total++
		}
	}

	val := []string{strconv.Itoa(total)}
	return val, ""
}
