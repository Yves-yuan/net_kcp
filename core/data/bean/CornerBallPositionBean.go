/**
 * Auto generated, do not edit it
 *
 * CornerBallPosition
 */
package bean

import (
	. "server/core/datastream"
)

type CornerBallPositionBean struct {
	id      int32   // 规则ID
	comment string  // 备注
	offsetX float32 // X偏移
	offsetZ float32 // Z偏移
}

func (c *CornerBallPositionBean) Id() int32 {
	return c.id

}
func (c *CornerBallPositionBean) Comment() string {
	return c.comment

}
func (c *CornerBallPositionBean) OffsetX() float32 {
	return c.offsetX

}
func (c *CornerBallPositionBean) OffsetZ() float32 {
	return c.offsetZ

}

func (c *CornerBallPositionBean) LoadData(dataStream *DataInputStream) {
	if dataStream != nil {
		c.id, _ = dataStream.ReadInt32()
		c.comment, _ = dataStream.ReadUTF()
		c.offsetX, _ = dataStream.ReadFloat32()
		c.offsetZ, _ = dataStream.ReadFloat32()
	}
}
