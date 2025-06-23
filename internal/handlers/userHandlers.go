package handlers

import (
	"context"
	"domashka/internal/userService"
	"domashka/internal/web/users"
	"strconv"
)

type UserHandlers struct {
	service userService.UserService
}

func NewUserHandlers(s userService.UserService) *UserHandlers {
	return &UserHandlers{service: s}
}

func (u UserHandlers) GetUsers(ctx context.Context, request users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := u.service.GetAllUsers()
	if err != nil {
		return nil, err
	}
	response := users.GetUsers200JSONResponse{}

	for _, user := range allUsers {
		user := users.Users{
			Id:       &user.ID,
			Email:    &user.Email,
			Password: &user.Password,
		}
		response = append(response, user)
	}
	return response, nil
}

func (u UserHandlers) PostUsers(ctx context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	PatchMail := *request.Body.Email
	PatchPassword := *request.Body.Password
	PostUSer, err := u.service.CreateUser(PatchMail, PatchPassword)
	if err != nil {
		return nil, err
	}
	response := users.PostUsers201JSONResponse{
		Id:       &PostUSer.ID,
		Email:    &PostUSer.Email,
		Password: &PostUSer.Password,
	}
	return response, nil
}

func (u UserHandlers) DeleteUsersId(ctx context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	userID := strconv.Itoa(request.Id)
	err := u.service.DeleteUser(userID)
	if err != nil {
		return nil, err
	}
	return users.DeleteUsersId204Response{}, nil
}

func (u UserHandlers) PatchUsersId(ctx context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	userID := strconv.Itoa(request.Id)
	PatchMail := *request.Body.Email
	PatchPassword := *request.Body.Password

	UpdatedUser, err := u.service.UpdateUser(userID, PatchMail, PatchPassword)

	if err != nil {
		return nil, err
	}

	response := users.PatchUsersId200JSONResponse{
		Id:       &UpdatedUser.ID,
		Email:    &UpdatedUser.Email,
		Password: &UpdatedUser.Password,
	}

	return response, nil
}
