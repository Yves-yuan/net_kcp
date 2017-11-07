/**
 * Auto generated, do not edit it
 *
 * Sliding
 */

package container

import (
	. "server/core/data/bean"
	. "server/core/datastream"
)

type SlidingContainer struct {
	list []SlidingBean
	maps map[int32]SlidingBean
}

func (c *SlidingContainer) LoadDataFromBin() {
	c.list = []SlidingBean{}
	c.maps = make(map[int32]SlidingBean)
	path := "bin/" + "SlidingBean" + ".bytes"
	dataStream := NewDataStream(path)
	if dataStream != nil {
		for dataStream.Available() {
			var bean SlidingBean
			bean.LoadData(dataStream)
			c.list = append(c.list, bean)
			c.maps[bean.Index()] = bean
		}
	} else {

	}
}

func (c *SlidingContainer) List() []SlidingBean {
	return c.list
}

func (c *SlidingContainer) GetBean(key int32) (*SlidingBean, bool) {
	val, isPresent := c.maps[key]
	return &val, isPresent
}
