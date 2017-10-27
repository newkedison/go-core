package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)
	ba := NewByteArray()
	assert.Equal(len(*ba), 0)
	ba = NewByteArray(10)
	assert.Equal(len(*ba), 10)
	ba = NewByteArray([]byte{0x01, 0x02})
	assert.Equal(len(*ba), 2)
	tmp := []byte{0x00}
	ba = NewByteArray(tmp)
	assert.Equal(len(*ba), 1)
	(*ba)[0] = 0xFF
	assert.EqualValues(tmp[0], 0xFF)

	assert.Panics(func() { NewByteArray("a") })
}

func TestToStringEx(t *testing.T) {
	assert := assert.New(t)
	var ba ByteArray
	assert.Equal(ba.ToStringEx(false, "", "", ""), "")
	assert.Equal(ba.ToStringEx(false, ",", "0x", ""), "")
	assert.Equal(ba.ToStringEx(true, ",", "0x", "XXX"), "[0]")
	ba = append(ba, 0x00)
	assert.Equal(ba.ToStringEx(false, "", "", ""), "00")
	assert.Equal(ba.ToStringEx(false, ",", "0x", ""), "0x00")
	assert.Equal(ba.ToStringEx(true, ",", "0x", ""), "[1]0x00")
	assert.Equal(ba.ToStringEx(true, ",", "", "H"), "[1]00H")
	assert.Equal(ba.ToStringEx(true, ",", "<<", ">>"), "[1]<<00>>")
	ba = append(ba, []byte{0x01, 0x02}...)
	assert.Equal(ba.ToStringEx(false, "", "", ""), "000102")
	assert.Equal(ba.ToStringEx(false, ",", "0x", ""), "0x00,0x01,0x02")
	assert.Equal(ba.ToStringEx(true, ",", "0x", ""), "[3]0x00,0x01,0x02")
	assert.Equal(ba.ToStringEx(true, ",", "", "H"), "[3]00H,01H,02H")
	assert.Equal(ba.ToStringEx(true, "|", "<<", ">>"), "[3]<<00>>|<<01>>|<<02>>")
}

func TestToString(t *testing.T) {
	assert := assert.New(t)
	var ba ByteArray
	assert.Equal(ba.ToString(), "[0]")
	ba = append(ba, 0x00)
	assert.Equal(ba.ToString(), "[1]00")
	ba = ByteArray([]byte{0x00, 0x11, 0x22, 0xFF})
	assert.Equal(ba.ToString(), "[4]00 11 22 FF")
}

func TestLen(t *testing.T) {
	assert := assert.New(t)
	var ba ByteArray
	assert.Equal(ba.Len(), 0)
	ba = append(ba, 0x00)
	assert.Equal(ba.Len(), 1)
	ba = ByteArray([]byte{0x00, 0x11, 0x22, 0xFF})
	assert.Equal(ba.Len(), 4)
}

func TestAppendByte(t *testing.T) {
	assert := assert.New(t)
	var ba ByteArray
	assert.Equal(ba.ToString(), "[0]")
	ba.AppendByte(0x00)
	assert.Equal(ba.ToString(), "[1]00")
}

func TestAppendBytes(t *testing.T) {
	assert := assert.New(t)
	var ba ByteArray
	assert.Equal(ba.ToString(), "[0]")
	ba.AppendBytes([]byte{0x00, 0x01, 0x02})
	assert.Equal(ba.ToString(), "[3]00 01 02")
}

func TestAppendIntAsByte(t *testing.T) {
	assert := assert.New(t)
	var ba ByteArray
	assert.Equal(ba.ToString(), "[0]")
	ba.AppendIntAsByte(0)
	ba.AppendIntAsByte(255)
	ba.AppendIntAsByte(256)
	ba.AppendIntAsByte(65538)
	assert.Equal(ba.ToString(), "[4]00 FF 00 02")
}

func TestAppendString(t *testing.T) {
	assert := assert.New(t)
	var ba ByteArray
	assert.Equal(ba.ToString(), "[0]")
	ba.AppendString("AAA")
	assert.Equal(ba.ToString(), "[3]41 41 41")
}

