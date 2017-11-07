/**
 * Auto generated, do not edit it
 *
 * PlayerAnimIDName
 */

package container

import (
	. "server/core/data/bean"
	. "server/core/datastream"
)

type PlayerAnimIDNameContainer struct {
	list []PlayerAnimIDNameBean
	maps map[int32]PlayerAnimIDNameBean
}

func (c *PlayerAnimIDNameContainer) LoadDataFromBin() {
	c.list = []PlayerAnimIDNameBean{}
	c.maps = make(map[int32]PlayerAnimIDNameBean)
	path := "bin/" + "PlayerAnimIDNameBean" + ".bytes"
	dataStream := NewDataStream(path)
	if dataStream != nil {
		for dataStream.Available() {
			var bean PlayerAnimIDNameBean
			bean.LoadData(dataStream)
			c.list = append(c.list, bean)
			c.maps[bean.Id()] = bean
		}
	} else {

	}
}

func (c *PlayerAnimIDNameContainer) List() []PlayerAnimIDNameBean {
	return c.list
}

func (c *PlayerAnimIDNameContainer) GetBean(key int32) (*PlayerAnimIDNameBean, bool) {
	val, isPresent := c.maps[key]
	return &val, isPresent
}
