/**
 * Auto generated, do not edit it
 *
 * PlayerAnimIDName
 */
package bean

import (
	. "server/core/datastream"
)

type PlayerAnimIDNameBean struct {
	id   int32  // 动画ID
	name string // 动画名称
}

func (c *PlayerAnimIDNameBean) Id() int32 {
	return c.id

}
func (c *PlayerAnimIDNameBean) Name() string {
	return c.name

}

func (c *PlayerAnimIDNameBean) LoadData(dataStream *DataInputStream) {
	if dataStream != nil {
		c.id, _ = dataStream.ReadInt32()
		c.name, _ = dataStream.ReadUTF()
	}
}
