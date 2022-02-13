package main

import (
	"planwey/configuration"
	"planwey/database"
	"planwey/web"
)

var Logger = configuration.Configuration.Logger

func main() {
	Logger.Print("Starting migration")
	err := database.Migrate()
	if err != nil {
		Logger.Panicf("Couldn't migrate error: %v", err)
	}
	Logger.Print("Migration completed")
	web.Run()
}
