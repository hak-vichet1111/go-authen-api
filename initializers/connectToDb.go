package initializers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
    // Read env variables loaded by godotenv
    host := os.Getenv("DB_HOST")
    if host == "" {
        host = "127.0.0.1"
    }
    port := os.Getenv("DB_PORT")
    if port == "" {
        port = "3306"
    }
    user := os.Getenv("DB_USER")
    pass := os.Getenv("DB_PASS")
    name := os.Getenv("DB_NAME")

    // Build DSN: user:password@tcp(host:port)/dbname?params
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, name)

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
	// Store db in DB
    DB = db
    log.Println("Database connected successfully")
}