/**
 * Auto generated, do not edit it
 *
 * GlobalConstant
 */

package container

import (
	. "server/core/data/bean"
	. "server/core/datastream"
)

type GlobalConstantContainer struct {
	list []GlobalConstantBean
	maps map[int32]GlobalConstantBean
}

func (c *GlobalConstantContainer) LoadDataFromBin() {
	c.list = []GlobalConstantBean{}
	c.maps = make(map[int32]GlobalConstantBean)
	path := "bin/" + "GlobalConstantBean" + ".bytes"
	dataStream := NewDataStream(path)
	if dataStream != nil {
		for dataStream.Available() {
			var bean GlobalConstantBean
			bean.LoadData(dataStream)
			c.list = append(c.list, bean)
			c.maps[bean.Id()] = bean
		}
	} else {

	}
}

func (c *GlobalConstantContainer) List() []GlobalConstantBean {
	return c.list
}

func (c *GlobalConstantContainer) GetBean(key int32) (*GlobalConstantBean, bool) {
	val, isPresent := c.maps[key]
	return &val, isPresent
}
