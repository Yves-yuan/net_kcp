/**
 * Auto generated, do not edit it
 *
 * Intercept
 */
package bean

import (
	. "server/core/datastream"
)

type InterceptBean struct {
	index          int32   // 动画索引
	screen         int32   // 筛选条件
	animID         int32   // 动作ID
	animName       string  // 动作名称
	touchTime      float32 // 触球时间（毫秒）
	mirror         int32   // 是否镜像
	robDistance    float32 // 铲球距离
	interceptSpeed float32 // 铲断后球的速度
}

func (c *InterceptBean) Index() int32 {
	return c.index

}
func (c *InterceptBean) Screen() int32 {
	return c.screen

}
func (c *InterceptBean) AnimID() int32 {
	return c.animID

}
func (c *InterceptBean) AnimName() string {
	return c.animName

}
func (c *InterceptBean) TouchTime() float32 {
	return c.touchTime

}
func (c *InterceptBean) Mirror() int32 {
	return c.mirror

}
func (c *InterceptBean) RobDistance() float32 {
	return c.robDistance

}
func (c *InterceptBean) InterceptSpeed() float32 {
	return c.interceptSpeed

}

func (c *InterceptBean) LoadData(dataStream *DataInputStream) {
	if dataStream != nil {
		c.index, _ = dataStream.ReadInt32()
		c.screen, _ = dataStream.ReadInt32()
		c.animID, _ = dataStream.ReadInt32()
		c.animName, _ = dataStream.ReadUTF()
		c.touchTime, _ = dataStream.ReadFloat32()
		c.mirror, _ = dataStream.ReadInt32()
		c.robDistance, _ = dataStream.ReadFloat32()
		c.interceptSpeed, _ = dataStream.ReadFloat32()
	}
}
