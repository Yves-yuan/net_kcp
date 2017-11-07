/**
 * Auto generated, do not edit it
 *
 * OpenBallPosition
 */

package container

import (
	. "server/core/data/bean"
	. "server/core/datastream"
)

type OpenBallPositionContainer struct {
	list []OpenBallPositionBean
	maps map[int32]OpenBallPositionBean
}

func (c *OpenBallPositionContainer) LoadDataFromBin() {
	c.list = []OpenBallPositionBean{}
	c.maps = make(map[int32]OpenBallPositionBean)
	path := "bin/" + "OpenBallPositionBean" + ".bytes"
	dataStream := NewDataStream(path)
	if dataStream != nil {
		for dataStream.Available() {
			var bean OpenBallPositionBean
			bean.LoadData(dataStream)
			c.list = append(c.list, bean)
			c.maps[bean.Id()] = bean
		}
	} else {

	}
}

func (c *OpenBallPositionContainer) List() []OpenBallPositionBean {
	return c.list
}

func (c *OpenBallPositionContainer) GetBean(key int32) (*OpenBallPositionBean, bool) {
	val, isPresent := c.maps[key]
	return &val, isPresent
}
