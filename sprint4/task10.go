package sprint4

import (
	"hash/fnv"
	"strconv"
	"strings"
)

func Task10(res *[]byte) string {

	lines := strings.Split(string(*res), "\n")

	arr1 := strings.Split(strings.Replace(lines[1], "\r", "", -1), " ")
	arr2 := strings.Split(strings.Replace(lines[3], "\r", "", -1), " ")
	//maximum := 0

	slice1 := make([]int, len(arr1))
	for i := 0; i < len(arr1); i++ {
		slice1[i], _ = strconv.Atoi(arr1[i])
	}
	slice2 := make([]int, len(arr2))
	for i := 0; i < len(arr2); i++ {
		slice2[i], _ = strconv.Atoi(arr2[i])
	}

	if len(arr1) == 0 && len(arr2) == 0 {
		return "0\n"
	}

	m := getAllSequences(slice1)
	left, right := 0, len(arr2)

	//step := 1
	//mid := len(arr2) / 2
	hashAllSlice2 := hash10(slice2)
	if v, ok := m[hashAllSlice2]; ok {
		return strconv.Itoa(v) + "\n"
	}

	//for left < right {
	//
	//	hashWithoutFirst := removeFirstElementHash(slice2[left], hashAllSlice2)
	//	if v, ok := m[hashWithoutFirst]; ok {
	//		return strconv.Itoa(v) + "\n"
	//	}
	//
	//	hashWithoutLast := removeLastElementHash(slice2[right-1], hashAllSlice2)
	//	if v, ok := m[hashWithoutLast]; ok {
	//		return strconv.Itoa(v) + "\n"
	//	}
	//
	//	hashAllSlice2 = removeFirstAndLastFromHash(hashAllSlice2, slice2[left], slice2[right-1])
	//	if v, ok := m[hashAllSlice2]; ok {
	//		return strconv.Itoa(v) + "\n"
	//	}

	maximum := hashWindow(hashAllSlice2, slice2[left:right], m)

	//left++
	//right--
	//}

	//return strconv.Itoa(maximum) + "\n"
	return maximum
}

func hashWindow(hashAllSlice2 uint32, slice2 []int, m map[uint32]int) string {

	if len(slice2) == 0 {
		return "0\n"
	}

	hashWithoutFirst := removeFirstElementHash(slice2[0], hashAllSlice2)
	if v, ok := m[hashWithoutFirst]; ok {
		return strconv.Itoa(v) + "\n"
	}

	hashWithoutLast := removeLastElementHash(slice2[len(slice2)-1], hashAllSlice2)
	if v, ok := m[hashWithoutLast]; ok {
		return strconv.Itoa(v) + "\n"
	}

	//hashAllSlice2 = removeFirstAndLastFromHash(hashAllSlice2, slice2[0], slice2[len(slice2)-1])
	//if v, ok := m[hashAllSlice2]; ok {
	//	return strconv.Itoa(v) + "\n"
	//}

	hashWindow(hashWithoutLast, slice2[:len(slice2)-1], m)
	hashWindow(hashWithoutFirst, slice2[1:], m)

	return ""
}

func removeFirstElementHash(firstElemInt int, hashVal uint32) uint32 {
	firstElemStr := strconv.Itoa(firstElemInt)
	h := fnv.New32a()
	h.Write([]byte(firstElemStr))
	hashVal ^= h.Sum32()
	return hashVal
}

func removeLastElementHash(lastElemInt int, hashVal uint32) uint32 {
	lastElemStr := strconv.Itoa(lastElemInt)
	h := fnv.New32a()
	h.Write([]byte(lastElemStr))
	hashVal ^= h.Sum32()
	return hashVal
}

func removeFirstAndLastFromHash(hashVal uint32, firstElem int, lastElem int) uint32 {
	firstElemStr := strconv.Itoa(firstElem)
	lastElemStr := strconv.Itoa(lastElem)
	firstH := fnv.New32a()
	firstH.Write([]byte(firstElemStr))
	lastH := fnv.New32a()
	lastH.Write([]byte(lastElemStr))

	hashVal ^= firstH.Sum32()
	hashVal ^= lastH.Sum32()

	return hashVal
}

func hash10(arr []int) uint32 {
	h := fnv.New32a()
	for _, num := range arr {
		h.Write([]byte(strconv.Itoa(num)))
	}
	return h.Sum32()
}

func addElementToHash(hashVal uint32, elem int) uint32 {
	elemStr := strconv.Itoa(elem)
	h := fnv.New32a()
	h.Write([]byte(elemStr))
	hashVal ^= h.Sum32()
	return hashVal
}

func getAllSequences(arr []int) map[uint32]int {

	m := map[uint32]int{}

	for i := 0; i < len(arr); i++ {
		h := hash10(arr[i : i+1])
		m[h] = 1
		for j := i + 1; j < len(arr); j++ {
			h = addElementToHash(h, arr[j])
			m[h] = len(arr[i : j+1])
		}
	}

	return m
}
