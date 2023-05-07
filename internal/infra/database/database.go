package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func InitDb() (*gorm.DB, error) {
	config := &DBConfig{
		Host:     "categories-db",
		Port:     5432,
		User:     "postgres",
		Password: "password123",
		DBName:   "meetup",
	}

	connectionString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s",
		config.Host, config.Port, config.User, config.DBName, config.Password)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Println(err)
		panic("failed to connect database")
	}

	return db, nil
}
