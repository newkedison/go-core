package core

import (
	"bytes"
	"errors"
	"github.com/newkedison/core/algorithm"
	"math"
	"reflect"
	"strconv"
)

type Number float64

const (
	MaxNumber Number = Number(math.MaxFloat64)
	MinNumber Number = Number(-math.MaxFloat64)
	// Ref: https://en.wikipedia.org/wiki/Double-precision_floating-point_format#IEEE_754_double-precision_binary_floating-point_format:_binary64
	// The 53-bit significand precision gives from 15 to 17 significant decimal digits precision (2^−53 ≈ 1.11 × 10^−16). If a decimal string with at most 15 significant digits is converted to IEEE 754 double-precision representation, and then converted back to a decimal string with the same number of digits, the final result should match the original string. If an IEEE 754 double-precision number is converted to a decimal string with at least 17 significant digits, and then converted back to double-precision representation, the final result must match the original number.
	MaxIntNumber  int64  = 1e16
	MinIntNumber  int64  = -1e16
	MaxUintNumber uint64 = 1e16
)

var (
	NumberType reflect.Type = reflect.TypeOf(Number(0))
)

func Float32Epsilon() float32 {
	return math.Nextafter32(1.0, 2.0) - 1.0
}

func Float64Epsilon() float64 {
	return math.Nextafter(1.0, 2.0) - 1.0
}

func NewNumber(d interface{}) Number {
	switch v := d.(type) {
	case int64:
		if v > MaxIntNumber {
			panic("Assign an int64 bigger than " +
				strconv.FormatInt(MaxIntNumber, 10) +
				" will lost significant digits, " +
				"used float64 instead")
		}
		if v < MinIntNumber {
			panic("Assign an int64 smaller than " +
				strconv.FormatInt(MinIntNumber, 10) +
				" will lost significant digits, " +
				"used float64 instead")
		}
	case uint64:
		if v > uint64(MaxIntNumber) {
			panic("Assign an uint64 bigger than " +
				strconv.FormatUint(MaxUintNumber, 10) +
				" will lost significant digits, " +
				"used float64 instead")
		}
	}
	if reflect.TypeOf(d).ConvertibleTo(NumberType) {
		return reflect.ValueOf(d).Convert(NumberType).Interface().(Number)
	}
	panic("Invalid number type")
}

func (v Number) ToByte() byte {
	return byte(uint64(v))
}

func (v Number) ToInt8() int8 {
	return int8(int64(v))
}

func (v Number) ToUint8() uint8 {
	return uint8(uint64(v))
}

func (v Number) ToInt16() int16 {
	return int16(int64(v))
}

func (v Number) ToUint16() uint16 {
	return uint16(uint64(v))
}

func (v Number) ToInt32() int32 {
	return int32(int64(v))
}

func (v Number) ToUint32() uint32 {
	return uint32(uint64(v))
}

func (v Number) ToInt64() int64 {
	return int64(v)
}

func (v Number) ToUint64() uint64 {
	return uint64(v)
}

func (v Number) ToFloat32() float32 {
	return float32(v)
}

func (v Number) ToFloat64() float64 {
	return float64(v)
}

var serializeVersion int32 = 0

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (v Number) MarshalBinary() ([]byte, error) {
	buf := new(bytes.Buffer)
	buf.Write(MarshalSimpleType(serializeVersion))
	buf.Write(MarshalSimpleType(float64(v)))
	return algorithm.AppendCrc16(buf.Bytes()), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
func (v *Number) UnmarshalBinary(data []byte) (err error) {
	defer SetErrorWhenNotEnoughDataErrorPanic(
		"core.Number.UnmarshalBinary", &err)()
	CheckBufferSize(data, 4)
	var ver int32
	offset := 0
	offset += UnmashalSimpleType(&ver, data)
	switch ver {
	case 0:
		CheckBufferSize(data, 4+8+2)
		if !algorithm.VerifyCrc16(data[:4+8+2]) {
			return errors.New("core.Number.UnmarshalBinary: CRC check fail")
		}
		UnmashalSimpleType((*float64)(v), data[4:4+8])
	default:
		return errors.New("core.Number.UnmarshalBinary: version error")
	}
	return nil
}
