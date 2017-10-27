package core

import (
	"bytes"
	"errors"
	_ "fmt"
	"github.com/newkedison/core/algorithm"
	"strconv"
)

type ByteArray []byte

func NewByteArray(args ...interface{}) *ByteArray {
	size := 0
	if len(args) > 0 {
		switch v := args[0].(type) {
		case int:
			size = v
		case []byte:
			ret := ByteArray(v)
			return &ret
		default:
			panic("Invalid type of 1st parameter for ByteArray.New, must be int")
		}
	}
	ret := make(ByteArray, size)
	return &ret
}

func (ba ByteArray) ToStringEx(
	withLen bool, sep string, prefix string, suffix string) string {
	var buffer bytes.Buffer
	if withLen {
		buffer.WriteString("[")
		buffer.WriteString(strconv.Itoa(len(ba)))
		buffer.WriteString("]")
	}
	for _, b := range ba {
		buffer.WriteString(prefix)
		buffer.WriteString(ByteToHexString(b))
		buffer.WriteString(suffix)
		buffer.WriteString(sep)
	}
	result := buffer.String()
	if len(ba) > 0 && len(sep) > 0 {
		return result[:len(result)-len(sep)]
	}
	return result
}

func (ba ByteArray) ToString() string {
	return ba.ToStringEx(true, " ", "", "")
}

func (ba ByteArray) Len() int {
	return len(ba)
}

func (ba *ByteArray) AppendByte(b byte) {
	*ba = append(*ba, b)
}

func (ba *ByteArray) AppendBytes(arr []byte) {
	*ba = append(*ba, arr...)
}

func (ba *ByteArray) AppendIntAsByte(i int) {
	*ba = append(*ba, byte(i&0xFF))
}

func (ba *ByteArray) AppendString(s string) {
	*ba = append(*ba, []byte(s)...)
}

func (ba *ByteArray) AddCrc16() {
	*ba = algorithm.AppendCrc16([]byte(*ba))
}

func (ba ByteArray) Crc8() byte {
	return algorithm.Crc8([]byte(ba))
}

func (ba ByteArray) Crc16() uint16 {
	return algorithm.Crc16([]byte(ba))
}

func (ba ByteArray) Crc32() uint32 {
	return algorithm.Crc32([]byte(ba))
}

func (ba ByteArray) Clone() ByteArray {
	return append(ByteArray{}, ba...)
}

func (ba *ByteArray) Assign(data []byte) {
	*ba = ByteArray(data)
}

func (ba *ByteArray) AssignByCopy(data []byte) {
	*ba = append(ByteArray{}, data...)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (ba ByteArray) MarshalBinary() ([]byte, error) {
	buf := new(bytes.Buffer)
	buf.Write(MarshalSimpleType(uint32(len(ba))))
	buf.Write([]byte(ba))
	return buf.Bytes(), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
func (ba *ByteArray) UnmarshalBinary(data []byte) error {
	buf := data
	if len(buf) == 0 {
		return errors.New("ByteArray.UnmarshalBinary: no data")
	}
	if len(buf) < 4 {
		return errors.New("ByteArray.UnmarshalBinary: invalid length")
	}
	var l uint32
	offset := 0
	offset += UnmashalSimpleType(&l, data)
	if uint32(len(buf)) < 4+l {
		return errors.New("ByteArray.UnmarshalBinary: invalid length")
	}
	ba.AssignByCopy(data[offset : offset+int(l)])
	return nil
}
