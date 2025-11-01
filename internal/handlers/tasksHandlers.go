package handlers

import (
	"Gorilla/internal/tasks"
	"context"
)

type TasksHandlers struct {
	service tasks.TasksService
}

func NewTasksHandlers(service tasks.TasksService) TasksHandlers {
	return TasksHandlers{service: service}
}

// GetTasks implements tasks.StrictServerInterface.
func (h TasksHandlers) GetTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.service.GetTasks()

	if err != nil {
		return nil, err
	}

	response := tasks.GetTasks200JSONResponse{}

	for _, task := range allTasks {
		tsk := tasks.TaskModel{
			Id:     &task.ID,
			Task:   &task.Task,
			Status: &task.Status,
		}
		response = append(response, tsk)
	}

	return response, nil
}

// GetTasksId implements tasks.StrictServerInterface.
func (h TasksHandlers) GetTasksId(ctx context.Context, request tasks.GetTasksIdRequestObject) (tasks.GetTasksIdResponseObject, error) {
	id := request.Id

	task, err := h.service.GetTask(int64(id))

	if err != nil {
		return nil, err
	}

	return tasks.GetTasksId200JSONResponse{
		Id:     &task.ID,
		Task:   &task.Task,
		Status: &task.Status,
	}, nil
}

// PostTasks implements tasks.StrictServerInterface.
func (h TasksHandlers) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskReq := request.Body

	taskToCreate := tasks.Task{
		Task:   *taskReq.Task,
		Status: *taskReq.Status,
	}

	createdTask, err := h.service.CreateTask(&taskToCreate)

	if err != nil {
		return nil, err
	}

	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		Status: &createdTask.Status,
	}

	return response, nil
}

// PatchTasksId implements tasks.StrictServerInterface.
func (h TasksHandlers) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	id := request.Id

	taskReq := request.Body

	taskToUpdate := tasks.Task{
		Task:   *taskReq.Task,
		Status: *taskReq.Status,
	}

	updatedTask, err := h.service.UpdateTask(int64(id), &taskToUpdate)

	if err != nil {
		return nil, err
	}

	response := tasks.PatchTasksId200JSONResponse{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Task,
		Status: &updatedTask.Status,
	}

	return response, nil
}

// DeleteTasksId implements tasks.StrictServerInterface.
func (h TasksHandlers) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	id := request.Id

	if err := h.service.DeleteTask(int64(id)); err != nil {
		return nil, err
	}

	return nil, nil
}
