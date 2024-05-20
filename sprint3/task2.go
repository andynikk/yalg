package sprint3

import (
	"strconv"
	"strings"
)

type Cech struct {
	keys   map[int]string
	input  []string
	output []string
}

func (c *Cech) genBinary(prefix string, idx int) {

	if idx == len(c.input) {
		c.output = append(c.output, prefix)
		return
	}

	key := c.input[idx]
	for i := 0; i < len(key); i++ {
		c.genBinary(prefix+string(key[i]), idx+1)
	}
}

func Task2(res *[]byte) string {

	keys := map[int]string{}
	keys[2] = "abc"
	keys[3] = "def"
	keys[4] = "ghi"
	keys[5] = "jkl"
	keys[6] = "mno"
	keys[7] = "pqrs"
	keys[8] = "tuv"
	keys[9] = "wxyz"

	c := Cech{keys: keys}

	lines := strings.Split(string(*res), "\n")
	line := strings.Replace(lines[0], "\r", "", -1)

	for i := 0; i < len(line); i++ {
		k, _ := strconv.Atoi(string(line[i]))
		c.input = append(c.input, keys[k])
	}

	if len(c.input) > 0 {
		c.genBinary("", 0)
	}

	return strings.Join(c.output, " ")
}
