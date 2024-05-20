package main

import (
	"fmt"
	"testing"
)

func TestTask(t *testing.T) {
	res := []byte("6 3\n")
	out := "13\n"
	answer := jumpingStairs(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("7 7\n")
	out = "32\n"
	answer = jumpingStairs(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("2 2\n")
	out = "1\n"
	answer = jumpingStairs(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("62 44\n")
	out = "535806680\n"
	answer = jumpingStairs(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	fmt.Println("===============================")
	fmt.Println("===============================")
}
