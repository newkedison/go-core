package core

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestMarshalSimple(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(MarshalSimpleType(byte(0x55)), []byte{0x55})
	assert.Equal(MarshalSimpleType(int8(-1)), []byte{0xFF})
	assert.Equal(MarshalSimpleType(uint8(255)), []byte{0xFF})
	assert.Equal(MarshalSimpleType(int16(-1)), []byte{0xFF, 0xFF})
	assert.Equal(MarshalSimpleType(uint16(65535)), []byte{0xFF, 0xFF})
	assert.Equal(MarshalSimpleType(1), []byte{0x01, 0x00, 0x00, 0x00})
	assert.Equal(MarshalSimpleType(-1), []byte{0xFF, 0xFF, 0xFF, 0xFF})
	assert.Equal(MarshalSimpleType(uint(1)), []byte{0x01, 0x00, 0x00, 0x00})
	assert.Equal(MarshalSimpleType(int32(-1)), []byte{0xFF, 0xFF, 0xFF, 0xFF})
	assert.Equal(MarshalSimpleType(uint32(4294967295)),
		[]byte{0xFF, 0xFF, 0xFF, 0xFF})
	assert.Equal(MarshalSimpleType(int64(-1)),
		[]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF})
	assert.Equal(MarshalSimpleType(uint64(18446744073709551615)),
		[]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF})
	assert.Equal(MarshalSimpleType(float32(1.234)),
		[]byte{0xB6, 0xF3, 0x9D, 0x3F})
	assert.Equal(MarshalSimpleType(float64(1.234)),
		[]byte{0x58, 0x39, 0xB4, 0xC8, 0x76, 0xBE, 0xF3, 0x3F})
	assert.Panics(func() { MarshalSimpleType("I will panic...") })
}

func TestUnmashalSimpleType(t *testing.T) {
	assert := assert.New(t)
	var b byte
	offset := UnmashalSimpleType(&b, []byte{0xAA})
	assert.EqualValues(offset, 1)
	assert.EqualValues(b, 0xAA)
	var i8 int8
	offset = UnmashalSimpleType(&i8, []byte{0xFF})
	assert.EqualValues(offset, 1)
	assert.EqualValues(i8, -1)
	var u8 uint8
	offset = UnmashalSimpleType(&u8, []byte{0xFF})
	assert.EqualValues(offset, 1)
	assert.EqualValues(u8, 0xFF)
	var i16 int16
	offset = UnmashalSimpleType(&i16, []byte{0xFF, 0xFF})
	assert.EqualValues(offset, 2)
	assert.EqualValues(i16, -1)
	var u16 uint16
	offset = UnmashalSimpleType(&u16, []byte{0xFF, 0xFF})
	assert.EqualValues(offset, 2)
	assert.EqualValues(u16, 0xFFFF)
	var i32 int32
	offset = UnmashalSimpleType(&i32, []byte{0xFF, 0xFF, 0xFF, 0xFF})
	assert.EqualValues(offset, 4)
	assert.EqualValues(i32, -1)
	var u32 uint32
	offset = UnmashalSimpleType(&u32, []byte{0xFF, 0xFF, 0xFF, 0xFF})
	assert.EqualValues(offset, 4)
	assert.EqualValues(u32, 0xFFFFFFFF)
	var i int
	offset = UnmashalSimpleType(&i, []byte{0xFF, 0xFF, 0xFF, 0xFF})
	assert.EqualValues(offset, 4)
	assert.EqualValues(i, -1)
	var u uint
	offset = UnmashalSimpleType(&u, []byte{0xFF, 0xFF, 0xFF, 0xFF})
	assert.EqualValues(offset, 4)
	assert.EqualValues(u, 0xFFFFFFFF)
	var i64 int64
	offset = UnmashalSimpleType(&i64,
		[]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF})
	assert.EqualValues(offset, 8)
	assert.EqualValues(i64, -1)
	var u64 uint64
	offset = UnmashalSimpleType(&u64,
		[]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF})
	assert.EqualValues(offset, 8)
	assert.EqualValues(u64, uint64(0xFFFFFFFFFFFFFFFF))
	epsilon32 := math.Nextafter32(1, 2) - 1
	epsilon64 := math.Nextafter(1, 2) - 1
	var f32 float32
	offset = UnmashalSimpleType(&f32, []byte{0xB6, 0xF3, 0x9D, 0x3F})
	assert.EqualValues(offset, 4)
	assert.InEpsilon(f32, 1.234, float64(epsilon32))
	var f64 float64
	offset = UnmashalSimpleType(&f64,
		[]byte{0x58, 0x39, 0xB4, 0xC8, 0x76, 0xBE, 0xF3, 0x3F})
	assert.EqualValues(offset, 8)
	assert.InEpsilon(f64, 1.234, epsilon64)
	assert.Panics(func() { UnmashalSimpleType(1, []byte{}) })
	assert.Panics(func() { UnmashalSimpleType(1.1, []byte{}) })
	assert.Panics(func() { UnmashalSimpleType("", []byte{}) })
}

