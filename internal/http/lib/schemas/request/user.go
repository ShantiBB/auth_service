package request

type UserCreateRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=5"`
}

type UserUpdateRequest struct {
	Username string `json:"username"`
	Email    string `json:"email" validate:"email"`
}

type UserUpdateRoleRequest struct {
	Role string `json:"role" validate:"required"`
}

type UserUpdateStatusRequest struct {
	IsActive bool `json:"is_active" validate:"required"`
}
