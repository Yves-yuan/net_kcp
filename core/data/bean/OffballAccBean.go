/**
 * Auto generated, do not edit it
 *
 * OffballAcc
 */
package bean

import (
	. "server/core/datastream"
)

type OffballAccBean struct {
	index          int32   // 动作索引
	name           string  // 动作名称
	length         float32 // 动画时长
	step1Begin     float32 // 第一步开始
	step1End       float32 // 第一步结束
	step2End       float32 // 第二步结束
	step3End       float32 // 第三步结束
	step4End       float32 // 第四步结束
	step1BeginFoot int32   // 第一步开始落地脚(左0，右1)
	step1EndFoot   int32   // 第一步结束落地脚(左0，右1)
	step2EndFoot   int32   // 第二步结束落地脚(左0，右1)
	step3EndFoot   int32   // 第三步结束落地脚(左0，右1)
	step4EndFoot   int32   // 第四步结束落地脚(左0，右1)
}

func (c *OffballAccBean) Index() int32 {
	return c.index

}
func (c *OffballAccBean) Name() string {
	return c.name

}
func (c *OffballAccBean) Length() float32 {
	return c.length

}
func (c *OffballAccBean) Step1Begin() float32 {
	return c.step1Begin

}
func (c *OffballAccBean) Step1End() float32 {
	return c.step1End

}
func (c *OffballAccBean) Step2End() float32 {
	return c.step2End

}
func (c *OffballAccBean) Step3End() float32 {
	return c.step3End

}
func (c *OffballAccBean) Step4End() float32 {
	return c.step4End

}
func (c *OffballAccBean) Step1BeginFoot() int32 {
	return c.step1BeginFoot

}
func (c *OffballAccBean) Step1EndFoot() int32 {
	return c.step1EndFoot

}
func (c *OffballAccBean) Step2EndFoot() int32 {
	return c.step2EndFoot

}
func (c *OffballAccBean) Step3EndFoot() int32 {
	return c.step3EndFoot

}
func (c *OffballAccBean) Step4EndFoot() int32 {
	return c.step4EndFoot

}

func (c *OffballAccBean) LoadData(dataStream *DataInputStream) {
	if dataStream != nil {
		c.index, _ = dataStream.ReadInt32()
		c.name, _ = dataStream.ReadUTF()
		c.length, _ = dataStream.ReadFloat32()
		c.step1Begin, _ = dataStream.ReadFloat32()
		c.step1End, _ = dataStream.ReadFloat32()
		c.step2End, _ = dataStream.ReadFloat32()
		c.step3End, _ = dataStream.ReadFloat32()
		c.step4End, _ = dataStream.ReadFloat32()
		c.step1BeginFoot, _ = dataStream.ReadInt32()
		c.step1EndFoot, _ = dataStream.ReadInt32()
		c.step2EndFoot, _ = dataStream.ReadInt32()
		c.step3EndFoot, _ = dataStream.ReadInt32()
		c.step4EndFoot, _ = dataStream.ReadInt32()
	}
}
