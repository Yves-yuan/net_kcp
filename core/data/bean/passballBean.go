/**
 * Auto generated, do not edit it
 *
 * PassBall
 */
package bean

import (
	. "server/core/datastream"
)

type PassBallBean struct {
	animationIndex  int32   // 动画索引
	isLeftFoot      int32   // 是否左脚
	isFloor         int32   // 是否地面传球
	isIdle          int32   // 是否静止
	isLight         int32   // 是否小力度
	combKey         string  // 前面5个组合成的key
	animationId     int32   // 动作id
	animationName   string  // 动作名称
	dribbleTime     float32 // 踢到球时间
	mirror          int32   // 镜像
	dribbleDistance float32 // 踢到球距离
}

func (c *PassBallBean) AnimationIndex() int32 {
	return c.animationIndex

}
func (c *PassBallBean) IsLeftFoot() int32 {
	return c.isLeftFoot

}
func (c *PassBallBean) IsFloor() int32 {
	return c.isFloor

}
func (c *PassBallBean) IsIdle() int32 {
	return c.isIdle

}
func (c *PassBallBean) IsLight() int32 {
	return c.isLight

}
func (c *PassBallBean) CombKey() string {
	return c.combKey

}
func (c *PassBallBean) AnimationId() int32 {
	return c.animationId

}
func (c *PassBallBean) AnimationName() string {
	return c.animationName

}
func (c *PassBallBean) DribbleTime() float32 {
	return c.dribbleTime

}
func (c *PassBallBean) Mirror() int32 {
	return c.mirror

}
func (c *PassBallBean) DribbleDistance() float32 {
	return c.dribbleDistance

}

func (c *PassBallBean) LoadData(dataStream *DataInputStream) {
	if dataStream != nil {
		c.animationIndex, _ = dataStream.ReadInt32()
		c.isLeftFoot, _ = dataStream.ReadInt32()
		c.isFloor, _ = dataStream.ReadInt32()
		c.isIdle, _ = dataStream.ReadInt32()
		c.isLight, _ = dataStream.ReadInt32()
		c.combKey, _ = dataStream.ReadUTF()
		c.animationId, _ = dataStream.ReadInt32()
		c.animationName, _ = dataStream.ReadUTF()
		c.dribbleTime, _ = dataStream.ReadFloat32()
		c.mirror, _ = dataStream.ReadInt32()
		c.dribbleDistance, _ = dataStream.ReadFloat32()
	}
}
