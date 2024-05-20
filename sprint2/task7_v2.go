package sprint2

//
//import (
//	"strconv"
//	"strings"
//)
//
//type StackMax struct {
//	Stack       []int
//	StackMaxVal []int
//}
//
//func (s *StackMax) Remove(idx int) {
//	tmpStackMaxValUp := s.StackMaxVal[:idx]
//	tmpStackMaxValDown := s.StackMaxVal[idx:]
//
//	s.StackMaxVal = []int{}
//	s.StackMaxVal = append(s.StackMaxVal, tmpStackMaxValUp...)
//	s.StackMaxVal = append(s.StackMaxVal, tmpStackMaxValDown...)
//}
//
//func (s *StackMax) RemoveMiddle(x int) {
//
//	switch x-s.StackMaxVal[0] <= x-s.StackMaxVal[len(s.StackMaxVal)-1] {
//	case false:
//		for i := len(s.StackMaxVal) - 1; i >= 0; i-- {
//			if s.StackMaxVal[i] <= x {
//				s.Remove(i)
//
//				break
//			}
//		}
//	default:
//		for i := 0; i < len(s.StackMaxVal); i++ {
//			if s.StackMaxVal[i] <= x {
//				s.Remove(i)
//
//				break
//			}
//		}
//	}
//
//}
//
//func (s *StackMax) Insert(idx int, val int) {
//	tmpStackMaxValUp := s.StackMaxVal[:idx]
//	tmpStackMaxValDown := s.StackMaxVal[idx:]
//
//	s.StackMaxVal = []int{}
//	s.StackMaxVal = append(s.StackMaxVal, tmpStackMaxValUp...)
//	s.StackMaxVal = append(s.StackMaxVal, val)
//	s.StackMaxVal = append(s.StackMaxVal, tmpStackMaxValDown...)
//}
//
//func (s *StackMax) InsertMax(x int) {
//
//	if x <= s.StackMaxVal[len(s.StackMaxVal)-1] {
//		s.StackMaxVal = append(s.StackMaxVal, x)
//
//		return
//	}
//
//	if x >= s.StackMaxVal[0] {
//		tmpStackMax := s.StackMaxVal[:]
//
//		s.StackMaxVal = []int{x}
//		s.StackMaxVal = append(s.StackMaxVal, tmpStackMax...)
//
//		return
//	}
//
//	switch x-s.StackMaxVal[0] <= x-s.StackMaxVal[len(s.StackMaxVal)-1] {
//	case true:
//		for i := len(s.StackMaxVal) - 1; i >= 0; i-- {
//			if s.StackMaxVal[i] <= x {
//				s.Insert(i, x)
//
//				break
//			}
//		}
//	default:
//		for i := 0; i < len(s.StackMaxVal); i++ {
//			if s.StackMaxVal[i] <= x {
//				s.Insert(i, x)
//
//				break
//			}
//		}
//	}
//
//}
//
//func (s *StackMax) Push(x int) {
//
//	//switch len(s.Stack) {
//	//case 0:
//	//	s.Stack = append(s.Stack, x)
//	//	s.StackMaxVal = append(s.StackMaxVal, x)
//	//default:
//	//tmpStack := s.Stack[:]
//
//	//s.Stack = []int{x}
//	//s.Stack = append(s.Stack, tmpStack...)
//
//	s.Stack = append(s.Stack, x)
//	if len(s.StackMaxVal) == 0 {
//		s.StackMaxVal = append(s.StackMaxVal, x)
//	} else {
//		s.InsertMax(x)
//	}
//
//	//}
//
//}
//
//func (s *StackMax) Pop() string {
//	lenStack := len(s.Stack)
//	if lenStack == 0 {
//		return "error"
//	}
//
//	z := s.Stack[len(s.Stack)-1]
//	s.Stack = s.Stack[:len(s.Stack)-1]
//
//	if len(s.Stack) == 0 {
//		s.StackMaxVal = []int{}
//	} else if z == s.StackMaxVal[0] {
//		s.StackMaxVal = s.StackMaxVal[1:]
//	} else if z == s.StackMaxVal[len(s.StackMaxVal)-1] {
//		s.StackMaxVal = s.StackMaxVal[:len(s.StackMaxVal)-1]
//	} else {
//		s.RemoveMiddle(z)
//	}
//
//	return ""
//}
//
//func (s *StackMax) GetMax() string {
//
//	if len(s.Stack) == 0 {
//		return "None"
//	}
//
//	return strconv.Itoa(s.StackMaxVal[0])
//
//}
//
//func (s *StackMax) Top() string {
//	if len(s.Stack) == 0 {
//		return "error"
//	}
//
//	return strconv.Itoa(s.Stack[len(s.Stack)-1])
//}
//
//func Task7(res *[]byte) string {
//
//	lines := strings.Split(string(*res), "\n")
//	if len(lines) < 2 {
//		return ""
//	}
//
//	val := ""
//	s := StackMax{}
//
//	for l := 1; l < len(lines); l++ {
//		line := strings.Split(lines[l], " ")
//
//		switch line[0] {
//		case "get_max":
//			val = val + s.GetMax() + "\n"
//		case "pop":
//			v := s.Pop()
//			if v == "error" {
//				val = val + v + "\n"
//			}
//		case "push":
//			z, err := strconv.ParseInt(line[1], 0, 0)
//			if err != nil {
//				panic(err)
//			}
//			s.Push(int(z))
//		case "top":
//			val = val + s.Top() + "\n"
//		default:
//			continue
//		}
//	}
//
//	return val
//}
