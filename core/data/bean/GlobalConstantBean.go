/**
 * Auto generated, do not edit it
 *
 * GlobalConstant
 */
package bean

import (
	. "server/core/datastream"
)

type GlobalConstantBean struct {
	id          int32   // 索引
	value       int32   // 整形值
	valuef      float32 // 浮点值
	values      string  // 字符串值
	description string  // 描述
}

func (c *GlobalConstantBean) Id() int32 {
	return c.id

}
func (c *GlobalConstantBean) Value() int32 {
	return c.value

}
func (c *GlobalConstantBean) Valuef() float32 {
	return c.valuef

}
func (c *GlobalConstantBean) Values() string {
	return c.values

}
func (c *GlobalConstantBean) Description() string {
	return c.description

}

func (c *GlobalConstantBean) LoadData(dataStream *DataInputStream) {
	if dataStream != nil {
		c.id, _ = dataStream.ReadInt32()
		c.value, _ = dataStream.ReadInt32()
		c.valuef, _ = dataStream.ReadFloat32()
		c.values, _ = dataStream.ReadUTF()
		c.description, _ = dataStream.ReadUTF()
	}
}
