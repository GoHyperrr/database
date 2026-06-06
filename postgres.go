package database

import (
	"github.com/GoHyperrr/hyperrr/pkg/db"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	db.RegisterDialect("postgres", func(dsn string) gorm.Dialector {
		return postgres.Open(dsn)
	})
}
