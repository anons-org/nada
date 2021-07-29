package Utils

import (
	"encoding/binary"
	"math"
)

func Float64frombytes(bytes []byte) float64 {
	bits := binary.BigEndian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return float
}

func Float64bytes(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, bits)
	return bytes
}

//func Uint16(b []byte) uint16 {
//	_ = b[1]
//	return uint16(b[0]) | uint16(b[1])<<8
//}
//
//
//func BytesToUint32(b []byte) uint32 {
//	_ = b[3] // bounds check hint to compiler; see golang.org/issue/14808
//	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
//}
//
//
//func Uint64(b []byte) uint64 {
//	_ = b[7]
//	return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |        uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56
//}
//
//
