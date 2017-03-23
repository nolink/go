package rtda

import "math"

type LocalVars []Slot

func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}
	return nil
}

func (self LocalVars) SetInt(index uint, val int32) {
	self[index].num = val
}

func (self LocalVars) GetInt(index uint) int32 {
	return self[index].num
}

func (self LocalVars) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	self[index].num = int32(bits)
}

func (self LocalVars) GetFloat(index uint) float32 {
	return math.Float32frombits(uint32(self[index].num))
}

func (self LocalVars) SetLong(index uint, val int64) {
	self[index].num = int32(val)
	self[index+1].num = int32(val >> 32)
}

func (self LocalVars) GetLong(index uint) int64 {
	low := uint32(self[index].num)
	high := uint32(self[index+1].num)
	return int64(high)<<32 | int64(low)
}

func (self LocalVars) SetDouble(index uint, val float64) {
	long_val := math.Float64bits(val)
	self.SetLong(index, int64(long_val))
}

func (self LocalVars) GetDouble(index uint) float64 {
	long_val := self.GetLong(index)
	return math.Float64frombits(uint64(long_val))
}

func (self LocalVars) SetRef(index uint, val *Object) {
	self[index].ref = val
}

func (self LocalVars) GetRef(index uint) *Object {
	return self[index].ref
}
