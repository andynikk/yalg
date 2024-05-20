package sprint3

import (
	"reflect"
)

func merge_sort(arr []int, lf int, rg int) {
	if rg-lf <= 1 {
		return
	}
	merge_sort(arr, lf, (rg+lf)/2)
	merge_sort(arr, (rg+lf)/2, rg)

	res := merge(arr, lf, (rg+lf)/2, rg)
	idx := 0
	for i := lf; i < rg; i++ {
		arr[i] = res[idx]
		idx++
	}
}

func merge(arr []int, left int, mid int, right int) (result []int) {
	first := arr[left:mid]
	second := arr[mid:right]

	result = make([]int, len(first)+len(second))
	f, s, r := 0, 0, 0
	for f < len(first) && s < len(second) {
		if first[f] <= second[s] {
			result[r] = first[f]
			f++
		} else {
			result[r] = second[s]
			s++
		}
		r++
	}
	for f < len(first) {
		result[r] = first[f]
		f++
		r++
	}
	for s < len(second) {
		result[r] = second[s]
		s++
		r++
	}
	return result
}

func Task11() {
	a := []int{1, 4, 9, 2, 10, 11}
	b := merge(a, 0, 3, 6)

	expected := []int{1, 2, 4, 9, 10, 11}
	if !reflect.DeepEqual(b, expected) {
		panic("WA. Merge")
	}

	c := []int{1, 4, 2, 10, 1, 2}
	merge_sort(c, 0, 6)
	expected = []int{1, 1, 2, 2, 4, 10}
	if !reflect.DeepEqual(c, expected) {
		panic("WA. MergeSort")
	}
}
