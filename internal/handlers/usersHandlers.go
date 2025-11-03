package handlers

import (
	"Gorilla/internal/users"
	"context"
)

type UsersHandlers struct {
	service users.UsersService
}

// GetUsersIdTasks implements users.StrictServerInterface.
func (h UsersHandlers) GetUsersIdTasks(ctx context.Context, request users.GetUsersIdTasksRequestObject) (users.GetUsersIdTasksResponseObject, error) {
	id := request.Id
	allTasksForUsers, err := h.service.GetTasksForUser(int64(id))

	if err != nil {
		return nil, err
	}

	response := users.GetUsersIdTasks200JSONResponse{}

	for _, user := range allTasksForUsers {
		usrTsk := users.TaskModel{
			Id:     &user.ID,
			Task:   &user.Task,
			Status: &user.Status,
			UserId: &user.UserID,
		}
		response = append(response, usrTsk)
	}

	return response, nil
}

func NewUsersHandlers(service users.UsersService) UsersHandlers {
	return UsersHandlers{service: service}
}

// GetUsers implements users.StrictServerInterface.
func (h UsersHandlers) GetUsers(ctx context.Context, request users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := h.service.GetUsers()

	if err != nil {
		return nil, err
	}

	response := users.GetUsers200JSONResponse{}

	for _, user := range allUsers {
		usr := users.UserModel{
			Id:        &user.ID,
			Email:     &user.Email,
			Password:  &user.Password,
			CreatedAt: &user.Created_At,
			UpdatedAt: &user.Updated_At,
			DeletedAt: &user.Deleted_At,
		}
		response = append(response, usr)
	}

	return response, nil
}

// GetUsersId implements users.StrictServerInterface.
func (h UsersHandlers) GetUsersId(ctx context.Context, request users.GetUsersIdRequestObject) (users.GetUsersIdResponseObject, error) {
	id := request.Id

	user, err := h.service.GetUser(int64(id))

	if err != nil {
		return nil, err
	}

	return users.GetUsersId200JSONResponse{
		Id:        &user.ID,
		Email:     &user.Email,
		Password:  &user.Password,
		CreatedAt: &user.Created_At,
		UpdatedAt: &user.Updated_At,
		DeletedAt: &user.Deleted_At,
	}, nil
}

// PostUsers implements users.StrictServerInterface.
func (h UsersHandlers) PostUsers(ctx context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	userReq := request.Body

	userToCreate := users.User{
		Email:    *userReq.Email,
		Password: *userReq.Password,
	}

	createdUser, err := h.service.CreateUser(&userToCreate)

	if err != nil {
		return nil, err
	}

	response := users.PostUsers201JSONResponse{
		Id:        &createdUser.ID,
		Email:     &createdUser.Email,
		Password:  &createdUser.Password,
		CreatedAt: &createdUser.Created_At,
		UpdatedAt: nil,
		DeletedAt: nil,
	}

	return response, nil
}

// PatchUsersId implements users.StrictServerInterface.
func (h UsersHandlers) PatchUsersId(ctx context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	id := request.Id

	userReq := request.Body

	userToUpdate := users.User{
		Email:    *userReq.Email,
		Password: *userReq.Password,
	}

	updatedTask, err := h.service.UpdateUser(int64(id), &userToUpdate)

	if err != nil {
		return nil, err
	}

	response := users.PatchUsersId200JSONResponse{
		Id:        &updatedTask.ID,
		Email:     &updatedTask.Email,
		Password:  &updatedTask.Password,
		CreatedAt: &updatedTask.Created_At,
		UpdatedAt: &updatedTask.Updated_At,
		DeletedAt: nil,
	}

	return response, nil
}

// DeleteUsersId implements users.StrictServerInterface.
func (h UsersHandlers) DeleteUsersId(ctx context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	id := request.Id

	if err := h.service.DeleteUser(int64(id)); err != nil {
		return nil, err
	}

	return nil, nil
}
