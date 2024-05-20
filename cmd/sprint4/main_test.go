package main

import (
	"fmt"
	"testing"
	"yalg/sprint4"
)

func TestFuncSprint1(t *testing.T) {
	//testTask1(t)
	//testTask2(t)
	//testTask3(t)
	//testTask4(t)
	//testTask6(t)
	//testTask7(t)
	//testTask8(t)
	//testTask9(t)
	//testTask10(t)
	testFinal1(t)
	//testFinal2(t)
}

func testTask1(t *testing.T) {
	res := []byte("8\nвышивание крестиком\nрисование мелками на парте\nнастольный керлинг\nнастольный керлинг\nкухня африканского племени ужасмай\nтяжелая атлетика\nтаракановедение\nтаракановедение\n")
	out := "вышивание крестиком\nрисование мелками на парте\nнастольный керлинг\nкухня африканского племени ужасмай\nтяжелая атлетика\nтаракановедение\n"
	answer := sprint4.Task1(&res)
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
	res := []byte("2\n0 1\n")
	out := "2\n"
	answer := sprint4.Task2(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("3\n0 1 0\n")
	out = "2\n"
	answer = sprint4.Task2(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("0\n\n")
	out = "0\n"
	answer = sprint4.Task2(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("10\n0 0 1 0 0 0 1 0 0 1\n")
	out = "4\n"
	answer = sprint4.Task2(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("10\n0 0 1 0 1 1 1 0 0 0\n")
	out = "8\n"
	answer = sprint4.Task2(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("16\n0 0 1 0 0 1 1 1 1 1 0 0 0 0 0 1\n")
	out = "14\n"
	answer = sprint4.Task2(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("72\n1 1 1 1 0 1 1 0 0 0 0 0 1 1 0 0 1 1 0 0 1 0 0 1 0 0 1 1 0 0 1 0 0 0 0 1 1 1 0 0 0 1 0 1 1 1 0 1 1 0 0 1 0 0 0 1 0 1 0 0 0 1 1 1 1 1 1 0 1 0 0 0\n")
	out = "46\n"
	answer = sprint4.Task2(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("9\n0 0 1 1 1 0 0 1 1\n")
	out = "8\n"
	answer = sprint4.Task2(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	fmt.Println("===============================")
	fmt.Println("===============================")
}

func testTask3(t *testing.T) {
	res := []byte("mxyskaoghi\nqodfrgmslc\n")
	out := "YES\n"
	answer := sprint4.Task3(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("agg\nxdd\n")
	out = "YES\n"
	answer = sprint4.Task3(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("agg\nxda\n")
	out = "NO\n"
	answer = sprint4.Task3(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("aba\nxxx")
	out = "NO\n"
	answer = sprint4.Task3(&res)
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
	res := []byte("123\n100003\na\n")
	out := "97\n"
	answer := sprint4.Task4(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("123\n100003\nhash\n")
	out = "6080\n"
	answer = sprint4.Task4(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("123\n100003\nHaSH\n")
	out = "56156\n"
	answer = sprint4.Task4(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("1000\n1000009\nabcdefgh\n")
	out = "436420\n"
	answer = sprint4.Task4(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("3\n256\nABACB\n")
	out = "219\n"
	answer = sprint4.Task4(&res)
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
	res := []byte("1000\n1000009\nabcdefgh\n7\n1 1\n1 5\n2 3\n3 4\n4 4\n1 8\n5 8\n")
	out := "97\n225076\n98099\n99100\n100\n436420\n193195\n"
	answer := sprint4.Task6(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("100\n10\na\n1\n1 1\n")
	out = "7\n"
	answer = sprint4.Task6(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("594\n5597764\nvqvlemal\n5\n4 8\n1 3\n6 8\n3 6\n1 4")
	out = "3110970\n2517540\n4930266\n4639367\n815880\n"
	answer = sprint4.Task6(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("939\n9538360\ncqvkvmzcab\n4\n2 4\n7 8\n2 7\n3 9")
	out = "4361782\n114657\n8617049\n5919440\n"
	answer = sprint4.Task6(&res)
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
	res := []byte("8\n10\n2 3 2 4 1 10 3 0\n")
	out := "3\n0 3 3 4\n1 2 3 4\n2 2 3 3\n"
	answer := sprint4.Task7(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("6\n0\n1 0 -1 0 2 -2\n")
	out = "3\n-2 -1 1 2\n-2 0 0 2\n-1 0 0 1\n"
	answer = sprint4.Task7(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("5\n4\n1 1 1 1 1\n")
	out = "1\n1 1 1 1\n"
	answer = sprint4.Task7(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("20\n381350422\n-142721627 206575381 564131650 -648707194 995752548 393312490 435642014 -9069088 326565756 140484837 -533059899 488966447 0 -780815037 699133600 -268205879 -70733143 496260289 -777196728 85424651\n")
	out = "3\n-648707194 140484837 393312490 496260289\n-268205879 0 85424651 564131650\n-142721627 -9069088 206575381 326565756\n"
	answer = sprint4.Task7(&res)
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
	res := []byte("abcabcbb\n")
	out := "3\n"
	answer := sprint4.Task8(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("bbbbb\n")
	out = "1\n"
	answer = sprint4.Task8(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("pwwkew\n")
	out = "3\n"
	answer = sprint4.Task8(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}
	//
	res = []byte("ojodx\n")
	out = "4\n"
	answer = sprint4.Task8(&res)
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
	res := []byte("6\ntan eat tea ate nat bat")
	out := "0 4\n1 2 3\n5"
	answer := sprint4.Task9(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	//res = []byte("bbbbb\n")
	//out = "1\n"
	//answer = sprint4.Task8(&res)
	//fmt.Println(answer)
	//fmt.Println(out)
	//fmt.Println("------------------")
	//if answer != out {
	//	t.Error("Error test")
	//}
	//
	//res = []byte("pwwkew\n")
	//out = "3\n"
	//answer = sprint4.Task8(&res)
	//fmt.Println(answer)
	//fmt.Println(out)
	//fmt.Println("------------------")
	//if answer != out {
	//	t.Error("Error test")
	//}
	////
	//res = []byte("ojodx\n")
	//out = "4\n"
	//answer = sprint4.Task8(&res)
	//fmt.Println(answer)
	//fmt.Println(out)
	//fmt.Println("------------------")
	//if answer != out {
	//	t.Error("Error test")
	//}

	fmt.Println("===============================")
	fmt.Println("===============================")
}

func testTask10(t *testing.T) {
	res := []byte("5\n1 2 3 2 1\n5\n3 2 1 5 6\n")
	out := "3\n"
	answer := sprint4.Task10(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("5\n1 2 3 4 5\n3\n4 5 9\n")
	out = "2\n"
	answer = sprint4.Task10(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("30\n69 114 84 135 195 18 64 137 164 244 81 210 107 218 49 115 208 216 206 80 241 189 173 217 50 82 202 202 222 34\n30\n214 157 57 196 58 128 89 168 197 1 200 167 243 165 89 44 18 247 175 241 136 219 114 206 254 198 65 5 144 101\n")
	out = "1\n"
	answer = sprint4.Task10(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("30\n156 176 45 212 116 8 165 5 15 111 147 223 30 119 175 139 131 205 199 138 178 26 80 76 180 67 125 109 83 12\n30\n105 200 10 213 83 123 168 74 226 43 44 228 185 188 49 104 78 200 162 193 20 7 64 125 197 199 197 184 201 59\n")
	out = "1\n"
	answer = sprint4.Task10(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("1\n1\n1\n1\n")
	out = "1\n"
	answer = sprint4.Task10(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("1\n1\n1\n2\n")
	out = "0\n"
	answer = sprint4.Task10(&res)
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
	res := []byte("3\ni love coffee\ncoffee with milk and sugar\nfree tea for everyone\n3\ni like black coffee without milk\neveryone loves new year\nmary likes black coffee without milk\n")
	out := "1 2\n3\n2 1\n"
	answer := sprint4.Final1(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("6\nbuy flat in moscow\nrent flat in moscow\nsell flat in moscow\nwant flat in moscow like crazy\nclean flat in moscow on weekends\nrenovate flat in moscow\n1\nflat in moscow for crazy weekends\n")
	out = "4 5 1 2 3\n"
	answer = sprint4.Final1(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("3\ni like dfs and bfs\ni like dfs dfs\ni like bfs with bfs and bfs\n1\ndfs dfs dfs dfs bfs\n")
	out = "3 1 2\n"
	answer = sprint4.Final1(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("10\ntjegerxbyk pdvmj wulmqfrx\npndygsm dvjihmxr tcdtqsmfe\ntxamzxqzeq dxkxwq aua\nhsciljsrdo fipazun kngi\nxtkomk aua wulmqfrx ydkbncmzee\npndygsm cqvffye pyrhcxbcef\nszyc uffqhayg ccktodig\nntr wpvlifrgjg htywpe\nkngi tjegerxbyk zsnfd\ntqilkkd gq qc fipazun\n5\ndxkxwq htywpe\naua tjegerxbyk\nxtkomk tjegerxbyk\nszyc fipazun\nxtkomk tjegerxbyk\n")
	out = "3 8\n1 3 5 9\n1 5 9\n4 7 10\n1 5 9\n"
	answer = sprint4.Final1(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("10\ning ioizzlux knjklikxxq\nlntw qinx iltxhmrmfp\nkmsjgmfyu gic xivg pu\nnlbudxrvnv vrbs spdeuoqr\neyew sosx iecoujab xtokvidtwf\nspdeuoqr qpcmtqzs ioizzlux\njeqoseryf cvoafqw hnovlrgzp\nuanxvs gic ljymhdwgw\nspdeuoqr pu xivg mjotkrq\nrg uqblrnrgwr pu aeiv\n5\nuqblrnrgwr\nrg cvoafqw\nvrbs xivg pu\nkqweuyzgkt\nrg gic xivg\n")
	out = "10\n7 10\n3 9 4 10\n\n3 8 9 10\n"
	//out = "10\n7 10\n3 9 4 10\n3 8 9 10\n"
	answer = sprint4.Final1(&res)
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
	res := []byte("10\nget 1\nput 1 10\nput 2 4\nget 1\nget 2\ndelete 2\nget 2\nput 1 5\nget 1\ndelete 2\n")
	out := "None\n10\n4\n4\nNone\n5\nNone\n"
	answer := sprint4.Final2(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("8\nget 9\ndelete 9\nput 9 1\nget 9\nput 9 2\nget 9\nput 9 3\nget 9\n")
	out = "None\nNone\n1\n2\n3\n"
	answer = sprint4.Final2(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	res = []byte("15\nput 20 27\nget 20\nput 20 21\nget 20\nget 20\nget -1\nget 20\nget -3\ndelete 20\nget -29\nget -33\ndelete -29\nget 16\nget 14\nput 29 39\n")
	out = "27\n21\n21\nNone\n21\nNone\n21\nNone\nNone\nNone\nNone\nNone\n"
	answer = sprint4.Final2(&res)
	fmt.Println(answer)
	fmt.Println(out)
	fmt.Println("------------------")
	if answer != out {
		t.Error("Error test")
	}

	fmt.Println("===============================")
	fmt.Println("===============================")
}
