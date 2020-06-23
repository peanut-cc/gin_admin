package iutil

import "github.com/peanut-pg/gin_admin/pkg/unique"

var idFunc = func() string {
	return unique.NewSnowflakeID().String()
}

func InitID() {
	idFunc = func() string {
		return unique.MustUUID().String()
	}
}

// NewID Create unique id
func NewID() string {
	return idFunc()
}
