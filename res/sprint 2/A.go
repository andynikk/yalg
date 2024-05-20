/*
-- ПРИНЦИП РАБОТЫ --
Я реализовал интерфейс с добавлением и извлечением элементов с обоих концов (Дек).
Элементы можно добавлять как в начало, так и в конец очереди.
Извлекать из очереди элементы так же можно и с начала, и с конца.

При запуске устанавливается максимально значение очереди и начальная ячейка 0.

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
Из описания алгоритма следует, что элементы устанавливаются в зависимости от команды либо в начало очереди, либо в конец. В свойства очереди записывается текущая занятая ячейка. И в любой момент можно получить номер свободной ячейки с начала или конца.


Дек хранит данные в той очередности, в которой их помещали с начала или конца. Забирать данные можно так же с начала или конца.


-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
Добавление в дек стоит O(1), потому что добавление в начало или в конец стоит O(1), т.к. мы знаем последнюю занятую ячейку. И кладем в дек в одно действие.

Извлечение из дека стоит так же O(1), т.к. мы знаем последнюю занятую ячейку. И забираем из дека в одно действие.

Общая сложность будет O(n), где n это количество команд заданных в задании.

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
Дек содержит m элементов. Количество макимального m указывается в программе (как переменная m) и сразу резирвируется в памяти.

Дек содержит k элементов, занимает O(k) памяти.
Поэтому и мой дек будет потреблять O(m) памяти.
*/


//https://contest.yandex.ru/contest/22781/run-report/108936249/

package main

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

func main() {
	res, err := inputTxt()
	if err != nil {
		panic(err)
	}

	r := Final1(&res)
	
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

func Final1(res *[]byte) string {

	lines := strings.Split(string(*res), "\n")

	m, _ := strconv.Atoi(lines[1])

	commands := lines[2:]
	val := []string{}
	q := Queue1{Data: make([]int, m), Max: m, Last: 1}

	for c := 0; c < len(commands); c++ {
		command := strings.Split(commands[c], " ")
		if command[0] == "" {
			continue
		}

		switch command[0] {
		case "push_front":
			v, err := strconv.ParseInt(command[1], 0, 0)
			if err != nil {
				panic(err)
			}
			err = q.PushFront(int(v))
			if err != nil {
				val = append(val, err.Error())
			}
		case "push_back":
			v, err := strconv.ParseInt(command[1], 0, 0)
			if err != nil {
				panic(err)
			}
			err = q.PushBack(int(v))
			if err != nil {
				val = append(val, err.Error())
			}
		case "pop_front":
			v, err := q.PopFront()
			switch err {
			case nil:
				val = append(val, strconv.Itoa(v))
			default:
				val = append(val, err.Error())
			}
		case "pop_back":
			v, err := q.PopBack()
			switch err {
			case nil:
				val = append(val, strconv.Itoa(v))
			default:
				val = append(val, err.Error())
			}
		default:
			continue
		}
	}

	return strings.Join(val, "\n") + "\n"
}

type Queue1 struct {
	Data []int

	First int
	Last  int

	Filled int

	Max int
}

func (q *Queue1) PushFront(x int) error {

	if q.Filled == q.Max {
		return errors.New("error")
	}

	q.Data[q.First] = x

	q.First--
	if q.First < 0 {
		q.First = q.Max - 1
	}

	q.Filled++

	return nil
}

func (q *Queue1) PushBack(x int) error {
	if q.Filled == q.Max {
		return errors.New("error")
	}

	q.Data[q.Last] = x
	q.Last++
	if q.Last == q.Max {
		q.Last = 0
	}

	q.Filled++

	return nil
}

func (q *Queue1) PopFront() (int, error) {
	if q.Filled == 0 {
		return 0, errors.New("error")
	}

	idx := q.First + 1
	if idx == q.Max {
		idx = 0
	}

	v := q.Data[idx]
	q.Data[idx] = 0

	q.First = idx
	q.Filled--

	return v, nil
}

func (q *Queue1) PopBack() (int, error) {
	if q.Filled == 0 {
		return 0, errors.New("error")
	}

	idx := q.Last - 1
	if idx < 0 {
		idx = q.Max - 1
	}

	v := q.Data[idx]
	q.Data[idx] = 0

	q.Last = idx
	q.Filled--

	return v, nil
}
