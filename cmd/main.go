package main

import (
	"gogif/config"
	"gogif/db"
	"gogif/routes"
	"log"
)

func main() {
	// read configuration
	cfg := config.ReadConfig()

	// connect to database
	database, err := db.Connect(cfg.Database.DSN)

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// initialize router
	router := routes.Initialize(database)

	// start server
	log.Printf("Listening on port %s", cfg.Server.PORT)
	if err := router.Run(":" + cfg.Server.PORT); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
