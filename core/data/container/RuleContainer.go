/**
 * Auto generated, do not edit it
 *
 * Rule
 */

package container

import (
	. "server/core/data/bean"
	. "server/core/datastream"
)

type RuleContainer struct {
	list []RuleBean
	maps map[int32]RuleBean
}

func (c *RuleContainer) LoadDataFromBin() {
	c.list = []RuleBean{}
	c.maps = make(map[int32]RuleBean)
	path := "bin/" + "RuleBean" + ".bytes"
	dataStream := NewDataStream(path)
	if dataStream != nil {
		for dataStream.Available() {
			var bean RuleBean
			bean.LoadData(dataStream)
			c.list = append(c.list, bean)
			c.maps[bean.Id()] = bean
		}
	} else {

	}
}

func (c *RuleContainer) List() []RuleBean {
	return c.list
}

func (c *RuleContainer) GetBean(key int32) (*RuleBean, bool) {
	val, isPresent := c.maps[key]
	return &val, isPresent
}
