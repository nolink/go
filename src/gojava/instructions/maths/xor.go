package maths

import (
	"gojava/instructions/base"
	"gojava/rtda"
)

type IXOR struct{ base.NoOperandsInstruction }
type LXOR struct{ base.NoOperandsInstruction }

func (self *IXOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	stack.PushInt(v1 ^ v2)
}

func (self *LXOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	stack.PushLong(v1 ^ v2)
}
