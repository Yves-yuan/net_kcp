/**
 * Auto generated, do not edit it
 *
 * Offballfaststop
 */
package bean

import (
	. "server/core/datastream"
)

type OffballfaststopBean struct {
	index         int32  // 索引
	name          string // 动画名称
	id            int32  // 动画ID
	steeringAngle int32  // 转向角度
	footScale     string // 脚比例
}

func (c *OffballfaststopBean) Index() int32 {
	return c.index

}
func (c *OffballfaststopBean) Name() string {
	return c.name

}
func (c *OffballfaststopBean) Id() int32 {
	return c.id

}
func (c *OffballfaststopBean) SteeringAngle() int32 {
	return c.steeringAngle

}
func (c *OffballfaststopBean) FootScale() string {
	return c.footScale

}

func (c *OffballfaststopBean) LoadData(dataStream *DataInputStream) {
	if dataStream != nil {
		c.index, _ = dataStream.ReadInt32()
		c.name, _ = dataStream.ReadUTF()
		c.id, _ = dataStream.ReadInt32()
		c.steeringAngle, _ = dataStream.ReadInt32()
		c.footScale, _ = dataStream.ReadUTF()
	}
}
