package main

import (
	"fmt"
	"testing"
)

func TestTask(t *testing.T) {

	res := []byte("5\n4 9 2 4 6\n7\n9 4 0 0 2 8 4\n")
	out := "3\n1 3 4\n2 5 7\n"
	answer := horoscopes(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("4\n1 1 1 1\n2\n2 2\n")
	out = "0\n"
	answer = horoscopes(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("8\n1 2 1 9 1 2 1 9\n5\n9 9 1 9 9\n")
	out = "3\n3 4 8\n3 4 5\n"
	answer = horoscopes(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	fmt.Println("===============================")
	fmt.Println("===============================")

}
