package sprint4

import (
	"fmt"
	"math/rand"
)

func createRandomString() string {
	letters := []rune("abcdefghijklmnopqrstuvwxyz")
	b := make([]rune, 12)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func hash(s string) int {
	a := 1000
	m := 123987123

	h := 0
	k := 1

	for i := 0; i < len(s); i++ {
		h += k * int(s[len(s)-1-i])
		k = (k * a) % m
		h = h % m
	}

	return h
}

func Task5() {

	m := map[int]string{}

	for {
		str := createRandomString()
		h := hash(str)
		if v, ok := m[h]; ok {
			fmt.Printf("%s\n%s", str, v)
			break
		}
		m[h] = str
	}

}
