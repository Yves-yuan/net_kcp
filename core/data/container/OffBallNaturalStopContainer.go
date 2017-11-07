/**
 * Auto generated, do not edit it
 *
 * OffBallNaturalStop
 */

package container

import (
	. "server/core/data/bean"
	. "server/core/datastream"
)

type OffBallNaturalStopContainer struct {
	list []OffBallNaturalStopBean
	maps map[int32]OffBallNaturalStopBean
}

func (c *OffBallNaturalStopContainer) LoadDataFromBin() {
	c.list = []OffBallNaturalStopBean{}
	c.maps = make(map[int32]OffBallNaturalStopBean)
	path := "bin/" + "OffBallNaturalStopBean" + ".bytes"
	dataStream := NewDataStream(path)
	if dataStream != nil {
		for dataStream.Available() {
			var bean OffBallNaturalStopBean
			bean.LoadData(dataStream)
			c.list = append(c.list, bean)
			c.maps[bean.Index()] = bean
		}
	} else {

	}
}

func (c *OffBallNaturalStopContainer) List() []OffBallNaturalStopBean {
	return c.list
}

func (c *OffBallNaturalStopContainer) GetBean(key int32) (*OffBallNaturalStopBean, bool) {
	val, isPresent := c.maps[key]
	return &val, isPresent
}
