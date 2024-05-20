package main

import (
	"fmt"
	"strings"
	"testing"
	"yalg/sprint1"
)

func TestFuncSprint1(t *testing.T) {
	testFinal1(t)
	//testFinal2(t)
	//testTask1(t)
	//testTask2(t)
	//testTask3(t)
	//testTask4(t)
	//testTask5(t)
	//testTask6(t)
	//testTask7(t)
	//testTask8(t)
	//testTask9(t)
	//testTask10(t)
	//testTask11(t)
	//testTask12(t)
}

func testFinal1(t *testing.T) {
	res := []byte("5\n0 1 4 9 0")
	out := "0 1 2 1 0"
	arrOut := strings.Split(strings.Replace(out, "\r", "", -1), " ")
	answer, _ := sprint1.Final1(&res)
	fmt.Println(answer)
	fmt.Println(arrOut)
	fmt.Println("------------------")
	for i := 0; i < len(arrOut); i++ {
		if answer[i] != arrOut[i] {
			t.Error("Error test")
			break
		}
	}

	res = []byte("6\n0 7 9 4 8 20")
	out = "0 1 2 3 4 5"
	arrOut = strings.Split(strings.Replace(out, "\r", "", -1), " ")
	answer, _ = sprint1.Final1(&res)
	fmt.Println(answer)
	fmt.Println(arrOut)
	fmt.Println("------------------")
	for i := 0; i < len(arrOut); i++ {
		if answer[i] != arrOut[i] {
			t.Error("Error test")
			break
		}
	}

	res = []byte("9\n64 68 37 11 77 80 48 82 0")
	out = "8 7 6 5 4 3 2 1 0"
	arrOut = strings.Split(strings.Replace(out, "\r", "", -1), " ")

	answer, _ = sprint1.Final1(&res)
	fmt.Println(answer)
	fmt.Println(arrOut)
	fmt.Println("------------------")
	for i := 0; i < len(arrOut); i++ {
		if answer[i] != arrOut[i] {
			t.Error("Error test")
			break
		}
	}

	res = []byte("20\n5 8 9 12 15 26 30 0 0 55 0 0 67 0 76 80 82 0 0 98")
	out = "7 6 5 4 3 2 1 0 0 1 0 0 1 0 1 2 1 0 0 1"
	arrOut = strings.Split(strings.Replace(out, "\r", "", -1), " ")

	answer, _ = sprint1.Final1(&res)
	fmt.Println(answer)
	fmt.Println(arrOut)
	fmt.Println("------------------")
	for i := 0; i < len(arrOut); i++ {
		if answer[i] != arrOut[i] {
			t.Error("Error test")
			break
		}
	}

	res = []byte("9\n98 0 10 77 0 59 28 0 94")
	out = "1 0 1 1 0 1 1 0 1"
	arrOut = strings.Split(strings.Replace(out, "\r", "", -1), " ")

	answer, _ = sprint1.Final1(&res)
	fmt.Println(answer)
	fmt.Println(arrOut)
	fmt.Println("------------------")
	for i := 0; i < len(arrOut); i++ {
		if answer[i] != arrOut[i] {
			t.Error("Error test")
			break
		}
	}

	res = []byte("100\n0 152871953 35758611 0 23446050 761817124 672461861 39089193 0 0 188187704 70952507 0 940207165 444534846 414873160 164274761 858907213 703483471 385149009 0 951784545 0 0 429548140 437438577 945148040 847574158 245872275 686676631 260948138 220055723 394760373 608300732 279782591 507679937 0 437680330 268329165 0 526207708 383613661 675427551 1897698 195749095 614799594 453061364 447724793 262574330 142751994 0 982802179 743490824 761200905 765521674 589838606 779130650 489999132 614469924 534675750 795652391 532818235 0 7462720 796842821 987969359 0 820963159 0 639387490 729833831 725754736 559672176 265743221 506640772 0 727270808 849485724 679340445 857923486 351773599 231748526 620691377 753766836 684745656 522743742 168625601 821831113 900417993 894851534 0 553949137 692702162 0 299571161 94853602 215074276 62820842 0 404316147")
	out = "0 1 1 0 1 2 2 1 0 0 1 1 0 1 2 3 4 3 2 1 0 1 0 0 1 2 3 4 5 6 6 5 4 3 2 1 0 1 1 0 1 2 3 4 5 5 4 3 2 1 0 1 2 3 4 5 6 5 4 3 2 1 0 1 2 1 0 1 0 1 2 3 3 2 1 0 1 2 3 4 5 6 7 7 6 5 4 3 2 1 0 1 1 0 1 2 2 1 0 1"
	arrOut = strings.Split(strings.Replace(out, "\r", "", -1), " ")

	answer, _ = sprint1.Final1(&res)
	fmt.Println(answer)
	fmt.Println(arrOut)
	fmt.Println("------------------")
	for i := 0; i < len(arrOut); i++ {
		if answer[i] != arrOut[i] {
			t.Error("Error test")
			break
		}
	}

}

