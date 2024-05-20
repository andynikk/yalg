package main

import (
	"fmt"
	"testing"
)

func TestTask(t *testing.T) {
	res := []byte("abacaba\n3\nqueue 2\ndeque 0\nstack 7\n")
	out := "dequeabqueueacabastack"
	answer := insertingLines(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("kukareku\n2\np 1\nq 2\n")
	out = "kpuqkareku"
	answer = insertingLines(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	fmt.Println("===============================")
	fmt.Println("===============================")
}
