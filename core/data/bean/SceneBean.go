/**
 * Auto generated, do not edit it
 *
 * Scene
 */
package bean

import (
	. "server/core/datastream"
)

type SceneBean struct {
	id                         int32   // 场景id
	name                       string  // 名字
	accelerationOfGravity      float32 // 重力加速度
	verticalSpeedDecayFactor   float32 // 发生碰撞时，垂直方向速度衰减系数
	horizontalSpeedDecayFactor float32 // 发生碰撞时，水平方向速度衰减系数
	verticalMinSpeed           float32 // 垂直方向上发生反弹的最小速度
	frictionAcceleration       float32 // 球在地面滚动时球的减速度
	topZ                       float32 // 绿草地上边界z
	bottomZ                    float32 // 绿草地下边界z
	leftX                      float32 // 绿草地左边界x
	rightX                     float32 // 绿草地右边界x
	edgeTopZ                   float32 // 上边线界z
	edgeBottomZ                float32 // 下边线界z
	edgeLeftX                  float32 // 左边线界x
	edgeRightX                 float32 // 右边线界x
	goalBottomZ                float32 // 球门下边界z
	goalTopZ                   float32 // 球门上边界
	goalHeight                 float32 // 球门高
}

func (c *SceneBean) Id() int32 {
	return c.id

}
func (c *SceneBean) Name() string {
	return c.name

}
func (c *SceneBean) AccelerationOfGravity() float32 {
	return c.accelerationOfGravity

}
func (c *SceneBean) VerticalSpeedDecayFactor() float32 {
	return c.verticalSpeedDecayFactor

}
func (c *SceneBean) HorizontalSpeedDecayFactor() float32 {
	return c.horizontalSpeedDecayFactor

}
func (c *SceneBean) VerticalMinSpeed() float32 {
	return c.verticalMinSpeed

}
func (c *SceneBean) FrictionAcceleration() float32 {
	return c.frictionAcceleration

}
func (c *SceneBean) TopZ() float32 {
	return c.topZ

}
func (c *SceneBean) BottomZ() float32 {
	return c.bottomZ

}
func (c *SceneBean) LeftX() float32 {
	return c.leftX

}
func (c *SceneBean) RightX() float32 {
	return c.rightX

}
func (c *SceneBean) EdgeTopZ() float32 {
	return c.edgeTopZ

}
func (c *SceneBean) EdgeBottomZ() float32 {
	return c.edgeBottomZ

}
func (c *SceneBean) EdgeLeftX() float32 {
	return c.edgeLeftX

}
func (c *SceneBean) EdgeRightX() float32 {
	return c.edgeRightX

}
func (c *SceneBean) GoalBottomZ() float32 {
	return c.goalBottomZ

}
func (c *SceneBean) GoalTopZ() float32 {
	return c.goalTopZ

}
func (c *SceneBean) GoalHeight() float32 {
	return c.goalHeight

}

func (c *SceneBean) LoadData(dataStream *DataInputStream) {
	if dataStream != nil {
		c.id, _ = dataStream.ReadInt32()
		c.name, _ = dataStream.ReadUTF()
		c.accelerationOfGravity, _ = dataStream.ReadFloat32()
		c.verticalSpeedDecayFactor, _ = dataStream.ReadFloat32()
		c.horizontalSpeedDecayFactor, _ = dataStream.ReadFloat32()
		c.verticalMinSpeed, _ = dataStream.ReadFloat32()
		c.frictionAcceleration, _ = dataStream.ReadFloat32()
		c.topZ, _ = dataStream.ReadFloat32()
		c.bottomZ, _ = dataStream.ReadFloat32()
		c.leftX, _ = dataStream.ReadFloat32()
		c.rightX, _ = dataStream.ReadFloat32()
		c.edgeTopZ, _ = dataStream.ReadFloat32()
		c.edgeBottomZ, _ = dataStream.ReadFloat32()
		c.edgeLeftX, _ = dataStream.ReadFloat32()
		c.edgeRightX, _ = dataStream.ReadFloat32()
		c.goalBottomZ, _ = dataStream.ReadFloat32()
		c.goalTopZ, _ = dataStream.ReadFloat32()
		c.goalHeight, _ = dataStream.ReadFloat32()
	}
}
