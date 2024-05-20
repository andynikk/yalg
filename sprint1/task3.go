package sprint1

import (
	"slices"
	"strconv"
	"strings"
)

func Task3(res *[]byte) ([]string, string) {

	lines := strings.Split(string(*res), "\n")

	totalX, err := strconv.ParseInt(strings.Replace(lines[1], "\r", "", -1), 0, 0)
	if err != nil {
		panic(err)
	}

	totalY, err := strconv.ParseInt(strings.Replace(lines[0], "\r", "", -1), 0, 0)
	if err != nil {
		panic(err)
	}

	arr := lines[2 : 2+totalY]
	intArr := make([][]int, totalY)

	for i := 0; i < int(totalY); i++ {
		tmpArr := make([]int, totalX)
		arrRow := strings.Split(strings.Replace(arr[i], "\n", "", -1), " ")

		for j := 0; j < int(totalX); j++ {
			intVal, _ := strconv.ParseInt(arrRow[j], 0, 0)
			tmpArr[j] = int(intVal)
		}
		intArr[i] = tmpArr
	}

	x, err := strconv.ParseInt(strings.Replace(lines[totalY+3], "\r", "", -1), 0, 0)
	if err != nil {
		panic(err)
	}
	y, err := strconv.ParseInt(strings.Replace(lines[totalY+2], "\r", "", -1), 0, 0)
	if err != nil {
		panic(err)
	}

	neighbourX1 := x - 1
	neighbourX2 := x + 1

	neighbourY1 := y - 1
	neighbourY2 := y + 1

	valInt := []int{}
	if neighbourX1 >= 0 && neighbourX1 < totalX {
		valInt = append(valInt, intArr[y][neighbourX1])
	}

	if neighbourX2 >= 0 && neighbourX2 < totalX {
		valInt = append(valInt, intArr[y][neighbourX2])
	}

	if neighbourY1 >= 0 && neighbourY1 < totalY {
		valInt = append(valInt, intArr[neighbourY1][x])
	}

	if neighbourY2 >= 0 && neighbourY2 < totalY {
		valInt = append(valInt, intArr[neighbourY2][x])
	}
	slices.Sort(valInt)

	val := make([]string, len(valInt))
	for i := 0; i < len(valInt); i++ {
		val[i] = strconv.Itoa(valInt[i])
	}
	return val, " "
}
