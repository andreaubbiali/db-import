package main

import (
	"aubbiali/db-import/database"
	"log"
)

func main() {
	log.Println("Start service")

	db := database.SetupDatabase()
	// defer db.close
}
