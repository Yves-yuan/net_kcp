/**
 * Auto generated, do not edit it
 *
 * FastDribble
 */

package container

import (
	. "server/core/data/bean"
	. "server/core/datastream"
)

type FastDribbleContainer struct {
	list []FastDribbleBean
	maps map[string]FastDribbleBean
}

func (c *FastDribbleContainer) LoadDataFromBin() {
	c.list = []FastDribbleBean{}
	c.maps = make(map[string]FastDribbleBean)
	path := "bin/" + "FastDribbleBean" + ".bytes"
	dataStream := NewDataStream(path)
	if dataStream != nil {
		for dataStream.Available() {
			var bean FastDribbleBean
			bean.LoadData(dataStream)
			c.list = append(c.list, bean)
			c.maps[bean.ParamId()] = bean
		}
	} else {

	}
}

func (c *FastDribbleContainer) List() []FastDribbleBean {
	return c.list
}

func (c *FastDribbleContainer) GetBean(key string) (*FastDribbleBean, bool) {
	val, isPresent := c.maps[key]
	return &val, isPresent
}
