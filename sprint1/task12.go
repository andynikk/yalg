package sprint1

import (
	"strings"
)

func Task12(res *[]byte) ([]string, string) {

	lines := strings.Split(string(*res), "\n")

	s := strings.Replace(lines[0], " ", "", -1)
	t := strings.Replace(lines[1], " ", "", -1)

	for i := 0; i < len(s); i++ {
		l := string(s[i])
		t = strings.Replace(t, l, "", 1)
	}

	return []string{t}, ""
}
