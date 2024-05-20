package main

import (
	"fmt"
	"testing"
)

func TestTask(t *testing.T) {
	res := []byte("3\nRB\nR\n")
	out := "NO\n"
	answer := railways(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("4\nBBB\nRB\nB\n")
	out = "YES\n"
	answer = railways(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("5\nRRRB\nBRR\nBR\nR\n")
	out = "NO\n"
	answer = railways(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("10\nRRBRRBRRR\nBBBBBBRB\nBBRBRRR\nRRBRRR\nRBRRR\nBBRR\nRRR\nRR\nB\n")
	out = "YES\n"
	answer = railways(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	fmt.Println("===============================")
	fmt.Println("===============================")
}
