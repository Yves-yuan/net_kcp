/*		自动生成, 请勿修改	*/
package binary

type BSkill struct {
	CBaseBinaryMessage
	m_skillId string // 技能Id
}

func NewBSkill() *BSkill {
	return &BSkill{}
}

// 获取技能Id
func (b *BSkill) GetSkillId() string {
	return b.m_skillId
}

// 设置技能Id
func (b *BSkill) SetSkillId(v string) {
	b.m_skillId = v
}

// 实现IBaseBinaryMessage接口
func (b *BSkill) Write(buf *CIOBuffer) {
	b.WriteString(buf, b.m_skillId) // 技能Id
}
