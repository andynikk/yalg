/*
-- ПРИНЦИП РАБОТЫ --

Беру крайние точки массива left и right (для первого обхода это 0, len(массив)-1, для рекуривных вазовах указывается).
Получаю опорный элемент (pivot), это элемент по середине массива.

Обхожу полученный массив с left до pivot. При обходе, если элемент меньше элемента pivot, то увеличиваю left.
Нахожу в левой части элемент, который меньше элемента pivot.

Обхожу полученный массив с pivot+1 до right. При обходе, если элемент больше элемента pivot, то увеличиваю right.
Нахожу в правой части элемент, который больше элемента pivot.

Если в левой и правой части найдены элементы, то меняю их местами. увеличваю left, уменьшаю right.

После этого в левой части от pivot лежат элементы меньше pivot, но между собой они еще не отсортированы.
Для сортировки левой части надо проделать то же самое.
Так же в правой части от pivot лежат элементы больше pivot, но между собой они еще не отсортированы.
Для сортировки левой части надо проделать то же самое.

И т.д. пока left < right.

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
Результат постоянен, большие массивы сортирует быстро.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
Сложность быстрой сортировки в составляет O(n log n).

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
По памяти сложность O(logN), т.к. задействована в стеке из-за рекурсии

*/

package sprint3

import (
	"strconv"
	"strings"
)

type programmers struct {
	Data []programmer
}

type programmer struct {
	solved  int
	penalty int
	name    string
}

func ASmallerB(a, b programmer) bool {

	if a.solved != b.solved {
		return a.solved < b.solved
	}

	if a.penalty != b.penalty {
		return a.penalty > b.penalty
	}
	return a.name > b.name
}

func AMoreB(a, b programmer) bool {

	if a.solved != b.solved {
		return a.solved > b.solved
	}

	if a.penalty != b.penalty {
		return a.penalty < b.penalty
	}

	return a.name < b.name

}

func (p *programmers) quicksort(left, right int) {

	if right <= left {
		return
	}

	pivot := p.Data[(left+right)/2]
	i := left
	j := right
	for i <= j {
		for ASmallerB(p.Data[i], pivot) {
			i++
		}
		for AMoreB(p.Data[j], pivot) {
			j--
		}
		if i >= j {
			break
		}

		p.Data[i], p.Data[j] = p.Data[j], p.Data[i]
		i++
		j--
	}

	p.quicksort(left, j)
	p.quicksort(j+1, right)
}

func Final2(res *[]byte) string {

	lines := strings.Split(string(*res), "\n")

	commands := lines[1:]

	arr := programmers{Data: make([]programmer, len(commands))}

	for k, v := range commands {
		tmp := strings.Split(strings.Replace(v, "\r", "", -1), " ")

		v1, _ := strconv.Atoi(tmp[1])
		v2, _ := strconv.Atoi(tmp[2])
		v3 := tmp[0]

		arr.Data[k] = programmer{solved: v1, penalty: v2, name: v3}

	}

	arr.quicksort(0, len(arr.Data)-1)

	val := []string{}
	for v := len(arr.Data) - 1; v >= 0; v-- {
		val = append(val, arr.Data[v].name)
	}

	return strings.Join(val, "\n") + "\n"
}
