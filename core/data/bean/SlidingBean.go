/**
 * Auto generated, do not edit it
 *
 * Sliding
 */
package bean

import (
	. "server/core/datastream"
)

type SlidingBean struct {
	index           int32   // 动画索引
	screen          int32   // 筛选条件
	animID          int32   // 动作ID
	animName        string  // 动作名称
	touchTime       float32 // 触球时间（毫秒）
	mirror          int32   // 是否镜像
	slidingDistance float32 // 铲球距离
	slidingspeed    float32 // 滑铲后球的速度
	standbytime     float32 // 接待机时间
}

func (c *SlidingBean) Index() int32 {
	return c.index

}
func (c *SlidingBean) Screen() int32 {
	return c.screen

}
func (c *SlidingBean) AnimID() int32 {
	return c.animID

}
func (c *SlidingBean) AnimName() string {
	return c.animName

}
func (c *SlidingBean) TouchTime() float32 {
	return c.touchTime

}
func (c *SlidingBean) Mirror() int32 {
	return c.mirror

}
func (c *SlidingBean) SlidingDistance() float32 {
	return c.slidingDistance

}
func (c *SlidingBean) Slidingspeed() float32 {
	return c.slidingspeed

}
func (c *SlidingBean) Standbytime() float32 {
	return c.standbytime

}

func (c *SlidingBean) LoadData(dataStream *DataInputStream) {
	if dataStream != nil {
		c.index, _ = dataStream.ReadInt32()
		c.screen, _ = dataStream.ReadInt32()
		c.animID, _ = dataStream.ReadInt32()
		c.animName, _ = dataStream.ReadUTF()
		c.touchTime, _ = dataStream.ReadFloat32()
		c.mirror, _ = dataStream.ReadInt32()
		c.slidingDistance, _ = dataStream.ReadFloat32()
		c.slidingspeed, _ = dataStream.ReadFloat32()
		c.standbytime, _ = dataStream.ReadFloat32()
	}
}
