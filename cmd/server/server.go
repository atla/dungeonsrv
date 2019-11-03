package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/atla/dungeonsrv/pkg/server"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dungeonDataDir := os.Getenv("DUNGEON_DATA_DIR")

	log.Info("Starting dungeonsrv with data folder: ", dungeonDataDir)

	server := server.NewApp(dungeonDataDir)
	server.Run()
}
