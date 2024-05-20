package main

import (
	"fmt"
	"testing"
)

func TestTask(t *testing.T) {

	res := []byte("4\n1 5 7 1\n")
	out := "True"
	answer := sameAmounts(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("3\n2 10 9\n")
	out = "False"
	answer = sameAmounts(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("6\n7 9 3 4 6 7\n")
	out = "True"
	answer = sameAmounts(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	fmt.Println("===============================")
	fmt.Println("===============================")

}
