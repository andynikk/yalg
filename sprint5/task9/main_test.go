package main

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	res := []byte("2")
	out := "2"
	answer := Solution(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("3")
	out = "5"
	answer = Solution(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("4")
	out = "14"
	answer = Solution(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

}
