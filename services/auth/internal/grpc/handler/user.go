package handler

import (
	"context"
	"log/slog"

	userv1 "auth/api/user/v1"
	"auth/internal/grpc/utils/helper"
	"auth/internal/grpc/utils/mapper"
)

func (h *Handler) CreateUser(
	ctx context.Context,
	req *userv1.CreateUserRequest,
) (*userv1.CreateUserResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, helper.HandleValidationErr(err)
	}

	user, err := mapper.CreateUserRequestToDomain(req)
	if err != nil {
		return nil, helper.HandleDomainErr(err)
	}

	created, err := h.svc.CreateUser(ctx, user)
	if err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper.HandleDomainErr(err)
	}

	return &userv1.CreateUserResponse{
		User: mapper.CreateUserResponseToProto(created),
	}, nil
}

func (h *Handler) GetUsers(
	ctx context.Context,
	req *userv1.GetUsersRequest,
) (*userv1.GetUsersResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, helper.HandleValidationErr(err)
	}

	userList, err := h.svc.GetUsers(ctx, req.Page, req.Limit)
	if err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper.HandleDomainErr(err)
	}

	return &userv1.GetUsersResponse{
		Users:      mapper.UsersResponseToProto(userList.Users),
		TotalCount: userList.TotalCount,
		Page:       req.Page,
		Limit:      req.Limit,
	}, nil
}

func (h *Handler) GetUser(
	ctx context.Context,
	req *userv1.GetUserRequest,
) (*userv1.GetUserResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, helper.HandleValidationErr(err)
	}

	user, err := h.svc.GetUserByID(ctx, req.Id)
	if err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper.HandleDomainErr(err)
	}

	return &userv1.GetUserResponse{
		User: mapper.UserResponseToProto(user),
	}, nil
}

func (h *Handler) UpdateUser(
	ctx context.Context,
	req *userv1.UpdateUserRequest,
) (*userv1.UpdateUserResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, helper.HandleValidationErr(err)
	}

	user := mapper.UpdateUserRequestToDomain(req)
	if err := h.svc.UpdateUserByID(ctx, user); err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper.HandleDomainErr(err)
	}

	return &userv1.UpdateUserResponse{
		User: mapper.UpdateUserResponseToProto(user),
	}, nil
}

func (h *Handler) DeleteUser(
	ctx context.Context,
	req *userv1.DeleteUserRequest,
) (*userv1.DeleteUserResponse, error) {
	if err := h.validator.Validate(req); err != nil {
		return nil, helper.HandleValidationErr(err)
	}

	if err := h.svc.DeleteUserByID(ctx, req.Id); err != nil {
		slog.ErrorContext(ctx, "failed", slog.String("error", err.Error()))
		return nil, helper.HandleDomainErr(err)
	}

	return &userv1.DeleteUserResponse{
		Message: "success",
	}, nil
}
