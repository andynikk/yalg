package main

import (
	"fmt"
	"testing"
)

func TestFuncTask11(t *testing.T) {
	res := []byte("3\n-1 0\n1 0\n2 5\n3\n10 0\n20 0\n22 5\n")
	out := "3"
	answer := task11(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("3\n-1 0\n1 0\n0 5\n3\n10 0\n20 0\n20 5\n")
	out = "2"
	answer = task11(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("10\n-54 -23\n-63 -26\n15 -47\n-12 -29\n-60 -21\n67 -106\n66 -99\n48 -96\n9 -60\n-51 -60\n15\n-32 -40\n57 -114\n58 -99\n10 -26\n64 -82\n-44 -34\n62 -80\n62 -113\n56 -112\n6 -20\n60 -119\n29 -91\n12 -47\n25 -88\n-6 -50\n")
	out = "7"
	answer = task11(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("10\n-87 -73\n-68 -71\n31 -27\n50 -8\n44 -7\n-33 -106\n-74 -50\n-30 -37\n-42 -92\n-78 -44\n15\n-106 -55\n-37 -46\n-91 -40\n3 -56\n-22 -58\n-85 -49\n-111 -74\n43 -17\n60 -3\n-81 -73\n-40 -40\n21 -30\n-91 -35\n-31 -62\n-51 -103\n")
	out = "10"
	answer = task11(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("10\n-91 -21\n81 -16\n-100 -51\n-97 -69\n-3 -5\n9 -4\n-90 -66\n-93 -53\n-99 -68\n-92 -69\n15\n115 -17\n-3 -24\n-68 -31\n22 -15\n-69 -4\n-104 -59\n-56 3\n-50 -9\n-80 -7\n-76 -72\n-116 -62\n92 -33\n-121 -60\n-81 -77\n22 -44\n")
	out = "7"
	answer = task11(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}
}
