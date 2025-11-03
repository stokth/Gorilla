package main

import (
	"Gorilla/internal/db"
	"Gorilla/internal/handlers"
	"Gorilla/internal/tasks"
	"Gorilla/internal/users"
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
	t := handlers.NewTasksHandlers(taskService)

	strictTaskHandler := tasks.NewStrictHandler(t, nil)
	tasks.RegisterHandlers(c, strictTaskHandler)

	userRepo := users.NewUsersRepository(database)
	userService := users.NewUsersService(userRepo)
	u := handlers.NewUsersHandlers(userService)

	strictUserHandler := users.NewStrictHandler(u, nil)
	users.RegisterHandlers(c, strictUserHandler)

	c.Use(middleware.CORS())
	c.Use(middleware.Logger())

	if err := c.Start("localhost:8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
