package main

import (
	"slices"
	"strings"
)

// <template>
func lineReversal(res *[]byte) string {
	lines := strings.Split(string(*res), "\n")

	text := strings.Split(lines[0], " ")
	slices.Reverse(text)

	return strings.Join(text, " ") + "\n"
}