func testFinal2(t *testing.T) {
	res := []byte("3\n1231\n2..2\n2..2\n2..2")
	out := "2"

	answer, _ := sprint1.Final2(&res)
	fmt.Println(answer[0])
	fmt.Println(out)
	fmt.Println("------------------")
	if answer[0] != out {
		t.Error("Error test")
	}

	res = []byte("4\n1111\n9999\n1111\n9911")
	out = "1"

	answer, _ = sprint1.Final2(&res)
	fmt.Println(answer[0])
	fmt.Println(out)
	fmt.Println("------------------")
	if answer[0] != out {
		t.Error("Error test")
	}

	res = []byte("4\n1111\n1111\n1111\n1111")
	out = "0"

	answer, _ = sprint1.Final2(&res)
	fmt.Println(answer[0])
	fmt.Println(out)
	fmt.Println("------------------")
	if answer[0] != out {
		t.Error("Error test")
	}

}

func testTask1(t *testing.T) {
	res := []byte("-8 -5 -2 7")
	out := "-183"
	answer, _ := sprint1.Task1(&res)
	fmt.Println(answer[0])
	fmt.Println(out)
	fmt.Println("------------------")
	if answer[0] != out {
		t.Error("Error test")
	}

	res = []byte("8 2 9 -10")
	out = "40"
	answer, _ = sprint1.Task1(&res)
	fmt.Println(answer[0])
	fmt.Println(out)
	fmt.Println("------------------")
	if answer[0] != out {
		t.Error("Error test")
	}
}

func testTask2(t *testing.T) {
	res := []byte("1 2 -3")
	out := "FAIL"
	answer, _ := sprint1.Task2(&res)
	fmt.Println(answer[0])
	fmt.Println(out)
	fmt.Println("------------------")
	if answer[0] != out {
		t.Error("Error test")
	}

	res = []byte("7 11 7")
	out = "WIN"
	answer, _ = sprint1.Task2(&res)
	fmt.Println(answer[0])
	fmt.Println(out)
	fmt.Println("------------------")
	if answer[0] != out {
		t.Error("Error test")
	}

	res = []byte("6 -2 0")
	out = "WIN"
	answer, _ = sprint1.Task2(&res)
	fmt.Println(answer[0])
	fmt.Println(out)
	fmt.Println("------------------")
	if answer[0] != out {
		t.Error("Error test")
	}
}

