/**
 * Auto generated, do not edit it
 *
 * PlayerAnimConfig
 */
package bean

import (
	. "server/core/datastream"
)

type PlayerAnimConfigBean struct {
	name              string  // 动画名称
	moveSpeed         float32 // 移动速度（米/秒）
	kickBallTime      float32 // 触球时间（秒）
	kickBallDistance  float32 // 触球距离（米）
	kickBallTime2     float32 // 触球时间（秒）
	kickBallDistance2 float32 // 触球距离（米）
	frameCount        int32   // 帧数
}

func (c *PlayerAnimConfigBean) Name() string {
	return c.name

}
func (c *PlayerAnimConfigBean) MoveSpeed() float32 {
	return c.moveSpeed

}
func (c *PlayerAnimConfigBean) KickBallTime() float32 {
	return c.kickBallTime

}
func (c *PlayerAnimConfigBean) KickBallDistance() float32 {
	return c.kickBallDistance

}
func (c *PlayerAnimConfigBean) KickBallTime2() float32 {
	return c.kickBallTime2

}
func (c *PlayerAnimConfigBean) KickBallDistance2() float32 {
	return c.kickBallDistance2

}
func (c *PlayerAnimConfigBean) FrameCount() int32 {
	return c.frameCount

}

func (c *PlayerAnimConfigBean) LoadData(dataStream *DataInputStream) {
	if dataStream != nil {
		c.name, _ = dataStream.ReadUTF()
		c.moveSpeed, _ = dataStream.ReadFloat32()
		c.kickBallTime, _ = dataStream.ReadFloat32()
		c.kickBallDistance, _ = dataStream.ReadFloat32()
		c.kickBallTime2, _ = dataStream.ReadFloat32()
		c.kickBallDistance2, _ = dataStream.ReadFloat32()
		c.frameCount, _ = dataStream.ReadInt32()
	}
}
