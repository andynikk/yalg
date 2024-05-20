package main

import (
	"fmt"
	"testing"
)

func TestTask(t *testing.T) {

	res := []byte("5 15\n3 8 1 2 5\n")
	out := "15\n"
	answer := goldLeprechauns(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("5 19\n10 10 7 7 4\n")
	out = "18\n"
	answer = goldLeprechauns(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	fmt.Println("===============================")
	fmt.Println("===============================")

}
