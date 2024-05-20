package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"yalg/sprint2"
)

func TestFuncSprint1(t *testing.T) {
	//testTask1(t)
	//testTask6(t)
	//testTask7(t)
	//testTask8(t)
	//testTask9(t)
	//testTask10(t)
	//testTask11(t)
	//testTask12(t)
	//testFinal1(t)
	testFinal2(t)
}

func testTask1(t *testing.T) {
	res := []byte("4\n3\n1 2 3\n0 2 6\n7 4 1\n2 7 0\n")
	out := "1 0 7 2\n2 2 4 7\n3 6 1 0\n"
	answer := sprint2.Task1(&res)

	fmt.Println("[" + answer + "]")
	fmt.Println("[" + out + "]")
	fmt.Println("------------------")

	if answer != out {
		t.Error("Error test")
	}

	res = []byte("9\n5\n-7 -1 0 -4 -9\n5 -1 2 2 9\n3 1 -8 -1 -7\n9 0 8 -8 -1\n2 4 5 2 8\n-7 10 0 -4 -8\n-3 10 -7 10 3\n1 6 -7 -5 9\n-1 9 9 1 9\n")
	out = "-7 5 3 9 2 -7 -3 1 -1\n-1 -1 1 0 4 10 10 6 9\n0 2 -8 8 5 0 -7 -7 9\n-4 2 -1 -8 2 -4 10 -5 1\n-9 9 -7 -1 8 -8 3 9 9\n"
	answer = sprint2.Task1(&res)

	fmt.Println("[" + answer + "]")
	fmt.Println("[" + out + "]")
	fmt.Println("------------------")

	if answer != out {
		t.Error("Error test")
	}

	res = []byte("2\n9\n6 4 10 3 8 3 5 -5 -9\n-8 2 6 -1 3 -8 -5 10 -4\n")
	out = "6 -8\n4 2\n10 6\n3 -1\n8 3\n3 -8\n5 -5\n-5 10\n-9 -4\n"
	answer = sprint2.Task1(&res)

	fmt.Println("[" + answer + "]")
	fmt.Println("[" + out + "]")
	fmt.Println("------------------")

	if answer != out {
		t.Error("Error test")
	}

	//res = []byte("950\n764\n-426 -716 -361 985 -788 918 545 319 674 299 273 873 615 -725 -339 905 682 -982 -86 974 -514 -897 -481 981 -747 551 453 -803 786 798 -504 824 695 811 -721 -803 276 -697 -233 92 58 167 -427 -8 -852 -475 -351 -401 -263 957 -835 978 960 -446 119 721 219 228 -455 730 92 31 -493 690 565 -296 -657 -775 941 329 -669 408 -309 867 130 820 736 -123 -45 479 -634 -304 546 -829 -233 -871 202 679 590 -942 545 -958 -46 722 -631 -642 -329 -673 685 -291 731 -639 510 246 126 -828 -361 -440 795 588 649 -352 -230 -160 -40 -700 117 -803 -45 -820 -656 158 -131 -405 -117 274 198 -200 -205 887 -647 787 -64 680 470 -872 -350 -943 -575 7 162 -640 598 -705 130 72 763 467 844 548 -37 -275 -360 -872 -94 -48 384 21 -558 -862 -523 -438 -885 -229 -81 730 536 -549 -519 793 467 -713 -907 511 136 -253 778 114 439 -996 987 -593 851 -228 -934 499 -465 459 -614 -501 239 -448 -872 488 65 466 429 437 82 -769 -41 -833 554 -362 860 583 -513 -97 -321 -481 126 840 -531 -606 -174 -737 436 -992 -812 -974 425 -142 -264 738 -426 809 -663 -148 117 83\n")
	//out = "-426 501 -9 -327 848 644 384 -171 -118 998 498 235 742 -278 384 -813 557 746 99 -985 -287 -545 -297 -640 -58 -200 139 -572 -34 -253 -984 17 -30 679 -758 11 -724 579 -953 -440 -306 633 -51 -51 719 -177 -604 -845 200 870 -188 -128 -788 -6 -757 539 287 -316 -316 523 -288 -444 -861 -955 30 581 265 5 -3 152 -801 49 842 991 897 -846 73 -38 925 -660 466 227 -325 -951 -213 -770 629 -250 -585 -90 869 -205 822 971 268 -547 -549 275 -811 -737 657 -85 -342 322 19 -128 -454 303 -781 -567 830 60 -643 941 -574 594 952 -511 561 753 763 742 -191 -790 -596 -419 -327 572 -196 -718 -551 -707 -517 932 755 -896 -831 652 -157 118 185 -736 -206 311 371 -734 -606 -190 -952 -173 -4 -600 -724 -907 751 -525 468 882 827 563 583 -62 -204 230 289 959 45 -335 -14 -830 454 -241 793 -394 983 554 -361 677 -43 407 908 58 257 177 -871 235 -592 -721 140 -158 713 308 -844 365 949 -779 -691 -198 -929 956 -294 368 -63 840 -110 248 -913 662 259 -937 865 -719 -923 961 -720 663 -420 606 -970 -167 72 661 187 -170 -852 447 -889 -768 516 -945 425 -859 -13\n"
	//answer = sprint2.Task1(&res)
	//
	//fmt.Println("[" + answer + "]")
	//fmt.Println("[" + out + "]")
	//fmt.Println("------------------")
	//
	//if answer != out {
	//	t.Error("Error test")
	//}

	fmt.Println("===============================")
	fmt.Println("===============================")
}

