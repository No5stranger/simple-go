package interf

import "fmt"

type I interface {
	Write()
}

type Me struct {
	Thought string
}

func (m Me) Write() {
	fmt.Print(m.Thought)
}

func NewMe(tho string) *Me {
	return &Me{Thought: tho}
}

func IsInterface(i interface{}) bool {
	_, ok := i.(I)
	return ok
}
