/**
 * Auto generated, do not edit it
 *
 * OffBallNaturalStop
 */
package bean

import (
	. "server/core/datastream"
)

type OffBallNaturalStopBean struct {
	index         int32   // 索引
	name          string  // 动画名称
	id            int32   // 动画ID
	steeringAngle int32   // 转向角度
	isMirror      int32   // 是否镜像
	beginFoot     int32   // 动作开始时踏地脚(0左，1右)
	overFoot      int32   // 动作结束时踏地脚(0左，1右)
	step1End      float32 // 第一步结束时间
	step2End      float32 // 第二步结束时间
	step1Foot     int32   // 第一步结束踏地脚(0左，1右)
	step2Foot     int32   // 第二步结束踏地脚(0左，1右)
}

func (c *OffBallNaturalStopBean) Index() int32 {
	return c.index

}
func (c *OffBallNaturalStopBean) Name() string {
	return c.name

}
func (c *OffBallNaturalStopBean) Id() int32 {
	return c.id

}
func (c *OffBallNaturalStopBean) SteeringAngle() int32 {
	return c.steeringAngle

}
func (c *OffBallNaturalStopBean) IsMirror() int32 {
	return c.isMirror

}
func (c *OffBallNaturalStopBean) BeginFoot() int32 {
	return c.beginFoot

}
func (c *OffBallNaturalStopBean) OverFoot() int32 {
	return c.overFoot

}
func (c *OffBallNaturalStopBean) Step1End() float32 {
	return c.step1End

}
func (c *OffBallNaturalStopBean) Step2End() float32 {
	return c.step2End

}
func (c *OffBallNaturalStopBean) Step1Foot() int32 {
	return c.step1Foot

}
func (c *OffBallNaturalStopBean) Step2Foot() int32 {
	return c.step2Foot

}

func (c *OffBallNaturalStopBean) LoadData(dataStream *DataInputStream) {
	if dataStream != nil {
		c.index, _ = dataStream.ReadInt32()
		c.name, _ = dataStream.ReadUTF()
		c.id, _ = dataStream.ReadInt32()
		c.steeringAngle, _ = dataStream.ReadInt32()
		c.isMirror, _ = dataStream.ReadInt32()
		c.beginFoot, _ = dataStream.ReadInt32()
		c.overFoot, _ = dataStream.ReadInt32()
		c.step1End, _ = dataStream.ReadFloat32()
		c.step2End, _ = dataStream.ReadFloat32()
		c.step1Foot, _ = dataStream.ReadInt32()
		c.step2Foot, _ = dataStream.ReadInt32()
	}
}
