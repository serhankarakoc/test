package configs

import (
	"davet.link/models"

	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

func LoadDatabase() Database {
	return Database{
		DBUsername: os.Getenv("DB_USERNAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
	}
}

func GetDBConnectionString(database Database) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", database.DBUsername, database.DBPassword, database.DBHost, database.DBPort, database.DBName)
}

var DB *gorm.DB

func SetupDatabase() *gorm.DB {
	if DB != nil {
		return DB
	}

	ld := LoadDatabase()
	dsn := GetDBConnectionString(ld)

	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	err = DB.AutoMigrate(
		&models.User{},
		&models.Category{},
	)

	if err != nil {
		panic("Failed to migrate database")
	}

	return DB
}
