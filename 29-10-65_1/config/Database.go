package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

// database config
func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "1234",
		DBName:   "test1",
	}
	return &dbConfig
}

// database url
func DBStr(dbConfig *DBConfig) string {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DBName)
	return psqlInfo
	//"host=localhost user=postgres password=1234 dbname=test1 port=5432 sslmode=disable TimeZone=Asia/Shanghai"
}
