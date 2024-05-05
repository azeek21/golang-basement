package main

import (
	"log"
	"randomaliens/internal/models"
	"randomaliens/internal/repository"
)

func main() {
	db, err := repository.NewPostgresDb()
	if err != nil {
		log.Fatalln(err.Error())
	}
	db.AutoMigrate(&models.Session{}, &models.Anomaly{})
	log.Println("Successfully migrated")
}
