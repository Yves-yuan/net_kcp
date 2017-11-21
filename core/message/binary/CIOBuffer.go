package binary

import (
	"bytes"
	"encoding/binary"
	"math"
)

/*************************************************
  作    者: Yuan
  版    本: v1
  完成日期: 2017-07-11
  功能描述: CIOBuffer将各个基础类型写入到buffer
*************************************************/
type CIOBuffer struct {
	bytes.Buffer
	bo binary.ByteOrder
}

func NewIOBuffer() *CIOBuffer {
	return &CIOBuffer{bo: binary.BigEndian}
}

// 写入有符号16位整形数
func (b *CIOBuffer) WriteByte(v byte) {
	b.Write([]byte{v})
}

// 写入有符号8位整形数
func (b *CIOBuffer) WriteInt8(v int8) {
	b.WriteByte(byte(v))
}

// 写入有符号8位整形数
func (b *CIOBuffer) WriteUint8(v uint8) {
	b.WriteByte(byte(v))
}

// 写入无符号16位整形数
func (b *CIOBuffer) WriteUint16(v uint16) {
	data := make([]byte, 2)
	b.bo.PutUint16(data, v)
	b.Write(data)
}

// 写入有符号16位整形数
func (b *CIOBuffer) WriteInt16(v int16) {
	data := make([]byte, 2)
	b.bo.PutUint16(data, uint16(v))
	b.Write(data)
}

// 写入无符号32位整形数
func (b *CIOBuffer) WriteUint32(v uint32) {
	data := make([]byte, 4)
	b.bo.PutUint32(data, v)
	b.Write(data)
}

// 写入有符号32位整形数
func (b *CIOBuffer) WriteInt32(v int32) {
	data := make([]byte, 4)
	b.bo.PutUint32(data, uint32(v))
	b.Write(data)
}

// 写入无符号64位整形数
func (b *CIOBuffer) WriteUint64(v uint64) {
	data := make([]byte, 8)
	b.bo.PutUint64(data, v)
	b.Write(data)
}

// 写入有符号64位整形数
func (b *CIOBuffer) WriteInt64(v int64) {
	data := make([]byte, 8)
	b.bo.PutUint64(data, uint64(v))
	b.Write(data)
}

// 写入32位浮点数
func (b *CIOBuffer) WriteFloat32(v float32) {
	data := make([]byte, 4)
	b.bo.PutUint32(data, math.Float32bits(v))
	b.Write(data)
}

// 写入64位浮点数
func (b *CIOBuffer) WriteFloat64(v float64) {
	data := make([]byte, 8)
	b.bo.PutUint64(data, math.Float64bits(v))
	b.Write(data)
}