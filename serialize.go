package core

import (
	"bytes"
	"encoding"
	"encoding/binary"
	"math"
)

var defaultByteOrder binary.ByteOrder = binary.LittleEndian

func SetByteOrder(o binary.ByteOrder) {
	defaultByteOrder = o
}

func MarshalSimpleType(d interface{}) []byte {
	tmp := make([]byte, 8)
	switch v := d.(type) {
	case byte:
		return []byte{v}
	case int8:
		return []byte{byte(v)}
	case int16:
		defaultByteOrder.PutUint16(tmp, uint16(v))
		return tmp[:2]
	case uint16:
		defaultByteOrder.PutUint16(tmp, v)
		return tmp[:2]
	case int:
		defaultByteOrder.PutUint32(tmp, uint32(v))
		return tmp[:4]
	case uint:
		defaultByteOrder.PutUint32(tmp, uint32(v))
		return tmp[:4]
	case int32:
		defaultByteOrder.PutUint32(tmp, uint32(v))
		return tmp[:4]
	case uint32:
		defaultByteOrder.PutUint32(tmp, v)
		return tmp[:4]
	case int64:
		defaultByteOrder.PutUint64(tmp, uint64(v))
		return tmp
	case uint64:
		defaultByteOrder.PutUint64(tmp, v)
		return tmp
	case float32:
		return MarshalSimpleType(math.Float32bits(v))
	case float64:
		return MarshalSimpleType(math.Float64bits(v))
	}
	panic("MarshalSimpleType: Unknown type")
}

func UnmashalSimpleType(p interface{}, data []byte) int {
	switch v := p.(type) {
	case *byte:
		*v = data[0]
		return 1
	case *int8:
		*v = int8(data[0])
		return 1
	case *int16:
		*v = int16(defaultByteOrder.Uint16(data))
		return 2
	case *uint16:
		*v = defaultByteOrder.Uint16(data)
		return 2
	case *int:
		*v = int(int32(defaultByteOrder.Uint32(data)))
		return 4
	case *uint:
		*v = uint(defaultByteOrder.Uint32(data))
		return 4
	case *int32:
		*v = int32(defaultByteOrder.Uint32(data))
		return 4
	case *uint32:
		*v = defaultByteOrder.Uint32(data)
		return 4
	case *int64:
		*v = int64(defaultByteOrder.Uint64(data))
		return 8
	case *uint64:
		*v = defaultByteOrder.Uint64(data)
		return 8
	case *float32:
		*v = math.Float32frombits(defaultByteOrder.Uint32(data))
		return 4
	case *float64:
		*v = math.Float64frombits(defaultByteOrder.Uint64(data))
		return 8
	}
	panic("UnmarshalSimpleType: Unknown type")
}

func MarshalString(s string) []byte {
	var buf bytes.Buffer
	buf.Write(MarshalSimpleType(uint16(len(s))))
	buf.WriteString(s)
	return buf.Bytes()
}

func UnmarshalString(dest *string, data []byte) int {
	var len uint16
	offset := UnmashalSimpleType(&len, data)
	*dest = string(data[offset : offset+int(len)])
	return offset + int(len)
}

func MarshalObject(obj encoding.BinaryMarshaler) ([]byte, error) {
	binary, err := obj.MarshalBinary()
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	buf.Write(MarshalSimpleType(uint32(len(binary))))
	buf.Write(binary)
	return buf.Bytes(), nil
}

func UnmarshalObject(dest encoding.BinaryUnmarshaler, data []byte) int {
	var len uint32
	offset := UnmashalSimpleType(&len, data)
	dest.UnmarshalBinary(data[offset : offset+int(len)])
	return offset + int(len)
}
