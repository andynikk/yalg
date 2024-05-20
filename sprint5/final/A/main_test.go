package main

import (
	"fmt"
	"testing"
)

func TestFinalA(t *testing.T) {
	res := []byte("5\nalla 4 100\ngena 6 1000\ngosha 2 90\nrita 2 90\ntimofey 4 80")
	out := "gena\ntimofey\nalla\ngosha\nrita\n"
	answer := finalA(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("5\nalla 0 0\ngena 0 0\ngosha 0 0\nrita 0 0\ntimofey 0 0")
	out = "alla\ngena\ngosha\nrita\ntimofey\n"
	answer = finalA(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	fmt.Println("===============================")
	fmt.Println("===============================")
}
