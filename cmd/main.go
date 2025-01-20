package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"server/pkg/handlers"
	taskHandler "server/pkg/handlers/task"
	"server/pkg/models"
	"server/pkg/services/task"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := models.NewDB()
	if err != nil {
		log.Fatal("Error connecting to database")
	}

	if err = db.Migration(); err != nil {
		log.Fatal("Error migrating database: ", err)
	}

	server := handlers.NewServer(db, true)

	env := "dev"
	taskService := task.NewTaskService(db, env)
	taskHandler.NewTaskHandler(server, "tasks", taskService, nil)

	log.Fatal(server.Gin.Run(":" + os.Getenv("PORT")))
}
