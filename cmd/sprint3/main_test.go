package main

import (
	"fmt"
	"testing"
	"yalg/sprint3"
)

func TestFuncSprint1(t *testing.T) {
	//testTask1(t)
	//testTask2(t)
	//testTask10(t)
	//testTask8(t)
	//testTask12(t)
	//testTask3(t)
	//testTask4(t)
	//testTask5(t)
	//testTask6(t)
	//testTask7(t)
	//testTask9(t)
	//testTask13(t)
	//testTask14(t)
	//testTask15(t)
	//testTask16(t)
	testFinal2(t)
	//testFinal2V2(t)
}

func testTask1(t *testing.T) {
	res := []byte("3")
	out := "((()))\n(()())\n(())()\n()(())\n()()()\n"
	answer := sprint3.Task1(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	fmt.Println("===============================")
	fmt.Println("===============================")
}

func testTask2(t *testing.T) {
	res := []byte("23")
	out := "ad ae af bd be bf cd ce cf"
	answer := sprint3.Task2(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("92")
	out = "wa wb wc xa xb xc ya yb yc za zb zc"
	answer = sprint3.Task2(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("7884")
	out = "pttg ptth ptti ptug ptuh ptui ptvg ptvh ptvi putg puth puti puug puuh puui puvg puvh puvi pvtg pvth pvti pvug pvuh pvui pvvg pvvh pvvi qttg qtth qtti qtug qtuh qtui qtvg qtvh qtvi qutg quth quti quug quuh quui quvg quvh quvi qvtg qvth qvti qvug qvuh qvui qvvg qvvh qvvi rttg rtth rtti rtug rtuh rtui rtvg rtvh rtvi rutg ruth ruti ruug ruuh ruui ruvg ruvh ruvi rvtg rvth rvti rvug rvuh rvui rvvg rvvh rvvi sttg stth stti stug stuh stui stvg stvh stvi sutg suth suti suug suuh suui suvg suvh suvi svtg svth svti svug svuh svui svvg svvh svvi"
	answer = sprint3.Task2(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	fmt.Println("===============================")
	fmt.Println("===============================")
}

func testTask10(t *testing.T) {
	res := []byte("5\n4 3 9 2 1")
	out := "3 4 2 1 9\n3 2 1 4 9\n2 1 3 4 9\n1 2 3 4 9\n"
	answer := sprint3.Task10(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("5\n12 8 9 10 11")
	out = "8 9 10 11 12\n"
	answer = sprint3.Task10(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("2\n4 5\nПросто пример из условия")
	out = "4 5\n"
	answer = sprint3.Task10(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	fmt.Println("===============================")
	fmt.Println("===============================")
}

func testTask8(t *testing.T) {
	res := []byte("30\n9 18 1 65 51 16 6 43 6 36 7 11 64 5 1 76 15 11 11 15 57 95 3 49 92 78 83 51 10 3\n")
	//out := "9 95 92 83 78 7 76 6 6 65 64 57 5 51 51 49 43 36 3 3 18 16 15 15 11 11 11 1 1 10\n"
	out := "995928378776666564575515149433633181615151111111110\n"
	answer := sprint3.Task8(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("3\n1 783 2\n")
	out = "78321\n"
	answer = sprint3.Task8(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("5\n2 4 5 2 10\n")
	out = "542210\n"
	answer = sprint3.Task8(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("6\n9 10 1 1 1 6\n")
	out = "9611110\n"
	answer = sprint3.Task8(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("100\n336 945 741 532 433 309 685 325 572 24 1000 631 675 831 807 134 473 1000 634 748 587 1000 852 353 262 305 243 27 697 919 906 844 938 670 71 843 803 72 180 781 491 379 590 407 233 498 617 533 28 493 542 1000 521 85 373 646 89 799 787 728 148 38 443 559 426 1000 916 866 1000 1000 610 925 882 888 965 317 104 256 976 9 83 415 244 239 1000 68 305 771 734 571 523 574 1000 495 139 94 757 709 573 412\n")
	out = "99769659494593892591991690689888882866858528448438383180780379978778177175774874173472872717096976868567567064663463161761059058757457357257155954253353252352149849549349147344343342641541240738379373353336325317309305305282726225624424324239233180148139134104100010001000100010001000100010001000\n"
	answer = sprint3.Task8(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("100\n82 468 941 181 287 861 291 515 263 424 470 620 954 894 565 69 148 587 823 57 730 389 921 1000 447 1000 748 104 831 943 174 24 340 1000 150 937 324 919 748 271 980 575 392 779 222 316 944 1000 160 501 319 436 26 828 348 211 825 857 486 1000 419 509 409 679 576 700 418 810 674 83 785 251 947 868 964 384 497 192 1000 998 756 649 269 290 197 30 95 796 642 980 474 122 443 707 839 213 1000 530 263 193\n")
	out = "99898098096495954947944943941937921919894868861857839838318288282582381079678577975674874873070770069679674649642620587576575755655305155095014974864744704684474434364244194184093923893843483403243193163029129028727126926326326251242222132111971931921811741601501481221041000100010001000100010001000\n"
	answer = sprint3.Task8(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
		for i := 0; i < len(answer); i++ {
			v1 := out[i]
			v2 := answer[i]

			if v1 != v2 {
				t.Error(i)
			}

		}
	}

	fmt.Println("===============================")
	fmt.Println("===============================")
}

func testTask3(t *testing.T) {
	res := []byte("abc\nahbgdcu\n")
	out := "True\n"
	answer := sprint3.Task3(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("abcp\nahpc\n")
	out = "False\n"
	answer = sprint3.Task3(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("ijha\nhmrqvftefyixinahlzgbkidroxiptbbkjmtwpsujevkulgrjiwiwzyhngulrodiwyg\n")
	out = "False\n"
	answer = sprint3.Task3(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	fmt.Println("===============================")
	fmt.Println("===============================")
}

func testTask4(t *testing.T) {
	res := []byte("2\n1 2\n3\n2 1 3\n")
	out := "2\n"
	answer := sprint3.Task4(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("3\n2 1 3\n2\n1 1\n")
	out = "1\n"
	answer = sprint3.Task4(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("10\n5 10 1 5 4 4 10 1 5 10\n9\n4 4 8 5 9 6 1 7 4\n")
	out = "7\n"
	answer = sprint3.Task4(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	fmt.Println("===============================")
	fmt.Println("===============================")
}

func testTask5(t *testing.T) {
	res := []byte("3 300\n999 999 999")
	out := "0"
	answer := sprint3.Task5(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("3 1000\n350 999 200")
	out = "2"
	answer = sprint3.Task5(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	fmt.Println("===============================")
	fmt.Println("===============================")
}

func testTask6(t *testing.T) {
	res := []byte("4\n6 3 3 2")
	out := "8\n"
	answer := sprint3.Task6(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("6\n5 3 7 2 8 3")
	out = "20\n"
	answer = sprint3.Task6(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("6\n5 3 7 2 8 2")
	out = "20\n"
	answer = sprint3.Task6(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	fmt.Println("===============================")
	fmt.Println("===============================")
}

func testTask7(t *testing.T) {
	res := []byte("7\n0 2 1 2 0 0 1")
	out := "0 0 0 1 1 2 2"
	answer := sprint3.Task7(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("5\n2 1 2 0 1")
	out = "0 1 1 2 2"
	answer = sprint3.Task7(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("6\n2 1 1 2 0 2")
	out = "0 1 1 2 2 2"
	answer = sprint3.Task7(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("0\n\n")
	out = "0 1 1 2 2 2"
	answer = sprint3.Task7(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	fmt.Println("===============================")
	fmt.Println("===============================")
}

func testTask9(t *testing.T) {
	res := []byte("7\n1 2 3 1 2 3 4\n3\n")
	out := "1 2 3"
	answer := sprint3.Task9(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("6\n1 1 1 2 2 3\n1\n")
	out = "1"
	answer = sprint3.Task9(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("6\n1 1 1 2 2 3\n3\n")
	out = "1 2 3"
	answer = sprint3.Task9(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("9\n1 1 1 2 2 3 7 5 5\n3\n")
	out = "1 2 5"
	answer = sprint3.Task9(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	fmt.Println("===============================")
	fmt.Println("===============================")
}

func testTask13(t *testing.T) {
	res := []byte("2\n1\n1 3\n2\n")
	out := "2\n"
	answer := sprint3.Task13(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("2\n2\n1 2\n3 4\n")
	out = "2.5\n"
	answer = sprint3.Task13(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("8\n10\n0 0 0 1 3 3 5 10\n4 4 5 7 7 7 8 9 9 10\n")
	out = "5\n"
	answer = sprint3.Task13(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	fmt.Println("===============================")
	fmt.Println("===============================")
}

func testTask14(t *testing.T) {
	res := []byte("4\n7 8\n7 8\n2 3\n6 10\n")
	out := "2 3\n6 10\n"
	answer := sprint3.Task14(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("4\n2 3\n5 6\n3 4\n3 4\n")
	out = "2 4\n5 6\n"
	answer = sprint3.Task14(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("6\n1 3\n3 5\n4 6\n5 6\n2 4\n7 10\n")
	out = "1 6\n7 10\n"
	answer = sprint3.Task14(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	//5 13
	//1 98
	res = []byte("15\n5 13\n11 85\n16 18\n2 37\n55 73\n33 61\n4 60\n39 98\n43 66\n26 61\n77 81\n67 72\n1 40\n50 98\n65 82\n")
	out = "1 98\n"
	answer = sprint3.Task14(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	fmt.Println("===============================")
	fmt.Println("===============================")
}

func testTask15(t *testing.T) {
	res := []byte("3\n2 3 4\n2")
	out := "1"
	answer := sprint3.Task15(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("3\n1 3 1\n1")
	out = "0"
	answer = sprint3.Task15(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("3\n1 3 5\n3")
	out = "4"
	answer = sprint3.Task15(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("4\n9 6 10 3\n3")
	out = "3"
	answer = sprint3.Task15(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("8\n9 1 10 3 4 6 2 7\n6")
	out = "2"
	answer = sprint3.Task15(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	fmt.Println("===============================")
	fmt.Println("===============================")
}

func testTask16(t *testing.T) {
	res := []byte("4\n0 1 3 2")
	out := "3"
	answer := sprint3.Task16(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("8\n3 6 7 4 1 5 0 2")
	out = "1"
	answer = sprint3.Task16(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("5\n1 0 2 3 4")
	out = "4"
	answer = sprint3.Task16(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("7\n3 2 0 1 4 6 5")
	out = "3"
	answer = sprint3.Task16(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("5\n0 4 1 2 3")
	out = "2"
	answer = sprint3.Task16(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("10\n2 7 0 3 6 4 1 5 8 9")
	out = "3"
	answer = sprint3.Task16(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	fmt.Println("===============================")
	fmt.Println("===============================")
}

func testFinal2(t *testing.T) {
	res := []byte("5\nalla 4 100\ngena 6 1000\ngosha 2 90\nrita 2 90\ntimofey 4 80")
	out := "gena\ntimofey\nalla\ngosha\nrita\n"
	answer := sprint3.Final2(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("5\nalla 0 0\ngena 0 0\ngosha 0 0\nrita 0 0\ntimofey 0 0")
	out = "alla\ngena\ngosha\nrita\ntimofey\n"
	answer = sprint3.Final2(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("11\nixetulueozem 55 26\nedrnkztanvrpyjvso 65 7\ne 89 9\nayrcrulcwh 99 14\nnsgnysqm 15 32\newdponbpcmtgfabnvo 65 15\nsfkropatfwkna 95 59\nkzibjralr 2 12\nhnpmykspichx 34 87\narmaholwvkttg 9 47\nvswomwpuhpqzxstltlw 10 40")
	out = "ayrcrulcwh\nsfkropatfwkna\ne\nedrnkztanvrpyjvso\newdponbpcmtgfabnvo\nixetulueozem\nhnpmykspichx\nnsgnysqm\nvswomwpuhpqzxstltlw\narmaholwvkttg\nkzibjralr\n"
	answer = sprint3.Final2(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	fmt.Println("===============================")
	fmt.Println("===============================")
}

func testFinal2V2(t *testing.T) {
	res := []byte("5\nalla 4 100\ngena 6 1000\ngosha 2 90\nrita 2 90\ntimofey 4 80")
	out := "gena\ntimofey\nalla\ngosha\nrita\n"
	answer := sprint3.Final2(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("5\nalla 0 0\ngena 0 0\ngosha 0 0\nrita 0 0\ntimofey 0 0")
	out = "alla\ngena\ngosha\nrita\ntimofey\n"
	answer = sprint3.Final2(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("11\nixetulueozem 55 26\nedrnkztanvrpyjvso 65 7\ne 89 9\nayrcrulcwh 99 14\nnsgnysqm 15 32\newdponbpcmtgfabnvo 65 15\nsfkropatfwkna 95 59\nkzibjralr 2 12\nhnpmykspichx 34 87\narmaholwvkttg 9 47\nvswomwpuhpqzxstltlw 10 40")
	out = "ayrcrulcwh\nsfkropatfwkna\ne\nedrnkztanvrpyjvso\newdponbpcmtgfabnvo\nixetulueozem\nhnpmykspichx\nnsgnysqm\nvswomwpuhpqzxstltlw\narmaholwvkttg\nkzibjralr\n"
	answer = sprint3.Final2(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("13\ntufhdbi 76 58\nrqyoazgbmv 59 78\nqvgtrlkmyrm 35 27\ntgcytmfpj 70 27\nxvf 84 19\njzpnpgpcqbsmczrgvsu 30 3\nevjphqnevjqakze 92 15\nwwzwv 87 8\ntfpiqpwmkkduhcupp 1 82\ntzamkyqadmybky 5 81\namotrxgba 0 6\neasfsifbzkfezn 100 28\nkivdiy 70 47")
	out = "easfsifbzkfezn\nevjphqnevjqakze\nwwzwv\nxvf\ntufhdbi\ntgcytmfpj\nkivdiy\nrqyoazgbmv\nqvgtrlkmyrm\njzpnpgpcqbsmczrgvsu\ntzamkyqadmybky\ntfpiqpwmkkduhcupp\namotrxgba\n"
	answer = sprint3.Final2(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	fmt.Println("===============================")
	fmt.Println("===============================")
}
