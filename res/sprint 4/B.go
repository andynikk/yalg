/*

https://contest.yandex.ru/contest/24414/run-report/111726175/

-- ПРИНЦИП РАБОТЫ --
	Я реализовал три функции put, get, delete. Эти функции обеспечивают базовую функциональность для работы с данными в хэш-таблице: добавление, получение и удаление значений по ключу.
	Функция put добавляет и/или обновляет значение в хэш-таблице по заданному ключу. Сначала она вычисляет хэш-ключа с помощью функции hashFinal2. Затем она находит соответствующую корзину в хэш-таблице и проверяет, есть ли уже в ней элемент с таким ключом. Если есть, то значение обновляется. Если нет, то новая пара ключ-значение добавляется в корзину.
	Функция get получает значение из хэш-таблицы по заданному ключу. Она также вычисляет хэш-ключа с помощью hashFinal2, затем находит соответствующую корзину и ищет в ней элемент с заданным ключом. Если элемент найден, то функция возвращает соответствующее значение. В противном случае возвращается пустая строка.
	Функция delete удаляет значение из хэш-таблицы по заданному ключу. Она также вычисляет хэш-ключа, находит соответствующую корзину и ищет в ней элемент с заданным ключом. Если элемент найден, то он удаляется из корзины. Если элемент не найден, то ничего не происходит.


-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
Команды правильно извлекаются из строк и обрабатываются. Функции put, get и delete работают правильно, изменяя и извлекая значения по ключу.
Функция put корректна, если элемент успешно добавлен и его значение можно прочитать с помощью функции get.
Функция get корректна, если она возвращает правильное значение для существующего ключа и -1 для отсутствующего.
Функция delete корректна, если элемент успешно удален и его значение больше не доступно через функцию get.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
Функция put работает за O(n). Т.к. hashFinal2, поиск корзины, добавление/обновление элемента в корзине работают, каждые, за константное время O(1). Но общая сложность по функции зависит от количества операций put.
Функция get работает за O(n). Т.к. hashFinal2, поиск корзины, поиск элемента в корзине работают, каждые, за константное время O(1). Но общая сложность по фунуции зависит от количества операций get.
Функция delete работает за O(n). Т.к. hashFinal2, поиск корзины, поиск и удаление элемента из корзины работают, каждые, за константное время O(1). Но общая сложность по функции зависит от количества операций delete.
Итоговая сложность O(n), где n - это время выполнения всех операций, эти операций не зависит от размера хэш-таблицы и остается постоянным в среднем случае.

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
Память, выделенная под хэш-таблицу (HashTable), зависит от размера size. Память, выделенная под слайс result в функции Final2, зависит от количества операций.
Итоговая сложность по памяти O(s+r), где s размер хэш-таблицы и r количество операций.

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

	r := Final2(&res)

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

type pair struct {
	key   int
	value int
}

type HashTable struct {
	size    int
	buckets [][]pair
}

func newHashTable(size int) *HashTable {
	return &HashTable{
		size:    size,
		buckets: make([][]pair, size),
	}
}

func hashFinal2(key string) int {
	hash := 0
	for i := 0; i < len(key); i++ {
		hash = (hash << 5) + hash + int(key[i])
	}
	return hash
}

func (ht *HashTable) put(key string, value int) {
	index := hashFinal2(key) % ht.size
	for i, p := range ht.buckets[index] {
		if p.key == hashFinal2(key) {
			ht.buckets[index][i] = pair{hashFinal2(key), value}
			return
		}
	}
	ht.buckets[index] = append(ht.buckets[index], pair{hashFinal2(key), value})
}

func (ht *HashTable) get(key string) (int, bool) {
	index := hashFinal2(key) % ht.size
	for _, p := range ht.buckets[index] {
		if p.key == hashFinal2(key) {
			return p.value, true
		}
	}
	return 0, false
}

func (ht *HashTable) delete(key string) (int, bool) {
	index := hashFinal2(key) % ht.size
	for i, p := range ht.buckets[index] {
		if p.key == hashFinal2(key) {
			//ht.buckets[index] = append(ht.buckets[index][:i], ht.buckets[index][i+1:]...)
			ht.buckets[index][i], ht.buckets[index][len(ht.buckets[index])-1] = ht.buckets[index][len(ht.buckets[index])-1], ht.buckets[index][i]
			ht.buckets[index] = ht.buckets[index][:len(ht.buckets[index])-1]
			return p.value, true
		}
	}

	return -1, false
}

func Final2(res *[]byte) string {

	lines := strings.Split(string(*res), "\n")

	commands := lines[1:]
	val := []string{}
	//q := Queue10{}

	ht := newHashTable(1<<12 + 1)
	for c := 0; c < len(commands); c++ {
		command := strings.Split(commands[c], " ")
		if command[0] == "" {
			continue
		}

		switch command[0] {
		case "get":
			v, found := ht.get(command[1])
			if !found {
				val = append(val, "None")
				continue
			}
			val = append(val, strconv.Itoa(v))
		case "put":
			//k, _ := strconv.Atoi(command[1])
			v, _ := strconv.Atoi(command[2])

			ht.put(command[1], v)
		case "delete":
			v, found := ht.delete(command[1])
			if !found {
				val = append(val, "None")
				continue
			}
			val = append(val, strconv.Itoa(v))
		default:
			continue
		}
	}

	return strings.Join(val, "\n") + "\n"
}
