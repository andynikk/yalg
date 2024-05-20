package sprint3

import (
	"sort"
	"strconv"
	"strings"
)

func Task14(res *[]byte) string {

	lines := strings.Split(string(*res), "\n")
	n, _ := strconv.Atoi(lines[0])
	strArr := lines[1 : 1+n]

	arr := make([][]int, len(strArr))
	for k, val := range strArr {
		v := strings.Split(strings.Replace(val, "\r", "", -1), " ")
		v1, _ := strconv.Atoi(v[0])
		v2, _ := strconv.Atoi(v[1])

		arr[k] = []int{v1, v2}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0]
	})

	notZero := 0
	for i := 0; i < len(arr); i++ {
		if i == 0 {
			continue
		}

		g1 := arr[notZero][0]
		g2 := arr[notZero][1]

		ng1 := arr[i][0]
		ng2 := arr[i][1]

		change := false
		if ng1 < g1 && g1 <= ng2 && ng2 <= g2 {
			arr[notZero][0] = ng1
			g1 = ng1
			change = true
		}

		if ng2 > g2 && g1 <= ng1 && ng1 <= g2 {
			arr[notZero][1] = ng2
			g2 = ng2
			change = true
		}

		if ng1 >= g1 && ng2 <= g2 {
			change = true
		}

		if change {
			arr[i] = []int{0, 0}
		} else {
			notZero = i
		}
	}

	val := []string{}
	for _, v := range arr {
		if v[0] == 0 && v[1] == 0 {
			continue
		}
		val = append(val, strconv.Itoa(v[0])+" "+strconv.Itoa(v[1]))
	}

	return strings.Join(val, "\n") + "\n"
}