func TestMarshalString(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(MarshalString("AAA"), []byte{0x03, 0x00, 0x41, 0x41, 0x41})
	assert.Equal(MarshalString("你好"),
		[]byte{0x06, 0x00, 0xe4, 0xbd, 0xa0, 0xe5, 0xa5, 0xbd}) // UTF-8
}

func TestUnmashalString(t *testing.T) {
	assert := assert.New(t)
	var s string
	assert.Panics(func() { UnmarshalString(&s, nil) })
	assert.Panics(func() { UnmarshalString(&s, []byte{0x00}) })
	assert.Panics(func() { UnmarshalString(&s, []byte{0x02, 0x00, 0x00}) })
	n := UnmarshalString(&s, []byte{0x03, 0x00, 0x41, 0x41, 0x41})
	assert.Equal(n, 5)
	assert.Equal(s, "AAA")
	n = UnmarshalString(
		&s, []byte{0x06, 0x00, 0xe4, 0xbd, 0xa0, 0xe5, 0xa5, 0xbd})
	assert.Equal(n, 8)
	assert.Equal(s, "你好")
}

type mashalableObject int

func (v mashalableObject) MarshalBinary() ([]byte, error) {
	if int(v) == 0 {
		return nil, errors.New("0 is error")
	}
	return MarshalSimpleType(int(v)), nil
}

func (v *mashalableObject) UnmarshalBinary(data []byte) error {
	var i int
	UnmashalSimpleType(&i, data)
	*v = mashalableObject(i)
	return nil
}

func TestMarshalObject(t *testing.T) {
	assert := assert.New(t)
	data, err := MarshalObject(mashalableObject(1))
	assert.Equal(data, []byte{0x04, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00})
	assert.Nil(err)
	data, err = MarshalObject(mashalableObject(0))
	assert.Nil(data)
	assert.Error(err)
}

func TestUnmashalObject(t *testing.T) {
	assert := assert.New(t)
	var o mashalableObject
	assert.Equal(UnmarshalObject(
		&o, []byte{0x04, 0x00, 0x00, 0x00, 0x0A, 0x00, 0x00, 0x00}), 8)
	assert.EqualValues(int(o), 10)
}

type myByteOrder struct{}

func (myByteOrder) Uint16(b []byte) uint16 {
	return 0
}

func (myByteOrder) PutUint16(b []byte, v uint16) {
}

func (myByteOrder) Uint32(b []byte) uint32 {
	return 42
}

func (myByteOrder) PutUint32(b []byte, v uint32) {
	b[0] = 0xAA
	b[1] = 0xAA
	b[2] = 0xAA
}

func (myByteOrder) Uint64(b []byte) uint64 {
	return 0
}

func (myByteOrder) PutUint64(b []byte, v uint64) {
}

func (myByteOrder) String() string { return "myByteOrder" }

func TestSetByteOrder(t *testing.T) {
	assert := assert.New(t)
	SetByteOrder(myByteOrder{})
	assert.Equal(MarshalSimpleType(uint32(0)), []byte{0xAA, 0xAA, 0xAA, 0x00})
	var v uint32
	assert.EqualValues(UnmashalSimpleType(&v, []byte{0xAA, 0xFF, 0xAA, 0xFF}), 4)
	assert.EqualValues(v, 42)
	data, err := MarshalObject(mashalableObject(1))
	assert.Equal(data, []byte{0xAA, 0xAA, 0xAA, 0x00, 0xAA, 0xAA, 0xAA, 0x00})
	assert.Nil(err)
}

func TestSetErrorWhenNotEnoughDataErrorPanic(t *testing.T) {
	assert := assert.New(t)
	var err error = errors.New("hello")
	defer func() {
		r := recover()
		assert.Nil(r)
		assert.Error(err, "aaabbb")
	}()
	defer SetErrorWhenNotEnoughDataErrorPanic("aaa", &err)()
	panic(NotEnoughDataError("bbb"))
}

func TestSetErrorWhenNotEnoughDataErrorPanic2(t *testing.T) {
	assert := assert.New(t)
	var err error = errors.New("hello")
	defer func() {
		r := recover()
		assert.NotNil(r)
		assert.IsType(r, 0)
		assert.Error(err, "hello")
	}()
	defer SetErrorWhenNotEnoughDataErrorPanic("aaa", &err)()
	panic(42)
}
