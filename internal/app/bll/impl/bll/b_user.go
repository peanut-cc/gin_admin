package bll

import (
	"github.com/peanut-pg/gin_admin/internal/app/model"
)

// User 用户管理
type User struct {
	TransModel model.ITrans
	UserModel  model.IUser
	//UserRoleModel model.IUserRole
	//RoleModel     model.IRole
}
