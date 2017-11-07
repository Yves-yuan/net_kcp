/**
 * Auto generated, do not edit it
 *
 * Intercept
 */

package container

import (
	. "server/core/data/bean"
	. "server/core/datastream"
)

type InterceptContainer struct {
	list []InterceptBean
	maps map[int32]InterceptBean
}

func (c *InterceptContainer) LoadDataFromBin() {
	c.list = []InterceptBean{}
	c.maps = make(map[int32]InterceptBean)
	path := "bin/" + "InterceptBean" + ".bytes"
	dataStream := NewDataStream(path)
	if dataStream != nil {
		for dataStream.Available() {
			var bean InterceptBean
			bean.LoadData(dataStream)
			c.list = append(c.list, bean)
			c.maps[bean.Index()] = bean
		}
	} else {

	}
}

func (c *InterceptContainer) List() []InterceptBean {
	return c.list
}

func (c *InterceptContainer) GetBean(key int32) (*InterceptBean, bool) {
	val, isPresent := c.maps[key]
	return &val, isPresent
}
