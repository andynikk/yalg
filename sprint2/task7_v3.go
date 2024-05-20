package sprint2

import (
	"strconv"
	"strings"
)

const MinInt int = -9223372036854775808

func Task7(res *[]byte) string {

	lines := strings.Split(string(*res), "\n")
	if len(lines) < 2 {
		return ""
	}

	val := []string{}
	s := StackMax{Max: MinInt, M: map[int]int{}}

	for l := 1; l < len(lines); l++ {
		line := strings.Split(lines[l], " ")

		switch line[0] {
		case "get_max":
			val = append(val, s.GetMax())
		case "pop":
			v := s.Pop()
			if v == "error" {
				val = append(val, s.Pop())
			}
		case "push":
			z, err := strconv.ParseInt(line[1], 0, 0)
			if err != nil {
				panic(err)
			}
			s.Push(int(z))
		case "top":
			val = append(val, s.Top())
		default:
			continue
		}
	}
	//None
	//error
	//None
	//None
	//None
	//None
	//1
	//1
	//None
	//None
	//None
	//-3
	//-3
	//9
	//9
	//0
	//4
	//4
	//4
	//4
	//4
	//4
	//4

	return strings.Join(val, "\n") + "\n"
}

type StackMax struct {
	Stack []int
	M     map[int]int
	Max   int
}

func (s *StackMax) Push(x int) {

	s.Stack = append(s.Stack, x)

	if x > s.Max {
		s.Max = x
	}

	v, ok := s.M[x]
	if ok {
		s.M[x] = v + 1
	} else {
		s.M[x] = 1
	}
}

func (s *StackMax) Pop() string {
	if len(s.Stack) == 0 {
		return "error"
	}

	i := s.Stack[len(s.Stack)-1]
	s.Stack = s.Stack[:len(s.Stack)-1]

	switch len(s.Stack) {
	case 0:
		s.Max = MinInt
		s.M[i] = s.M[i] - 1
	default:

		s.M[i] = s.M[i] - 1

		if i == s.Max && s.M[i] == 0 {
			for j := i - 1; true; j-- {
				if v, ok := s.M[j]; ok && v > 0 {
					//s.Max = slices.Max(s.Stack)
					s.Max = j
					break
				}
			}
		}
	}

	return ""
}

func (s *StackMax) GetMax() string {
	if len(s.Stack) == 0 {
		return "None"
	}

	return strconv.Itoa(s.Max)
}

func (s *StackMax) Top() string {
	if len(s.Stack) == 0 {
		return "error"
	}

	return strconv.Itoa(s.Stack[len(s.Stack)-1])
}
