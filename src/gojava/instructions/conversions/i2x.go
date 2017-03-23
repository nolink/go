package conversions

import (
	"gojava/instructions/base"
	"gojava/rtda"
)

type I2D struct{ base.NoOperandsInstruction }
type I2F struct{ base.NoOperandsInstruction }
type I2L struct{ base.NoOperandsInstruction }
type I2B struct{ base.NoOperandsInstruction }
type I2C struct{ base.NoOperandsInstruction }
type I2S struct{ base.NoOperandsInstruction }

func (self *I2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	f := stack.PopInt()
	d := float64(f)
	stack.PushDouble(d)
}

func (self *I2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	f := stack.PopInt()
	i := float32(f)
	stack.PushFloat(i)
}

func (self *I2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	f := stack.PopInt()
	d := int64(f)
	stack.PushLong(d)
}

func (self *I2B) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	f := stack.PopInt()
	d := int8(f)
	stack.PushInt(int32(d))
}

func (self *I2C) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	f := stack.PopInt()
	d := uint8(f)
	stack.PushInt(int32(d))
}

func (self *I2S) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	f := stack.PopInt()
	d := int16(f)
	stack.PushInt(int32(d))
}
