package handler

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"

	"auth_service/api/http/lib/validation"
	"auth_service/api/http/schemas"
	"auth_service/internal/entity"
	"auth_service/utils/password"
)

type UserHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	// GetAll(w http.ResponseWriter, r *http.Request)
	// Update(w http.ResponseWriter, r *http.Request)
	// Delete(w http.ResponseWriter, r *http.Request)
}

func (h Handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req schemas.UserCreateRequest

	if err := render.DecodeJSON(r.Body, &req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, "Bad request")
		return
	}

	validate := validator.New()
	if err := validate.Struct(&req); err != nil {
		var validateErr validator.ValidationErrors
		errors.As(err, &validateErr)

		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, schemas.ValidateErrorResponse{
			Errors: validation.FormatValidationErrors(validateErr),
		})
		return
	}

	hashPassword, err := password.HashPassword(req.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, "Error hashing password")
		return
	}

	newUser := &entity.UserCreate{
		Username:    req.Username,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		Description: req.Description,
		Password:    hashPassword,
	}

	createdUser, err := h.UserService.Create(ctx, *newUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	render.JSON(w, r, schemas.UserResponse{
		ID:          createdUser.ID,
		Username:    createdUser.Username,
		FirstName:   createdUser.FirstName,
		LastName:    createdUser.LastName,
		Email:       createdUser.Email,
		Description: createdUser.Description,
	})

	if password.VerifyPassword(req.Password, hashPassword) {
		println("Password verification succeeded")
	}
}
