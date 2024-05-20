package sprint4

import (
	"strconv"
	"strings"
)

func slidingWindows(prefHash, hash []int, first, last, mod int) int {
	h := prefHash[last] - (prefHash[first-1] * hash[last-first+1])
	if h < 0 {
		h %= mod
		h += mod
	}
	return h
}

func Task6(res *[]byte) string {

	lines := strings.Split(string(*res), "\n")

	a, _ := strconv.Atoi(lines[0])
	m, _ := strconv.Atoi(lines[1])
	s := lines[2]
	sections := lines[4:]

	val := []string{}

	hash := make([]int, len(s)+1)
	hash[0] = 1
	for i := 1; i <= len(s); i++ {
		hash[i] = (hash[i-1] * a) % m
	}

	prefHash := make([]int, len(s)+1)
	prefHash[0] = 0
	for i := 1; i <= len(s); i++ {
		prefHash[i] = (prefHash[i-1]*a + int(s[i-1])) % m
	}

	inputs := len(sections)
	for i := 0; i < inputs; i++ {
		section := strings.Split(sections[i], " ")
		if len(section) <= 1 {
			continue
		}

		g1, _ := strconv.Atoi(section[0])
		g2, _ := strconv.Atoi(section[1])

		h := slidingWindows(prefHash, hash, g1, g2, m)
		val = append(val, strconv.Itoa(h))
	}

	return strings.Join(val, "\n") + "\n"
}
