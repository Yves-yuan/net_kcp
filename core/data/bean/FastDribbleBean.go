/**
 * Auto generated, do not edit it
 *
 * FastDribble
 */
package bean

import (
	. "server/core/datastream"
)

type FastDribbleBean struct {
	id                   int32   // 动作索引
	foot                 int32   // 惯用脚
	paramId              string  // 综合key
	animId               int32   // 动画ID
	animationName        string  // 动画名字
	dribbleDistance      float32 // 趟球距离
	dribbleTime          float32 // 趟球时间
	isMirror             int32   // 是否镜像
	straightBeginPercent float32 // 切换下一动作百分比
}

func (c *FastDribbleBean) Id() int32 {
	return c.id

}
func (c *FastDribbleBean) Foot() int32 {
	return c.foot

}
func (c *FastDribbleBean) ParamId() string {
	return c.paramId

}
func (c *FastDribbleBean) AnimId() int32 {
	return c.animId

}
func (c *FastDribbleBean) AnimationName() string {
	return c.animationName

}
func (c *FastDribbleBean) DribbleDistance() float32 {
	return c.dribbleDistance

}
func (c *FastDribbleBean) DribbleTime() float32 {
	return c.dribbleTime

}
func (c *FastDribbleBean) IsMirror() int32 {
	return c.isMirror

}
func (c *FastDribbleBean) StraightBeginPercent() float32 {
	return c.straightBeginPercent

}

func (c *FastDribbleBean) LoadData(dataStream *DataInputStream) {
	if dataStream != nil {
		c.id, _ = dataStream.ReadInt32()
		c.foot, _ = dataStream.ReadInt32()
		c.paramId, _ = dataStream.ReadUTF()
		c.animId, _ = dataStream.ReadInt32()
		c.animationName, _ = dataStream.ReadUTF()
		c.dribbleDistance, _ = dataStream.ReadFloat32()
		c.dribbleTime, _ = dataStream.ReadFloat32()
		c.isMirror, _ = dataStream.ReadInt32()
		c.straightBeginPercent, _ = dataStream.ReadFloat32()
	}
}
