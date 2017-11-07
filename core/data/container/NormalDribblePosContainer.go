/**
 * Auto generated, do not edit it
 *
 * NormalDribblePos
 */

package container

import (
	. "server/core/data/bean"
	. "server/core/datastream"
)

type NormalDribblePosContainer struct {
	list []NormalDribblePosBean
	maps map[string]NormalDribblePosBean
}

func (c *NormalDribblePosContainer) LoadDataFromBin() {
	c.list = []NormalDribblePosBean{}
	c.maps = make(map[string]NormalDribblePosBean)
	path := "bin/" + "NormalDribblePosBean" + ".bytes"
	dataStream := NewDataStream(path)
	if dataStream != nil {
		for dataStream.Available() {
			var bean NormalDribblePosBean
			bean.LoadData(dataStream)
			c.list = append(c.list, bean)
			c.maps[bean.AnimIndex()] = bean
		}
	} else {

	}
}

func (c *NormalDribblePosContainer) List() []NormalDribblePosBean {
	return c.list
}

func (c *NormalDribblePosContainer) GetBean(key string) (*NormalDribblePosBean, bool) {
	val, isPresent := c.maps[key]
	return &val, isPresent
}
