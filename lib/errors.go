package lib

import "fmt"

type ParseError struct {
	msg string
	idx int
}

func (p ParseError) Error() string {
	return fmt.Sprintf("%v at %v", p.msg, p.idx)
}
