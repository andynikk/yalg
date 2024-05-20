package sprint4

import (
	"strings"
)

func Task1(res *[]byte) string {

	//aa := math.Mod(1499, 137)
	//_ = aa

	lines := strings.Split(string(*res), "\n")

	commands := lines[1:]
	arr := []string{}

	for _, v := range commands {
		if v == "" {
			continue
		}
		find := false
		for _, a := range arr {
			if a == v {
				find = true
				break
			}
		}
		if !find {
			arr = append(arr, v)
		}
	}

	return strings.Join(arr, "\n") + "\n"
}
