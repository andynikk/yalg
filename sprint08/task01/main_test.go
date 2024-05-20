package main

import (
	"fmt"
	"testing"
)

func TestTask(t *testing.T) {
	res := []byte("6\n7 1 5 3 6 4")
	out := "7\n"
	answer := maxProfit(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("5\n1 2 3 4 5")
	out = "4\n"
	answer = maxProfit(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("6\n1 12 12 16 1 8")
	out = "22\n"
	answer = maxProfit(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	fmt.Println("===============================")
	fmt.Println("===============================")
}
