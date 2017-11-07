/**
 * Auto generated, do not edit it
 *
 * RecvBallAnimCfg
 */
package bean

import (
	. "server/core/datastream"
)

type RecvBallAnimCfgBean struct {
	index     int32   // 动画索引
	part      int32   // 接球部位，0脚部，1胸部
	direction int32   // 接球方向，0前，1左，2右
	move      int32   // 是否移动
	high      int32   // 是否高空球
	animID    int32   // 动作ID
	animName  string  // 动作名称
	touchTime float32 // 触球时间（毫秒）
	mirror    int32   // 是否镜像
}

func (c *RecvBallAnimCfgBean) Index() int32 {
	return c.index

}
func (c *RecvBallAnimCfgBean) Part() int32 {
	return c.part

}
func (c *RecvBallAnimCfgBean) Direction() int32 {
	return c.direction

}
func (c *RecvBallAnimCfgBean) Move() int32 {
	return c.move

}
func (c *RecvBallAnimCfgBean) High() int32 {
	return c.high

}
func (c *RecvBallAnimCfgBean) AnimID() int32 {
	return c.animID

}
func (c *RecvBallAnimCfgBean) AnimName() string {
	return c.animName

}
func (c *RecvBallAnimCfgBean) TouchTime() float32 {
	return c.touchTime

}
func (c *RecvBallAnimCfgBean) Mirror() int32 {
	return c.mirror

}

func (c *RecvBallAnimCfgBean) LoadData(dataStream *DataInputStream) {
	if dataStream != nil {
		c.index, _ = dataStream.ReadInt32()
		c.part, _ = dataStream.ReadInt32()
		c.direction, _ = dataStream.ReadInt32()
		c.move, _ = dataStream.ReadInt32()
		c.high, _ = dataStream.ReadInt32()
		c.animID, _ = dataStream.ReadInt32()
		c.animName, _ = dataStream.ReadUTF()
		c.touchTime, _ = dataStream.ReadFloat32()
		c.mirror, _ = dataStream.ReadInt32()
	}
}
