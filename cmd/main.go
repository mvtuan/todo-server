package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"server/pkg/handlers"
	authHandler "server/pkg/handlers/auth"
	taskHandler "server/pkg/handlers/task"
	userHandler "server/pkg/handlers/user"
	"server/pkg/models"
	"server/pkg/services/auth"
	"server/pkg/services/task"
	"server/pkg/services/user"
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
	userService := user.NewUserService(db, env)
	userHandler.NewUserHandler(server, "users", userService, nil)

	authService := auth.NewAuthService(db, env)
	authHandler.NewAuthHandler(server, "auth", authService, nil)

	taskService := task.NewTaskService(db, env)
	taskHandler.NewTaskHandler(server, "tasks", taskService, nil)

	log.Fatal(server.Gin.Run(":" + os.Getenv("PORT")))
}
