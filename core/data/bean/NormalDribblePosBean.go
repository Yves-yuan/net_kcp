/**
 * Auto generated, do not edit it
 *
 * NormalDribblePos
 */
package bean

import (
	. "server/core/datastream"
)

type NormalDribblePosBean struct {
	animIndex string  // 索引ID(动画ID_使用脚_动作类型_Pose编号)
	animId    int32   // 动画ID
	animName  string  // 动画名字
	mirror    int32   // 是否镜像
	foot      int32   // 使用脚
	angle     float32 // 动作转向角度
	posIndex  int32   // Pose编号
	startTime int32   // Pose开始帧
	endTime   int32   // Pose结束帧
	kickTime  float32 // 踢球时间
}

func (c *NormalDribblePosBean) AnimIndex() string {
	return c.animIndex

}
func (c *NormalDribblePosBean) AnimId() int32 {
	return c.animId

}
func (c *NormalDribblePosBean) AnimName() string {
	return c.animName

}
func (c *NormalDribblePosBean) Mirror() int32 {
	return c.mirror

}
func (c *NormalDribblePosBean) Foot() int32 {
	return c.foot

}
func (c *NormalDribblePosBean) Angle() float32 {
	return c.angle

}
func (c *NormalDribblePosBean) PosIndex() int32 {
	return c.posIndex

}
func (c *NormalDribblePosBean) StartTime() int32 {
	return c.startTime

}
func (c *NormalDribblePosBean) EndTime() int32 {
	return c.endTime

}
func (c *NormalDribblePosBean) KickTime() float32 {
	return c.kickTime

}

func (c *NormalDribblePosBean) LoadData(dataStream *DataInputStream) {
	if dataStream != nil {
		c.animIndex, _ = dataStream.ReadUTF()
		c.animId, _ = dataStream.ReadInt32()
		c.animName, _ = dataStream.ReadUTF()
		c.mirror, _ = dataStream.ReadInt32()
		c.foot, _ = dataStream.ReadInt32()
		c.angle, _ = dataStream.ReadFloat32()
		c.posIndex, _ = dataStream.ReadInt32()
		c.startTime, _ = dataStream.ReadInt32()
		c.endTime, _ = dataStream.ReadInt32()
		c.kickTime, _ = dataStream.ReadFloat32()
	}
}
