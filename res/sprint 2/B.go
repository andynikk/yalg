/*
-- ПРИНЦИП РАБОТЫ --
Для реализации этой задачи я использую один цикл для внесения данных. 
И в этом же цикле идет чтение и удаление данных. Циклом я обхожу всю строку с n(0) по n(len(n)-1). 
Символы, не входящие в список действий («+-*/»), попадают в очередь FIFO. 
Как только попадаю на символ входящий в список действий, то произвожу арифметическую операцию с последними двумя символами из очереди и удаляю их из очереди. 

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ –
Из описания алгоритма следует, что элементы считываются по порядку и записываются в очередь. 
И в порядке попадания попарно считываются. Производится арифметическая операция. 
Итоговое значение меняется после каждой арифметической операции.


-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
Добавление в очередь стоит O(1). Кладем в очередь в одно действие.
Извлечение из очереди стоит так же O(1), т.к. я беру строго снизу, без перебора.
Общая сложность равна O(n+ь), где n количество вставок в очередь, m количество чтений из очереди

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
Очередь может содержать максимум n-1 элементов.


Очередь содержит k элементов, занимает O(k) памяти.
Поэтому и моя очередь будет потреблять O(n) памяти.
*/


//https://contest.yandex.ru/contest/22781/run-report/109000228/

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

func Final2(res *[]byte) string {

	lines := strings.Split(string(*res), "\n")
	commands := lines[0]
	val := []string{}
	
	specSymbol := "+-*/"
	hash := []int{}

	command := strings.Split(commands, " ")

	for i := 0; i < len(command); i++ {
		symbol := command[i]
		if strings.Contains(specSymbol, symbol) {		
			v1 := hash[len(hash)-2]
			v2 := hash[len(hash)-1]

			v := 0
			switch symbol {
			case "+":
				v = v1 + v2
			case "-":
				v = v1 - v2
			case "*":
				v = v1 * v2
			case "/":

				v = floorDiv(v1, v2)
			}

			hash = hash[0 : len(hash)-2]
			hash = append(hash, v)
			continue
		}

		h, _ := strconv.Atoi(symbol)
		hash = append(hash, h)

	}

	val = append(val, strconv.Itoa(hash[len(hash)-1]))

	return strings.Join(val, "\n") + "\n"
}

func floorDiv(x, y int) int {
	q := x / y

	if x^y < 0 && q*y != x {
		return q - 1
	}
	return q

}
