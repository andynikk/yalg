package main

import (
	"fmt"
	"testing"
)

func TestTask(t *testing.T) {
	res := []byte("one two three\n")
	out := "three two one\n"
	answer := lineReversal(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("hello\n")
	out = "hello\n"
	answer = lineReversal(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("may the force be with you\n")
	out = "you with be force the may\n"
	answer = lineReversal(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	fmt.Println("===============================")
	fmt.Println("===============================")
}
