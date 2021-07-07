package main

import (
	"fmt"

	"example.com/models"
	"example.com/routes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectDB() *gorm.DB {
	// Connect to DB
	dsn := "host=portto_api_db_1 user=user password=password dbname=user port=5432 sslmode=disable TimeZone=Asia/Taipei"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect database")
	}
	// Get generic database object sql.DB to use its functions
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("failed to get DB interface")
	}

	// Set connection pool
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(10)

	return db
}

func main() {
	db := connectDB()
	models.Init(db)
	route := routes.Init()
	route.Run(":3000")
}
