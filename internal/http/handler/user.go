package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"auth_service/internal/domain/models"
	"auth_service/internal/http/lib/validation"
	schemas2 "auth_service/internal/http/schemas"
	"auth_service/package/utils/password"
)

type UserService interface {
	CreateUser(ctx context.Context, user models.UserCreate) (*models.User, error)
	GetUser(ctx context.Context, id int64) (*models.User, error)
	GetAllUsers(ctx context.Context) ([]models.User, error)
	UpdateUser(ctx context.Context, user models.User) (*models.User, error)
	DeleteUser(ctx context.Context, id int64) error
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req schemas2.UserCreateRequest

	if err := render.DecodeJSON(r.Body, &req); err != nil {
		h.sendError(w, r, http.StatusBadRequest, "Bad request")
		return
	}

	if errResp := validation.CheckErrors(&req); errResp != nil {
		h.sendError(w, r, http.StatusBadRequest, errResp)
		return
	}

	hashPassword, err := password.HashPassword(req.Password)
	if err != nil {
		h.sendError(w, r, http.StatusBadRequest, "Error hashing password")
		return
	}

	newUser := h.UserCreateRequestToEntity(&req, hashPassword)
	createdUser, err := h.svc.CreateUser(ctx, *newUser)
	if err != nil {
		h.sendError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	userResponse := h.UserEntityToResponse(createdUser)
	h.sendJSON(w, r, http.StatusCreated, userResponse)
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	paramID := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(paramID, 10, 64)
	if err != nil {
		h.sendError(w, r, http.StatusBadRequest, "Invalid user ID")
		return
	}

	user, err := h.svc.GetUser(ctx, id)
	if err != nil {
		h.sendError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	userResponse := h.UserEntityToResponse(user)
	h.sendJSON(w, r, http.StatusOK, userResponse)
}

func (h *Handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	users, err := h.svc.GetAllUsers(ctx)
	if err != nil {
		h.sendError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	usersResp := make([]schemas2.UserResponse, 0, len(users))
	for _, user := range users {
		userResponse := h.UserEntityToResponse(&user)
		usersResp = append(usersResp, *userResponse)
	}
	h.sendJSON(w, r, http.StatusOK, usersResp)
}
