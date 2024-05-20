/*
-- ПРИНЦИП РАБОТЫ --
Для начала я решил найти индекс элемента с наименьшим значением. Реализовал это бинарным поиском.
После того, как я нашел индекс минимального значения, я проверил в какой части от минимального значения лежит искомое значение.
Взял эту часть и на ее базе реализовал поиск разделением массива на две части. Взял основной массив разделил его на две части.
Если в массиве нечетное количество элементов, то большее количество во воторой (правой) части массива. Условием больше-меньше
определяю в какой части массива лежит искомый элемент и повотряю. Суммирую значение индекса подмоссива.
Если в финальном массиве 0 элементов, то возвращаю -1.Если в финальной части 1 элемент и он равен искомому то возвращаю 0.
Обратно отдаю получившийся индекс

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
Из описания алгоритма следует, что элемент с необходимым значением находится. Если искомое значение отсутсвует, то возвращается -1


-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
Поиск индекса элемента O(log N) где N это количество элемнтов массива, я не обхожу весь массив.
Максимально я два раза обхожу весь массив. И с каждым шагом я беру только половину от полученного подмассива.

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
В памяти будет занимать n элементов массива и меннятся не будет.

*/

package sprint3

func search(arr []int, k int) int {
	if len(arr) == 0 {
		return -1
	}
	if len(arr) == 1 {
		if arr[0] == k {
			return 0
		} else {
			return -1
		}
	}

	mid := len(arr) / 2

	idx1 := 0
	idx2 := mid

	valIdxL1 := arr[idx1]
	valIdxL2 := arr[max(0, idx2-1)]

	left := k >= valIdxL1 && k <= valIdxL2

	if !left {
		idx1 = mid
		idx2 = len(arr)
	}

	idx := search(arr[idx1:idx2], k)
	if idx != -1 {
		idx += idx1
	}

	return idx
}

func indexes(arr []int) (int, int) {
	if len(arr) == 1 {
		return 0, 1
	}

	mid := len(arr) / 2
	idx1, idx2 := 0, mid

	valIdxL1, valIdxL2 := arr[idx1], arr[idx2-1]
	valIdxR1, valIdxR2 := arr[idx2], arr[len(arr)-1]

	left := (valIdxL1 <= valIdxR1 && valIdxL2 <= valIdxR1 && valIdxL1 <= valIdxR2 && valIdxR2 <= valIdxR2) ||
		(valIdxR1 <= valIdxL1 && valIdxR1 > valIdxL2)

	if !left {
		idx1, idx2 = mid, len(arr)
	}

	return idx1, idx2
}

func indexMinVal(arr []int, k *int) {

	if len(arr) == 1 {
		return
	}

	idx1, idx2 := indexes(arr)
	*k += idx1
	indexMinVal(arr[idx1:idx2], k)
}

func brokenSearch(arr []int, k int) int {

	idx1, idx2 := indexes(arr)
	indexMinVal(arr[idx1:idx2], &idx1)

	valL1, valL2 := arr[0], arr[max(idx1-1, 0)]

	idx2 = len(arr)
	if k >= min(valL1, valL2) && k < max(valL1, valL2) {
		idx2, idx1 = idx1, 0
	}

	idx := search(arr[idx1:idx2], k)
	if idx == -1 {
		return -1
	}
	return idx + idx1
}

func testFinal1() {
	arr := []int{19, 21, 100, 101, 1, 4, 5, 7, 12}
	if brokenSearch(arr, 5) != 6 {
		println("test 1 error")
	} else {
		println("test 1 OK")
	}

	arr = []int{5, 1}
	if brokenSearch(arr, 1) != 1 {
		println("test 2 error")
	} else {
		println("test 2 OK")
	}

	arr = []int{3, 5, 6, 7, 9, 1, 2}
	if brokenSearch(arr, 4) != -1 {
		println("test 3 error")
	} else {
		println("test 3 OK")
	}

	arr = []int{3, 6, 7}
	if brokenSearch(arr, 8) != -1 {
		println("test 4 error")
	} else {
		println("test 4 OK")
	}

	arr = []int{8158, 8164, 8228, 8296, 8394, 8426, 8719, 8728, 9182, 9220, 9388, 9453, 9512, 9544, 9962, 37, 265, 392, 444, 519, 549, 649, 910, 999, 1056, 1090, 1211, 1429, 1526, 1628, 1688, 1694, 1733, 1816, 1994, 2039, 2290, 2335, 2389, 2667, 2690, 2748, 2799, 2831, 2905, 2927, 2993, 3033, 3048, 3132, 3166, 3330, 3346, 3417, 3457, 3505, 3573, 3599, 367, 3691, 3839, 4009, 4013, 4151, 4192, 4219, 4305, 4548, 4799, 4862, 4866, 4869, 4976, 5190, 5401, 5452, 5477, 5553, 5938, 5945, 6041, 6099, 6132, 6163, 6437, 6524, 6780, 6801, 6888, 7101, 7187, 7220, 7228, 7346, 7387, 7546, 7762, 7981, 7983, 8120}
	if brokenSearch(arr, 9220) != 9 {
		println("test 5 error")
	} else {
		println("test 5 OK")
	}

	arr = []int{3271, 3298, 3331, 3397, 3407, 3524, 3584, 3632, 3734, 3797, 3942, 4000, 4180, 4437, 4464, 4481, 4525, 4608, 4645, 4803, 4804, 4884, 4931, 4965, 5017, 5391, 5453, 5472, 5671, 5681, 5959, 6045, 6058, 6301, 6529, 6621, 6961, 7219, 7291, 7372, 7425, 7517, 7600, 7731, 7827, 7844, 7987, 8158, 8169, 8265, 8353, 8519, 8551, 8588, 8635, 9209, 9301, 9308, 9336, 9375, 9422, 9586, 9620, 9752, 9776, 9845, 9906, 9918, 16, 25, 45, 152, 199, 309, 423, 614, 644, 678, 681, 725, 825, 830, 936, 1110, 1333, 1413, 1617, 1895, 1938, 2107, 2144, 2184, 2490, 2517, 2769, 2897, 2970, 3023, 3112, 3156}
	if brokenSearch(arr, 25) != 69 {
		println("test 6 error")
	} else {
		println("test 6 OK")
	}
	arr = []int{1}
	if brokenSearch(arr, 1) != 0 {
		println("test 7 error")
	} else {
		println("test 7 OK")
	}
	arr = []int{1, 5}
	if brokenSearch(arr, 1) != 0 {
		println("test 8 error")
	} else {
		println("test 8 OK")
	}

	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 0}
	if brokenSearch(arr, 1) != 0 {
		println("test 9 error")
	} else {
		println("test 9 OK")
	}
}

func Final1() {
	testFinal1()
}
