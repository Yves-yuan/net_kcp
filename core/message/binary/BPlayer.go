/*		自动生成, 请勿修改	*/
package binary

type BPlayer struct {
	CBaseBinaryMessage
	m_type       int16   // 类型
	m_mover      *BMover // 移动者
	m_profession int32   // 职业
	m_roleName   string  // 角色名字
	m_roleId     string  // 角色Id
	m_vipLevel   int32   // vip等级
	m_gender     int32   // 性别
}

func NewBPlayer() *BPlayer {
	return &BPlayer{}
}

// 获取类型
func (b *BPlayer) GetType() int16 {
	return b.m_type
}

// 设置类型
func (b *BPlayer) SetType(v int16) {
	b.m_type = v
}

// 获取移动者
func (b *BPlayer) GetMover() *BMover {
	return b.m_mover
}

// 设置移动者
func (b *BPlayer) SetMover(v *BMover) {
	b.m_mover = v
}

// 获取职业
func (b *BPlayer) GetProfession() int32 {
	return b.m_profession
}

// 设置职业
func (b *BPlayer) SetProfession(v int32) {
	b.m_profession = v
}

// 获取角色名字
func (b *BPlayer) GetRoleName() string {
	return b.m_roleName
}

// 设置角色名字
func (b *BPlayer) SetRoleName(v string) {
	b.m_roleName = v
}

// 获取角色Id
func (b *BPlayer) GetRoleId() string {
	return b.m_roleId
}

// 设置角色Id
func (b *BPlayer) SetRoleId(v string) {
	b.m_roleId = v
}

// 获取vip等级
func (b *BPlayer) GetVipLevel() int32 {
	return b.m_vipLevel
}

// 设置vip等级
func (b *BPlayer) SetVipLevel(v int32) {
	b.m_vipLevel = v
}

// 获取性别
func (b *BPlayer) GetGender() int32 {
	return b.m_gender
}

// 设置性别
func (b *BPlayer) SetGender(v int32) {
	b.m_gender = v
}

// 实现IBaseBinaryMessage接口
func (b *BPlayer) Write(buf *CIOBuffer) {
	b.WriteInt16(buf, b.m_type)       // 类型
	b.WriteBean(buf, b.m_mover)       // 移动者
	b.WriteInt32(buf, b.m_profession) // 职业
	b.WriteString(buf, b.m_roleName)  // 角色名字
	b.WriteString(buf, b.m_roleId)    // 角色Id
	b.WriteInt32(buf, b.m_vipLevel)   // vip等级
	b.WriteInt32(buf, b.m_gender)     // 性别
}
