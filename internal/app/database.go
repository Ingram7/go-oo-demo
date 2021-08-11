package app

import (
	"go-oo-demo/pkg/orm"
	"gorm.io/gorm"
)

type Database struct {
	config *DBConfig
	mode *Mode

	orm *orm.Orm
}

func newDatabase(config *DBConfig, mode *Mode) *Database {
	database := new(Database)
	database.config = config
	database.mode = mode
	return database
}

func (database *Database) init() {
	database.orm = orm.New(database.config.Username, database.config.Password, database.config.Host, database.config.Name, database.config.MaxIdle, database.config.MaxOpen, database.mode.IsDebug())
}

func (database *Database) db() *gorm.DB {
	return database.orm.DB
}

func (database *Database) clear() {
	database.orm.Clear()
}
