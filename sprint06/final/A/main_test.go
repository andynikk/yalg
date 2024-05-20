package main

import (
	"fmt"
	"testing"
)

func TestTask(t *testing.T) {
	res := []byte("4 4\n1 2 5\n1 3 6\n2 4 8\n3 4 3\n")
	out := "19\n"
	answer := expensiveNetwork(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("3 3\n1 2 1\n1 2 2\n2 3 1\n")
	out = "3\n"
	answer = expensiveNetwork(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("2 0\n")
	out = "Oops! I did it again\n"
	answer = expensiveNetwork(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("10 20\n9 10 4\n2 2 4\n4 2 8\n10 5 3\n1 10 6\n7 4 2\n10 10 6\n3 7 4\n8 9 4\n8 10 7\n6 10 10\n2 8 8\n3 8 1\n3 10 3\n9 5 8\n10 10 2\n1 8 1\n10 1 5\n3 6 10\n9 10 8\n")
	out = "69\n"
	answer = expensiveNetwork(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	fmt.Println("===============================")
	fmt.Println("===============================")
}
