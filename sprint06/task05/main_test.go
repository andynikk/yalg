package main

import (
	"fmt"
	"testing"
)

func TestTask(t *testing.T) {
	res := []byte("6 3\n1 2\n6 5\n2 3\n")
	out := "3\n1 2 3 \n4 \n5 6 \n"
	answer := task05(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("2 0\n")
	out = "2\n1 \n2 \n"
	answer = task05(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("4 3\n2 3\n2 1\n4 3\n")
	out = "1\n1 2 3 4 \n"
	answer = task05(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	fmt.Println("===============================")
	fmt.Println("===============================")
}