func testTask3(t *testing.T) {
	res := []byte("4\n3\n1 2 3\n0 2 6\n7 4 1\n2 7 0\n3\n0")
	out := "7 7"
	arrOut := strings.Split(strings.Replace(out, "\r", "", -1), " ")
	answer, _ := sprint1.Task3(&res)
	fmt.Println(answer)
	fmt.Println(arrOut)
	fmt.Println("------------------")
	if len(arrOut) == len(answer) {
		for i := 0; i < len(arrOut); i++ {
			if answer[i] != arrOut[i] {
				t.Error("Error test")
				break
			}
		}
	} else {
		t.Error("Error test")
	}

	res = []byte("4\n3\n1 2 3\n0 2 6\n7 4 1\n2 7 0\n0\n0")
	out = "0 2"
	arrOut = strings.Split(strings.Replace(out, "\r", "", -1), " ")
	answer, _ = sprint1.Task3(&res)
	fmt.Println(answer)
	fmt.Println(arrOut)
	fmt.Println("------------------")
	if len(arrOut) == len(answer) {
		for i := 0; i < len(arrOut); i++ {
			if answer[i] != arrOut[i] {
				t.Error("Error test")
				break
			}
		}
	} else {
		t.Error("Error test")
	}

	res = []byte("1\n1\n1\n0\n0")
	arrOut = []string{}
	answer, _ = sprint1.Task3(&res)
	fmt.Println(answer)
	fmt.Println(arrOut)
	fmt.Println("------------------")
	if len(arrOut) == len(answer) {
		for i := 0; i < len(arrOut); i++ {
			if answer[i] != arrOut[i] {
				t.Error("Error test")
				break
			}
		}
	} else {
		t.Error("Error test")
	}

	res = []byte("3\n10\n4 -10 4 -9 9 5 -7 1 4 -3\n-3 0 -1 -6 -6 2 3 3 4 0\n-1 -5 1 -9 -9 -6 3 -1 -10 -7\n1\n0")
	out = "-1 0 4"
	arrOut = strings.Split(strings.Replace(out, "\r", "", -1), " ")
	answer, _ = sprint1.Task3(&res)
	fmt.Println(answer)
	fmt.Println(arrOut)
	fmt.Println("------------------")
	if len(arrOut) == len(answer) {
		for i := 0; i < len(arrOut); i++ {
			if answer[i] != arrOut[i] {
				t.Error("Error test")
				break
			}
		}
	} else {
		t.Error("Error test")
	}

	res = []byte("2\n1\n1\n9\n0\n0")
	out = "9"
	arrOut = strings.Split(strings.Replace(out, "\r", "", -1), " ")
	answer, _ = sprint1.Task3(&res)
	fmt.Println(answer)
	fmt.Println(arrOut)
	fmt.Println("------------------")
	if len(arrOut) == len(answer) {
		for i := 0; i < len(arrOut); i++ {
			if answer[i] != arrOut[i] {
				t.Error("Error test")
				break
			}
		}
	} else {
		t.Error("Error test")
	}

	res = []byte("1\n1\n1\n0\n0")
	arrOut = []string{}
	answer, _ = sprint1.Task3(&res)
	fmt.Println(answer)
	fmt.Println(arrOut)
	fmt.Println("------------------")
	if len(arrOut) == len(answer) {
		for i := 0; i < len(arrOut); i++ {
			if answer[i] != arrOut[i] {
				t.Error("Error test")
				break
			}
		}
	} else {
		t.Error("Error test")
	}

	res = []byte("8\n9\n3 3 -9 7 -5 8 -6 -10 -4\n5 -2 -6 -9 8 -4 5 -5 0\n-9 -3 3 2 1 -4 -6 3 -9\n-7 1 -2 4 -2 1 -5 4 -8\n-2 5 5 7 -7 2 3 -4 -4\n-1 7 -10 7 4 5 -7 1 5\n-1 3 0 -8 -10 -2 5 1 7\n10 4 -9 5 3 -1 7 10 -5\n3\n0")
	out = "-9 -2 1"
	arrOut = strings.Split(strings.Replace(out, "\r", "", -1), " ")
	answer, _ = sprint1.Task3(&res)
	fmt.Println(answer)
	fmt.Println(arrOut)
	fmt.Println("------------------")
	if len(arrOut) == len(answer) {
		for i := 0; i < len(arrOut); i++ {
			if answer[i] != arrOut[i] {
				t.Error("Error test")
				break
			}
		}
	} else {
		t.Error("Error test")
	}

	res = []byte("7\n3\n6 -6 4\n1 4 9\n7 -2 -8\n8 0 6\n9 -2 -5\n-3 -2 -8\n10 -10 3\n6\n1")
	out = "-2 3 10"
	arrOut = strings.Split(strings.Replace(out, "\r", "", -1), " ")
	answer, _ = sprint1.Task3(&res)
	fmt.Println(answer)
	fmt.Println(arrOut)
	fmt.Println("------------------")
	if len(arrOut) == len(answer) {
		for i := 0; i < len(arrOut); i++ {
			if answer[i] != arrOut[i] {
				t.Error("Error test")
				break
			}
		}
	} else {
		t.Error("Error test")
	}

}

