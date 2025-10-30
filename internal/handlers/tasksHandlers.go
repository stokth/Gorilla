package handlers

import (
	"Gorilla/internal/tasks"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TasksHandlers struct {
	service tasks.TasksService
}

func NewTasksHandlers(service tasks.TasksService) TasksHandlers {
	return TasksHandlers{service: service}
}

func (h *TasksHandlers) GetTask(c echo.Context) error {
	tasks, err := h.service.GetTasks()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch tasks"})
	}

	return c.JSON(http.StatusOK, tasks)
}

func (h *TasksHandlers) GetTaskById(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	task, err := h.service.GetTask(idInt)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch task"})
	}

	return c.JSON(http.StatusOK, task)
}

func (h *TasksHandlers) PostTask(c echo.Context) error {
	var req tasks.TaskRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	tsk, err := h.service.CreateTask(tasks.Task{
		Task:   req.Task,
		Status: req.Status,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create task"})
	}

	return c.JSON(http.StatusCreated, tsk)
}

func (h *TasksHandlers) PatchTask(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	var req tasks.TaskRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	updadeTask, err := h.service.UpdateTask(idInt, tasks.Task{
		Task:   req.Task,
		Status: req.Status,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update task"})
	}

	return c.JSON(http.StatusOK, updadeTask)
}

func (h *TasksHandlers) DeleteTask(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	if err := h.service.DeleteTask(idInt); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete task"})
	}

	return c.NoContent(http.StatusNoContent)
}
