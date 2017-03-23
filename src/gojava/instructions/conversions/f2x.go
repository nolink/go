package conversions

import (
	"gojava/instructions/base"
	"gojava/rtda"
)

type F2D struct{ base.NoOperandsInstruction }
type F2I struct{ base.NoOperandsInstruction }
type F2L struct{ base.NoOperandsInstruction }

func (self *F2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	d := float64(f)
	stack.PushDouble(d)
}

func (self *F2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	i := int32(f)
	stack.PushInt(i)
}

func (self *F2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	d := int64(f)
	stack.PushLong(d)
}
