package main

import (
	"fmt"
	"testing"
)

func TestFinalA(t *testing.T) {
	res := []byte("5\nalla 4 100\ngena 6 1000\ngosha 2 90\nrita 2 90\ntimofey 4 80")
	out := "gena\ntimofey\nalla\ngosha\nrita\n"
	answer := ""
	//answer := finalB(&res)
	finalB()
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	_ = res

	fmt.Println("===============================")
	fmt.Println("===============================")
}
