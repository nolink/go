package classfile

type ConstantStringInfo struct {
	cp       ConstantPool
	strIndex uint16
}

func (self *ConstantStringInfo) readInfo(reader *ClassReader) {
	self.strIndex = reader.readUint16()
}

func (self *ConstantStringInfo) String() string {
	return self.cp.getUtf8(self.strIndex)
}
