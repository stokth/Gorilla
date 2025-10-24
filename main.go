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

func main() {
	c := echo.New()

	c.Use(middleware.CORS())
	c.Use(middleware.Logger())

	c.GET("/tasks", getTask)

	c.Start("localhost:8080")
}
