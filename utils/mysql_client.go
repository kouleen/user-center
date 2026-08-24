package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var mysqlDB *gorm.DB

func init() {
	if os.Getenv("MYSQL_USERNAME") == "" {
		log.Fatal("MYSQL_USERNAME env variable not set")
	}
	if os.Getenv("MYSQL_PASSWORD") == "" {
		log.Fatal("MYSQL_PASSWORD env variable not set")
	}
	if os.Getenv("MYSQL_HOST") == "" {
		log.Fatal("MYSQL_HOST env variable not set")
	}
	if os.Getenv("MYSQL_PORT") == "" {
		log.Fatal("MYSQL_PORT env variable not set")
	}
	if os.Getenv("MYSQL_DATABASE") == "" {
		log.Fatal("MYSQL_DATABASE env variable not set")
	}
	if err := InitMysql(os.Getenv("MYSQL_USERNAME"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE")); err != nil {
		log.Fatal(err)
	}
}

func InitMysql(username, password, host, port, database string) error {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		return err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(30 * time.Second)
	sqlDB.SetConnMaxIdleTime(30 * time.Second)
	mysqlDB = db
	return nil
}

func GetMysqlDDB() *gorm.DB {
	return mysqlDB
}
