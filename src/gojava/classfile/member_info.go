package classfile

type MemberInfo struct {
	cp          ConstantPool
	accessFlags uint16
	nameIndex   uint16
	descIndex   uint16
	attributes  []AttributeInfo
}

func (self *MemberInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}

func (self *MemberInfo) AccessFlags() uint16 {
	return self.accessFlags
}

func (self *MemberInfo) Descriptor() string {
	return self.cp.getUtf8(self.descIndex)
}

func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memCnt := reader.readUint16()
	members := make([]*MemberInfo, memCnt)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:          cp,
		accessFlags: reader.readUint16(),
		nameIndex:   reader.readUint16(),
		descIndex:   reader.readUint16(),
		attributes:  readAttributes(reader, cp),
	}
}
