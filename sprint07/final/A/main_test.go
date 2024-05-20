package main

import (
	"fmt"
	"testing"
)

func TestTask(t *testing.T) {

	res := []byte("5\n4 2 9 1 13\n")
	out := "3\n1 3 5\n"
	answer := journey(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("6\n1 2 4 8 16 32\n")
	out = "6\n1 2 3 4 5 6\n"
	answer = journey(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	fmt.Println("===============================")
	fmt.Println("===============================")

}
