package database

import (
	"github.com/GoHyperrr/mdk"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	mdk.RegisterDialect("postgres", func(dsn string) gorm.Dialector {
		return postgres.Open(dsn)
	})
}
