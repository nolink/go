package classfile

type ConstantPool []ConstantInfo

type ConstantInfo interface {
	readInfo(reader *ClassReader)
}

func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUint16())
	cp := make([]ConstantInfo, cpCount)
	for i := 1; i < cpCount; i++ {
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}
	return cp
}

func (self ConstantPool) getConstantInfo(idx uint16) ConstantInfo {
	if cpInfo := self[idx]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index!")
}

func (self ConstantPool) getClassName(index uint16) string {
	cpInfo := self.getConstantInfo(index).(*ConstantClassInfo)
	return self.getUtf8(cpInfo.classIndex)
}

func (self ConstantPool) getNameAndType(index uint16) (string, string) {
	cpInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := self.getUtf8(cpInfo.nameIndex)
	_type := self.getUtf8(cpInfo.descIndex)
	return name, _type
}

func (self ConstantPool) getUtf8(index uint16) string {
	cpInfo := self.getConstantInfo(index).(*ConstantUtf8Info)
	return cpInfo.val
}

func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	tag := reader.readUint8()
	c := newConstantInfo(tag, cp)
	c.readInfo(reader)
	return c
}

func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	switch tag {
	case CONSTANT_Integer:
		return &ConstantIntegerInfo{}
	case CONSTANT_Float:
		return &ConstantFloatInfo{}
	case CONSTANT_Long:
		return &ConstantLongInfo{}
	case CONSTANT_Double:
		return &ConstantDoubleInfo{}
	case CONSTANT_Utf8:
		return &ConstantUtf8Info{}
	case CONSTANT_String:
		return &ConstantStringInfo{}
	case CONSTANT_Class:
		return &ConstantClassInfo{}
	case CONSTANT_Fieldref:
		return &ConstantFieldrefInfo{}
	case CONSTANT_Methodref:
		return &ConstantMemberrefInfo{}
	case CONSTANT_InterfaceMethodref:
		return &COnstantInterfaceMethodrefInfo{}
	case CONSTANT_NameAndType:
		return &ConstantNameAndTypeInfo{}
	case CONSTANT_MethodType:
		return &ConstantMethodTypeInfo{}
	case CONSTANT_MethodHandle:
		return &ConstantMethodHandelInfo{}
	case CONSTANT_InvokeDynamic:
		return &ConstantInvokeDynamicInfo{}
	default:
		panic("java.lang.ClassFormatError: constant pool tag!")
	}
}
