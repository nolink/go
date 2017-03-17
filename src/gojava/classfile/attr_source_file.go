package classfile

type SourceFileAttribute struct {
	cp           ConstantPool
	srcFileIndex uint16
}

func (self *SourceFileAttribute) readInfo(reader *ClassReader) {
	self.srcFileIndex = reader.readUint16()
}

func (self *SourceFileAttribute) FileName() string {
	return self.cp.getUtf8(self.srcFileIndex)
}
