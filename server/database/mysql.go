package database

import (
	"fmt"
	"log"

	"github.com/sujin/todo-app/config"
	"github.com/sujin/todo-app/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        config.DBUser, config.DBPassword, config.DB_HOST, config.DB_PORT, config.DBName)

    var err error
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}) // ✅ assign to package-level DB
    if err != nil {
        log.Fatal("Failed to connect:", err)
    }

    fmt.Println("DB connected")

    DB.AutoMigrate(&models.User{}, &models.Todo{})
}

