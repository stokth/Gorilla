package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Task struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}

var tasks = []Task{}

type TaskRequest struct {
	Task string `json:"task"`
}

func getTask(c echo.Context) error {
	return c.JSON(http.StatusOK, tasks)
}

func postTask(c echo.Context) error {
	var req TaskRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	tasks = append(tasks, Task{ID: len(tasks) + 1, Task: req.Task})
	return c.JSON(http.StatusCreated, tasks)
}

func main() {
	c := echo.New()

	c.Use(middleware.CORS())
	c.Use(middleware.Logger())

	c.GET("/tasks", getTask)
	c.POST("/tasks", postTask)

	c.Start("localhost:8080")
}
