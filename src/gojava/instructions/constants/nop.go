package constants

import (
	"gojava/instructions/base"
	"gojava/rtda"
)

type NOP struct{ base.NoOperandsInstruction }

func (self *NOP) Execute(frame *rtda.Frame) {}
