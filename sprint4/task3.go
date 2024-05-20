package sprint4

import (
	"strings"
)

func Task3(res *[]byte) string {

	lines := strings.Split(string(*res), "\n")
	str1 := lines[0]
	str2 := lines[1]

	if len(str1) != len(str2) {
		return "NO\n"
	}

	m := map[string]string{}
	mCheck := map[string]string{}
	for i := 0; i < len(str1); i++ {
		s1 := string(str1[i])
		s2 := string(str2[i])

		if v, ok := m[s1]; ok && v != s2 {
			return "NO\n"
		}
		if v, ok := mCheck[s2]; ok && v != s1 {
			return "NO\n"
		}
		m[s1] = s2
		mCheck[s2] = s1
	}

	return "YES\n"
}
