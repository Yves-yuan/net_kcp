/**
 * Auto generated, do not edit it
 *
 * CatchBall
 */

package container

import (
	. "server/core/data/bean"
	. "server/core/datastream"
)

type CatchBallContainer struct {
	list []CatchBallBean
	maps map[int32]CatchBallBean
}

func (c *CatchBallContainer) LoadDataFromBin() {
	c.list = []CatchBallBean{}
	c.maps = make(map[int32]CatchBallBean)
	path := "bin/" + "CatchBallBean" + ".bytes"
	dataStream := NewDataStream(path)
	if dataStream != nil {
		for dataStream.Available() {
			var bean CatchBallBean
			bean.LoadData(dataStream)
			c.list = append(c.list, bean)
			c.maps[bean.Id()] = bean
		}
	} else {

	}
}

func (c *CatchBallContainer) List() []CatchBallBean {
	return c.list
}

func (c *CatchBallContainer) GetBean(key int32) (*CatchBallBean, bool) {
	val, isPresent := c.maps[key]
	return &val, isPresent
}
