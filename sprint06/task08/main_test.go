package main

import (
	"fmt"
	"testing"
)

func TestTask(t *testing.T) {
	res := []byte("6 8\n2 6\n1 6\n3 1\n2 5\n4 3\n3 2\n1 2\n1 4\n")
	out := "0 11\n1 6\n8 9\n7 10\n2 3\n4 5\n"
	answer := task08(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("3 2\n1 2\n2 3\n")
	out = "0 5\n1 4\n2 3\n"
	answer = task08(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	fmt.Println("===============================")
	fmt.Println("===============================")
}
