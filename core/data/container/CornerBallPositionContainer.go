/**
 * Auto generated, do not edit it
 *
 * CornerBallPosition
 */

package container

import (
	. "server/core/data/bean"
	. "server/core/datastream"
)

type CornerBallPositionContainer struct {
	list []CornerBallPositionBean
	maps map[int32]CornerBallPositionBean
}

func (c *CornerBallPositionContainer) LoadDataFromBin() {
	c.list = []CornerBallPositionBean{}
	c.maps = make(map[int32]CornerBallPositionBean)
	path := "bin/" + "CornerBallPositionBean" + ".bytes"
	dataStream := NewDataStream(path)
	if dataStream != nil {
		for dataStream.Available() {
			var bean CornerBallPositionBean
			bean.LoadData(dataStream)
			c.list = append(c.list, bean)
			c.maps[bean.Id()] = bean
		}
	} else {

	}
}

func (c *CornerBallPositionContainer) List() []CornerBallPositionBean {
	return c.list
}

func (c *CornerBallPositionContainer) GetBean(key int32) (*CornerBallPositionBean, bool) {
	val, isPresent := c.maps[key]
	return &val, isPresent
}
