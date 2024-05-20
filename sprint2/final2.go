package sprint2

import (
	"strconv"
	"strings"
)

func Final2(res *[]byte) string {

	lines := strings.Split(string(*res), "\n")
	commands := lines[0]
	val := []string{}

	specSymbol := "+-*/"
	hash := []int{}

	command := strings.Split(commands, " ")

	for i := 0; i < len(command); i++ {
		symbol := command[i]
		if strings.Contains(specSymbol, symbol) {
			v1 := hash[len(hash)-2]
			v2 := hash[len(hash)-1]

			v := 0
			switch symbol {
			case "+":
				v = v1 + v2
			case "-":
				v = v1 - v2
			case "*":
				v = v1 * v2
			case "/":
				v = floorDiv(v1, v2)
			}

			hash = hash[0 : len(hash)-2]
			hash = append(hash, v)
			continue
		}

		h, _ := strconv.Atoi(symbol)
		hash = append(hash, h)

	}

	val = append(val, strconv.Itoa(hash[len(hash)-1]))

	return strings.Join(val, "\n") + "\n"
}

func floorDiv(x, y int) int {
	q := x / y

	if x^y < 0 && q*y != x {
		return q - 1
	}
	return q

}
