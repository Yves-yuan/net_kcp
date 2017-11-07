/**
 * Auto generated, do not edit it
 *
 * Scene
 */

package container

import (
	. "server/core/data/bean"
	. "server/core/datastream"
)

type SceneContainer struct {
	list []SceneBean
	maps map[int32]SceneBean
}

func (c *SceneContainer) LoadDataFromBin() {
	c.list = []SceneBean{}
	c.maps = make(map[int32]SceneBean)
	path := "bin/" + "SceneBean" + ".bytes"
	dataStream := NewDataStream(path)
	if dataStream != nil {
		for dataStream.Available() {
			var bean SceneBean
			bean.LoadData(dataStream)
			c.list = append(c.list, bean)
			c.maps[bean.Id()] = bean
		}
	} else {

	}
}

func (c *SceneContainer) List() []SceneBean {
	return c.list
}

func (c *SceneContainer) GetBean(key int32) (*SceneBean, bool) {
	val, isPresent := c.maps[key]
	return &val, isPresent
}
