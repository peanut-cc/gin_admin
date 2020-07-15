package bll

import (
	"context"

	"github.com/google/wire"
	"github.com/peanut-pg/gin_admin/internal/app/bll"
	"github.com/peanut-pg/gin_admin/internal/app/model"
	"github.com/peanut-pg/gin_admin/internal/app/schema"
	"github.com/peanut-pg/gin_admin/pkg/errors"
)

var _ bll.IUser = (*User)(nil)

// UserSet 注入User
var UserSet = wire.NewSet(wire.Struct(new(User), "*"), wire.Bind(new(bll.IUser), new(*User)))

// User 用户管理
type User struct {
	TransModel    model.ITrans
	UserModel     model.IUser
	UserRoleModel model.IUserRole
	RoleModel     model.IRole
}

// Query 查询数据
func (a *User) Query(ctx context.Context, params schema.UserQueryParam, opts ...schema.UserQueryOptions) (*schema.UserQueryResult, error) {
	return a.UserModel.Query(ctx, params, opts...)
}

// QueryShow 查询显示项数据
func (a *User) QueryShow(ctx context.Context, params schema.UserQueryParam, opts ...schema.UserQueryOptions) (*schema.UserShowQueryResult, error) {
	result, err := a.UserModel.Query(ctx, params, opts...)
	if err != nil {
		return nil, err
	} else if result == nil {
		return nil, nil
	}

	userRoleResult, err := a.UserRoleModel.Query(ctx, schema.UserRoleQueryParam{
		UserIDs: result.Data.ToIDs(),
	})
	if err != nil {
		return nil, err
	}

	roleResult, err := a.RoleModel.Query(ctx, schema.RoleQueryParam{
		IDs: userRoleResult.Data.ToRoleIDs(),
	})
	if err != nil {
		return nil, err
	}

	return result.ToShowResult(userRoleResult.Data.ToUserIDMap(), roleResult.Data.ToMap()), nil
}

// Get 查询指定数据
func (a *User) Get(ctx context.Context, id string, opts ...schema.UserQueryOptions) (*schema.User, error) {
	item, err := a.UserModel.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	userRoleResult, err := a.UserRoleModel.Query(ctx, schema.UserRoleQueryParam{
		UserID: id,
	})
	if err != nil {
		return nil, err
	}
	item.UserRoles = userRoleResult.Data

	return item, nil
}
