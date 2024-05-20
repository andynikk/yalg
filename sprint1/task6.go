package sprint1

import (
	"strings"
	"unicode"
)

func Task6(res *[]byte) ([]string, string) {

	lines := strings.Split(string(*res), "\n")
	str := lines[0]

	nonLetter := func(c rune) bool { return !unicode.IsLetter(c) && !unicode.IsNumber(c) }
	words := strings.FieldsFunc(str, nonLetter)
	str = strings.Join(words, "")

	r := false
	lenHalfLine := len(str) / 2
	for i := 0; i <= lenHalfLine; i++ {
		l1 := strings.ToLower(string(str[i]))
		l2 := strings.ToLower(string(str[len(str)-1-i]))

		if l1 != l2 {
			break
		}
		r = true

	}

	val := []string{func() string {
		if r {
			return "True"
		}
		return "False"
	}()}
	return val, ""
}
