package bll

import (
	"github.com/peanut-pg/gin_admin/internal/app/model"
)

// Role 角色管理
type Role struct {
	//Enforcer      *casbin.SyncedEnforcer
	TransModel model.ITrans
	RoleModel  model.IRole
	//RoleMenuModel model.IRoleMenu
	UserModel model.IUser
}
