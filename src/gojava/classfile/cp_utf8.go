package classfile

type ConstantUtf8Info struct {
	val string
}

func (self *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	self.val = string(bytes)
}
