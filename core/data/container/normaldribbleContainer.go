/**
 * Auto generated, do not edit it
 *
 * NormalDribble
 */

package container

import (
	. "server/core/data/bean"
	. "server/core/datastream"
)

type NormalDribbleContainer struct {
	list []NormalDribbleBean
	maps map[string]NormalDribbleBean
}

func (c *NormalDribbleContainer) LoadDataFromBin() {
	c.list = []NormalDribbleBean{}
	c.maps = make(map[string]NormalDribbleBean)
	path := "bin/" + "NormalDribbleBean" + ".bytes"
	dataStream := NewDataStream(path)
	if dataStream != nil {
		for dataStream.Available() {
			var bean NormalDribbleBean
			bean.LoadData(dataStream)
			c.list = append(c.list, bean)
			c.maps[bean.ParamId()] = bean
		}
	} else {

	}
}

func (c *NormalDribbleContainer) List() []NormalDribbleBean {
	return c.list
}

func (c *NormalDribbleContainer) GetBean(key string) (*NormalDribbleBean, bool) {
	val, isPresent := c.maps[key]
	return &val, isPresent
}
