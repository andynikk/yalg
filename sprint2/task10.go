package sprint2

import (
	"strconv"
	"strings"
)

func Task10(res *[]byte) string {

	lines := strings.Split(string(*res), "\n")

	commands := lines[1:]
	val := []string{}
	q := Queue10{}

	for c := 0; c < len(commands); c++ {
		command := strings.Split(commands[c], " ")
		if command[0] == "" {
			continue
		}

		switch command[0] {
		case "get":
			val = append(val, q.Get())
		case "put":
			z, err := strconv.ParseInt(command[1], 0, 0)
			if err != nil {
				panic(err)
			}
			q.Put(int(z))
		case "size":
			val = append(val, q.SizeQ())
		default:
			continue
		}
	}

	return strings.Join(val, "\n") + "\n"
}

type Queue10 struct {
	Data int
	Q    *Queue10
	Size int
}

func (q *Queue10) Get() string {

	if q.Size == 0 {
		return "error"
	}

	i := strconv.Itoa(q.Data)
	m := q.Size - 1

	if m == 0 {
		q.Data = 0
	} else {
		*q = *q.Q
	}
	q.Size = m

	return i
}

func (q *Queue10) Add(x int) {

	if q.Q == nil {
		q.Q = &Queue10{Data: x}
	} else {
		q.Q.Add(x)
	}
}

func (q *Queue10) Put(x int) string {

	q.Size = q.Size + 1
	if q.Size == 1 {
		q.Data = x
		return ""
	}

	q.Add(x)

	return ""
}

func (q *Queue10) SizeQ() string {
	return strconv.Itoa(q.Size)
}
