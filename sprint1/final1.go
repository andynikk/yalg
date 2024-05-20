package sprint1

import (
	"strconv"
	"strings"
)

func Task1(res *[]byte) []string {

	lines := strings.Split(string(*res), "\n")

	arr := strings.Split(strings.Replace(lines[1], "\r", "", -1), " ")
	lines = nil

	previousZero := -1
	curZero := -1

	for i := 0; i < len(arr); i++ {

		switch arr[i] {
		case "0":
			if curZero != -1 {
				previousZero = curZero
			}
			curZero = i
		default:
			arr[i] = strconv.Itoa(i - curZero)
			continue
		}

		for j := previousZero + 1; j < i; j++ {
			differenceBefore := j - previousZero
			if previousZero < 0 {
				differenceBefore = curZero - j
			}
			differenceAfter := curZero - j

			mMin := 0
			if differenceBefore < 0 && differenceAfter >= 0 {
				mMin = differenceAfter
			} else if differenceBefore >= 0 && differenceAfter < 0 {
				mMin = differenceBefore
			} else {
				mMin = min(differenceBefore, differenceAfter)
			}

			arr[j] = strconv.Itoa(mMin)
			curZero = i
		}
	}
	return arr
}
