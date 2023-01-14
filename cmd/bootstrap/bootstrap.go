package bootstrap

import (
	"log"
	"os"
	"regexp"

	"github.com/bypepe77/secret-app-backend/internal/common/database"
	"github.com/bypepe77/secret-app-backend/internal/server"
	"github.com/joho/godotenv"
)

const status = "dev"

func Run() error {
	if status != "production" {
		loadEnv()
	}

	db := database.DatabaseConnectionEnt()
	config := server.NewConfig(os.Getenv("APP_NAME"), os.Getenv("PORT"))
	server := server.NewServer(config, db)

	return server.Run()
}

const projectDirName = "secret-app-backend"

func loadEnv() {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`)

	if err != nil {
		log.Fatalf("Error loading .env file", err)
	}
}
