/**
 * Auto generated, do not edit it
 *
 * CatchBall
 */
package bean

import (
	. "server/core/datastream"
)

type CatchBallBean struct {
	id            int32   // 动画ID
	priority      int32   // 优先级
	animType      int32   // 动作类型(停球 11, 带球 12, 传球 13, 射门 14)
	preMovement   int32   // 前序移动状态(静止 0，移动 1，静止或移动 2)
	turnAngle     float32 // 带球（接球转向）角度
	ballHAngle    float32 // 来球水平方向与人朝向夹角（正方向来球为0）
	ballVMovement int32   // 来球垂直方向的运动(平高 1, 下落 2，反弹 4; 组合情况就将对应的值加起来，如：    平高+下落为3，平高+反弹为5，下落+反弹为6，平高+下落+反弹为7)
	animName      string  // 动作名称
	animId        int32   // 动作对应的ID
	reboundHAngle float32 // 球反弹的水平角度（相对于玩家当前朝向的角度）
	reboundVAngle float32 // 球反弹的垂直角度
	reboundSpeed  float32 // 球反弹的速度
}

func (c *CatchBallBean) Id() int32 {
	return c.id

}
func (c *CatchBallBean) Priority() int32 {
	return c.priority

}
func (c *CatchBallBean) AnimType() int32 {
	return c.animType

}
func (c *CatchBallBean) PreMovement() int32 {
	return c.preMovement

}
func (c *CatchBallBean) TurnAngle() float32 {
	return c.turnAngle

}
func (c *CatchBallBean) BallHAngle() float32 {
	return c.ballHAngle

}
func (c *CatchBallBean) BallVMovement() int32 {
	return c.ballVMovement

}
func (c *CatchBallBean) AnimName() string {
	return c.animName

}
func (c *CatchBallBean) AnimId() int32 {
	return c.animId

}
func (c *CatchBallBean) ReboundHAngle() float32 {
	return c.reboundHAngle

}
func (c *CatchBallBean) ReboundVAngle() float32 {
	return c.reboundVAngle

}
func (c *CatchBallBean) ReboundSpeed() float32 {
	return c.reboundSpeed

}

func (c *CatchBallBean) LoadData(dataStream *DataInputStream) {
	if dataStream != nil {
		c.id, _ = dataStream.ReadInt32()
		c.priority, _ = dataStream.ReadInt32()
		c.animType, _ = dataStream.ReadInt32()
		c.preMovement, _ = dataStream.ReadInt32()
		c.turnAngle, _ = dataStream.ReadFloat32()
		c.ballHAngle, _ = dataStream.ReadFloat32()
		c.ballVMovement, _ = dataStream.ReadInt32()
		c.animName, _ = dataStream.ReadUTF()
		c.animId, _ = dataStream.ReadInt32()
		c.reboundHAngle, _ = dataStream.ReadFloat32()
		c.reboundVAngle, _ = dataStream.ReadFloat32()
		c.reboundSpeed, _ = dataStream.ReadFloat32()
	}
}
