package classfile

type ConstantValueAttribute struct {
	constValueIndex uint16
}

func (self *ConstantValueAttribute) readInfo(reader *ClassReader) {
	self.constValueIndex = reader.readUint16()
}

func (self *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return self.constValueIndex
}
