package sprint4

import (
	"sort"
	"strconv"
	"strings"
)

func Task7(res *[]byte) string {

	lines := strings.Split(string(*res), "\n")

	//n, _ := strconv.Atoi(lines[0])
	s, _ := strconv.Atoi(lines[1])
	str := lines[2]

	strArr := strings.Split(str, " ")
	intArr := make([]int, len(strArr))

	for i := 0; i < len(intArr); i++ {
		intArr[i], _ = strconv.Atoi(strArr[i])
	}
	sort.Ints(intArr)

	firstTwins := map[int][][2]int{}
	for i := 0; i < len(intArr); i++ {
		for j := i + 1; j < len(intArr); j++ {
			key := intArr[i] + intArr[j]
			value := [2]int{i, j}
			firstTwins[key] = append(firstTwins[key], value)
		}
	}

	arr := [][4]int{}

	fours := make(map[[4]int]struct{})
	target := 0
	for k, fsVAl := range firstTwins {
		for _, fs := range fsVAl {
			target = s - k
			if stVal, ok := firstTwins[target]; ok {
				for _, st := range stVal {
					if st[0] <= fs[1] {
						continue
					}
					if _, ok = fours[[4]int{intArr[fs[0]], intArr[fs[1]], intArr[st[0]], intArr[st[1]]}]; ok {
						continue
					}
					four := [4]int{intArr[fs[0]], intArr[fs[1]], intArr[st[0]], intArr[st[1]]}
					arr = append(arr, four)
					fours[four] = struct{}{}
				}
			}
		}
	}

	sort.Slice(arr, func(i, j int) bool {
		if arr[i][0] != arr[j][0] {
			return arr[i][0] < arr[j][0]
		}
		if arr[i][1] != arr[j][1] {
			return arr[i][1] < arr[j][1]
		}
		if arr[i][2] != arr[j][2] {
			return arr[i][2] < arr[j][2]
		}
		if arr[i][3] != arr[j][3] {
			return arr[i][3] < arr[j][3]
		}
		return true
	})

	val := make([]string, len(fours))
	podVal := make([]string, 4)
	for i, k := range arr {
		podVal[0] = strconv.Itoa(k[0])
		podVal[1] = strconv.Itoa(k[1])
		podVal[2] = strconv.Itoa(k[2])
		podVal[3] = strconv.Itoa(k[3])

		val[i] = strings.Join(podVal, " ")
	}

	return strconv.Itoa(len(val)) + "\n" + strings.Join(val, "\n") + "\n"
}
