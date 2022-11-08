// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flatbuffers_test

import (
	"testing"

	. "github.com/frankban/quicktest"
	flatbuffers "github.com/google/flatbuffers/go"
	fb "github.com/golibraries/encoding/flatbuffers"
)

type Access struct {
	_tab flatbuffers.Table
}

func GetRootAsAccess(buf []byte, offset flatbuffers.UOffsetT) *Access {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Access{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsAccess(buf []byte, offset flatbuffers.UOffsetT) *Access {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Access{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Access) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Access) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Access) Type() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Access) MutateType(n int64) bool {
	return rcv._tab.MutateInt64Slot(4, n)
}

func (rcv *Access) Account(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *Access) AccountLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Access) AccountBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Access) MutateAccount(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func (rcv *Access) Enable() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *Access) MutateEnable(n bool) bool {
	return rcv._tab.MutateBoolSlot(8, n)
}

func AccessStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func AccessAddType(builder *flatbuffers.Builder, type_ int64) {
	builder.PrependInt64Slot(0, type_, 0)
}
func AccessAddAccount(builder *flatbuffers.Builder, account flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(account), 0)
}
func AccessStartAccountVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func AccessAddEnable(builder *flatbuffers.Builder, enable bool) {
	builder.PrependBoolSlot(2, enable, false)
}
func AccessEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}


func TestFlatbuffers(t *testing.T) {
	c := New(t)
	c.Run("TestFlatbuffers", func(c *C) {
		builder := flatbuffers.NewBuilder(0)
		AccessStart(builder)
		AccessAddType(builder, 1)
		// cannot nested vector
		// AccessAddAccount(builder, builder.CreateByteVector([]byte{1, 2, 3}))
		AccessAddEnable(builder, true)
		builder.Finish(AccessEnd(builder))
		data, err := fb.Marshal(builder)
		c.Assert(err, IsNil)
		var access Access
		c.Assert(fb.Unmarshal(data, &access), IsNil)
		c.Assert(access.Type(), Equals, int64(1))
		c.Assert(access.Enable(), Equals, true)
	})
}