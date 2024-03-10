package models

import (
	"fmt"
	"github.com/techoc/fanqie-novel-api/pkg/global"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

var (
	db  *gorm.DB
	err error
)

// Setup initializes the database instance
func Setup() {
	switch global.DatabaseConf.Type {
	case "mysql":
		initMysqlDB()
	case "sqlite":
		initSqliteDB()
	default:
		log.Fatalf("models.Setup err,please set database type")
	}

	// 自动迁移
	err = db.AutoMigrate(&Author{}, &Book{}, &Chapter{}, &Label{})
	if err != nil {
		return
	}
}

func initDB() {
	switch global.DatabaseConf.Type {
	case "sqlite":
		log.Println("数据库类型为sqlite")
		initSqliteDB()

	case "mysql":
		log.Println("数据库类型为mysql")
		initMysqlDB()
	}
}

func initMysqlDB() {

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		global.DatabaseConf.User,
		global.DatabaseConf.Password,
		global.DatabaseConf.Host,
		global.DatabaseConf.Name,
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   global.DatabaseConf.TablePrefix,
			SingularTable: true,
		},
	})

	if err != nil {
		log.Fatalf("mysql init err: %v", err)
	}
}

func initSqliteDB() {
	db, err = gorm.Open(sqlite.Open(global.DatabaseConf.Name+".db"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   global.DatabaseConf.TablePrefix,
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatalf("sqlite init err: %v", err)
	}
}
