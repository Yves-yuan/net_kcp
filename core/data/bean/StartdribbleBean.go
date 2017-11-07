/**
 * Auto generated, do not edit it
 *
 * Startdribble
 */
package bean

import (
	. "server/core/datastream"
)

type StartdribbleBean struct {
	motionIndex              int32   // 动画索引
	isRightFoot              int32   // 是否右脚
	combKey                  string  // 前面2个组合成的key
	animationId              int32   // 动作id
	animationName            string  // 动作名称
	dribbleTime              float32 // 踢到球时间
	mirror                   int32   // 镜像
	dribbleDistance          float32 // 踢到球距离
	straightBeginTimePercent float32 // 接下个跑步动作的百分比
}

func (c *StartdribbleBean) MotionIndex() int32 {
	return c.motionIndex

}
func (c *StartdribbleBean) IsRightFoot() int32 {
	return c.isRightFoot

}
func (c *StartdribbleBean) CombKey() string {
	return c.combKey

}
func (c *StartdribbleBean) AnimationId() int32 {
	return c.animationId

}
func (c *StartdribbleBean) AnimationName() string {
	return c.animationName

}
func (c *StartdribbleBean) DribbleTime() float32 {
	return c.dribbleTime

}
func (c *StartdribbleBean) Mirror() int32 {
	return c.mirror

}
func (c *StartdribbleBean) DribbleDistance() float32 {
	return c.dribbleDistance

}
func (c *StartdribbleBean) StraightBeginTimePercent() float32 {
	return c.straightBeginTimePercent

}

func (c *StartdribbleBean) LoadData(dataStream *DataInputStream) {
	if dataStream != nil {
		c.motionIndex, _ = dataStream.ReadInt32()
		c.isRightFoot, _ = dataStream.ReadInt32()
		c.combKey, _ = dataStream.ReadUTF()
		c.animationId, _ = dataStream.ReadInt32()
		c.animationName, _ = dataStream.ReadUTF()
		c.dribbleTime, _ = dataStream.ReadFloat32()
		c.mirror, _ = dataStream.ReadInt32()
		c.dribbleDistance, _ = dataStream.ReadFloat32()
		c.straightBeginTimePercent, _ = dataStream.ReadFloat32()
	}
}
