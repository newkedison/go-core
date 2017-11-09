package core

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestFloat32Epsilon(t *testing.T) {
	assert.True(t, Float32Epsilon() < float32(1e-6))
}

func TestFloat64Epsilon(t *testing.T) {
	assert.True(t, Float64Epsilon() < float64(1e-15))
}

func TestNewNumber(t *testing.T) {
	assert := assert.New(t)
	n := NewNumber(byte(3))
	assert.EqualValues(n, 3)
	n = NewNumber(1.23)
	assert.EqualValues(n, 1.23)
	assert.Panics(func() { NewNumber("string is not allow") })
	assert.Panics(func() { NewNumber(MaxIntNumber + 1) })
	assert.Panics(func() { NewNumber(MinIntNumber - 1) })
	assert.Panics(func() { NewNumber(uint64(MaxIntNumber + 1)) })
}

func TestNumberToByte(t *testing.T) {
	assert := assert.New(t)
	assert.EqualValues(NewNumber(0).ToByte(), 0)
	assert.EqualValues(NewNumber(128).ToByte(), 128)
	assert.EqualValues(NewNumber(255).ToByte(), 255)
	assert.EqualValues(NewNumber(256).ToByte(), 0)
	assert.EqualValues(NewNumber(256.0).ToByte(), 0)
	assert.EqualValues(NewNumber(65536).ToByte(), 0)
	assert.EqualValues(NewNumber(math.MaxInt8).ToByte(), 0x7F)
	assert.EqualValues(NewNumber(math.MinInt8).ToByte(), 0x80)
	assert.EqualValues(NewNumber(math.MaxUint8).ToByte(), 0xFF)
	assert.EqualValues(NewNumber(math.MaxInt16).ToByte(), 0xFF)
	assert.EqualValues(NewNumber(math.MinInt16).ToByte(), 0)
	assert.EqualValues(NewNumber(math.MaxUint16).ToByte(), 0xFF)
	assert.EqualValues(NewNumber(math.MaxInt32).ToByte(), 0xFF)
	assert.EqualValues(NewNumber(math.MinInt32).ToByte(), 0)
	assert.EqualValues(NewNumber(math.MaxUint32).ToByte(), 0xFF)
	assert.EqualValues(NewNumber(math.MaxFloat32).ToByte(), 0)
	assert.EqualValues(NewNumber(math.MaxFloat64).ToByte(), 0)
}

func TestNumberToInt8(t *testing.T) {
	assert := assert.New(t)
	assert.EqualValues(NewNumber(0).ToInt8(), 0)
	assert.EqualValues(NewNumber(128).ToInt8(), -128)
	assert.EqualValues(NewNumber(255).ToInt8(), -1)
	assert.EqualValues(NewNumber(256).ToInt8(), 0)
	assert.EqualValues(NewNumber(256.0).ToInt8(), 0)
	assert.EqualValues(NewNumber(65536).ToInt8(), 0)
	assert.EqualValues(NewNumber(math.MaxInt8).ToInt8(), 0x7F)
	assert.EqualValues(NewNumber(math.MinInt8).ToInt8(), -128)
	assert.EqualValues(NewNumber(math.MaxUint8).ToInt8(), -1)
	assert.EqualValues(NewNumber(math.MaxInt16).ToInt8(), -1)
	assert.EqualValues(NewNumber(math.MinInt16).ToInt8(), 0)
	assert.EqualValues(NewNumber(math.MaxUint16).ToInt8(), -1)
	assert.EqualValues(NewNumber(math.MaxInt32).ToInt8(), -1)
	assert.EqualValues(NewNumber(math.MinInt32).ToInt8(), 0)
	assert.EqualValues(NewNumber(math.MaxUint32).ToInt8(), -1)
	assert.EqualValues(NewNumber(math.MaxFloat32).ToInt8(), 0)
	assert.EqualValues(NewNumber(math.MaxFloat64).ToInt8(), 0)
}