func testTask6(t *testing.T) {
	res := []byte("8\nget_max\npush 7\npop\npush -2\npush -1\npop\nget_max\nget_max\n")
	out := "None\n-2\n-2\n"
	answer := sprint2.Task6(&res)

	fmt.Println("[" + answer + "]")
	fmt.Println("[" + out + "]")
	fmt.Println("------------------")

	if answer != out {
		t.Error("Error test")
	}

	res = []byte("7\nget_max\npop\npop\npop\npush 10\nget_max\npush -9\n")
	out = "None\nerror\nerror\nerror\n10\n"
	answer = sprint2.Task6(&res)
	fmt.Println("[" + answer + "]")
	fmt.Println("[" + out + "]")
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("10\npop\npop\npush 9\npush -7\nget_max\npop\npop\npush -8\nget_max\nget_max\n")
	out = "error\nerror\n9\n-8\n-8\n"
	answer = sprint2.Task6(&res)
	fmt.Println("[" + answer + "]")
	fmt.Println("[" + out + "]")
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	fmt.Println("===============================")
	fmt.Println("===============================")
}

func testTask7(t *testing.T) {
	res := []byte("13\npop\npop\ntop\npush 4\npush -5\ntop\npush 7\npop\npop\nget_max\ntop\npop\nget_max\n")
	out := "error\nerror\nerror\n-5\n4\n4\nNone\n"
	answer := sprint2.Task7(&res)

	fmt.Println(answer)
	//fmt.Println("")
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("10\nget_max\npush -6\npop\npop\nget_max\npush 2\nget_max\npop\npush -2\npush -6\n")
	out = "None\nerror\nNone\n2\n"
	answer = sprint2.Task7(&res)
	fmt.Println(answer)
	//fmt.Println("")
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("12\npop\nget_max\npop\npop\npop\npush -6\npop\nget_max\npop\npop\npush -6\nget_max\n")
	out = "error\nNone\nerror\nerror\nerror\nNone\nerror\nerror\n-6\n"
	answer = sprint2.Task7(&res)
	fmt.Println(answer)
	//fmt.Println("")
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("17\nget_max\npop\npush -3\nget_max\npop\nget_max\npush 7\nget_max\npop\npush 9\nget_max\npush 8\nget_max\npush -9\nget_max\npush -10\nget_max\n")
	out = "None\nerror\n-3\nNone\n7\n9\n9\n9\n9\n"
	answer = sprint2.Task7(&res)
	fmt.Println(answer)
	//fmt.Println("")
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("50\nget_max\npop\npush -7\npop\nget_max\nget_max\nget_max\nget_max\npush 1\nget_max\nget_max\npush 3\npop\npop\nget_max\nget_max\nget_max\npush -3\nget_max\nget_max\npush 0\npush 9\nget_max\nget_max\npop\nget_max\npush 4\nget_max\nget_max\npush -3\npush -4\nget_max\npush 9\npush 9\npop\npop\npush -1\npush -4\npush 3\npush 10\npop\nget_max\nget_max\npop\npush -9\nget_max\npop\nget_max\npop\npop\n")
	out = "None\nerror\nNone\nNone\nNone\nNone\n1\n1\nNone\nNone\nNone\n-3\n-3\n9\n9\n0\n4\n4\n4\n4\n4\n4\n4\n"
	answer = sprint2.Task7(&res)
	fmt.Println(answer)
	//fmt.Println("")
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	fmt.Println("===============================")
	fmt.Println("===============================")
}

