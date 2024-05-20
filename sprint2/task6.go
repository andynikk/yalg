package sprint2

// import (
//
//	"slices"
//	"strconv"
//	"strings"
//
// )
//
// const minInt int = -9223372036854775808
//
//	type StackMax struct {
//		stack []int
//		max   int
//	}
//
// func (s *StackMax) push(x int) {
//
//		//s.stack = append(s.stack, copy([]int{x}, s.stack))
//
//		s.stack = append(s.stack, x)
//		for i := len(s.stack) - 2; i >= 0; i-- {
//			s.stack[i+1] = s.stack[i]
//		}
//		s.stack[0] = x
//
//		if x > s.max {
//			s.max = x
//		}
//	}
//
//	func (s *StackMax) pop() string {
//		if len(s.stack) == 0 {
//			return "error"
//		}
//
//		z := s.stack[0]
//		for i := 1; i < len(s.stack); i++ {
//			s.stack[i-1] = s.stack[i]
//		}
//		s.stack = s.stack[:len(s.stack)-1]
//
//		switch len(s.stack) {
//		case 0:
//			s.max = minInt
//		default:
//			if z == s.max {
//				s.max = slices.Max(s.stack)
//			}
//		}
//
//		return ""
//	}
//
// func (s *StackMax) get_max() string {
//
//	if len(s.stack) == 0 {
//		return "None"
//	}
//	return strconv.Itoa(s.max)
//
// }
func Task6(res *[]byte) string {
	//
	//	lines := strings.Split(string(*res), "\n")
	//	if len(lines) < 2 {
	//		return ""
	//	}
	//
	//	//val := []string{}
	//	s := StackMax{}
	//
	//	val := ""
	//	for l := 1; l < len(lines); l++ {
	//		line := strings.Split(lines[l], " ")
	//
	//		switch line[0] {
	//		case "get_max":
	//			val = val + s.get_max() + "\n"
	//		case "pop":
	//			v := s.pop()
	//			if v == "error" {
	//				val = val + v + "\n"
	//			}
	//		case "push":
	//			z, err := strconv.ParseInt(line[1], 0, 0)
	//			if err != nil {
	//				panic(err)
	//			}
	//			s.push(int(z))
	//		default:
	//			continue
	//		}
	//	}
	//
	//	return val //strings.Join(val, "\n") + "\n"
	return ""
}