func TestNumberToUint8(t *testing.T) {
	assert := assert.New(t)
	assert.EqualValues(NewNumber(0).ToUint8(), 0)
	assert.EqualValues(NewNumber(128).ToUint8(), 128)
	assert.EqualValues(NewNumber(255).ToUint8(), 255)
	assert.EqualValues(NewNumber(256).ToUint8(), 0)
	assert.EqualValues(NewNumber(256.0).ToUint8(), 0)
	assert.EqualValues(NewNumber(65536).ToUint8(), 0)
	assert.EqualValues(NewNumber(math.MaxInt8).ToUint8(), 0x7F)
	assert.EqualValues(NewNumber(math.MinInt8).ToUint8(), 0x80)
	assert.EqualValues(NewNumber(math.MaxUint8).ToUint8(), 0xFF)
	assert.EqualValues(NewNumber(math.MaxInt16).ToUint8(), 0xFF)
	assert.EqualValues(NewNumber(math.MinInt16).ToUint8(), 0)
	assert.EqualValues(NewNumber(math.MaxUint16).ToUint8(), 0xFF)
	assert.EqualValues(NewNumber(math.MaxInt32).ToUint8(), 0xFF)
	assert.EqualValues(NewNumber(math.MinInt32).ToUint8(), 0)
	assert.EqualValues(NewNumber(math.MaxUint32).ToUint8(), 0xFF)
	assert.EqualValues(NewNumber(math.MaxFloat32).ToUint8(), 0)
	assert.EqualValues(NewNumber(math.MaxFloat64).ToUint8(), 0)
}

func TestNumberToInt16(t *testing.T) {
	assert := assert.New(t)
	assert.EqualValues(NewNumber(0).ToInt16(), 0)
	assert.EqualValues(NewNumber(128).ToInt16(), 128)
	assert.EqualValues(NewNumber(255).ToInt16(), 255)
	assert.EqualValues(NewNumber(256).ToInt16(), 256)
	assert.EqualValues(NewNumber(256.0).ToInt16(), 256)
	assert.EqualValues(NewNumber(65536).ToInt16(), 0)
	assert.EqualValues(NewNumber(math.MaxInt8).ToInt16(), math.MaxInt8)
	assert.EqualValues(NewNumber(math.MinInt8).ToInt16(), math.MinInt8)
	assert.EqualValues(NewNumber(math.MaxUint8).ToInt16(), math.MaxUint8)
	assert.EqualValues(NewNumber(math.MaxInt16).ToInt16(), math.MaxInt16)
	assert.EqualValues(NewNumber(math.MinInt16).ToInt16(), math.MinInt16)
	assert.EqualValues(NewNumber(math.MaxUint16).ToInt16(), -1)
	assert.EqualValues(NewNumber(math.MaxInt32).ToInt16(), -1)
	assert.EqualValues(NewNumber(math.MinInt32).ToInt16(), 0)
	assert.EqualValues(NewNumber(math.MaxUint32).ToInt16(), -1)
	assert.EqualValues(NewNumber(math.MaxFloat32).ToInt16(), 0)
	assert.EqualValues(NewNumber(math.MaxFloat64).ToInt16(), 0)
}

func TestNumberToUint16(t *testing.T) {
	assert := assert.New(t)
	assert.EqualValues(NewNumber(0).ToUint16(), 0)
	assert.EqualValues(NewNumber(128).ToUint16(), 128)
	assert.EqualValues(NewNumber(255).ToUint16(), 255)
	assert.EqualValues(NewNumber(256).ToUint16(), 256)
	assert.EqualValues(NewNumber(256.0).ToUint16(), 256)
	assert.EqualValues(NewNumber(65536).ToUint16(), 0)
	assert.EqualValues(NewNumber(math.MaxInt8).ToUint16(), math.MaxInt8)
	assert.EqualValues(NewNumber(math.MinInt8).ToUint16(), 0xFF80)
	assert.EqualValues(NewNumber(math.MaxUint8).ToUint16(), math.MaxUint8)
	assert.EqualValues(NewNumber(math.MaxInt16).ToUint16(), math.MaxInt16)
	assert.EqualValues(NewNumber(math.MinInt16).ToUint16(), 0x8000)
	assert.EqualValues(NewNumber(math.MaxUint16).ToUint16(), math.MaxUint16)
	assert.EqualValues(NewNumber(math.MaxInt32).ToUint16(), 0xFFFF)
	assert.EqualValues(NewNumber(math.MinInt32).ToUint16(), 0)
	assert.EqualValues(NewNumber(math.MaxUint32).ToUint16(), 0xFFFF)
	assert.EqualValues(NewNumber(math.MaxFloat32).ToUint16(), 0)
	assert.EqualValues(NewNumber(math.MaxFloat64).ToUint16(), 0)
}

