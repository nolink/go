package maths

import (
	"gojava/instructions/base"
	"gojava/rtda"
)

type INEG struct{ base.NoOperandsInstruction }
type LNEG struct{ base.NoOperandsInstruction }
type DNEG struct{ base.NoOperandsInstruction }
type FNEG struct{ base.NoOperandsInstruction }

func (self *INEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	stack.PushInt(-v2)
}

func (self *LNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	stack.PushLong(-v2)
}

func (self *DNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	stack.PushDouble(-v2)
}

func (self *FNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	stack.PushFloat(-v2)
}
