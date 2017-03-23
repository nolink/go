package maths

import (
	"gojava/instructions/base"
	"gojava/rtda"
)

type ISUB struct{ base.NoOperandsInstruction }
type LSUB struct{ base.NoOperandsInstruction }
type DSUB struct{ base.NoOperandsInstruction }
type FSUB struct{ base.NoOperandsInstruction }

func (self *ISUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	stack.PushInt(v1 - v2)
}

func (self *LSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	stack.PushLong(v1 - v2)
}

func (self *DSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	stack.PushDouble(v1 - v2)
}

func (self *FSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	stack.PushFloat(v1 - v2)
}
