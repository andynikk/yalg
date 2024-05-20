package main

import (
	"strings"
)

const (
	B = iota
	R
)

func lengthenRoad(road [2]int, roads map[[2]int]int) {

	newRoadSection := [2]int{road[0], road[1]}
	for {
		previousRoadSection := [2]int{newRoadSection[0], newRoadSection[1] - 1}
		if _, ok := roads[previousRoadSection]; !ok {
			break
		}

		newRoadSection = [2]int{newRoadSection[0] - 1, newRoadSection[1]}
		if newRoadSection[0] == 0 {
			break
		}

		roads[newRoadSection] += 1
	}
}

func railways(res *[]byte) string {
	lines := strings.Split(string(*res), "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	lines = lines[1:]

	roads := map[int]map[[2]int]int{}

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			typeRoad := R
			if lines[i][j] == 66 {
				typeRoad = B
			}

			typedRoad := roads[typeRoad]
			if typedRoad == nil {
				typedRoad = map[[2]int]int{}
			}

			road := [2]int{i + 1, j + i + 2}
			typedRoad[road] = 1
			lengthenRoad(road, typedRoad)

			roads[typeRoad] = typedRoad
		}
	}

	for road, _ := range roads[B] {
		if _, ok := roads[R][road]; ok {
			return "NO\n"
		}
	}

	return "YES\n"
}