func testTask4(t *testing.T) {
	res := []byte("7\n-1 -10 -8 0 2 0 5")
	out := "3"
	arrOut := strings.Split(strings.Replace(out, "\r", "", -1), " ")
	answer, _ := sprint1.Task4(&res)
	fmt.Println(answer)
	fmt.Println(arrOut)
	fmt.Println("------------------")
	if len(arrOut) == len(answer) {
		for i := 0; i < len(arrOut); i++ {
			if answer[i] != arrOut[i] {
				t.Error("Error test")
				break
			}
		}
	} else {
		t.Error("Error test")
	}

	res = []byte("5\n1 2 5 4 8")
	out = "2"
	arrOut = strings.Split(strings.Replace(out, "\r", "", -1), " ")
	answer, _ = sprint1.Task4(&res)
	fmt.Println(answer)
	fmt.Println(arrOut)
	fmt.Println("------------------")
	if len(arrOut) == len(answer) {
		for i := 0; i < len(arrOut); i++ {
			if answer[i] != arrOut[i] {
				t.Error("Error test")
				break
			}
		}
	} else {
		t.Error("Error test")
	}

}
func testTask5(t *testing.T) {
	res := []byte("19\ni love segment tree")
	out := "segment 7"
	arrOut := strings.Split(strings.Replace(out, "\r", "", -1), "\n")
	answer, _ := sprint1.Task5(&res)
	fmt.Println(answer)
	fmt.Println(arrOut)
	fmt.Println("------------------")
	if len(arrOut) == len(answer) {
		for i := 0; i < len(arrOut); i++ {
			if answer[i] != arrOut[i] {
				t.Error("Error test")
				break
			}
		}
	} else {
		t.Error("Error test")
	}

	res = []byte("21\nfrog jumps from river")
	out = "jumps 5"
	arrOut = strings.Split(strings.Replace(out, "\r", "", -1), " ")
	answer, _ = sprint1.Task5(&res)
	fmt.Println(answer)
	fmt.Println(arrOut)
	fmt.Println("------------------")
	if len(arrOut) == len(answer) {
		for i := 0; i < len(arrOut); i++ {
			if answer[i] != arrOut[i] {
				t.Error("Error test")
				break
			}
		}
	} else {
		t.Error("Error test")
	}

}

func testTask6(t *testing.T) {
	res := []byte("A man, a plan, a canal: Panama")
	out := "True"
	arrOut := strings.Split(strings.Replace(out, "\r", "", -1), "\n")
	answer, _ := sprint1.Task6(&res)
	fmt.Println(answer)
	fmt.Println(arrOut)
	fmt.Println("------------------")
	if len(arrOut) == len(answer) {
		for i := 0; i < len(arrOut); i++ {
			if answer[i] != arrOut[i] {
				t.Error("Error test")
				break
			}
		}
	} else {
		t.Error("Error test")
	}

	res = []byte("zo")
	out = "False"
	arrOut = strings.Split(strings.Replace(out, "\r", "", -1), "\n")
	answer, _ = sprint1.Task6(&res)
	fmt.Println(answer)
	fmt.Println(arrOut)
	fmt.Println("------------------")
	if len(arrOut) == len(answer) {
		for i := 0; i < len(arrOut); i++ {
			if answer[i] != arrOut[i] {
				t.Error("Error test")
				break
			}
		}
	} else {
		t.Error("Error test")
	}

	res = []byte("10 ten animals I slam in a net 012")
	out = "False"
	arrOut = strings.Split(strings.Replace(out, "\r", "", -1), "\n")
	answer, _ = sprint1.Task6(&res)
	fmt.Println(answer)
	fmt.Println(arrOut)
	fmt.Println("------------------")
	if len(arrOut) == len(answer) {
		for i := 0; i < len(arrOut); i++ {
			if answer[i] != arrOut[i] {
				t.Error("Error test")
				break
			}
		}
	} else {
		t.Error("Error test")
	}

	res = []byte("12321!")
	out = "True"
	arrOut = strings.Split(strings.Replace(out, "\r", "", -1), "\n")
	answer, _ = sprint1.Task6(&res)
	fmt.Println(answer)
	fmt.Println(arrOut)
	fmt.Println("------------------")
	if len(arrOut) == len(answer) {
		for i := 0; i < len(arrOut); i++ {
			if answer[i] != arrOut[i] {
				t.Error("Error test")
				break
			}
		}
	} else {
		t.Error("Error test")
	}

}

