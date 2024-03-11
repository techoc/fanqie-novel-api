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
	initDB()

	// 自动迁移
	err = db.AutoMigrate(&Author{}, &Book{}, &Chapter{}, &Label{})
	if err != nil {
		return
	}
}

func initDB() {
	switch global.DatabaseConf.Type {
	case "sqlite":
		log.Println("database type is sqlite")
		initSqliteDB()

	case "mysql":
		log.Println("database type is mysql")
		initMysqlDB()
	default:
		log.Fatalf("models.Setup err,please set database type")
	}
}

func initMysqlDB() {
	log.Println("init mysql database")
	if global.DatabaseConf.DSN != "" {
		db, err = gorm.Open(mysql.Open(global.DatabaseConf.DSN), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: global.DatabaseConf.TablePrefix,
			},
		})
		if err != nil {
			log.Fatalf("mysql init err: %v", err)
		}
	} else {
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

}

func initSqliteDB() {
	log.Println("init sqlite database")
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
