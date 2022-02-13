package database

// Database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"planwey/configuration"
)

// Connect

var DbConfig = configuration.Configuration.Database

var Db *gorm.DB = initiateDB()

func initiateDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		DbConfig.Host, DbConfig.User, DbConfig.Password,
		DbConfig.Name, DbConfig.Port, DbConfig.SslMode,
		configuration.Configuration.Timezone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Couldn't connect to database")
	}
	return db
}
