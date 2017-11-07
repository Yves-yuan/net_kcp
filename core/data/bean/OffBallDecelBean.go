/**
 * Auto generated, do not edit it
 *
 * OffBallDecel
 */
package bean

import (
	. "server/core/datastream"
)

type OffBallDecelBean struct {
	index             int32   // 动作索引
	name              string  // 动作名称
	id                int32   // 动作ID
	mirrorName        string  // 镜像动作名称
	mirrorID          int32   // 镜像动作ID
	fromNormalPercent float32 // 常速移动进入的百分比
	beginFoot         int32   // 动作开始时踏地脚(0左，1右)
	overFoot          int32   // 动作结束时踏地脚(0左，1右)
	step1End          float32 // 第一步结束百分比
	step2End          float32 // 第二步结束百分比
	step3End          float32 // 第三部结束百分比
	step4End          float32 // 第四步结束百分比
	step1Foot         int32   // 第一步结束踏地脚(0左，1右)
	step2Foot         int32   // 第二步结束踏地脚(0左，1右)
	step3Foot         int32   // 第三步结束踏地脚(0左，1右)
	step4Foot         int32   // 第四步结束踏地脚(0左，1右)
}

func (c *OffBallDecelBean) Index() int32 {
	return c.index

}
func (c *OffBallDecelBean) Name() string {
	return c.name

}
func (c *OffBallDecelBean) Id() int32 {
	return c.id

}
func (c *OffBallDecelBean) MirrorName() string {
	return c.mirrorName

}
func (c *OffBallDecelBean) MirrorID() int32 {
	return c.mirrorID

}
func (c *OffBallDecelBean) FromNormalPercent() float32 {
	return c.fromNormalPercent

}
func (c *OffBallDecelBean) BeginFoot() int32 {
	return c.beginFoot

}
func (c *OffBallDecelBean) OverFoot() int32 {
	return c.overFoot

}
func (c *OffBallDecelBean) Step1End() float32 {
	return c.step1End

}
func (c *OffBallDecelBean) Step2End() float32 {
	return c.step2End

}
func (c *OffBallDecelBean) Step3End() float32 {
	return c.step3End

}
func (c *OffBallDecelBean) Step4End() float32 {
	return c.step4End

}
func (c *OffBallDecelBean) Step1Foot() int32 {
	return c.step1Foot

}
func (c *OffBallDecelBean) Step2Foot() int32 {
	return c.step2Foot

}
func (c *OffBallDecelBean) Step3Foot() int32 {
	return c.step3Foot

}
func (c *OffBallDecelBean) Step4Foot() int32 {
	return c.step4Foot

}

func (c *OffBallDecelBean) LoadData(dataStream *DataInputStream) {
	if dataStream != nil {
		c.index, _ = dataStream.ReadInt32()
		c.name, _ = dataStream.ReadUTF()
		c.id, _ = dataStream.ReadInt32()
		c.mirrorName, _ = dataStream.ReadUTF()
		c.mirrorID, _ = dataStream.ReadInt32()
		c.fromNormalPercent, _ = dataStream.ReadFloat32()
		c.beginFoot, _ = dataStream.ReadInt32()
		c.overFoot, _ = dataStream.ReadInt32()
		c.step1End, _ = dataStream.ReadFloat32()
		c.step2End, _ = dataStream.ReadFloat32()
		c.step3End, _ = dataStream.ReadFloat32()
		c.step4End, _ = dataStream.ReadFloat32()
		c.step1Foot, _ = dataStream.ReadInt32()
		c.step2Foot, _ = dataStream.ReadInt32()
		c.step3Foot, _ = dataStream.ReadInt32()
		c.step4Foot, _ = dataStream.ReadInt32()
	}
}
