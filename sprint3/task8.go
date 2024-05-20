package sprint3

import (
	"strconv"
	"strings"
)

func Task8(res *[]byte) string {

	lines := strings.Split(string(*res), "\n")
	n, _ := strconv.Atoi(lines[0])

	arr := strings.Split(strings.Replace(lines[1], "\r", "", -1), " ")

	for i := 1; i < n; i++ {
		for j := 0; j < len(arr)-1; j++ {
			v1 := arr[j]
			v2 := arr[j+1]

			if v1[0] == v2[0] {
				arr[j], arr[j+1] = MaxString(v1, v2)
			} else if v1 < v2 {
				arr[j], arr[j+1] = v2, v1
			}
		}
	}

	println(strings.Join(arr, " ") + "\n")
	return strings.Join(arr, "") + "\n"
}

func MaxString(a, b string) (string, string) {

	maxLen := max(len(a), len(b))
	minLen := min(len(a), len(b))

	sp := ""
	for i := 0; i < maxLen-minLen; i++ {

		sp += string(a[0])
	}

	a1 := a
	if len(a) == minLen {
		a1 = a + sp
	}

	b1 := b
	if len(b) == minLen {
		b1 = b + sp
	}

	if a1 == b1 {
		if string(a[len(a)-1]) < string(b[len(b)-1]) {
			return b, a
		}
		return a, b
	} else if a1 > b1 {
		return a, b
	}
	return b, a
}