func TestNumberToInt32(t *testing.T) {
	assert := assert.New(t)
	assert.EqualValues(NewNumber(0).ToInt32(), 0)
	assert.EqualValues(NewNumber(128).ToInt32(), 128)
	assert.EqualValues(NewNumber(255).ToInt32(), 255)
	assert.EqualValues(NewNumber(256).ToInt32(), 256)
	assert.EqualValues(NewNumber(256.0).ToInt32(), 256)
	assert.EqualValues(NewNumber(65536).ToInt32(), 65536)
	assert.EqualValues(NewNumber(math.MaxInt8).ToInt32(), math.MaxInt8)
	assert.EqualValues(NewNumber(math.MinInt8).ToInt32(), math.MinInt8)
	assert.EqualValues(NewNumber(math.MaxUint8).ToInt32(), math.MaxUint8)
	assert.EqualValues(NewNumber(math.MaxInt16).ToInt32(), math.MaxInt16)
	assert.EqualValues(NewNumber(math.MinInt16).ToInt32(), math.MinInt16)
	assert.EqualValues(NewNumber(math.MaxUint16).ToInt32(), math.MaxUint16)
	assert.EqualValues(NewNumber(math.MaxInt32).ToInt32(), math.MaxInt32)
	assert.EqualValues(NewNumber(math.MinInt32).ToInt32(), math.MinInt32)
	assert.EqualValues(NewNumber(math.MaxUint32).ToInt32(), -1)
	assert.EqualValues(NewNumber(math.MaxFloat32).ToInt32(), 0)
	assert.EqualValues(NewNumber(math.MaxFloat64).ToInt32(), 0)
}

func TestNumberToUint32(t *testing.T) {
	assert := assert.New(t)
	assert.EqualValues(NewNumber(0).ToUint32(), 0)
	assert.EqualValues(NewNumber(128).ToUint32(), 128)
	assert.EqualValues(NewNumber(255).ToUint32(), 255)
	assert.EqualValues(NewNumber(256).ToUint32(), 256)
	assert.EqualValues(NewNumber(256.0).ToUint32(), 256)
	assert.EqualValues(NewNumber(65536).ToUint32(), 65536)
	assert.EqualValues(NewNumber(math.MaxInt8).ToUint32(), 0x7F)
	assert.EqualValues(NewNumber(math.MinInt8).ToUint32(), 0xFFFFFF80)
	assert.EqualValues(NewNumber(math.MaxUint8).ToUint32(), math.MaxUint8)
	assert.EqualValues(NewNumber(math.MaxInt16).ToUint32(), 0x7FFF)
	assert.EqualValues(NewNumber(math.MinInt16).ToUint32(), 0xFFFF8000)
	assert.EqualValues(NewNumber(math.MaxUint16).ToUint32(), math.MaxUint16)
	assert.EqualValues(NewNumber(math.MaxInt32).ToUint32(), math.MaxInt32)
	assert.EqualValues(NewNumber(math.MinInt32).ToUint32(), 0x80000000)
	assert.EqualValues(NewNumber(math.MaxUint32).ToUint32(), math.MaxUint32)
	assert.EqualValues(NewNumber(math.MaxFloat32).ToUint32(), 0)
	assert.EqualValues(NewNumber(math.MaxFloat64).ToUint32(), 0)
}

