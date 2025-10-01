package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"auth_service/internal/domain/models"
	"auth_service/internal/http/lib/schemas"
	"auth_service/internal/http/lib/schemas/request"
	"auth_service/internal/http/lib/schemas/response"
	"auth_service/package/utils/password"
)

type UserService interface {
	UserCreate(ctx context.Context, user models.UserCreate) (*models.User, error)
	UserGetByID(ctx context.Context, id int64) (*models.User, error)
	UserList(ctx context.Context) ([]models.User, error)
	UserUpdateByID(ctx context.Context, user *models.User) (*models.User, error)
	UserDeleteByID(ctx context.Context, id int64) error
}

func (h *Handler) UserCreate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req request.UserCreateRequest

	if ok := h.ParseJSON(w, r, &req); !ok {
		return
	}

	hashPassword, err := password.HashPassword(req.Password)
	if err != nil {
		errMsg := schemas.NewErrorResponse("Error hashing password")
		h.sendError(w, r, http.StatusBadRequest, errMsg)
		return
	}

	newUser := h.UserCreateRequestToEntity(&req, hashPassword)
	createdUser, err := h.svc.UserCreate(ctx, *newUser)
	if err != nil {
		errMsg := schemas.NewErrorResponse("Error creating user")
		h.sendError(w, r, http.StatusInternalServerError, errMsg)
		return
	}

	userResponse := h.UserEntityToResponse(createdUser)
	h.sendJSON(w, r, http.StatusCreated, userResponse)
}

func (h *Handler) UserGetByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	paramID := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(paramID, 10, 64)
	if err != nil {
		errMsg := schemas.NewErrorResponse("Invalid user ID")
		h.sendError(w, r, http.StatusBadRequest, errMsg)
		return
	}

	user, err := h.svc.UserGetByID(ctx, id)
	if err != nil {
		errMsg := schemas.NewErrorResponse("Error retrieving user")
		h.sendError(w, r, http.StatusInternalServerError, errMsg)
		return
	}

	userResponse := h.UserEntityToResponse(user)
	h.sendJSON(w, r, http.StatusOK, userResponse)
}

func (h *Handler) UserList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	users, err := h.svc.UserList(ctx)
	if err != nil {
		errMsg := schemas.NewErrorResponse("Error retrieving users")
		h.sendError(w, r, http.StatusInternalServerError, errMsg)
		return
	}

	usersResp := make([]response.UserResponse, 0, len(users))
	for _, user := range users {
		userResponse := h.UserEntityToResponse(&user)
		usersResp = append(usersResp, *userResponse)
	}
	h.sendJSON(w, r, http.StatusOK, usersResp)
}

func (h *Handler) UserUpdateByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	paramID := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(paramID, 10, 64)
	if err != nil {
		errMsg := schemas.NewErrorResponse("Invalid user ID")
		h.sendError(w, r, http.StatusBadRequest, errMsg)
		return
	}

	var req request.UserUpdateRequest
	if ok := h.ParseJSON(w, r, &req); !ok {
		return
	}

	user := h.UserUpdateRequestToEntity(&req, id)
	userToUpdate, err := h.svc.UserUpdateByID(ctx, user)
	if err != nil {
		errMsg := schemas.NewErrorResponse("Error updating user")
		h.sendError(w, r, http.StatusInternalServerError, errMsg)
		return
	}

	userResponse := h.UserEntityToResponse(userToUpdate)

	h.sendJSON(w, r, http.StatusOK, userResponse)
}

func (h *Handler) UserDeleteByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	paramID := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(paramID, 10, 64)
	if err != nil {
		errMsg := schemas.NewErrorResponse("Invalid user ID")
		h.sendError(w, r, http.StatusBadRequest, errMsg)
		return
	}

	if err = h.svc.UserDeleteByID(ctx, id); err != nil {
		errMsg := schemas.NewErrorResponse("Error deleting user")
		h.sendError(w, r, http.StatusInternalServerError, errMsg)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
