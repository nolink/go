package classfile

type ConstantMethodTypeInfo struct {
	descIndex uint16
}

func (self *ConstantMethodTypeInfo) readInfo(reader *ClassReader) {
	self.descIndex = reader.readUint16()
}

type ConstantMethodHandelInfo struct {
	refKind  uint8
	refIndex uint16
}

func (self *ConstantMethodHandelInfo) readInfo(reader *ClassReader) {
	self.refKind = reader.readUint8()
	self.refIndex = reader.readUint16()
}

type ConstantInvokeDynamicInfo struct {
	bootMethodAttrIndex uint16
	nameAndTypeIndex    uint16
}

func (self *ConstantInvokeDynamicInfo) readInfo(reader *ClassReader) {
	self.bootMethodAttrIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}
