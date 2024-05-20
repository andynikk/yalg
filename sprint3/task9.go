package sprint3

import (
	"slices"
	"strconv"
	"strings"
)

func Task9(res *[]byte) string {

	lines := strings.Split(string(*res), "\n")
	e, _ := strconv.Atoi(lines[0])
	if e == 0 {
		return ""
	}

	d := strings.Split(strings.Replace(lines[1], "\r", "", -1), " ")
	k, _ := strconv.Atoi(lines[2])

	m := MaxInArr(d) + 1

	arrMax := []int{}
	idxs := make([]int, m)

	maxVUZ := 0
	for i := 0; i < len(d); i++ {
		v, _ := strconv.Atoi(d[i])
		idxs[v]++

		if idxs[v] > maxVUZ {
			maxVUZ = idxs[v]
		}
	}

	tmpMax := make([]int, maxVUZ+1)

	for i := 0; i < len(idxs); i++ {
		v := idxs[i]
		if v > 0 && tmpMax[v] == 0 {
			arrMax = append(arrMax, v)
		}
		tmpMax[v]++
	}
	slices.Sort(arrMax)

	val := []string{}

	for j := len(arrMax) - 1; j >= 0; j-- {
		if len(val) == k {
			break
		}

		for i := 0; i < len(idxs); i++ {
			if idxs[i] != arrMax[j] {
				continue
			}
			val = append(val, strconv.Itoa(i))
			if len(val) == k {
				break
			}
		}
	}

	return strings.Join(val, " ")
}
