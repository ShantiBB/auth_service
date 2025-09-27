package handler

import "auth_service/internal/service"

type Handler struct {
	UserService service.UserService
}

func New(userService service.UserService) *Handler {
	return &Handler{
		UserService: userService,
	}
}