func TestNumberToInt64(t *testing.T) {
	assert := assert.New(t)
	assert.EqualValues(NewNumber(0).ToInt64(), 0)
	assert.EqualValues(NewNumber(128).ToInt64(), 128)
	assert.EqualValues(NewNumber(255).ToInt64(), 255)
	assert.EqualValues(NewNumber(256).ToInt64(), 256)
	assert.EqualValues(NewNumber(256.0).ToInt64(), 256)
	assert.EqualValues(NewNumber(65536).ToInt64(), 65536)
	assert.EqualValues(NewNumber(math.MaxInt8).ToInt64(), math.MaxInt8)
	assert.EqualValues(NewNumber(math.MinInt8).ToInt64(), math.MinInt8)
	assert.EqualValues(NewNumber(math.MaxUint8).ToInt64(), math.MaxUint8)
	assert.EqualValues(NewNumber(math.MaxInt16).ToInt64(), math.MaxInt16)
	assert.EqualValues(NewNumber(math.MinInt16).ToInt64(), math.MinInt16)
	assert.EqualValues(NewNumber(math.MaxUint16).ToInt64(), math.MaxUint16)
	assert.EqualValues(NewNumber(math.MaxInt32).ToInt64(), math.MaxInt32)
	assert.EqualValues(NewNumber(math.MinInt32).ToInt64(), math.MinInt32)
	assert.EqualValues(NewNumber(math.MaxUint32).ToInt64(), math.MaxUint32)
}

func TestNumberToUint64(t *testing.T) {
	assert := assert.New(t)
	assert.EqualValues(NewNumber(0).ToUint64(), 0)
	assert.EqualValues(NewNumber(128).ToUint64(), 128)
	assert.EqualValues(NewNumber(255).ToUint64(), 255)
	assert.EqualValues(NewNumber(256).ToUint64(), 256)
	assert.EqualValues(NewNumber(256.0).ToUint64(), 256)
	assert.EqualValues(NewNumber(65536).ToUint64(), 65536)
	assert.EqualValues(NewNumber(math.MaxInt8).ToUint64(), 0x7F)
	assert.EqualValues(
		NewNumber(math.MinInt8).ToUint64(), uint64(0xFFFFFFFFFFFFFF80))
	assert.EqualValues(NewNumber(math.MaxUint8).ToUint64(), math.MaxUint8)
	assert.EqualValues(NewNumber(math.MaxInt16).ToUint64(), 0x7FFF)
	assert.EqualValues(
		NewNumber(math.MinInt16).ToUint64(), uint64(0xFFFFFFFFFFFF8000))
	assert.EqualValues(NewNumber(math.MaxUint16).ToUint64(), math.MaxUint16)
	assert.EqualValues(NewNumber(math.MaxInt32).ToUint64(), math.MaxInt32)
	assert.EqualValues(
		NewNumber(math.MinInt32).ToUint64(), uint64(0xFFFFFFFF80000000))
	assert.EqualValues(NewNumber(math.MaxUint32).ToUint64(), math.MaxUint32)
}

func TestNumberToFloat32(t *testing.T) {
	assert := assert.New(t)
	assert.EqualValues(NewNumber(0).ToFloat32(), 0)
	assert.EqualValues(NewNumber(128).ToFloat32(), 128)
	assert.EqualValues(NewNumber(255).ToFloat32(), 255)
	assert.EqualValues(NewNumber(256).ToFloat32(), 256)
	assert.EqualValues(NewNumber(256.0).ToFloat32(), 256)
	assert.EqualValues(NewNumber(65536).ToFloat32(), 65536)
	assert.EqualValues(NewNumber(math.MaxInt8).ToFloat32(), math.MaxInt8)
	assert.EqualValues(NewNumber(math.MinInt8).ToFloat32(), math.MinInt8)
	assert.EqualValues(NewNumber(math.MaxUint8).ToFloat32(), math.MaxUint8)
	assert.EqualValues(NewNumber(math.MaxInt16).ToFloat32(), math.MaxInt16)
	assert.EqualValues(NewNumber(math.MinInt16).ToFloat32(), math.MinInt16)
	assert.EqualValues(NewNumber(math.MaxUint16).ToFloat32(), math.MaxUint16)
	assert.InEpsilon(NewNumber(math.MaxInt32).ToFloat32(), math.MaxInt32,
		float64(Float32Epsilon()))
	assert.EqualValues(NewNumber(math.MinInt32).ToFloat32(), math.MinInt32)
	assert.InEpsilon(NewNumber(math.MaxUint32).ToFloat32(), math.MaxUint32,
		float64(Float32Epsilon()))
}

