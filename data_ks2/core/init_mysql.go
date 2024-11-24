package core

import (
	"data_ks2/global"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMySql() *gorm.DB {
	if global.Config.Mysql.Host == "" {
		return nil
	}
	dsn := global.Config.Mysql.Dsn()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, //禁用迁移期间的外键约束
	})
	if err != nil {
		fmt.Println("------！！！数据库连接失败！！！--------")
		return nil
	} else {
		fmt.Println("---------数据库连接成功-----------")
	}
	sqlDB, _ := db.DB()                     //这一行代码的作用是从GORM的 *gorm.DB 实例中获取底层的 *sql.DB 实例
	sqlDB.SetMaxIdleConns(10)               // 设置空闲连接数
	sqlDB.SetMaxOpenConns(100)              // 设置最大连接数
	sqlDB.SetConnMaxLifetime(time.Hour * 4) // 设置连接最大生命周期
	return db
}
