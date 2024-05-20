package main

import (
	"fmt"
	"testing"
)

func TestTask(t *testing.T) {

	res := []byte("abacaba\nabaabc\n")
	out := "2\n"
	answer := distanceLowenstein(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("innokentiy\ninnnokkentia\n")
	out = "3\n"
	answer = distanceLowenstein(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("r\nx\n")
	out = "1\n"
	answer = distanceLowenstein(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	fmt.Println("===============================")
	fmt.Println("===============================")

}
