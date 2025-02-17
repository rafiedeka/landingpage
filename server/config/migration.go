package config

import (
	"landingpage/user"
	"log"
)

func RunMigrations() {
	db := GetDB()

	err := db.AutoMigrate(&user.User{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Database migrated successfully")
}
