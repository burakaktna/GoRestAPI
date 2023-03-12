package config

import (
	"fmt"
	"time"

	"github.com/burakaktna/GoRestAPI/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "mysql:mysql@tcp(db)/mysql?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDb, err := db.DB()
	if err != nil {
		panic(err)
	}

	start := time.Now()
	for sqlDb.Ping() != nil {
		if time.Since(start) > 10*time.Second {
			fmt.Println("db connection timeout")
			break
		}
	}

	fmt.Println("Connected to database: ", sqlDb.Ping() == nil)

	db.AutoMigrate(&models.Article{}, &models.Tag{}, &models.ArticleTag{})

	DB = db
}