func testTask8(t *testing.T) {
	res := []byte("{[()]}")
	out := "True"
	answer := sprint2.Task8(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("()")
	out = "True"
	answer = sprint2.Task8(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("(){}[]")
	out = "True"
	answer = sprint2.Task8(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("({})({[]}[])")
	out = "True"
	answer = sprint2.Task8(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("(((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((()))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))){}\n")
	out = "True"
	answer = sprint2.Task8(&res)
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
	res := []byte("8\n2\npeek\npush 5\npush 2\npeek\nsize\nsize\npush 1\nsize\n")
	out := "None\n5\n2\n2\nerror\n2\n"
	answer := sprint2.Task9(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("10\n1\npush 1\nsize\npush 3\nsize\npush 1\npop\npush 1\npop\npush 3\npush 3\n")
	out = "1\nerror\n1\nerror\n1\n1\nerror\n"
	answer = sprint2.Task9(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("17\n8\npush -82\npush -25\npush -57\npush -24\nsize\npush 12\npush 21\npush 62\npush 64\npush -90\nsize\npop\npeek\npush -10\npush 60\npush 67\nsize\n")
	out = "4\nerror\n8\n-82\n-25\nerror\nerror\n8\n"
	answer = sprint2.Task9(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("42\n2\npush -5\npeek\npop\npeek\nsize\npush -92\npush 91\nsize\npop\npop\npush -8\npush 50\npop\npeek\nsize\npeek\npush 30\npeek\npeek\npop\nsize\npeek\nsize\npush 96\npush -9\npush -74\npeek\npush 22\npush -36\npeek\npop\npeek\npush -1\npop\npush 55\npush -83\npush -8\npush -57\npop\npeek\nsize\npush 80\n")
	out = "-5\n-5\nNone\n0\n2\n-92\n91\n-8\n50\n1\n50\n50\n50\n50\n1\n30\n1\nerror\nerror\n30\nerror\nerror\n30\n30\n96\n96\nerror\nerror\nerror\n-1\n55\n1\n"
	answer = sprint2.Task9(&res)
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
	res := []byte("10\nput -34\nput -23\nget\nsize\nget\nsize\nget\nget\nput 80\nsize\n")
	out := "-34\n1\n-23\n0\nerror\nerror\n1\n"
	answer := sprint2.Task10(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("6\nput -66\nput 98\nsize\nsize\nget\nget\n")
	out = "2\n2\n-66\n98\n"
	answer = sprint2.Task10(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("9\nget\nsize\nput 74\nget\nsize\nput 90\nsize\nsize\nsize\n")
	out = "error\n0\n74\n0\n1\n1\n1\n"
	answer = sprint2.Task10(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("26\nget\nsize\nsize\nget\nsize\nsize\nput 26\nsize\nsize\nsize\nsize\nput 6\nsize\nput 18\nsize\nget\nget\nsize\nget\nget\nsize\nsize\nsize\nget\nsize\nsize\n")
	out = "error\n0\n0\nerror\n0\n0\n1\n1\n1\n1\n2\n3\n26\n6\n1\n18\nerror\n0\n0\n0\nerror\n0\n0\n"
	answer = sprint2.Task10(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	fmt.Println("===============================")
	fmt.Println("===============================")
}

func testTask11(t *testing.T) {
	res := []byte("3")
	out := "3\n"
	answer := sprint2.Task11(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("0")
	out = "1\n"
	answer = sprint2.Task11(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("18")
	out = "4181\n"
	answer = sprint2.Task11(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	fmt.Println("===============================")
	fmt.Println("===============================")
}

func testTask12(t *testing.T) {
	res := []byte("3 1")
	out := "3\n"
	answer := sprint2.Task12(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("10 1")
	out = "9\n"
	answer = sprint2.Task12(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("1 5")
	out = "1\n"
	answer = sprint2.Task12(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("0 5")
	out = "1\n"
	answer = sprint2.Task12(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	fmt.Println("===============================")
	fmt.Println("===============================")
}

func testFinal1(t *testing.T) {
	res := []byte("4\n4\npush_front 861\npush_front -819\npop_back\npop_back\n")
	out := "861\n-819\n"
	answer := sprint2.Final1(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("7\n10\npush_front -855\npush_front 0\npop_back\npop_back\npush_back 844\npop_back\npush_back 823\n")
	out = "-855\n0\n844\n"
	answer = sprint2.Final1(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("6\n6\npush_front -201\npush_back 959\npush_back 102\npush_front 20\npop_front\npop_back\n")
	out = "20\n102\n"
	answer = sprint2.Final1(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	//res := []byte("213\n34\npop_back\npop_back\npop_back\npop_front\npop_back\npop_back\npush_front 284\npop_back\npush_front 799\npop_front\npop_back\npop_back\npop_front\npush_back 792\npush_front 45\npush_front 374\npop_back\npush_back 740\npop_front\npop_back\npush_front 607\npush_front -492\npush_back 661\npush_front -551\npop_front\npush_back 582\npop_front\npush_front 475\npop_front\npush_back -946\npop_back\npush_front -145\npush_back 226\npush_back 515\npop_front\npush_back 222\npush_back 273\npush_back -545\npush_front 413\npop_front\npush_front -272\npop_back\npush_front 950\npop_back\npush_back -764\npop_back\npop_back\npop_back\npop_front\npop_front\npush_back -160\npush_back 411\npop_back\npush_back -385\npush_front 636\npop_back\npop_front\npush_front 784\npush_back 197\npush_front 9\npop_back\npop_front\npush_back 484\npop_front\npop_front\npush_front -838\npop_front\npush_front -207\npush_front -507\npush_front 79\npush_front -697\npop_front\npush_front -316\npush_front -257\npush_front 73\npop_front\npop_front\npush_front -550\npush_front 521\npush_front 291\npush_back 813\npop_back\npop_fron\n")
	//out := "error\nerror\nerror\nerror\nerror\nerror\n284\n799\nerror\nerror\nerror\n792\n374\n740\n-551\n-492\n475\n-946\n-145\n413\n-545\n273\n-764\n222\n515\n950\n-272\n411\n-385\n636\n197\n9\n784\n607\n-838\n-697\n73\n-257\n813\n291\n484\n-160\n226\n964\n395\n714\n541\n60\n286\n207\n-624\n297\n521\n820\n-879\n-550\n-383\n-216\n666\n812\n-316\n410\n-409\n989\n692\n-929\n176\n79\n-507\n-55\n849\n740\n-582\n-411\n138\n-116\n528\n-765\n-580\n-464\n169\n341\n730\n-694\n120\n-921\n-823\n988\n-962\n-240\n-967\n159\n372\n-144\n833\n-726\n548\n289\n-572\n57\n"
	//answer := sprint2.Final1(&res)
	//fmt.Println(answer)
	//fmt.Println(out)
	//fmt.Println("------------------")
	//if answer != out {
	//	t.Error("Error test")
	//}

	fmt.Println("===============================")
	fmt.Println("===============================")
}

func testFinal2(t *testing.T) {
	res := []byte("2 1 + 3 *\n")
	out := "9\n"
	answer := sprint2.Final2(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("7 2 + 4 * 2 +\n")
	out = "38\n"
	answer = sprint2.Final2(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("10 2 4 * -\n")
	out = "2\n"
	answer = sprint2.Final2(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("2 5 - 4 /\n")
	out = "-1\n"
	answer = sprint2.Final2(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("4 13 5 / +\n")
	out = "6\n"
	answer = sprint2.Final2(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("-4 3 * 3 -5 - / -7 -1 - -8 -10 - + *\n")
	out = "8\n"
	answer = sprint2.Final2(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("0 10 * -8 10 / + -9 4 / -10 5 * - * 1 0 - 6 -3 * - 7 3 / 10 -6 - - - /\n")
	out = "-2\n"
	answer = sprint2.Final2(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	fmt.Println("===============================")
	fmt.Println("===============================")
}

func BenchmarkTask7(b *testing.B) {

	lines := []string{}
	lines = append(lines, "9999")
	for i := 1; i <= 5000; i++ {
		lines = append(lines, "push "+strconv.Itoa(i))
	}
	for i := 1; i <= 5000; i++ {
		lines = append(lines, "pop")
	}

	val := ""
	s := sprint2.StackMax{Max: sprint2.MinInt, M: map[int]int{}}

	for l := 1; l < len(lines); l++ {
		line := strings.Split(lines[l], " ")

		switch line[0] {
		case "get_max":
			val = val + s.GetMax() + "\n"
		case "pop":
			v := s.Pop()
			if v == "error" {
				val = val + v + "\n"
			}
		case "push":
			z, err := strconv.ParseInt(line[1], 0, 0)
			if err != nil {
				panic(err)
			}
			s.Push(int(z))
		case "top":
			val = val + s.Top() + "\n"
		default:
			continue
		}
	}
	//None
	//error
	//None
	//None
	//None
	//None
	//1
	//1
	//None
	//None
	//None
	//-3
	//-3
	//9
	//9
	//0
	//4
	//4
	//4
	//4
	//4
	//4
	//4

}
