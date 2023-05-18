package stackXy

import "fmt"
import "temp/chess/glob/Xy"

type Stack struct {
	Slice []Xy.Xy
}

func (s *Stack) Push(elem Xy.Xy) {
	s.Slice = append(s.Slice, elem)
}
func (s *Stack) Top() (Xy.Xy, error) {
	if len(s.Slice) < 1 {
		return Xy.Xy{0, 0}, NegativeIndexError{
			functionName: "Stack.Top()",
			index:        len(s.Slice) - 1,
		}
	}
	return s.Slice[len(s.Slice)-1], nil
}
func (s *Stack) Pop() error {
	if len(s.Slice) < 1 {
		return NegativeIndexError{
			functionName: "Stack.Pop()",
			index:        len(s.Slice) - 1,
		}
	}
	s.Slice = s.Slice[:len(s.Slice)-1]
	return nil
}
func (s *Stack) Empty() bool {
	if len(s.Slice) == 0 {
		return true
	}
	return false
}
func (s *Stack) Size() int {
	return len(s.Slice)
}

type NegativeIndexError struct {
	functionName string
	index        int
}

func (s NegativeIndexError) Error() string {
	return fmt.Sprintf("%v: negative index [%v]", s.functionName, s.index)
}
