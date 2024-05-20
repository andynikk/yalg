/*
https://contest.yandex.ru/contest/24414/run-report/111840901/

-- ПРИНЦИП РАБОТЫ --
	При обходе строк документов, я раскладываю документ на массив строк, в котором каждый элемент это слово.
Пробегаюсь по строке и добавляю в хэш-таблицу docs. Где ключ — это слово. Значение — это двухмерный массив.
Где первый элемент значения — это индекс массива слов документа. Второй это количество повторов этого слова в документе.
При прохождении строки запросов я ищу в слово в хэш-таблице docs. Если нахожу слово, то прохожу по значению (массив - индекс, количество повторений)).
Кладем в массив interimResult. Где индекс — это номер документа по порядку, значение - количество повторов.
Полученный результат перекалываем в массив resultSearchString. Где индекс ничего не значит, значение массив двухмерный массив. Где первое значение - индекс документа, второе значение - сумма повторов. Максимум пять раз проходим массив resultSearchString в поиске максимального значения. После прохода удаляем найденный элемент из массива resultSearchString.
Получаем результат по строке запроса. Склеиваем его с предыдущим результатом.


-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
Все слова документа находятся, количество повторяй считается. Никаких исключений нет. Слова не теряются.
Максимум проходит циклом 5 раз по результату, что гарантирует выполнения условия, выводить не больше пяти результатов.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
n - количество строк документов
w - количество слов в строке документа
m - количество строк поиска
g - количество слов в строке поиска

Обход строк документов: O(nw)
Обход слов в строках документов: O(w)
Обход строк поиска: O(mgn)
Обход слов строк поиска: O(gn)
Обход документов для каждого слова поиска: O(n)
Обход промежуточного результат: O(n)
Обход массива с готовыми данными для формирования финального текстового результата : O(5n) или O(n)
Итоговая сложность: O(nw + mgn)

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
n - количество строк документов
w - количество слов в строке документа
m - количество строк поиска
g - количество слов в строке поиска

Хэш-таблица docs, где хранятся слова документов и количество повторов  O(nw)
Хэш-таблица processedSearchWords, где хранятся слова поиска (защита от повторных поисков) O(mg)
Слайс interimResult, для хранения промежуточных данных O(n)
Слайс resultSearchString, для хранения данных для формирования итоговых данных O(n)
Строка txtFinalValue, финальный результат O(n)
Итоговая сложность: O(nw + 2n + m + mg) ~ O(nw+mg)

*/

package main

import (
	"os"
	"slices"
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
	countDocs, _ := strconv.Atoi(lines[0])

	//docs - словарь документов. key - хеш слова, value - [индекс документа, кол-во повторов]
	docs := map[string][][2]int{}

	txtFinalValue := ""

	i := 1
	for {
		c := lines[i]
		if c == "" {
			continue
		}

		_, err := strconv.Atoi(lines[i])
		if err == nil {
			break
		}
		documentWords := strings.Fields(c)

		//храним повторы в строке документа.
		uniqueWords := map[string]int{}
		for j := 0; j < len(documentWords); j++ {
			//h := hashFinal1(arr[j])
			h := documentWords[j]

			//увеличиваем кол-во повторов по слову
			uniqueWords[h]++

			if uniqueWords[h] > 1 {
				//получаем индекс предыдущего документа и слова в словаре docs
				idx := slices.Index(docs[h], [2]int{i - 1, uniqueWords[h] - 1})
				//Увеличиваем значение слов-повторов
				docs[h][idx][1] = uniqueWords[h]
				continue
			}

			//если у слова нет повтора, то просто апендим массив [индекс документа, кол-во повторов = 1]
			docs[h] = append(docs[h], [2]int{i - 1, 1})
		}
		i++
	}
	i++

	counterSearchString := 0
	for o := i; o < len(lines); o++ {
		searchString := lines[o]
		if searchString == "" {
			continue
		}

		//z отрезает повторные слова в строке поиска
		processedSearchWords := map[string]struct{}{}
		searchWords := strings.Fields(searchString)

		//создаем промежуточный результат, где [индекс документа, сумма повторов]
		interimResult := make([]int, countDocs)
		for j := 0; j < len(searchWords); j++ {
			//h := hashFinal1(arr[j])
			h := searchWords[j]

			//если слово из строки поиска уже было, то пропускаем
			if _, ok := processedSearchWords[h]; ok {
				continue
			}

			processedSearchWords[h] = struct{}{}

			//Из мапы всех документов docs получаем массив массивов по слову поиска и рапределяем.
			//В массиве promResult индекс - индекс документа, значение - сумма повторов по слову
			for p := 0; p < len(docs[h]); p++ {
				idx := docs[h][p][0]
				interimResult[idx] += docs[h][p][1]
			}
		}

		//создаем массив где индекс ничего не значит, и кладем в него массив [2]int где, первое значение - индекс документа, второе значение - сумма повторов
		resultSearchString := make([][2]int, 0)
		for p := 0; p < len(interimResult); p++ {
			if interimResult[p] == 0 {
				continue
			}
			//resultSearchString = append(resultSearchString, [2]int{p, interimResult[p]})
			resultSearchString = append(resultSearchString, [2]int{interimResult[p], p})
		}

		//Выбираем топ-5 найденных слов
		arrayValuesSearchString := []string{} //4 5 1 2 3
		lenRresultSearchString := len(resultSearchString)
		for p := 0; p < min(lenRresultSearchString, 5); p++ {
			maxValue := [2]int{}
			maxValue, resultSearchString = findMaxAndRemove(resultSearchString)
			arrayValuesSearchString = append(arrayValuesSearchString, strconv.Itoa(maxValue[1]+1))
		}
		txtFinalValue += strings.Join(arrayValuesSearchString, " ") + "\n"
		counterSearchString++
	}

	return txtFinalValue
}

func findMaxAndRemove(arr [][2]int) (maxValue [2]int, newArr [][2]int) {
	low := 0
	high := len(arr) - 1

	delElement := -1
	maxValue = [2]int{0, 0}
	for low <= high {

		if arr[low][0] > maxValue[0] || (arr[low][0] == maxValue[0] && arr[low][1] < maxValue[1]) {
			maxValue = arr[low]
			delElement = low
		}
		if arr[high][0] > maxValue[0] || (arr[high][0] == maxValue[0] && arr[high][1] < maxValue[1]) {
			maxValue = arr[high]
			delElement = high
		}

		low++
		high--
	}

	switch delElement {
	case -1:
		arr = nil
	default:
		arr[delElement], arr[len(arr)-1] = arr[len(arr)-1], arr[delElement]
		arr = arr[:len(arr)-1]
	}
	return maxValue, arr
}
