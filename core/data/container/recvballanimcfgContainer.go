/**
 * Auto generated, do not edit it
 *
 * RecvBallAnimCfg
 */

package container

import (
	. "server/core/data/bean"
	. "server/core/datastream"
)

type RecvBallAnimCfgContainer struct {
	list []RecvBallAnimCfgBean
	maps map[int32]RecvBallAnimCfgBean
}

func (c *RecvBallAnimCfgContainer) LoadDataFromBin() {
	c.list = []RecvBallAnimCfgBean{}
	c.maps = make(map[int32]RecvBallAnimCfgBean)
	path := "bin/" + "RecvBallAnimCfgBean" + ".bytes"
	dataStream := NewDataStream(path)
	if dataStream != nil {
		for dataStream.Available() {
			var bean RecvBallAnimCfgBean
			bean.LoadData(dataStream)
			c.list = append(c.list, bean)
			c.maps[bean.Index()] = bean
		}
	} else {

	}
}

func (c *RecvBallAnimCfgContainer) List() []RecvBallAnimCfgBean {
	return c.list
}

func (c *RecvBallAnimCfgContainer) GetBean(key int32) (*RecvBallAnimCfgBean, bool) {
	val, isPresent := c.maps[key]
	return &val, isPresent
}
