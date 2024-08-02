package main

import (
	"e-rt/app/entities"
	"e-rt/config"
	"e-rt/database"
	"log"
)

func main() {
	cfg := config.GetConfig()
	db := database.NewPostgresDatabase(cfg)
	dbMigrate(db)
}

func dbMigrate(db database.Database) {
	err := db.GetDb().Migrator().CreateTable(&entities.User{}, &entities.Role{}, &entities.RoleGroup{}, &entities.Menu{})
	if err != nil {
		log.Printf("Error when creating table %v", err)
	}
}
