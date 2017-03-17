package classfile

type MarkersAttribute struct{}
type DeprecatedAttribute struct{ MarkersAttribute }
type SyntheticAttribute struct{ MarkersAttribute }

func (self *MarkersAttribute) readInfo(reader *ClassReader) {}
