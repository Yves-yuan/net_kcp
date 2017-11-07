/**
 * Auto generated, do not edit it
 *
 * OffballAcc
 */

package container

import (
	. "server/core/data/bean"
	. "server/core/datastream"
)

type OffballAccContainer struct {
	list []OffballAccBean
	maps map[int32]OffballAccBean
}

func (c *OffballAccContainer) LoadDataFromBin() {
	c.list = []OffballAccBean{}
	c.maps = make(map[int32]OffballAccBean)
	path := "bin/" + "OffballAccBean" + ".bytes"
	dataStream := NewDataStream(path)
	if dataStream != nil {
		for dataStream.Available() {
			var bean OffballAccBean
			bean.LoadData(dataStream)
			c.list = append(c.list, bean)
			c.maps[bean.Index()] = bean
		}
	} else {

	}
}

func (c *OffballAccContainer) List() []OffballAccBean {
	return c.list
}

func (c *OffballAccContainer) GetBean(key int32) (*OffballAccBean, bool) {
	val, isPresent := c.maps[key]
	return &val, isPresent
}
