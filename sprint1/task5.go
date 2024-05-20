package sprint1

import (
	"strconv"
	"strings"
)

func Task5(res *[]byte) ([]string, string) {

	lines := strings.Split(string(*res), "\n")

	longWord := struct {
		word string
		len  int
	}{}

	arr := strings.Split(strings.Replace(lines[1], "\r", "", -1), " ")
	for i := 0; i < len(arr); i++ {

		if longWord.len < len(arr[i]) {
			longWord.word = arr[i]
			longWord.len = len(arr[i])
		}

	}

	val := []string{longWord.word, strconv.Itoa(longWord.len)}
	return val, "\n"
}
