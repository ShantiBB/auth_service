package handler

import (
	"context"
	"errors"
	"net/http"

	"auth_service/internal/config"
	"auth_service/internal/domain/entity"
	"auth_service/internal/http/lib/schemas"
	"auth_service/internal/http/lib/schemas/request"
	"auth_service/internal/http/lib/schemas/response"
	"auth_service/package/utils/errs"
	"auth_service/package/utils/password"
)

type TokenService interface {
	RegisterByEmail(ctx context.Context, email, password string, cfg *config.Config) (*entity.Token, error)
	LoginByEmail(ctx context.Context, email, pass string, cfg *config.Config) (*entity.Token, error)
	RefreshToken(token *entity.Token, cfg *config.Config) (*entity.Token, error)
}

func (h *Handler) RegisterByEmail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req request.Register
	if ok := h.ParseJSON(w, r, &req); !ok {
		return
	}

	hashPassword, err := password.HashPassword(req.Password)
	if err != nil {
		errMsg := schemas.NewErrorResponse("Error hashing password")
		h.sendError(w, r, http.StatusBadRequest, errMsg)
		return
	}

	tokens, err := h.svc.RegisterByEmail(ctx, req.Email, hashPassword, h.cfg)
	if err != nil {
		if errors.Is(err, errs.UniqueUserField) {
			errMsg := schemas.NewErrorResponse("Email or username already exists")
			h.sendError(w, r, http.StatusConflict, errMsg)
			return
		}
		errMsg := schemas.NewErrorResponse("Error registering user")
		h.sendError(w, r, http.StatusInternalServerError, errMsg)
		return
	}

	tokensResp := response.Token{
		Access:  tokens.Access,
		Refresh: tokens.Refresh,
	}
	h.sendJSON(w, r, http.StatusCreated, tokensResp)
}

func (h *Handler) LoginByEmail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req request.LoginByEmail
	if ok := h.ParseJSON(w, r, &req); !ok {
		return
	}

	tokens, err := h.svc.LoginByEmail(ctx, req.Email, req.Password, h.cfg)
	if err != nil {
		if errors.Is(err, errs.InvalidCredentials) || errors.Is(err, errs.UserNotFound) {
			errMsg := schemas.NewErrorResponse("Invalid email or password")
			h.sendError(w, r, http.StatusUnauthorized, errMsg)
			return
		}
		errMsg := schemas.NewErrorResponse("Error logging in user")
		h.sendError(w, r, http.StatusInternalServerError, errMsg)
		return
	}

	h.sendJSON(w, r, http.StatusOK, response.Token{
		Access:  tokens.Access,
		Refresh: tokens.Refresh,
	})
}

func (h *Handler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	var req request.RefreshToken
	if ok := h.ParseJSON(w, r, &req); !ok {
		return
	}

	token := &entity.Token{Refresh: req.RefreshToken}
	tokens, err := h.svc.RefreshToken(token, h.cfg)
	if err != nil {
		if errors.Is(err, errs.InvalidToken) {
			errMsg := schemas.NewErrorResponse(errs.InvalidToken.Error())
			h.sendError(w, r, http.StatusUnauthorized, errMsg)
			return
		}
		errMsg := schemas.NewErrorResponse("Error refreshing token")
		h.sendError(w, r, http.StatusInternalServerError, errMsg)
		return
	}

	h.sendJSON(w, r, http.StatusOK, response.Token{
		Access:  tokens.Access,
		Refresh: tokens.Refresh,
	})
}
