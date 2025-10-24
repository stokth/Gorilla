package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Task struct {
	ID     int    `json:"id"`
	Task   string `json:"task"`
	Status string `json:"status"`
}

var tasks = []Task{}

type TaskRequest struct {
	Task   string `json:"task"`
	Status string `json:"status"`
}

func getTask(c echo.Context) error {
	return c.JSON(http.StatusOK, tasks)
}

func postTask(c echo.Context) error {
	var req TaskRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	tasks = append(tasks, Task{ID: len(tasks) + 1, Task: req.Task, Status: req.Status})
	return c.JSON(http.StatusCreated, tasks)
}

func patchTask(c echo.Context) error {
	id := c.Param("id")

	var req TaskRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	for i, task := range tasks {
		if task.ID == idInt {
			tasks[i].Task = req.Task
			tasks[i].Status = req.Status
			return c.JSON(http.StatusOK, tasks[i])
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{"error": "Task not found"})
}

func deleteTask(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	for i, task := range tasks {
		if task.ID == idInt {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{"error": "Task not found"})
}

func main() {
	c := echo.New()

	c.Use(middleware.CORS())
	c.Use(middleware.Logger())

	c.GET("/tasks", getTask)
	c.POST("/tasks", postTask)
	c.PATCH("/tasks/:id", patchTask)
	c.DELETE("/tasks/:id", deleteTask)

	c.Start("localhost:8080")
}
