package sprint3

import (
	"strconv"
	"strings"
)

const sOpen = "("
const sClose = ")"

type arr struct {
	g []string
}

func genBinary(n int, prefix string, a *arr) {
	if n == 0 {
		g := append(a.g, prefix)
		a.g = g
	} else {
		genBinary(n-1, prefix+sOpen, a)
		genBinary(n-1, prefix+sClose, a)
	}
}

func Task1(res *[]byte) string {

	lines := strings.Split(string(*res), "\n")
	n, err := strconv.Atoi(strings.Replace(lines[0], "\r", "", -1))
	if err != nil {
		panic(err)
	}
	n *= 2

	val := ""

	a := arr{g: []string{}}
	genBinary(n, "", &a)

	for i := 0; i < len(a.g); i++ {
		if (string(a.g[i][0]) != sOpen || string(a.g[i][len(a.g[i])-1]) != sClose) ||
			(strings.Count(a.g[i], sOpen) != strings.Count(a.g[i], sClose)) {

			continue
		}

		open := 0
		closed := 0

		keep := false

		for j := 0; j < len(a.g[i]); j++ {
			symbol := string(a.g[i][j])
			switch symbol {
			case sOpen:
				open++
			default:
				closed++
			}

			if open < closed {
				keep = true
				break
			}

		}
		if keep {
			continue
		}

		val += a.g[i] + "\n"
	}

	return val
}
