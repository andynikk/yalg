/*
https://contest.yandex.ru/contest/25597/run-report/114061269/

-- ПРИНЦИП РАБОТЫ --
	В функции sameAmounts я посчитал сумму (sum) всех входящих элементов. Если сумма нечетная, то сразу возвращаю "False". Далее я создал массив dp размером sum/2+1 для хранения результатов вычислений.
	Функция dividedTwo принимает массив dp, входной массив inputData, а также n и m - размеры входного массива и массива dp соответственно. Функция стремится определить, можно ли разделить множество элементов inputData на две равные части так, чтобы сумма элементов в каждой части была равна m.

	Двойной цикл перебирает все элементы массива inputData и рассматривает возможность добавления каждого элемента для достижения целевой суммы m.
	Внутренний цикл выполняется в обратном порядке от m до значения текущего элемента inputData[i].
	В каждой итерации обновляется значение dp[j], указывающее, можно ли получить сумму j с использованием элементов из inputData.
	Значение dp[j] становится истинным (true), если либо уже можно получить сумму j, либо можно получить сумму j - inputData[i] и добавить к ней элемент inputData[i].
	Таким образом, алгоритм обновляет информацию о возможности получить каждую сумму от 0 до m.

	Если последий элемент массива dp равен true, это значит, что можно разделить множество элементов inputData на две равные части так, чтобы сумма элементов в каждой части была равна m

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
	Перед началом работы алгоритма все элементы массива dp устанавливаются в false, за исключением dp[0], которое устанавливается в true. Это соответствует случаю, когда сумма равна 0, и разделение на две равные части возможно.
	На определенном шаге алгоритма значение dp[j] становится true. Это означает, что можно получить сумму j с использованием элементов из inputData.
   	На каждом шаге алгоритма я рассматриваю возможность добавления очередного элемента inputData[i] к текущей сумме j.
	Если dp[j-inputData[i]] равно true, то это означает, что можно получить сумму j - inputData[i] с использованием элементов из inputData.
	Таким образом, установка значения dp[j] = dp[j] || dp[j-inputData[i]] позволяет обновить информацию о возможности получить сумму j.
	По завершении работы алгоритма значение dp[sum/2] указывает на возможность разделения множества на две равные части с одинаковыми суммами.

	Таким образом, код работает корректно благодаря использованию динамического программирования и правильному обновлению значений в массиве dp.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
	n - это длина входного массива
	m - максимальная сумма, которую я рассматриваю для разделения массива

	Итоговая сложность кода составляет O(n * m).

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
	n - это длина входного массива
	m - максимальная сумма, которую я рассматриваю для разделения массива

	Пространственная сложность этого кода составляет O(m). Я используем массив dp размером m+1, чтобы хранить промежуточные результаты. И сам, входящий массив О(n).

	Итоговая сложность кода составляет O(m + n).

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

	r := sameAmounts(&res)

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

func dividedTwo(dp []bool, inputData []int, n, m int) bool {

	for i := 0; i < n; i++ {
		for j := m; j >= inputData[i]; j-- {
			dp[j] = dp[j] || dp[j-inputData[i]]
		}
	}

	return dp[m]

}

func sameAmounts(res *[]byte) string {

	lines := strings.Split(string(*res), "\n")

	n, _ := strconv.Atoi(lines[0])
	inputs := strings.Split(lines[1], " ")
	intInputs := make([]int, len(inputs))

	sum := 0
	for i := range inputs {
		intInputs[i], _ = strconv.Atoi(inputs[i])
		sum += intInputs[i]
	}

	if sum%2 != 0 {
		return "False"
	}

	dp := make([]bool, sum/2+1)
	dp[0] = true

	result := strconv.FormatBool(dividedTwo(dp, intInputs, n, sum/2))
	result = strings.ToUpper(result[:1]) + result[1:]

	return result

}