/**
 * Auto generated, do not edit it
 *
 * Startdribble
 */

package container

import (
	. "server/core/data/bean"
	. "server/core/datastream"
)

type StartdribbleContainer struct {
	list []StartdribbleBean
	maps map[string]StartdribbleBean
}

func (c *StartdribbleContainer) LoadDataFromBin() {
	c.list = []StartdribbleBean{}
	c.maps = make(map[string]StartdribbleBean)
	path := "bin/" + "StartdribbleBean" + ".bytes"
	dataStream := NewDataStream(path)
	if dataStream != nil {
		for dataStream.Available() {
			var bean StartdribbleBean
			bean.LoadData(dataStream)
			c.list = append(c.list, bean)
			c.maps[bean.CombKey()] = bean
		}
	} else {

	}
}

func (c *StartdribbleContainer) List() []StartdribbleBean {
	return c.list
}

func (c *StartdribbleContainer) GetBean(key string) (*StartdribbleBean, bool) {
	val, isPresent := c.maps[key]
	return &val, isPresent
}
