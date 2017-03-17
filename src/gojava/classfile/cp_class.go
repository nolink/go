package classfile

type ConstantClassInfo struct {
	cp         ConstantPool
	classIndex uint16
}

func (self *ConstantClassInfo) readInfo(reader *ClassReader) {
	self.classIndex = reader.readUint16()
}

func (self *ConstantClassInfo) String() string {
	return self.cp.getUtf8(self.classIndex)
}
