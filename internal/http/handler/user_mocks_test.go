package handler

import (
	"errors"

	"github.com/stretchr/testify/mock"

	"auth_service/internal/mocks"
	"auth_service/package/utils/errs"
)

var mockUserCreateSuccess = func(m *mocks.Service) {
	m.On("UserCreate", mock.Anything, mock.Anything).Return(&userMock, nil)
}

var mockUserCreateConflict = func(m *mocks.Service) {
	m.On("UserCreate", mock.Anything, mock.Anything).Return(nil, errs.UniqueUserField)
}

var mockUserCreateServerError = func(m *mocks.Service) {
	m.On("UserCreate", mock.Anything, mock.Anything).Return(nil, errors.New("database error"))
}

var mockNoSetup = func(m *mocks.Service) {}
