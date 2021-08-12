package orm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

type Orm struct {
	DB *gorm.DB
}

func New(username string, password string, host string, name string, maxIdle int, maxOpen int, isDebug bool) *Orm {

	orm := new(Orm)

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username,
		password,
		host,
		name,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
		Logger: orm.getLogger(isDebug),
		PrepareStmt: true,
	})

	if err != nil {
		log.Fatalf("gorm init error: %s", err.Error())
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("gorm db error: %s", err.Error())
	}
	sqlDB.SetMaxIdleConns(maxIdle)           //  设置空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(maxOpen)          //  设置打开数据库连接的最大数量。
	sqlDB.SetConnMaxLifetime(time.Hour) // 设置了连接可复用的最大时间

	orm.DB = db
	return orm
}

func (o *Orm) getLogLevel(isDebug bool) logger.LogLevel {
	if isDebug {
		return logger.Info
	}

	return logger.Silent
}

func (o *Orm) getLogger(isDebug bool) logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: o.getSlowThreshold(),
			LogLevel:      o.getLogLevel(isDebug),
			Colorful:      true, // 禁用彩色打印
		},
	)
}

func (o *Orm) getSlowThreshold() time.Duration {
	return time.Second
}

func (o *Orm) Clear() {

	sqlDB, err := o.DB.DB()
	if err != nil {
		log.Printf("gorm db error: %s", err.Error())
	}

	if err := sqlDB.Close(); err != nil {
		log.Printf("gorm close error: %s", err.Error())
	}
}
