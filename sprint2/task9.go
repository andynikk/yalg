package sprint2

import (
	"strconv"
	"strings"
)

func Task9(res *[]byte) string {

	lines := strings.Split(string(*res), "\n")

	count, err := strconv.ParseInt(strings.Replace(lines[1], "\r", "", -1), 0, 0)
	if err != nil {
		panic(err)
	}

	commands := lines[1:]
	val := []string{}
	s := Stack9{Max: int(count)}

	for c := 1; c < len(commands); c++ {
		command := strings.Split(commands[c], " ")
		if command[0] == "" {
			continue
		}

		switch command[0] {
		case "peek":
			val = append(val, s.Peek())
		case "pop":
			val = append(val, s.Pop())
		case "push":
			z, err := strconv.ParseInt(command[1], 0, 0)
			if err != nil {
				panic(err)
			}
			v := s.Push(int(z))
			if v == "error" {
				val = append(val, v)
			}
		case "size":
			val = append(val, strconv.Itoa(s.Size()))
		default:
			continue
		}
	}

	//8
	//push -82
	//push -25
	//push -57
	//push -24
	//size
	//push 12
	//push 21
	//push 62
	//push 64
	//push -90
	//size
	//pop
	//peek
	//push -10
	//push 60
	//push 67
	//size

	//4	error	8	64	-82	error	error	8

	//4	error	8	-82	-25	error	error	8

	return strings.Join(val, "\n") + "\n"
}

type Stack9 struct {
	S   []int
	Max int
}

func (s *Stack9) Push(x int) string {

	if len(s.S) == s.Max {
		return "error"
	}

	s.S = append(s.S, x)
	return ""
}

func (s *Stack9) Pop() string {
	if len(s.S) == 0 {
		return "None"
	}

	//i := s.S[len(s.S)-1]
	//s.S = s.S[:len(s.S)-1]

	i := s.S[0]
	s.S = s.S[1:len(s.S)]

	return strconv.Itoa(i)
}

func (s *Stack9) Peek() string {

	if len(s.S) == 0 {
		return "None"
	}

	return strconv.Itoa(s.S[0])

}

func (s *Stack9) Size() int {

	return len(s.S)

}
