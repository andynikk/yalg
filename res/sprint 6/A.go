/*
https://contest.yandex.ru/contest/25070/run-report/113477795/

-- ПРИНЦИП РАБОТЫ --
	Я создал алгоритм построения минимального остовного дерева на основе алгоритма Крускала.
	Сначала определил структуры данных. Тип Edge для представления ребра графа. Тип UnionFind для реализации структуры данных для оптимизации поиска корневого элемента.
	Функция expensiveNetwork преобразует входные данные в удобный формат, создает массив ребер и вызывает функцию minSpanningTree для построения минимального остовного дерева.
	Функция minSpanningTree принимает количество вершин n и массив ребер edges, сортирует ребра по весу в убывающем порядке, затем проходит по отсортированным ребрам и добавляет их к остовному дереву, если они не образуют цикл.

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
	Я сортирую все ребра графа по весу в убывающем порядке. Это позволяет рассматривать ребра по порядку от самого легкого к самому тяжелому. Ребра были отсортированы верно, иначе возникнут ошибки в построении остовного дерева.
	После сортировки ребер я прохожу по отсортированным ребрам и добавляю их к остовному дереву. Ребро добавляется к остовному дереву только в том случае, если оно не образует цикл с уже добавленными ребрами. Это проверяется с помощью структуры данных UnionFind. На каждом шаге добавляется ребро минимального веса, которое не образует цикл, и таким образом формируется минимальное остовное дерево.
	После добавления всех подходящих ребер я завершаю работу. Для корректности алгоритма важно, чтобы количество добавленных ребер равнялось n-1, где n - количество вершин в графе. Иначе полученная структура не является остовным деревом.


-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
 E - количество ребер в графе.
 Сложность сортировки всех ребер графа по весу составляет O(E log E).
 Проверка наличия цикла и объединение множеств имеет амортизированную сложность близкую к O(1) за операцию.

Итоговая сложность: O(E log E)

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
 E - количество ребер в графе.
 V - количество вершин в графе.

 Для хранения ребер графа и структуры данных требуется O(E) памяти.
 Для хранения вершин графа требуется O(V) памяти.

Итоговая сложность: O(E+V)

*/

package main

import (
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	res, err := inputTxt()
	if err != nil {
		panic(err)
	}

	r := expensiveNetwork(&res)

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

type Edge struct {
	u, v, w int
}

type DSU struct {
	parent []int
}

func newDSU(n int) *DSU {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return &DSU{parent}
}

func (dsu *DSU) find(x int) int {
	if dsu.parent[x] != x {
		dsu.parent[x] = dsu.find(dsu.parent[x])
	}
	return dsu.parent[x]
}

func (dsu *DSU) union(x, y int) {
	dsu.parent[dsu.find(x)] = dsu.find(y)
}

func minSpanningTree(n int, edges []Edge) string {
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].w > edges[j].w
	})

	dsu := newDSU(n)
	totalWeight := 0
	numEdges := 0

	for _, edge := range edges {
		if dsu.find(edge.u) != dsu.find(edge.v) {
			dsu.union(edge.u, edge.v)
			totalWeight += edge.w
			numEdges++
		}
	}

	result := "Oops! I did it again" + "\n"
	if numEdges == n-1 {
		result = strconv.Itoa(totalWeight) + "\n"
	}

	return result
}

func expensiveNetwork(res *[]byte) string {
	lines := strings.Split(string(*res), "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	nm := strings.Split(strings.Replace(lines[0], "\r", "", -1), " ")
	n, _ := strconv.Atoi(nm[0])
	m, _ := strconv.Atoi(nm[1])

	lines = lines[1:]

	edges := make([]Edge, m)
	del := 0
	for i := 0; i < m; i++ {
		line := strings.Split(strings.Replace(lines[i], "\r", "", -1), " ")
		if line[0] == line[1] {
			del++
			continue
		}

		u, _ := strconv.Atoi(line[0])
		v, _ := strconv.Atoi(line[1])
		w, _ := strconv.Atoi(line[2])

		edges[i-del] = Edge{u - 1, v - 1, w}
	}

	edges = edges[:len(edges)-del]
	return minSpanningTree(n, edges)
}
