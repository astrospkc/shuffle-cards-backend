package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var DB *gorm.DB

func ConnectDB(){
   dsn := "host=localhost user=gorm password=gorm dbname=shuffle port=9920 sslmode=disable TimeZone=Asia/Shanghai"
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to the database")
    }
}
