package main

import (
	"fmt"
	"testing"
)

func TestTask(t *testing.T) {

	res := []byte("2 3\n101\n110\n")
	out := "3\n"
	answer := fieldFlowers(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("3 3\n100\n110\n001\n")
	out = "2\n"
	answer = fieldFlowers(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("1 1\n1\n\n")
	out = "1\n"
	answer = fieldFlowers(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	fmt.Println("===============================")
	fmt.Println("===============================")

}
