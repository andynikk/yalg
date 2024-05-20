package sprint2

//
//import (
//	"strconv"
//	"strings"
//)
//
//const MinInt int = -9223372036854775808
//
//type MaxNode struct {
//	Data int
//	Next *MaxNode
//}
//
//type StackMax struct {
//	Stack   []int
//	M       map[int]int
//	Max     int
//	MaxNode MaxNode
//}
//
//func SetNode(prev MaxNode, v int) MaxNode {
//
//	switch prev.Next {
//	case nil:
//		newNoda := MaxNode{Data: v}
//
//		if v < prev.Data {
//			prev.Next = &newNoda
//			return prev
//		} else {
//			newNoda.Next = &prev
//			return newNoda
//		}
//
//	default:
//
//		if prev.Data >= v && prev.Next.Data <= v {
//			newNoda := MaxNode{Data: v}
//
//			newNoda.Next = prev.Next
//			prev.Next = &newNoda
//
//			return prev
//
//		} else if prev.Data <= v {
//			newNoda := MaxNode{Data: v}
//			newNoda.Next = &prev
//
//			return newNoda
//
//		} else {
//			*prev.Next = SetNode(*prev.Next, v)
//		}
//	}
//
//	return prev
//}
//
//func DelNode(prev MaxNode, v int) MaxNode {
//
//	if prev.Data == v {
//		return *prev.Next
//	}
//
//	*prev.Next = DelNode(*prev.Next, v)
//
//	return prev
//}
//
//func (s *StackMax) Push(x int) {
//
//	s.Stack = append(s.Stack, x)
//	//for i := len(s.Stack) - 2; i >= 0; i-- {
//	//	s.Stack[i+1] = s.Stack[i]
//	//}
//
//	if len(s.Stack) > 1 {
//
//		ss := s.Stack[:len(s.Stack)-1]
//		s.Stack = []int{}
//		s.Stack = append(s.Stack, x)
//		s.Stack = append(s.Stack, ss...)
//
//	}
//	s.Stack[0] = x
//	_, ok := s.M[x]
//	if !ok {
//		s.M[x] = 1
//	} else {
//		s.M[x] = s.M[x] + 1
//	}
//
//	if x > s.Max {
//		s.Max = x
//	}
//
//	//s.MaxNode = SetNode(s.MaxNode, x)
//
//}
//
//func (s *StackMax) Pop() string {
//	lenStack := len(s.Stack)
//	if lenStack == 0 {
//		return "error"
//	}
//
//	z := s.Stack[0]
//
//	//for i := 1; i < len(s.Stack); i++ {
//	//	s.Stack[i-1] = s.Stack[i]
//	//}
//	//s.Stack = s.Stack[:len(s.Stack)-1]
//
//	s.Stack = s.Stack[1:]
//
//	//s.MaxNode = DelNode(s.MaxNode, z)
//	//s.Max = s.MaxNode.Data
//
//	s.M[z] = s.M[z] - 1
//	//if len(s.Stack) == 0 {
//	//	s.Max = MinInt
//	//	return ""
//	//}
//
//	if z == s.Max {
//		//keys := make([]int, 0, len(s.M))
//		//for k := range s.M {
//		//	keys = append(keys, k)
//		//}
//		//sort.Ints(keys)
//		//for i := len(keys) - 1; i >= 0; i-- {
//		//	if s.M[keys[i]] > 0 {
//		//		s.Max = keys[i]
//		//		break
//		//	}
//		//}
//
//		//max1 := MinInt
//		//for k, v := range s.M {
//		//	if v > 0 && k > max1 {
//		//		max1 = k
//		//	}
//		//}
//		//s.Max = max1
//
//	}
//
//	//_ = z
//	//s.Stack = s.Stack[:len(s.Stack)-1]
//	//lenStack--
//	//
//	//switch lenStack {
//	//case 0:
//	//	s.Max = MinInt
//	//case 1:
//	//	s.Max = s.Stack[0]
//	//default:
//	//	if z == s.Max {
//	//		mx := MinInt
//	//		sm := 1
//	//		newLenStack := lenStack / 2
//	//
//	//		if lenStack%2 != 0 {
//	//			mx = s.Stack[lenStack-1]
//	//			newLenStack = (lenStack - 1) / 2
//	//			sm = 2
//	//		}
//	//
//	//		for i := 0; i < newLenStack; i++ {
//	//
//	//			v1 := s.Stack[i]
//	//			v2 := s.Stack[lenStack-sm-i]
//	//			v := max(v1, v2)
//	//
//	//			if v > mx {
//	//				mx = v
//	//			}
//	//
//	//			if z == mx {
//	//				break
//	//			}
//	//		}
//	//
//	//		s.Max = mx
//	//	}
//	//}
//	//
//	return ""
//}
//
//func (s *StackMax) Get_max() string {
//
//	if len(s.Stack) == 0 {
//		return "None"
//	}
//
//	max1 := MinInt
//	for k, v := range s.M {
//		if v > 0 && k > max1 {
//			max1 = k
//		}
//	}
//
//	return strconv.Itoa(max1)
//
//}
//
//func (s *StackMax) Top() string {
//	if len(s.Stack) == 0 {
//		return "error"
//	}
//
//	return strconv.Itoa(s.Stack[0])
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
//	s := StackMax{Max: MinInt, MaxNode: MaxNode{Data: MinInt}, M: map[int]int{}}
//
//	for l := 1; l < len(lines); l++ {
//		line := strings.Split(lines[l], " ")
//
//		switch line[0] {
//		case "get_max":
//			val = val + s.Get_max() + "\n"
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
//	//error
//	//error
//	//error
//	//-5
//	//4
//	//4
//	//None
//
//	return val //strings.Join(val, "\n") + "\n"
//}
