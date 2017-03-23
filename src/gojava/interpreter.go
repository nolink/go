package main

import (
	"fmt"
	"gojava/classfile"
	"gojava/instructions"
	"gojava/instructions/base"
	"gojava/rtda"
)

func interpret(methodInfo *classfile.MemberInfo) {
	attrInfo := methodInfo.CodeAttribute()
	maxLocals := attrInfo.MaxLocals()
	maxStack := attrInfo.MaxStack()
	code := attrInfo.Code()
	thread := rtda.NewThread()
	frame := thread.NewFrame(maxLocals, maxStack)
	thread.PushFrame(frame)
	defer catchErr(frame)
	loop(thread, code)
}

func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars: %v\n", frame.LocalVars())
		fmt.Printf("OperandStack: %v\n", frame.OperandStack())
		panic(r)
	}
}

func loop(thread *rtda.Thread, code []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}
	for {
		pc := frame.NextPc()
		thread.SetPc(pc)
		reader.Reset(code, pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPc(reader.Pc())
		fmt.Printf("pc: %2d, inst: %T %v\n", pc, inst, inst)
		inst.Execute(frame)
	}
}
