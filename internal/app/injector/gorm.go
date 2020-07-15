package injector

import (
	"github.com/jinzhu/gorm"
	"github.com/peanut-pg/gin_admin/internal/app/config"
)

func NewGormDB() (*gorm.DB, func(), error) {
	cfg := config.C

}
