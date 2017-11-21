/*
	日期: 2017-07-10
	作者: Yuan
	功能: 作为二进制传输协议的基类，定义了将各种基础类型(int,uint32,uint64,uint16)等写入缓冲区的函数
*/
package binary

import "server/cmd/battle/math/position"

type CBaseBinaryMessage struct{}

// 写入有符号16位整形数
func (b *CBaseBinaryMessage) WriteByte(buf *CIOBuffer, v byte) {
	buf.WriteByte(v)
}

// 写入有符号8位整形数
func (b *CBaseBinaryMessage) WriteInt8(buf *CIOBuffer, v int8) {
	buf.WriteByte(byte(v))
}

// 写入有符号8位整形数
func (b *CBaseBinaryMessage) WriteUint8(buf *CIOBuffer, v uint8) {
	buf.WriteByte(byte(v))
}

// 写入无符号16位整形数
func (b *CBaseBinaryMessage) WriteUint16(buf *CIOBuffer, v uint16) {
	buf.WriteUint16(v)
}

// 写入有符号16位整形数
func (b *CBaseBinaryMessage) WriteInt16(buf *CIOBuffer, v int16) {
	buf.WriteInt16(v)
}

// 写入无符号32位整形数
func (b *CBaseBinaryMessage) WriteUint32(buf *CIOBuffer, v uint32) {
	buf.WriteUint32(v)
}

// 写入有符号32位整形数
func (b *CBaseBinaryMessage) WriteInt32(buf *CIOBuffer, v int32) {
	buf.WriteInt32(v)
}

// 写入无符号64位整形数
func (b *CBaseBinaryMessage) WriteUint64(buf *CIOBuffer, v uint64) {
	buf.WriteUint64(v)
}

// 写入有符号64位整形数
func (b *CBaseBinaryMessage) WriteInt64(buf *CIOBuffer, v int64) {
	buf.WriteInt64(v)
}

// 写入32位浮点数
func (b *CBaseBinaryMessage) WriteFloat32(buf *CIOBuffer, v float32) {
	buf.WriteFloat32(v)
}

// 写入64位浮点数
func (b *CBaseBinaryMessage) WriteFloat64(buf *CIOBuffer, v float64) {
	buf.WriteFloat64(v)
}

// 写入字符串
func (b *CBaseBinaryMessage) WriteString(buf *CIOBuffer, v string) {
	data := []byte(v)
	length := len(data)

	// 写入字符串长度
	buf.WriteUint16(uint16(length))

	// 写入数据
	buf.Write(data)
}

// 写入坐标
func (b *CBaseBinaryMessage) WriteVec3f(buf *CIOBuffer, v *position.CVec3f) {
	buf.WriteFloat32(v.X())
	buf.WriteFloat32(v.Y())
	buf.WriteFloat32(v.Z())
}

// 写入Bean
func (b *CBaseBinaryMessage) WriteBean(buf *CIOBuffer, v IBaseBinaryMessage) {
	v.Write(buf)
}
