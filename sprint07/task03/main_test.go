package main

import (
	"fmt"
	"testing"
)

func TestTask(t *testing.T) {
	res := []byte("10\n3\n8 1\n2 10\n4 5\n")
	out := "36\n"
	answer := goldRush(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("10000\n1\n4 20\n")
	out = "80\n"
	answer = goldRush(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	fmt.Println("===============================")
	fmt.Println("===============================")
}
