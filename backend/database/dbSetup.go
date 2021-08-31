package database

import (
	"fmt"
	"github.com/simpleittools/simplepsa/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Conn will control the connection to the database
func Conn() {
	dbType := config.GoDotEnvVariable("DB_TYPE")
	dbAddress := config.GoDotEnvVariable("DB_ADDRESS")
	dbPort := config.GoDotEnvVariable("DB_PORT")
	dbUname := config.GoDotEnvVariable("DB_USERNAME")
	dbPassword := config.GoDotEnvVariable("DB_PASSWORD")
	dbName := config.GoDotEnvVariable("DB_NAME")
	timeZone := config.GoDotEnvVariable("TIMEZONE")
	sslMode := config.GoDotEnvVariable("SSL_MODE")
	if dbType == "mysql" {
		dsn := dbUname + ":" + dbPassword + "@tcp(127.0.0.1:3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

		conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			panic("Could not connect to the DB")
		} else {
			fmt.Println("Connected to the MySql Database")
		}
		DB = conn
	} else if dbType == "sqlite" {
		conn, err := gorm.Open(sqlite.Open(dbName+".db"), &gorm.Config{})

		if err != nil {
			panic("Could not connect to the DB")
		} else {
			fmt.Println("Connected to the SQLite Database")
		}
		DB = conn
	} else if dbType == "pgsql" {
		dsn := "host=" + dbAddress +
			" user=" + dbUname +
			" password=" + dbPassword +
			" dbname=" + dbName +
			" port=" + dbPort +
			" sslmode=" + sslMode +
			" TimeZone=" + timeZone
		conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			panic("Could not connect to the DB")
		} else {
			fmt.Println("Connected to the PostgreSQL Database")
		}
		DB = conn
	}

	// AutoMigrate will migrate the models into the table data
	//conn.AutoMigrate(&models.Client{}, &models.Device{}, &models.DeviceType{}, &models.Manufacturer{})
}
