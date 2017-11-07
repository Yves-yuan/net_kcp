/*		自动生成, 请勿修改	*/
package binary

type BMover struct {
	CBaseBinaryMessage
	m_entity *BEntity // 防御者
	m_speed  int32    // 速度
}

func NewBMover() *BMover {
	return &BMover{}
}

// 获取防御者
func (b *BMover) GetEntity() *BEntity {
	return b.m_entity
}

// 设置防御者
func (b *BMover) SetEntity(v *BEntity) {
	b.m_entity = v
}

// 获取速度
func (b *BMover) GetSpeed() int32 {
	return b.m_speed
}

// 设置速度
func (b *BMover) SetSpeed(v int32) {
	b.m_speed = v
}

// 实现IBaseBinaryMessage接口
func (b *BMover) Write(buf *CIOBuffer) {
	b.WriteBean(buf, b.m_entity) // 防御者
	b.WriteInt32(buf, b.m_speed) // 速度
}
