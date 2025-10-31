package main

import (
	"Gorilla/internal/db"
	"Gorilla/internal/handlers"
	"Gorilla/internal/tasks"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	c := echo.New()

	taskRepo := tasks.NewTasksRepository(database)
	taskService := tasks.NewTasksService(taskRepo)
	h := handlers.NewTasksHandlers(taskService)

	strictHandler := tasks.NewStrictHandler(h, nil)
	tasks.RegisterHandlers(c, strictHandler)

	c.Use(middleware.CORS())
	c.Use(middleware.Logger())

	if err := c.Start("localhost:8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
