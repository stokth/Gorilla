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

	c.Use(middleware.CORS())
	c.Use(middleware.Logger())

	c.GET("/tasks", h.GetTask)
	c.GET("/tasks/:id", h.GetTaskById)
	c.POST("/tasks", h.PostTask)
	c.PATCH("/tasks/:id", h.PatchTask)
	c.DELETE("/tasks/:id", h.DeleteTask)

	c.Start("localhost:8080")
}
