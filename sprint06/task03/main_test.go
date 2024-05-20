package main

import (
	"fmt"
	"testing"
)

func TestTask(t *testing.T) {
	res := []byte("4 4\n3 2\n4 3\n1 4\n1 2\n3\n")
	out := "3 2 1 4 "
	answer := task03(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("2 1\n1 2\n1\n")
	out = "1 2 "
	answer = task03(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	fmt.Println("===============================")
	fmt.Println("===============================")
}
