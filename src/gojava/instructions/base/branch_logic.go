package base

import "gojava/rtda"

func Branch(frame *rtda.Frame, offset int) {
	pc := frame.Thread().Pc()
	nextPc := pc + offset
	frame.SetNextPc(nextPc)
}