func TestAddCrc16(t *testing.T) {
	assert := assert.New(t)
	var ba ByteArray
	assert.Equal(ba.ToString(), "[0]")
	ba.AddCrc16()
	assert.Equal(ba.ToString(), "[2]FF FF")
	ba = []byte{0x10}
	ba.AddCrc16()
	assert.Equal(ba.ToString(), "[3]10 BE 8C")
	ba.AddCrc16()
	assert.Equal(ba.ToString(), "[5]10 BE 8C 00 00")
}

func TestCrc8(t *testing.T) {
	assert := assert.New(t)
	assert.EqualValues(ByteArray{0x10}.Crc8(), 0x70)
}

func TestCrc16(t *testing.T) {
	assert := assert.New(t)
	assert.EqualValues(ByteArray{0x10}.Crc16(), 0x8CBE)
}

func TestCrc32(t *testing.T) {
	assert := assert.New(t)
	assert.EqualValues(ByteArray{0x10}.Crc32(), 0xCFB5FFE9)
}

func TestClone(t *testing.T) {
	assert := assert.New(t)
	ba := ByteArray([]byte{0x01, 0x02})
	ba2 := ba.Clone()
	assert.Equal(len(ba2), len(ba))
	assert.Equal(ba2[0], ba[0])
	ba2[0] += 1
	assert.NotEqual(ba2[0], ba[0])
}

func TestAssign(t *testing.T) {
	assert := assert.New(t)
	ba := ByteArray([]byte{0x01, 0x02})
	data := []byte{0xFF}
	ba.Assign(data)
	assert.Equal(len(ba), 1)
	assert.EqualValues(ba[0], 0xFF)
	ba[0] = 0xAA
	assert.EqualValues(data[0], 0xAA)
}

func TestAssignByCopy(t *testing.T) {
	assert := assert.New(t)
	ba := ByteArray([]byte{0x01, 0x02})
	data := []byte{0xFF}
	ba.AssignByCopy(data)
	assert.Equal(len(ba), 1)
	assert.EqualValues(ba[0], 0xFF)
	ba[0] = 0xAA
	assert.EqualValues(data[0], 0xFF)
}

func TestMarshalBinary(t *testing.T) {
	assert := assert.New(t)
	ba := ByteArray([]byte{0x01, 0x02})
	buf, err := ba.MarshalBinary()
	assert.NoError(err)
	assert.Equal(len(buf), 6)
	assert.EqualValues(buf[0], 0x02)
	assert.EqualValues(buf[1], 0x00)
	assert.EqualValues(buf[2], 0x00)
	assert.EqualValues(buf[3], 0x00)
	assert.EqualValues(buf[4], 0x01)
	assert.EqualValues(buf[5], 0x02)
}

func TestUnmarshalBinary(t *testing.T) {
	assert := assert.New(t)
	ba := ByteArray([]byte{0x01, 0x02})
	buf, err := ba.MarshalBinary()
	assert.NoError(err)
	assert.Equal(len(buf), 6)
	var ba2 ByteArray
	if err := ba2.UnmarshalBinary(nil); assert.Error(err) {
		assert.EqualError(err, "ByteArray.UnmarshalBinary: no data")
	}
	if err := ba2.UnmarshalBinary([]byte{}); assert.Error(err) {
		assert.EqualError(err, "ByteArray.UnmarshalBinary: no data")
	}
	if err := ba2.UnmarshalBinary([]byte{0x00}); assert.Error(err) {
		assert.EqualError(err, "ByteArray.UnmarshalBinary: invalid length")
	}
	if err := ba2.UnmarshalBinary(
		[]byte{0x01, 0x00, 0x00, 0x00}); assert.Error(err) {
		assert.EqualError(err, "ByteArray.UnmarshalBinary: invalid length")
	}
	err = ba2.UnmarshalBinary(buf)
	assert.NoError(err)
	assert.Equal(len(ba2), 2)
	assert.EqualValues(ba2[0], 0x01)
	assert.EqualValues(ba2[1], 0x02)
}
