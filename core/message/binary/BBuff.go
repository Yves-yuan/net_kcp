/*		自动生成, 请勿修改	*/
package binary

type BBuff struct {
	CBaseBinaryMessage
	m_buffId       int32 // buffId
	m_continueTime int32 // buff持续时间
}

func NewBBuff() *BBuff {
	return &BBuff{}
}

// 获取buffId
func (b *BBuff) GetBuffId() int32 {
	return b.m_buffId
}

// 设置buffId
func (b *BBuff) SetBuffId(v int32) {
	b.m_buffId = v
}

// 获取buff持续时间
func (b *BBuff) GetContinueTime() int32 {
	return b.m_continueTime
}

// 设置buff持续时间
func (b *BBuff) SetContinueTime(v int32) {
	b.m_continueTime = v
}

// 实现IBaseBinaryMessage接口
func (b *BBuff) Write(buf *CIOBuffer) {
	b.WriteInt32(buf, b.m_buffId)       // buffId
	b.WriteInt32(buf, b.m_continueTime) // buff持续时间
}
