package entity

// RoleMenu 角色菜单实体
type RoleMenu struct {
	Model
	RoleID   string `gorm:"column:role_id;size:36;index;default:'';not null;"`   // 角色ID
	MenuID   string `gorm:"column:menu_id;size:36;index;default:'';not null;"`   // 菜单ID
	ActionID string `gorm:"column:action_id;size:36;index;default:'';not null;"` // 动作ID
}

// TableName 表名
func (a RoleMenu) TableName() string {
	return a.Model.TableName("role_menu")
}
