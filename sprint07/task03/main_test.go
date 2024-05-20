package main

import (
	"fmt"
	"testing"
)

func TestTask(t *testing.T) {
	res := []byte("5\n9 10\n9.3 10.3\n10 11\n10.3 11.3\n11 12\n")
	out := "3\n9 10\n10 11\n11 12\n"
	answer := schedules(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("3\n9 10\n11 12.25\n12.15 13.3\n")
	out = "2\n9 10\n11 12.25\n"
	answer = schedules(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("7\n19 19\n7 14\n12 14\n8 22\n22 23\n5 21\n9 23\n")
	out = "3\n7 14\n19 19\n22 23\n"
	answer = schedules(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	fmt.Println("===============================")
	fmt.Println("===============================")
}
