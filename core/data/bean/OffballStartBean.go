/**
 * Auto generated, do not edit it
 *
 * OffballStart
 */
package bean

import (
	. "server/core/datastream"
)

type OffballStartBean struct {
	index      int32   // 动作索引
	name       string  // 动作名称
	id         int32   // 动作ID
	turnAngle  float32 // 转向角度(服务器用)
	breakRange float32 // 可打断区间
	accMirror  int32   // 后续加速动作是否镜像
	overFoot   int32   // 动作结束时落地脚（左0，右1）
}

func (c *OffballStartBean) Index() int32 {
	return c.index

}
func (c *OffballStartBean) Name() string {
	return c.name

}
func (c *OffballStartBean) Id() int32 {
	return c.id

}
func (c *OffballStartBean) TurnAngle() float32 {
	return c.turnAngle

}
func (c *OffballStartBean) BreakRange() float32 {
	return c.breakRange

}
func (c *OffballStartBean) AccMirror() int32 {
	return c.accMirror

}
func (c *OffballStartBean) OverFoot() int32 {
	return c.overFoot

}

func (c *OffballStartBean) LoadData(dataStream *DataInputStream) {
	if dataStream != nil {
		c.index, _ = dataStream.ReadInt32()
		c.name, _ = dataStream.ReadUTF()
		c.id, _ = dataStream.ReadInt32()
		c.turnAngle, _ = dataStream.ReadFloat32()
		c.breakRange, _ = dataStream.ReadFloat32()
		c.accMirror, _ = dataStream.ReadInt32()
		c.overFoot, _ = dataStream.ReadInt32()
	}
}
