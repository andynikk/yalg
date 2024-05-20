package main

import (
	"sort"
	"strconv"
	"strings"
)

type Line struct {
	index int
	text  string
}

func insertingLines(res *[]byte) string {
	lines := strings.Split(string(*res), "\n")

	text := lines[0]
	n, _ := strconv.Atoi(lines[1])

	arrSort := make([]Line, n)
	shifts := map[int]string{}

	for i := 2; i < n+2; i++ {
		insertData := strings.Split(lines[i], " ")

		insertIndex, _ := strconv.Atoi(insertData[1])

		shifts[insertIndex] = insertData[0]

		arrSort[i-2] = Line{insertIndex, insertData[0]}
	}

	sort.Slice(arrSort, func(i, j int) bool {
		return arrSort[i].index < arrSort[j].index
	})

	//increment := 0
	//for i := 0; i < len(arrSort); i++ {
	//	position := arrSort[i].index + increment
	//
	//	text = text[:position] + arrSort[i].text + text[position:]
	//	increment += len(arrSort[i].text)
	//}

	var pos int
	var result strings.Builder
	for _, val := range arrSort {
		for i := pos; i < val.index; i++ {
			result.WriteByte(text[i])
		}
		result.WriteString(val.text)
		pos = val.index
	}
	for i := pos; i < len(text); i++ {
		result.WriteByte(text[i])
	}
	return result.String()

	//return text
}
