/**
 * Auto generated, do not edit it
 *
 * Offballfaststop
 */

package container

import (
	. "server/core/data/bean"
	. "server/core/datastream"
)

type OffballfaststopContainer struct {
	list []OffballfaststopBean
	maps map[int32]OffballfaststopBean
}

func (c *OffballfaststopContainer) LoadDataFromBin() {
	c.list = []OffballfaststopBean{}
	c.maps = make(map[int32]OffballfaststopBean)
	path := "bin/" + "OffballfaststopBean" + ".bytes"
	dataStream := NewDataStream(path)
	if dataStream != nil {
		for dataStream.Available() {
			var bean OffballfaststopBean
			bean.LoadData(dataStream)
			c.list = append(c.list, bean)
			c.maps[bean.Index()] = bean
		}
	} else {

	}
}

func (c *OffballfaststopContainer) List() []OffballfaststopBean {
	return c.list
}

func (c *OffballfaststopContainer) GetBean(key int32) (*OffballfaststopBean, bool) {
	val, isPresent := c.maps[key]
	return &val, isPresent
}
