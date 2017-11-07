/**
 * Auto generated, do not edit it
 *
 * OffballStart
 */

package container

import (
	. "server/core/data/bean"
	. "server/core/datastream"
)

type OffballStartContainer struct {
	list []OffballStartBean
	maps map[int32]OffballStartBean
}

func (c *OffballStartContainer) LoadDataFromBin() {
	c.list = []OffballStartBean{}
	c.maps = make(map[int32]OffballStartBean)
	path := "bin/" + "OffballStartBean" + ".bytes"
	dataStream := NewDataStream(path)
	if dataStream != nil {
		for dataStream.Available() {
			var bean OffballStartBean
			bean.LoadData(dataStream)
			c.list = append(c.list, bean)
			c.maps[bean.Index()] = bean
		}
	} else {

	}
}

func (c *OffballStartContainer) List() []OffballStartBean {
	return c.list
}

func (c *OffballStartContainer) GetBean(key int32) (*OffballStartBean, bool) {
	val, isPresent := c.maps[key]
	return &val, isPresent
}
