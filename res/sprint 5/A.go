/*
https://contest.yandex.ru/contest/24810/run-report/113024682/

-- ПРИНЦИП РАБОТЫ --
	Сначала я создал кучу из исходного массива. Для этого я вызываем функцию heapify, которая превращает массив в кучу. Функция heapify начинает свою работу с последнего узла и двигается к корню дерева. Функция сравнивает текущий узел с его потомками, и если необходимо, меняет их местами, чтобы удовлетворить условие кучи (родитель больше детей).
	После построения кучи, я делаю сортировку. На каждой итерации сортировки я обмениваю максимальный элемент (который находится в корне кучи) с последним элементом массива и уменьшаю размер кучи на 1. После этого я вызываю heapify для корня кучи, чтобы восстановить свойство кучи. Этот процесс повторяется до тех пор, пока вся куча не будет отсортирована.

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
	Сначала я строю кучу из этого массива, что гарантирует, что максимальный элемент окажется в корне кучи. Затем я обмениваем этот максимальный элемент с последним элементом массива и уменьшаю размер кучи на 1. После этого вызываю heapify для корня кучи, чтобы восстановить свойство кучи. Массив размером n-1 элементов будет корректно отсортирован. Таким образом, после обмена максимального элемента с последним и вызова heapify, максимальный элемент окажется на своем месте в отсортированной части массива.Повторяя этот процесс для всех элементов до тех пор, пока вся куча не будет отсортирована, мы получаем отсортированный массив размером n.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
n - количество изночальных данных
	Построение кучи выполняется за время O(n). В цикле я вызываю функцию heapify для каждого узла, начиная с последнего уровня и двигаясь к корню дерева. Количество узлов, которые нужно обработать, равно n/2, так как у нас n узлов и половина из них являются листьями, и нам не нужно их рассматривать.
	Сортировка массива, которая состоит из n-1 итераций. На каждой итерации я обмениваю максимальный элемент (который находится в корне кучи) с последним элементом массива и затем вызываем heapify для уменьшенной кучи (без последнего элемента). Это занимает O(log n) времени для каждой итерации. Сложностьсортировки массива составляет O(n*log n).

Итоговая сложность: O(n + n*log n) = O(n*log n)

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
n - количество строк входящих данных

Куча (heap) O(n)
Результирующий массив sortedArray O(n). Но при увелечении sortedArray на 1, куча heap уменьшается на 1.
Итоговая сложность: O(n)

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

	r := finalA(&res)

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

type Node struct {
	value contestant
	left  *Node
	right *Node
}

type contestant struct {
	name    string
	solved  int
	penalty int
}

func heapify(arr []contestant, n, i int) {
	for {
		largest := i
		left := 2*i + 1
		right := 2*i + 2

		//не понял как я должен заменить сверку двух структур на больше-меньше
		//сортировкой и зачем (https://yourbasic.org/golang/how-to-sort-in-go/)
		if left < n && !ASmallerB(arr[i], arr[left]) {
			largest = left
		}

		if right < n && !ASmallerB(arr[largest], arr[right]) {
			largest = right
		}

		if largest == i {
			break
		}

		arr[i], arr[largest] = arr[largest], arr[i]
		i = largest
	}
}

func buildHeap(arr []contestant) {
	n := len(arr)

	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}
}

func sortHeap(arr []contestant) {
	n := len(arr)

	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
}

func ASmallerB(a, b contestant) bool {
	if a.solved != b.solved {
		return a.solved < b.solved
	}
	if a.penalty != b.penalty {
		return a.penalty > b.penalty
	}
	return a.name > b.name
}

func finalA(res *[]byte) string {
	lines := strings.Split(string(*res), "\n")

	commands := lines[1:]
	heap := make([]contestant, len(commands))

	for k, v := range commands {
		tmp := strings.Split(strings.Replace(v, "\r", "", -1), " ")

		solved, _ := strconv.Atoi(tmp[1])
		penalty, _ := strconv.Atoi(tmp[2])

		heap[k] = contestant{name: tmp[0], solved: solved, penalty: penalty}
		_ = k
	}

	buildHeap(heap)
	sortHeap(heap)

	result := make([]string, len(heap))
	for k, v := range heap {
		result[k] = v.name
	}

	return strings.Join(result, "\n") + "\n"
}
