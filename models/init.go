package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewGormDB() {
	dsn := `root:123456@tcp(192.168.47.99:3306)/go-admin?charset=utf8mb4&parseTime=True&loc=Local`
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//SkipDefaultTransaction:                   false,
		//NamingStrategy:                           nil,
		//FullSaveAssociations:                     false,
		//Logger:                                   nil,
		//NowFunc:                                  nil,
		//DryRun:                                   false,
		//PrepareStmt:                              false,
		//DisableAutomaticPing:                     false,
		DisableForeignKeyConstraintWhenMigrating: false,
		//IgnoreRelationshipsWhenMigrating:         false,
		//DisableNestedTransaction:                 false,
		//AllowGlobalUpdate:                        false,
		//QueryFields:                              false,
		//CreateBatchSize:                          0,
		//TranslateError:                           false,
		//ClauseBuilders:                           nil,
		//ConnPool:                                 nil,
		//Dialector:                                nil,
		//Plugins:                                  nil,
	})
	if err != nil {
		panic(fmt.Errorf("connect db fail: %w", err))
	}

	//自动建表
	err = db.AutoMigrate(&SysUser{}, &SysRole{}, &SysMenu{}, &SysLog{})

	DB = db
}
