package orm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

var orm *gorm.DB

func ConnectDatabase(dbUserName, dbPasword, dbHost, dbPort, dbName string) {
	dsn := dbUserName + ":" + dbPasword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatalln("databases connect failed")
		return
	} else {
		log.Println("databases connect success")
	}
	sqldb, err := db.DB()
	if err != nil {
		log.Fatalln("sql db error ", err)
		return
	}
	err = sqldb.Ping()
	if err != nil {
		log.Fatalln("db ping error")
	} else {
		log.Println("db ping success")
	}
	// 禁用默认表名的复数形式，如果置为 true，则 `User` 的默认表名是 `user`
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	sqldb.SetMaxIdleConns(10)

	// SetMaxOpenCons 设置数据库的最大连接数量。
	sqldb.SetMaxOpenConns(100)

	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	sqldb.SetConnMaxLifetime(time.Second)
	orm = db
	fmt.Println(dbName, "connect success")
}
