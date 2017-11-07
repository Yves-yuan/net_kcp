/*		自动生成, 请勿修改	*/
package binary

import "server/cmd/battle/position"

type BEntity struct {
	CBaseBinaryMessage
	m_entityId  int32            // 实体Id
	m_state     int32            // 实体状态
	m_position  *position.CVec3f // 当前位置
	m_direction *position.CVec3f // 方向
}

func NewBEntity() *BEntity {
	return &BEntity{}
}

// 获取实体Id
func (b *BEntity) GetEntityId() int32 {
	return b.m_entityId
}

// 设置实体Id
func (b *BEntity) SetEntityId(v int32) {
	b.m_entityId = v
}

// 获取实体状态
func (b *BEntity) GetState() int32 {
	return b.m_state
}

// 设置实体状态
func (b *BEntity) SetState(v int32) {
	b.m_state = v
}

// 获取当前位置
func (b *BEntity) GetPosition() *position.CVec3f {
	return b.m_position
}

// 设置当前位置
func (b *BEntity) SetPosition(v *position.CVec3f) {
	b.m_position = v
}

// 获取方向
func (b *BEntity) GetDirection() *position.CVec3f {
	return b.m_direction
}

// 设置方向
func (b *BEntity) SetDirection(v *position.CVec3f) {
	b.m_direction = v
}

// 实现IBaseBinaryMessage接口
func (b *BEntity) Write(buf *CIOBuffer) {
	b.WriteInt32(buf, b.m_entityId)  // 实体Id
	b.WriteInt32(buf, b.m_state)     // 实体状态
	b.WriteVec3f(buf, b.m_position)  // 当前位置
	b.WriteVec3f(buf, b.m_direction) // 方向
}
