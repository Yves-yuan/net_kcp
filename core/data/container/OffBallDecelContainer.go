/**
 * Auto generated, do not edit it
 *
 * OffBallDecel
 */

package container

import (
	. "server/core/data/bean"
	. "server/core/datastream"
)

type OffBallDecelContainer struct {
	list []OffBallDecelBean
	maps map[int32]OffBallDecelBean
}

func (c *OffBallDecelContainer) LoadDataFromBin() {
	c.list = []OffBallDecelBean{}
	c.maps = make(map[int32]OffBallDecelBean)
	path := "bin/" + "OffBallDecelBean" + ".bytes"
	dataStream := NewDataStream(path)
	if dataStream != nil {
		for dataStream.Available() {
			var bean OffBallDecelBean
			bean.LoadData(dataStream)
			c.list = append(c.list, bean)
			c.maps[bean.Index()] = bean
		}
	} else {

	}
}

func (c *OffBallDecelContainer) List() []OffBallDecelBean {
	return c.list
}

func (c *OffBallDecelContainer) GetBean(key int32) (*OffBallDecelBean, bool) {
	val, isPresent := c.maps[key]
	return &val, isPresent
}
