package sprint1

import (
	"slices"
	"strconv"
	"strings"
)

func Task10(res *[]byte) ([]string, string) {

	lines := strings.Split(string(*res), "\n")
	k, err := strconv.ParseInt(strings.Replace(lines[0], "\r", "", -1), 0, 0)
	if err != nil {
		panic(err)
	}

	intK := int(k)
	factors := []int{}

	for i := 2; i*i <= intK; i++ {
		for intK%i == 0 {
			factors = append(factors, i)
			intK /= i
		}
	}

	if intK > 1 {
		factors = append(factors, intK)
	}
	slices.Sort(factors)

	strFactors := []string{}
	for _, v := range factors {
		strFactors = append(strFactors, strconv.Itoa(v))
	}

	return strFactors, " "

}
