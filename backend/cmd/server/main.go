package main

import (
	"log"

	"github.com/candelatorrez/northwind-app/internal/api"
	"github.com/candelatorrez/northwind-app/internal/config"
	"github.com/candelatorrez/northwind-app/internal/database"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load("../.env"); err != nil {
		log.Println("warning: .env not found")
	}

	cfg := config.Load()

	db, err := database.Connect(
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
		cfg.DBUser,
		cfg.DBPassword,
	)

	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := db.DB()

	if err != nil {
		log.Fatal(err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("database connected")

	router := api.NewRouter()

	log.Printf("server running on :%s", cfg.AppPort)

	if err := router.Run(":" + cfg.AppPort); err != nil {
		log.Fatal(err)
	}
}
