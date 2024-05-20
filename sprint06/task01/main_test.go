package main

import (
	"fmt"
	"testing"
)

func TestTask(t *testing.T) {
	res := []byte("5 3\n1 3\n2 3\n5 2\n")
	out := "1 3 \n1 3 \n0 \n0 \n1 2 \n"
	answer := task01(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("10 55\n1 3\n1 5\n1 6\n1 9\n1 10\n2 5\n2 7\n2 8\n2 10\n3 1\n3 6\n3 7\n4 1\n4 2\n4 3\n4 5\n4 7\n4 8\n4 10\n5 3\n5 4\n5 6\n5 7\n5 8\n5 9\n5 10\n6 1\n6 2\n6 3\n6 5\n6 8\n6 9\n6 10\n7 1\n7 3\n7 8\n8 1\n8 2\n8 3\n8 4\n8 5\n8 6\n8 7\n8 9\n9 2\n9 4\n9 5\n9 6\n9 8\n10 1\n10 2\n10 4\n10 5\n10 6\n10 7\n")
	out = "5 3 5 6 9 10 \n4 5 7 8 10 \n3 1 6 7 \n7 1 2 3 5 7 8 10 \n7 3 4 6 7 8 9 10 \n7 1 2 3 5 8 9 10 \n3 1 3 8 \n8 1 2 3 4 5 6 7 9 \n5 2 4 5 6 8 \n6 1 2 4 5 6 7 \n"
	answer = task01(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	fmt.Println("===============================")
	fmt.Println("===============================")
}