func testTask7(t *testing.T) {
	res := []byte("5")
	out := "101"
	arrOut := strings.Split(strings.Replace(out, "\r", "", -1), "\n")
	answer, _ := sprint1.Task7(&res)
	fmt.Println(answer)
	fmt.Println(arrOut)
	fmt.Println("------------------")
	if len(arrOut) == len(answer) {
		for i := 0; i < len(arrOut); i++ {
			if answer[i] != arrOut[i] {
				t.Error("Error test")
				break
			}
		}
	} else {
		t.Error("Error test")
	}

	res = []byte("14")
	out = "1110"
	arrOut = strings.Split(strings.Replace(out, "\r", "", -1), "\n")
	answer, _ = sprint1.Task7(&res)
	fmt.Println(answer)
	fmt.Println(arrOut)
	fmt.Println("------------------")
	if len(arrOut) == len(answer) {
		for i := 0; i < len(arrOut); i++ {
			if answer[i] != arrOut[i] {
				t.Error("Error test")
				break
			}
		}
	} else {
		t.Error("Error test")
	}
}

func testTask8(t *testing.T) {
	res := []byte("1010\n1011")
	out := "10101"
	arrOut := strings.Split(strings.Replace(out, "\r", "", -1), "\n")
	answer, _ := sprint1.Task8(&res)
	fmt.Println(answer)
	fmt.Println(arrOut)
	fmt.Println("------------------")
	if len(arrOut) == len(answer) {
		for i := 0; i < len(arrOut); i++ {
			if answer[i] != arrOut[i] {
				t.Error("Error test")
				break
			}
		}
	} else {
		t.Error("Error test")
	}

	res = []byte("1\n1")
	out = "10"
	arrOut = strings.Split(strings.Replace(out, "\r", "", -1), "\n")
	answer, _ = sprint1.Task8(&res)
	fmt.Println(answer)
	fmt.Println(arrOut)
	fmt.Println("------------------")
	if len(arrOut) == len(answer) {
		for i := 0; i < len(arrOut); i++ {
			if answer[i] != arrOut[i] {
				t.Error("Error test")
				break
			}
		}
	} else {
		t.Error("Error test")
	}

	res = []byte("1101110101000001000100101000100101100010101100001111001101011111111\n10101001011000011110101111101000101111110111110011000101111000011000101001000010101000")
	out = "10101001011000100000011110010000111000011100110111110010001101111010100010101110100111"

	arrOut = strings.Split(strings.Replace(out, "\r", "", -1), "\n")
	answer, _ = sprint1.Task8(&res)
	fmt.Println(answer)
	fmt.Println(arrOut)
	fmt.Println("------------------")
	if len(arrOut) == len(answer) {
		for i := 0; i < len(arrOut); i++ {
			if answer[i] != arrOut[i] {
				t.Error("Error test")
				break
			}
		}
	} else {
		t.Error("Error test")
	}
}

func testTask9(t *testing.T) {
	res := []byte("15")
	out := "False"
	arrOut := strings.Split(strings.Replace(out, "\r", "", -1), "\n")
	answer, _ := sprint1.Task9(&res)
	fmt.Println(answer)
	fmt.Println(arrOut)
	fmt.Println("------------------")
	if len(arrOut) == len(answer) {
		for i := 0; i < len(arrOut); i++ {
			if answer[i] != arrOut[i] {
				t.Error("Error test")
				break
			}
		}
	} else {
		t.Error("Error test")
	}

	res = []byte("16")
	out = "True"
	arrOut = strings.Split(strings.Replace(out, "\r", "", -1), "\n")
	answer, _ = sprint1.Task9(&res)
	fmt.Println(answer)
	fmt.Println(arrOut)
	fmt.Println("------------------")
	if len(arrOut) == len(answer) {
		for i := 0; i < len(arrOut); i++ {
			if answer[i] != arrOut[i] {
				t.Error("Error test")
				break
			}
		}
	} else {
		t.Error("Error test")
	}

}

func testTask10(t *testing.T) {
	res := []byte("8")
	out := "2 2 2"
	arrOut := strings.Split(strings.Replace(out, " ", "\n", -1), "\n")
	answer, _ := sprint1.Task10(&res)
	fmt.Println(answer)
	fmt.Println(arrOut)
	fmt.Println("------------------")
	if len(arrOut) == len(answer) {
		for i := 0; i < len(arrOut); i++ {
			if answer[i] != arrOut[i] {
				t.Error("Error test")
				break
			}
		}
	} else {
		t.Error("Error test")
	}

	res = []byte("13")
	out = "13"
	arrOut = strings.Split(strings.Replace(out, " ", "\n", -1), "\n")
	answer, _ = sprint1.Task10(&res)
	fmt.Println(answer)
	fmt.Println(arrOut)
	fmt.Println("------------------")
	if len(arrOut) == len(answer) {
		for i := 0; i < len(arrOut); i++ {
			if answer[i] != arrOut[i] {
				t.Error("Error test")
				break
			}
		}
	} else {
		t.Error("Error test")
	}

	res = []byte("100")
	out = "2 2 5 5"
	arrOut = strings.Split(strings.Replace(out, " ", "\n", -1), "\n")
	answer, _ = sprint1.Task10(&res)
	fmt.Println(answer)
	fmt.Println(arrOut)
	fmt.Println("------------------")
	if len(arrOut) == len(answer) {
		for i := 0; i < len(arrOut); i++ {
			if answer[i] != arrOut[i] {
				t.Error("Error test")
				break
			}
		}
	} else {
		t.Error("Error test")
	}

}

