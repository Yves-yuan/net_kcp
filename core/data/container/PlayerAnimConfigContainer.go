/**
 * Auto generated, do not edit it
 *
 * PlayerAnimConfig
 */

package container

import (
	. "server/core/data/bean"
	. "server/core/datastream"
)

type PlayerAnimConfigContainer struct {
	list []PlayerAnimConfigBean
	maps map[string]PlayerAnimConfigBean
}

func (c *PlayerAnimConfigContainer) LoadDataFromBin() {
	c.list = []PlayerAnimConfigBean{}
	c.maps = make(map[string]PlayerAnimConfigBean)
	path := "bin/" + "PlayerAnimConfigBean" + ".bytes"
	dataStream := NewDataStream(path)
	if dataStream != nil {
		for dataStream.Available() {
			var bean PlayerAnimConfigBean
			bean.LoadData(dataStream)
			c.list = append(c.list, bean)
			c.maps[bean.Name()] = bean
		}
	} else {

	}
}

func (c *PlayerAnimConfigContainer) List() []PlayerAnimConfigBean {
	return c.list
}

func (c *PlayerAnimConfigContainer) GetBean(key string) (*PlayerAnimConfigBean, bool) {
	val, isPresent := c.maps[key]
	return &val, isPresent
}
