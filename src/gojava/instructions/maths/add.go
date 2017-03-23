package maths

import (
	"gojava/instructions/base"
	"gojava/rtda"
)

type IADD struct{ base.NoOperandsInstruction }
type LADD struct{ base.NoOperandsInstruction }
type DADD struct{ base.NoOperandsInstruction }
type FADD struct{ base.NoOperandsInstruction }

func (self *IADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	stack.PushInt(v1 + v2)
}

func (self *LADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	stack.PushLong(v1 + v2)
}

func (self *DADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	stack.PushDouble(v1 + v2)
}

func (self *FADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	stack.PushFloat(v1 + v2)
}