func testTask11(t *testing.T) {
	res := []byte("4\n1 2 0 0\n34")
	out := "1 2 3 4"
	arrOut := strings.Split(strings.Replace(out, " ", "\n", -1), "\n")
	answer, _ := sprint1.Task11(&res)
	fmt.Println(answer)
	fmt.Println(arrOut)
	fmt.Println("------------------")
	if len(arrOut) == len(answer) {
		for i := 0; i < len(arrOut); i++ {
			if answer[i] != arrOut[i] {
				t.Error("Error test")
				break
			}
		}
	} else {
		t.Error("Error test")
	}

	res = []byte("2\n9 5\n17")
	out = "1 1 2"
	arrOut = strings.Split(strings.Replace(out, " ", "\n", -1), "\n")
	answer, _ = sprint1.Task11(&res)
	fmt.Println(answer)
	fmt.Println(arrOut)
	fmt.Println("------------------")
	if len(arrOut) == len(answer) {
		for i := 0; i < len(arrOut); i++ {
			if answer[i] != arrOut[i] {
				t.Error("Error test")
				break
			}
		}
	} else {
		t.Error("Error test")
	}

	res = []byte("30\n2 0 4 3 3 2 1 8 1 9 6 0 0 1 3 3 8 9 0 8 3 8 6 3 7 9 4 0 2 6\n349")
	out = "2 0 4 3 3 2 1 8 1 9 6 0 0 1 3 3 8 9 0 8 3 8 6 3 7 9 4 3 7 5"
	arrOut = strings.Split(strings.Replace(out, " ", "\n", -1), "\n")
	answer, _ = sprint1.Task11(&res)
	fmt.Println(answer)
	fmt.Println(arrOut)
	fmt.Println("------------------")
	if len(arrOut) == len(answer) {
		for i := 0; i < len(arrOut); i++ {
			if answer[i] != arrOut[i] {
				t.Error("Error test")
				break
			}
		}
	} else {
		t.Error("Error test")
	}

}

func testTask12(t *testing.T) {
	res := []byte("abcd\nabcde")
	out := "e"
	arrOut := strings.Split(strings.Replace(out, " ", "\n", -1), "\n")
	answer, _ := sprint1.Task12(&res)
	fmt.Println(answer)
	fmt.Println(arrOut)
	fmt.Println("------------------")
	if len(arrOut) == len(answer) {
		for i := 0; i < len(arrOut); i++ {
			if answer[i] != arrOut[i] {
				t.Error("Error test")
				break
			}
		}
	} else {
		t.Error("Error test")
	}

	res = []byte("go\nogg")
	out = "g"
	arrOut = strings.Split(strings.Replace(out, " ", "\n", -1), "\n")
	answer, _ = sprint1.Task12(&res)
	fmt.Println(answer)
	fmt.Println(arrOut)
	fmt.Println("------------------")
	if len(arrOut) == len(answer) {
		for i := 0; i < len(arrOut); i++ {
			if answer[i] != arrOut[i] {
				t.Error("Error test")
				break
			}
		}
	} else {
		t.Error("Error test")
	}

	res = []byte("xtkpx\nxkctpx")
	out = "c"
	arrOut = strings.Split(strings.Replace(out, " ", "\n", -1), "\n")
	answer, _ = sprint1.Task12(&res)
	fmt.Println(answer)
	fmt.Println(arrOut)
	fmt.Println("------------------")
	if len(arrOut) == len(answer) {
		for i := 0; i < len(arrOut); i++ {
			if answer[i] != arrOut[i] {
				t.Error("Error test")
				break
			}
		}
	} else {
		t.Error("Error test")
	}

}