func TestNumberToFloat64(t *testing.T) {
	assert := assert.New(t)
	assert.EqualValues(NewNumber(0).ToFloat64(), 0)
	assert.EqualValues(NewNumber(128).ToFloat64(), 128)
	assert.EqualValues(NewNumber(255).ToFloat64(), 255)
	assert.EqualValues(NewNumber(256).ToFloat64(), 256)
	assert.EqualValues(NewNumber(256.0).ToFloat64(), 256)
	assert.EqualValues(NewNumber(65536).ToFloat64(), 65536)
	assert.EqualValues(NewNumber(math.MaxInt8).ToFloat64(), math.MaxInt8)
	assert.EqualValues(NewNumber(math.MinInt8).ToFloat64(), math.MinInt8)
	assert.EqualValues(NewNumber(math.MaxUint8).ToFloat64(), math.MaxUint8)
	assert.EqualValues(NewNumber(math.MaxInt16).ToFloat64(), math.MaxInt16)
	assert.EqualValues(NewNumber(math.MinInt16).ToFloat64(), math.MinInt16)
	assert.EqualValues(NewNumber(math.MaxUint16).ToFloat64(), math.MaxUint16)
	assert.EqualValues(NewNumber(math.MaxInt32).ToFloat64(), math.MaxInt32)
	assert.EqualValues(NewNumber(math.MinInt32).ToFloat64(), math.MinInt32)
	assert.EqualValues(NewNumber(math.MaxUint32).ToFloat64(), math.MaxUint32)
}

func TestNumberCompare(t *testing.T) {
	assert := assert.New(t)
	assert.True(NewNumber(0) < 1)
	assert.True(NewNumber(0) == 0.0)
	assert.False(NewNumber(0.0) > 1)
}

func TestNumberCalculate(t *testing.T) {
	assert := assert.New(t)
	assert.EqualValues(NewNumber(1)+1, 2)
	assert.EqualValues(2*NewNumber(1)+NewNumber(2.2)/NewNumber(2), 3.1)
}

func TestNumberToString(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(fmt.Sprint(NewNumber(0.1)), "0.1")
	assert.Equal(fmt.Sprint(NewNumber(2.0)), "2")
	assert.Equal(fmt.Sprintf("%.2f", NewNumber(1)), "1.00")
	assert.Equal(fmt.Sprintf("%02d", NewNumber(1).ToUint32()), "01")
}

func TestNumberMarshalBinary(t *testing.T) {
	assert := assert.New(t)
	buf, err := NewNumber(1).MarshalBinary()
	assert.Nil(err)
	assert.Equal(buf, []byte{0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xF0, 0x3F, 0x60, 0x12})
	buf, err = MarshalObject(NewNumber(1))
	assert.Nil(err)
	assert.Equal(buf, []byte{0x0E, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xF0, 0x3F, 0x60, 0x12})

}

func TestNumberUnmarshalBinary(t *testing.T) {
	assert := assert.New(t)
	var v Number
	err := v.UnmarshalBinary(nil)
	assert.NotNil(err)
	err = v.UnmarshalBinary([]byte{})
	assert.NotNil(err)
	err = v.UnmarshalBinary([]byte{0x00})
	assert.NotNil(err)
	err = v.UnmarshalBinary([]byte{0x00, 0x00, 0x00, 0x00})
	assert.NotNil(err)
	err = v.UnmarshalBinary([]byte{0x01, 0x00, 0x00, 0x00})
	assert.NotNil(err)
	err = v.UnmarshalBinary([]byte{0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xF0, 0x3F, 0x60, 0x00})
	assert.NotNil(err)
	err = v.UnmarshalBinary([]byte{0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xF0, 0x3F, 0x60, 0x12})
	assert.Nil(err)
	assert.EqualValues(v, 1)
	err = v.UnmarshalBinary([]byte{0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xF0, 0x3F, 0x60, 0x12, 0xAA})
	assert.Nil(err)
	assert.EqualValues(v, 1)
	v = 2
	assert.EqualValues(v, 2)
	n := UnmarshalObject(&v, []byte{0x0E, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xF0, 0x3F, 0x60, 0x12,
		0xAA, 0xFF})
	assert.EqualValues(n, 18)
	assert.EqualValues(v, 1)
}
