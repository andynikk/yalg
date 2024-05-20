package main

import (
	"sort"
	"strconv"
	"strings"
)

func chooseLessons(schedule [][2]float64) string {

	begin := schedule[0][0]
	end := schedule[len(schedule)-1][1]

	lessons := make([][][2]float64, 0) // lessons[]
	c := []int{}

	maximum := 0
	maxIndex := 0

	for i := 0; i < len(schedule); i++ {
		beginLesson := schedule[i][0]
		endLesson := schedule[i][1]

		if beginLesson < begin || endLesson > end {
			continue
		}

		for j := 0; j < len(lessons); j++ {
			if lessons[j][len(lessons[j])-1][1] <= beginLesson {
				lessons[j] = append(lessons[j], [2]float64{beginLesson, endLesson})

				if maximum < len(lessons[j]) {
					maximum = len(lessons[j])
					maxIndex = j
				}

			}
		}

		f := [][2]float64{}
		f = append(f, [2]float64{beginLesson, endLesson})
		c = append(c, 1)
		lessons = append(lessons, f)

	}

	str := ""
	for i := 0; i < len(lessons[maxIndex]); i++ {
		str += strconv.FormatFloat(lessons[maxIndex][i][0], 'f', -1, 64) + " " + strconv.FormatFloat(lessons[maxIndex][i][1], 'f', -1, 64) + "\n"
	}
	return strconv.Itoa(len(lessons[maxIndex])) + "\n" + str
}

// <template>
func schedules(res *[]byte) string {

	lines := strings.Split(string(*res), "\n")
	numberClasses, _ := strconv.Atoi(lines[0])

	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	lines = lines[1:]

	schedule := make([][2]float64, numberClasses)
	for i := 0; i < len(lines); i++ {
		line := strings.Split(strings.Replace(lines[i], "\r", "", -1), " ")

		begin, _ := strconv.ParseFloat(line[0], 64)
		end, _ := strconv.ParseFloat(line[1], 64)

		schedule[i] = [2]float64{begin, end}
	}
	sort.Slice(schedule, func(i, j int) bool {
		if schedule[i][1] == schedule[j][1] {
			return schedule[i][0] < schedule[j][0]
		}
		return schedule[i][1] < schedule[j][1]
	})

	result := chooseLessons(schedule)

	return result
}
