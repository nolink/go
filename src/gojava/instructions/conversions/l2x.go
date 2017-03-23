package conversions

import (
	"gojava/instructions/base"
	"gojava/rtda"
)

type L2D struct{ base.NoOperandsInstruction }
type L2F struct{ base.NoOperandsInstruction }
type L2I struct{ base.NoOperandsInstruction }

func (self *L2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	f := stack.PopLong()
	d := float64(f)
	stack.PushDouble(d)
}

func (self *L2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	f := stack.PopLong()
	i := float32(f)
	stack.PushFloat(i)
}

func (self *L2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	f := stack.PopLong()
	d := int32(f)
	stack.PushInt(d)
}
