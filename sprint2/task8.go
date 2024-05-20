package sprint2

import (
	"strings"
)

func Task8(res *[]byte) string {

	lines := strings.Split(string(*res), "\n")
	if len(lines) == 0 {
		return "False"
	}

	line := lines[0]
	if len(line)%2 != 0 {
		return "False"
	}

	m := map[string]string{}
	m["("] = ")"
	m["["] = "]"
	m["{"] = "}"

	symbolsOpen := "({["

	stack := []string{}
	for i := 0; i < len(line); i++ {
		symbol := string(line[i])
		thisOpenSymbol := strings.Contains(symbolsOpen, symbol)
		if !thisOpenSymbol && len(stack) == 0 {
			return "False"
		}

		if thisOpenSymbol {
			stack = append(stack, symbol)
			continue
		}

		prevSymbol := stack[len(stack)-1]
		if symbol == m[prevSymbol] {
			stack = stack[:len(stack)-1]
			continue
		}
	}

	if len(stack) == 0 {
		return "True"
	} else {
		return "False"
	}

}
