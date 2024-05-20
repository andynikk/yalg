package main

import (
	"fmt"
	"testing"
)

func TestTask(t *testing.T) {
	res := []byte("5 4\n2 1\n4 5\n4 3\n3 2\n2\n")
	out := "3\n"
	answer := task07(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("3 3\n3 1\n1 2\n2 3\n1\n")
	out = "1\n"
	answer = task07(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("6 8\n6 1\n1 3\n5 1\n3 5\n3 4\n6 5\n5 2\n6 2\n4\n")
	out = "3\n"
	answer = task07(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	fmt.Println("===============================")
	fmt.Println("===============================")
}
