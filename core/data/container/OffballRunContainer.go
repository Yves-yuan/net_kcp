/**
 * Auto generated, do not edit it
 *
 * OffballRun
 */

package container

import (
	. "server/core/data/bean"
	. "server/core/datastream"
)

type OffballRunContainer struct {
	list []OffballRunBean
	maps map[int32]OffballRunBean
}

func (c *OffballRunContainer) LoadDataFromBin() {
	c.list = []OffballRunBean{}
	c.maps = make(map[int32]OffballRunBean)
	path := "bin/" + "OffballRunBean" + ".bytes"
	dataStream := NewDataStream(path)
	if dataStream != nil {
		for dataStream.Available() {
			var bean OffballRunBean
			bean.LoadData(dataStream)
			c.list = append(c.list, bean)
			c.maps[bean.Index()] = bean
		}
	} else {

	}
}

func (c *OffballRunContainer) List() []OffballRunBean {
	return c.list
}

func (c *OffballRunContainer) GetBean(key int32) (*OffballRunBean, bool) {
	val, isPresent := c.maps[key]
	return &val, isPresent
}
