package rtda

type Frame struct {
	lower     *Frame
	localVars LocalVars
	opStack   *OperandStack
	thread    *Thread
	nextPc    int
}

func newFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
	return &Frame{
		thread:    thread,
		localVars: newLocalVars(maxLocals),
		opStack:   newOperandStack(maxStack),
	}
}

func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}

func (self *Frame) OperandStack() *OperandStack {
	return self.opStack
}

func (self *Frame) Thread() *Thread {
	return self.thread
}

func (self *Frame) SetNextPc(val int) {
	self.nextPc = val
}

func (self *Frame) NextPc() int {
	return self.nextPc
}
