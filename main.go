package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() {
	dsn := "host=localhost user=postgres password=root dbname=postgres port=5432 sslmode=disable"
	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&Task{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}

type Task struct {
	ID     int64  `gorm:"primaryKey" json:"id"`
	Task   string `json:"task"`
	Status string `json:"status"`
}

type TaskRequest struct {
	Task   string `json:"task"`
	Status string `json:"status"`
}

func getTask(c echo.Context) error {
	var tasks []Task

	if err := db.Find(&tasks).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch tasks"})
	}

	return c.JSON(http.StatusOK, tasks)
}

func postTask(c echo.Context) error {
	var req TaskRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	var ids int64

	if err := db.Model(&Task{}).Count(&ids).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to count tasks"})
	}

	ids++

	task := Task{
		ID:     ids,
		Task:   req.Task,
		Status: req.Status,
	}

	if err := db.Create(&task).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create task"})
	}
	return c.JSON(http.StatusCreated, task)
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

	var task Task
	if err := db.First(&task, "id = ?", idInt).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Task not found"})
	}

	task.Task = req.Task
	task.Status = req.Status

	if err := db.Save(&task).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update task"})
	}

	return c.JSON(http.StatusOK, task)
}

func deleteTask(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	if err := db.Delete(&Task{}, "id = ?", idInt).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete task"})
	}

	return c.NoContent(http.StatusNoContent)
}

func main() {
	initDB()

	c := echo.New()

	c.Use(middleware.CORS())
	c.Use(middleware.Logger())

	c.GET("/tasks", getTask)
	c.POST("/tasks", postTask)
	c.PATCH("/tasks/:id", patchTask)
	c.DELETE("/tasks/:id", deleteTask)

	c.Start("localhost:8080")
}
