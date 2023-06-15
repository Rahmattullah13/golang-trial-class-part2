package config

import (
	"fmt"
	"golang/trial-class/part2/entity"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnect() {
	dsn := "host=localhost user=postgres password=rahmat011099 dbname=ecommerce port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database", err.Error())
	} else {
		fmt.Println("connected to db")
		DB = db
	}

	db.AutoMigrate(&entity.Product{}, &entity.Order{})
}
