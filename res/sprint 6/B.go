/*
https://contest.yandex.ru/contest/25070/run-report/113399220/

-- ПРИНЦИП РАБОТЫ --
	Код реализует функцию railways, которая проверяет возможность построения железных дорог между городами с учетом ограничений.
	Функция railways разбивает входные данные на строки, строит граф дорог между городами на основе типов дорог ('B' или 'R') и вызывает checkRoads для проверки наличия циклов. В зависимости от результата, функция возвращает "YES" или "NO".
	Функция checkRoads проверяет все города в графе, вызывая dfs для каждой не посещенной вершины. Если обнаруживается цикл, функция возвращает false, иначе true.
	Функция dfs обходит в глубину, которая проверяет наличие циклов в графе, представленном в виде списка смежности. Она использует стек для хранения вершин и массив цветов для пометки вершин (0 - не посещена, 1 - посещена, 2 - завершена). Если обнаруживается цикл, функция возвращает true, иначе false.

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
	Функция railways разбивает входные данные на строки, строит граф дорог между городами на основе типов дорог ('B' или 'R') и вызывает checkRoads для проверки наличия циклов. Если checkRoads возвращает true, то функция возвращает "YES", иначе "NO". Поскольку checkRoads корректно проверяет наличие циклов, то и функция railways работает корректно.
	Функция checkRoads проверяет наличие циклов в графе, вызывая dfs для каждой не посещенной вершины. Если хотя бы один вызов dfs обнаруживает цикл, функция возвращает false, иначе true. Таким образом, эта функция корректно проверяет наличие циклов в графе.
	Функция dfs реализует алгоритм обхода в глубину для проверки наличия циклов в графе. После обхода всех вершин, если обнаруживается ребро к уже посещенной вершине, то это означает наличие цикла. В противном случае, циклов нет. Этот алгоритм корректно работает для обнаружения циклов в графе.


-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
 V - количество городов (вершин)
 E - дорог между городами (количество ребер).

 Парсинг входных данных и построение графа: O(E)
 Обход графа в функции checkRoads с использованием обхода в глубину: O(V + E)

 Итоговая сложность алгоритма: O(V + 2E).

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
 V - количество городов (вершин)
 E - дорог между городами (количество ребер).
 W - максимальная глубина рекурсии

 O(V + E), так как нужно хранить информацию о вершинах и ребрах.
 Дополнительная память для рекурсивных вызовов функции dfs: O(W)

 Итоговая сложность алгоритма: O(V + E + W).

*/

package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	res, err := inputTxt()
	if err != nil {
		panic(err)
	}

	r := railways(&res)

	err = outputStr(&r)
}

func inputTxt() ([]byte, error) {
	res, err := os.ReadFile("input.txt")
	return res, err
}

func outputStr(r *string) error {

	file, err := os.Create("output.txt")
	if err != nil {
		return err
	}
	file.WriteString(*r)
	return nil

}

const (
	B = iota
	R
)

func typeRoad(roadType uint8) int {
	if roadType == 'B' {
		return B
	}

	return R
}

func dfs(startNode int, roads [][]int, colors []int) bool {
	stack := []int{}
	stack = append(stack, startNode)

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if colors[node] == 0 {
			colors[node] = 1
			stack = append(stack, node)

			for i := 0; i < len(roads[node]); i++ {
				if colors[roads[node][i]] == 0 {
					stack = append(stack, roads[node][i])
				} else if colors[roads[node][i]] == 1 {
					return true
				}
			}
		} else if colors[node] == 1 {
			colors[node] = 2
		}
	}
	return false
}

func checkRoads(cities int, roads [][]int) bool {
	colors := make([]int, cities+1)
	for node := 1; node < len(colors); node++ {
		if colors[node] != 0 {
			continue
		}
		if dfs(node, roads, colors) {
			return false
		}
	}
	return true
}

func railways(res *[]byte) string {
	lines := strings.Split(string(*res), "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	n, _ := strconv.Atoi(lines[0])
	lines = lines[1:]

	roads := make([][]int, n+1)
	for i := 0; i < n-1; i++ {
		cityOut := i + 1
		for j := 0; j < len(lines[i]); j++ {
			cityIn := cityOut + j + 1
			if typeRoad(lines[i][j]) == B {
				roads[cityOut] = append(roads[cityOut], cityIn)
			} else {
				roads[cityIn] = append(roads[cityIn], cityOut)
			}
		}
	}

	if checkRoads(n, roads) {
		return "YES\n"
	} else {
		return "NO\n"
	}
}
