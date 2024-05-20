package sprint4

import (
	"slices"
	"strconv"
	"strings"
)

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

		//Выбираем топ-5 найденных слов, которые встречаются в документах
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

	delElement := -1
	maxValue = [2]int{0, 0}
	for i := 0; i < len(arr); i++ {

		if arr[i][0] > maxValue[0] || (arr[i][0] == maxValue[0] && arr[i][1] < maxValue[1]) {
			maxValue = arr[i]
			delElement = i
		}

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
