package infra

import (
	"fmt"
	"yk/internal/app/enterprise"

	"gorm.io/gorm/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Mysql *gorm.DB

func ConnMysql(config MysqlConf) {
	fmt.Println("init mysql...")

	// Connect
	dsn := config.UserName + ":" + config.Password + "@tcp(" + config.Address + ")/" + config.DataBase + "?charset=utf8mb4&parseTime=True&loc=Local"
	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}); err != nil {
		enterprise.Logger().Fatal(err.Error())
	} else {
		Mysql = db
		fmt.Println("init mysql success!")
	}
}
