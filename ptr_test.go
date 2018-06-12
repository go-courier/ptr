package ptr

import (
	github_com_stretchr_testify_assert "github.com/stretchr/testify/assert"
	testing "testing"
)

func TestBool(t *testing.T) {
	github_com_stretchr_testify_assert.Equal(t, true, *Bool(true))
	github_com_stretchr_testify_assert.Equal(t, false, *Bool(false))
}

func TestInt(t *testing.T) {
	github_com_stretchr_testify_assert.Equal(t, int(1), *Int(1))
}

func TestInt8(t *testing.T) {
	github_com_stretchr_testify_assert.Equal(t, int8(1), *Int8(1))
}

func TestInt16(t *testing.T) {
	github_com_stretchr_testify_assert.Equal(t, int16(1), *Int16(1))
}

func TestInt32(t *testing.T) {
	github_com_stretchr_testify_assert.Equal(t, int32(1), *Int32(1))
}

func TestInt64(t *testing.T) {
	github_com_stretchr_testify_assert.Equal(t, int64(1), *Int64(1))
}

func TestUint(t *testing.T) {
	github_com_stretchr_testify_assert.Equal(t, uint(1), *Uint(1))
}

func TestUint8(t *testing.T) {
	github_com_stretchr_testify_assert.Equal(t, uint8(1), *Uint8(1))
}

func TestUint16(t *testing.T) {
	github_com_stretchr_testify_assert.Equal(t, uint16(1), *Uint16(1))
}

func TestUint32(t *testing.T) {
	github_com_stretchr_testify_assert.Equal(t, uint32(1), *Uint32(1))
}

func TestUint64(t *testing.T) {
	github_com_stretchr_testify_assert.Equal(t, uint64(1), *Uint64(1))
}

func TestUintptr(t *testing.T) {
	github_com_stretchr_testify_assert.Equal(t, uintptr(1), *Uintptr(1))
}

func TestFloat32(t *testing.T) {
	github_com_stretchr_testify_assert.Equal(t, float32(1), *Float32(1))
}

func TestFloat64(t *testing.T) {
	github_com_stretchr_testify_assert.Equal(t, float64(1), *Float64(1))
}

func TestComplex64(t *testing.T) {
	github_com_stretchr_testify_assert.Equal(t, complex64(1), *Complex64(1))
}

func TestComplex128(t *testing.T) {
	github_com_stretchr_testify_assert.Equal(t, complex128(1), *Complex128(1))
}

func TestString(t *testing.T) {
	github_com_stretchr_testify_assert.Equal(t, "string", *String("string"))
}

func TestByte(t *testing.T) {
	github_com_stretchr_testify_assert.Equal(t, []uint8{
		98,
		121,
		116,
		101,
		115,
	}[0], *Byte([]uint8{
		98,
		121,
		116,
		101,
		115,
	}[0]))
}

func TestRune(t *testing.T) {
	github_com_stretchr_testify_assert.Equal(t, 'r', *Rune('r'))
}
