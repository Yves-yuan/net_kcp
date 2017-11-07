/**
 * Auto generated, do not edit it
 *
 * PassBall
 */

package container

import (
	. "server/core/data/bean"
	. "server/core/datastream"
)

type PassBallContainer struct {
	list []PassBallBean
	maps map[string]PassBallBean
}

func (c *PassBallContainer) LoadDataFromBin() {
	c.list = []PassBallBean{}
	c.maps = make(map[string]PassBallBean)
	path := "bin/" + "PassBallBean" + ".bytes"
	dataStream := NewDataStream(path)
	if dataStream != nil {
		for dataStream.Available() {
			var bean PassBallBean
			bean.LoadData(dataStream)
			c.list = append(c.list, bean)
			c.maps[bean.CombKey()] = bean
		}
	} else {

	}
}

func (c *PassBallContainer) List() []PassBallBean {
	return c.list
}

func (c *PassBallContainer) GetBean(key string) (*PassBallBean, bool) {
	val, isPresent := c.maps[key]
	return &val, isPresent
}
