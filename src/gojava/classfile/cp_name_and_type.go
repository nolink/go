package classfile

type ConstantNameAndTypeInfo struct {
	nameIndex uint16
	descIndex uint16
}

func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
	self.descIndex = reader.readUint16()
}
