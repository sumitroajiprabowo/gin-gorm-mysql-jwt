package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDatabase() *gorm.DB {

	err := godotenv.Load() // Load .env file
	if err != nil {
		log.Fatal("Error loading .env file") // Log error
	}

	dbUser := os.Getenv("DB_USER")         // Load the DB_USER from the .env file
	dbPassword := os.Getenv("DB_PASSWORD") // Load the DB_PASSWORD from the .env file
	dbHost := os.Getenv("DB_HOST")         // Load the DB_HOST from the .env file
	dbPort := os.Getenv("DB_PORT")         // Load the DB_PORT from the .env file
	dbName := os.Getenv("DB_NAME")         // Load the DB_NAME from the .env file

	// Create the connection string
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Open connection to the database
	db, err := gorm.Open(mysql.Open(dbURI), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	// SetConnIdleTime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(10 * time.Hour)

	return db

}

func CloseDatabaseConnection(db *gorm.DB) {

	// Get the underlying sql.DB instance from the gorm.DB instance.
	dbSQL, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	dbSQL.Close() // Close the database connection

}
