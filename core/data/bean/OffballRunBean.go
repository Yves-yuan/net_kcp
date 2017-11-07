/**
 * Auto generated, do not edit it
 *
 * OffballRun
 */
package bean

import (
	. "server/core/datastream"
)

type OffballRunBean struct {
	index          int32   // 动作索引
	length         float32 // 动画时长
	directSpeed    float32 // 直线速度
	turnSpeed      float32 // 大圈速度
	turnSharpSpeed float32 // 小圈速度
	step1Begin     float32 // 第一步开始
	step1End       float32 // 第一步结束
	step2End       float32 // 第二步结束
	step1BeginFoot int32   // 第一步开始落地脚(左0，右1)
	step1EndFoot   int32   // 第一步结束落地脚(左0，右1)
	step2EndFoot   int32   // 第二步结束落地脚(左0，右1)
}

func (c *OffballRunBean) Index() int32 {
	return c.index

}
func (c *OffballRunBean) Length() float32 {
	return c.length

}
func (c *OffballRunBean) DirectSpeed() float32 {
	return c.directSpeed

}
func (c *OffballRunBean) TurnSpeed() float32 {
	return c.turnSpeed

}
func (c *OffballRunBean) TurnSharpSpeed() float32 {
	return c.turnSharpSpeed

}
func (c *OffballRunBean) Step1Begin() float32 {
	return c.step1Begin

}
func (c *OffballRunBean) Step1End() float32 {
	return c.step1End

}
func (c *OffballRunBean) Step2End() float32 {
	return c.step2End

}
func (c *OffballRunBean) Step1BeginFoot() int32 {
	return c.step1BeginFoot

}
func (c *OffballRunBean) Step1EndFoot() int32 {
	return c.step1EndFoot

}
func (c *OffballRunBean) Step2EndFoot() int32 {
	return c.step2EndFoot

}

func (c *OffballRunBean) LoadData(dataStream *DataInputStream) {
	if dataStream != nil {
		c.index, _ = dataStream.ReadInt32()
		c.length, _ = dataStream.ReadFloat32()
		c.directSpeed, _ = dataStream.ReadFloat32()
		c.turnSpeed, _ = dataStream.ReadFloat32()
		c.turnSharpSpeed, _ = dataStream.ReadFloat32()
		c.step1Begin, _ = dataStream.ReadFloat32()
		c.step1End, _ = dataStream.ReadFloat32()
		c.step2End, _ = dataStream.ReadFloat32()
		c.step1BeginFoot, _ = dataStream.ReadInt32()
		c.step1EndFoot, _ = dataStream.ReadInt32()
		c.step2EndFoot, _ = dataStream.ReadInt32()
	}
}
