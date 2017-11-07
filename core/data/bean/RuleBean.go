/**
 * Auto generated, do not edit it
 *
 * Rule
 */
package bean

import (
	. "server/core/datastream"
)

type RuleBean struct {
	id                     int32   // 规则ID
	name                   string  // 名字
	halfTime               int32   // 半场时间(单位秒)
	minAddtionTime         int32   // 最低补时时间(秒)
	singleAddtionTime      int32   // 单次开球补时时间(秒)
	waitingTime            float32 // 动画等待时间(单位秒,及球出界后, 等待多少秒再播放动画)
	autoPassTime           float32 // 自动传球时间(发球时, 如果该时间内没有发球, 系统自动发球)
	cornerPlayerX          float32 // 角球发球时,
	cornerPlayerZ          float32 // 角球发球时, 人距离顶点的Z偏移值
	cornerBallX            float32 // 角球发球时, 球距离顶点的X偏移值
	cornerBallZ            float32 // 角球发球时, 球距离顶点的Z偏移值
	cornerObstacleRadius   float32 // 角球阻挡半径
	edgeBallZOffset        float32 // 边线球发球球Z轴偏移值
	edgeBallDefenderX      float32 // 边线球防守方球员X
	edgeBallDefenderZ      float32 // 边线球防守方球员Z
	edgeBallOffenseX       float32 // 边线球进攻方球员X
	edgeBallOffenseZ       float32 // 边线球进攻方球员Z
	edgeBallObstacleRadius float32 // 边线球阻挡半径
}

func (c *RuleBean) Id() int32 {
	return c.id

}
func (c *RuleBean) Name() string {
	return c.name

}
func (c *RuleBean) HalfTime() int32 {
	return c.halfTime

}
func (c *RuleBean) MinAddtionTime() int32 {
	return c.minAddtionTime

}
func (c *RuleBean) SingleAddtionTime() int32 {
	return c.singleAddtionTime

}
func (c *RuleBean) WaitingTime() float32 {
	return c.waitingTime

}
func (c *RuleBean) AutoPassTime() float32 {
	return c.autoPassTime

}
func (c *RuleBean) CornerPlayerX() float32 {
	return c.cornerPlayerX

}
func (c *RuleBean) CornerPlayerZ() float32 {
	return c.cornerPlayerZ

}
func (c *RuleBean) CornerBallX() float32 {
	return c.cornerBallX

}
func (c *RuleBean) CornerBallZ() float32 {
	return c.cornerBallZ

}
func (c *RuleBean) CornerObstacleRadius() float32 {
	return c.cornerObstacleRadius

}
func (c *RuleBean) EdgeBallZOffset() float32 {
	return c.edgeBallZOffset

}
func (c *RuleBean) EdgeBallDefenderX() float32 {
	return c.edgeBallDefenderX

}
func (c *RuleBean) EdgeBallDefenderZ() float32 {
	return c.edgeBallDefenderZ

}
func (c *RuleBean) EdgeBallOffenseX() float32 {
	return c.edgeBallOffenseX

}
func (c *RuleBean) EdgeBallOffenseZ() float32 {
	return c.edgeBallOffenseZ

}
func (c *RuleBean) EdgeBallObstacleRadius() float32 {
	return c.edgeBallObstacleRadius

}

func (c *RuleBean) LoadData(dataStream *DataInputStream) {
	if dataStream != nil {
		c.id, _ = dataStream.ReadInt32()
		c.name, _ = dataStream.ReadUTF()
		c.halfTime, _ = dataStream.ReadInt32()
		c.minAddtionTime, _ = dataStream.ReadInt32()
		c.singleAddtionTime, _ = dataStream.ReadInt32()
		c.waitingTime, _ = dataStream.ReadFloat32()
		c.autoPassTime, _ = dataStream.ReadFloat32()
		c.cornerPlayerX, _ = dataStream.ReadFloat32()
		c.cornerPlayerZ, _ = dataStream.ReadFloat32()
		c.cornerBallX, _ = dataStream.ReadFloat32()
		c.cornerBallZ, _ = dataStream.ReadFloat32()
		c.cornerObstacleRadius, _ = dataStream.ReadFloat32()
		c.edgeBallZOffset, _ = dataStream.ReadFloat32()
		c.edgeBallDefenderX, _ = dataStream.ReadFloat32()
		c.edgeBallDefenderZ, _ = dataStream.ReadFloat32()
		c.edgeBallOffenseX, _ = dataStream.ReadFloat32()
		c.edgeBallOffenseZ, _ = dataStream.ReadFloat32()
		c.edgeBallObstacleRadius, _ = dataStream.ReadFloat32()
	}
}
