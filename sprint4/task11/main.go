package main

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func task11(res *[]byte) string {
	lines := strings.Split(string(*res), "\n")
	countExit, _ := strconv.Atoi(lines[0])

	stop := 0

	//m := map[[2]int][4]int{}
	exits := make([][2]int, countExit)

	sch := 1
	for i := 1; i <= countExit; i++ {
		c := lines[i]
		if c == "" {
			continue
		}
		arr := strings.Fields(c)
		x, _ := strconv.Atoi(arr[0])
		y, _ := strconv.Atoi(arr[1])

		//m[[2]int{x, y}] = [4]int{x - 20, x + 20, y - 20, y + 20}
		exits[i-1] = [2]int{x, y}

		sch++
	}

	countStop, _ := strconv.Atoi(lines[sch])
	stops := make([][2]int, countStop)
	for i := sch + 1; i <= sch+countStop; i++ {
		c := lines[i]
		if c == "" {
			continue
		}

		arr := strings.Fields(c)

		x, _ := strconv.Atoi(arr[0])
		y, _ := strconv.Atoi(arr[1])

		stops[i-sch-1] = [2]int{x, y}
	}

	//stop = countStops(exits, stops)
	//sort.Slice(stops, func(i, j int) bool {
	//	if stops[i][0] != stops[j][1] {
	//		return stops[i][0] < stops[j][0]
	//	}
	//
	//	return stops[i][1] < stops[j][1]
	//})

	maximum := 0
	for i := 0; i < len(exits); i++ {
		locMaximum := 0

		grX := exits[i][0]
		grY := exits[i][1]

		for j := 0; j < len(stops); j++ {
			stopsX := stops[j][0]
			stopsY := stops[j][1]

			x := math.Pow(float64(grX)-float64(stopsX), 2)
			y := math.Pow(float64(grY)-float64(stopsY), 2)

			r := math.Sqrt(x + y)

			//if locMaximum != 0 && r > 20 {
			//	break
			//}

			if r <= 20 {
				locMaximum++
			}

		}

		if locMaximum > maximum {
			maximum = locMaximum
			stop = i + 1
		}
	}

	return strconv.Itoa(stop)
}

func distance(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(math.Pow(float64(x2-x1), 2) + math.Pow(float64(y2-y1), 2))
}

func countStops(metroExits, busStops [][2]int) int {
	sort.Slice(busStops, func(i, j int) bool {
		return distance(0, 0, busStops[i][0], busStops[i][1]) < distance(0, 0, busStops[j][0], busStops[j][1])
	})

	maxStops := 0
	bestExit := -1
	for i, exit := range metroExits {
		closestStopIndex := sort.Search(len(busStops), func(i int) bool {
			return distance(exit[0], exit[1], busStops[i][0], busStops[i][1]) <= 20
		})
		stops := len(busStops) - closestStopIndex
		if stops > maxStops {
			maxStops = stops
			bestExit = i
		}
	}
	return bestExit
}
