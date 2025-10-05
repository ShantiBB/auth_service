package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"auth_service/internal/config"
	"auth_service/internal/mocks"
)

func TestRegisterByEmail(t *testing.T) {
	cases := []struct {
		name           string
		requestBody    interface{}
		mockSetup      func(*mocks.Service)
		expectedStatus int
		checkResponse  func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name:           "Successful registration",
			requestBody:    registerReq,
			mockSetup:      mockRegisterSuccess,
			expectedStatus: http.StatusCreated,
			checkResponse:  checkTokenResponse,
		},
		{
			name:           "Invalid JSON in request body",
			requestBody:    "invalid json",
			mockSetup:      mockNoSetup,
			expectedStatus: http.StatusBadRequest,
			checkResponse:  checkInvalidJSONResponse,
		},
		{
			name:           "Email already exists",
			requestBody:    registerReq,
			mockSetup:      mockRegisterConflict,
			expectedStatus: http.StatusConflict,
			checkResponse:  checkRegisterConflictResponse,
		},
		{
			name:           "Internal server error during registration",
			requestBody:    registerReq,
			mockSetup:      mockRegisterServerError,
			expectedStatus: http.StatusInternalServerError,
			checkResponse:  checkRegisterServerErrorResponse,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			mockSvc := mocks.NewService(t)
			c.mockSetup(mockSvc)

			var body []byte
			if str, ok := c.requestBody.(string); ok {
				body = []byte(str)
			} else {
				body, _ = json.Marshal(c.requestBody)
			}

			req := httptest.NewRequest("POST", "/register", bytes.NewBuffer(body))
			w := httptest.NewRecorder()

			handler := &Handler{
				svc: mockSvc,
				cfg: &config.Config{},
			}
			handler.RegisterByEmail(w, req)

			assert.Equal(t, c.expectedStatus, w.Code)
			c.checkResponse(t, w)

			mockSvc.AssertExpectations(t)
		})
	}
}

func TestLoginByEmail(t *testing.T) {
	cases := []struct {
		name           string
		requestBody    interface{}
		mockSetup      func(*mocks.Service)
		expectedStatus int
		checkResponse  func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name:           "Successful login",
			requestBody:    loginReq,
			mockSetup:      mockLoginSuccess,
			expectedStatus: http.StatusOK,
			checkResponse:  checkTokenResponse,
		},
		{
			name:           "Invalid JSON in request body",
			requestBody:    "invalid json",
			mockSetup:      mockNoSetup,
			expectedStatus: http.StatusBadRequest,
			checkResponse:  checkInvalidJSONResponse,
		},
		{
			name:           "Invalid credentials",
			requestBody:    loginReq,
			mockSetup:      mockLoginInvalidCredentials,
			expectedStatus: http.StatusUnauthorized,
			checkResponse:  checkUnauthorizedResponse,
		},
		{
			name:           "User not found",
			requestBody:    loginReq,
			mockSetup:      mockLoginUserNotFound,
			expectedStatus: http.StatusUnauthorized,
			checkResponse:  checkUnauthorizedResponse,
		},
		{
			name:           "Internal server error during login",
			requestBody:    loginReq,
			mockSetup:      mockLoginServerError,
			expectedStatus: http.StatusInternalServerError,
			checkResponse:  checkLoginServerErrorResponse,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			mockSvc := mocks.NewService(t)
			c.mockSetup(mockSvc)

			var body []byte
			if str, ok := c.requestBody.(string); ok {
				body = []byte(str)
			} else {
				body, _ = json.Marshal(c.requestBody)
			}

			req := httptest.NewRequest("POST", "/login", bytes.NewBuffer(body))
			w := httptest.NewRecorder()

			handler := &Handler{
				svc: mockSvc,
				cfg: &config.Config{},
			}
			handler.LoginByEmail(w, req)

			assert.Equal(t, c.expectedStatus, w.Code)
			c.checkResponse(t, w)

			mockSvc.AssertExpectations(t)
		})
	}
}

func TestRefreshToken(t *testing.T) {
	cases := []struct {
		name           string
		requestBody    interface{}
		mockSetup      func(*mocks.Service)
		expectedStatus int
		checkResponse  func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name:           "Successful token refresh",
			requestBody:    refreshReq,
			mockSetup:      mockRefreshSuccess,
			expectedStatus: http.StatusOK,
			checkResponse:  checkTokenResponse,
		},
		{
			name:           "Invalid JSON in request body",
			requestBody:    "invalid json",
			mockSetup:      mockNoSetup,
			expectedStatus: http.StatusBadRequest,
			checkResponse:  checkInvalidJSONResponse,
		},
		{
			name:           "Invalid refresh token",
			requestBody:    refreshReq,
			mockSetup:      mockRefreshInvalidToken,
			expectedStatus: http.StatusUnauthorized,
			checkResponse:  checkUnauthorizedResponse,
		},
		{
			name:           "Internal server error during token refresh",
			requestBody:    refreshReq,
			mockSetup:      mockRefreshServerError,
			expectedStatus: http.StatusInternalServerError,
			checkResponse:  checkRefreshServerErrorResponse,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			mockSvc := mocks.NewService(t)
			c.mockSetup(mockSvc)

			var body []byte
			if str, ok := c.requestBody.(string); ok {
				body = []byte(str)
			} else {
				body, _ = json.Marshal(c.requestBody)
			}

			req := httptest.NewRequest("POST", "/refresh", bytes.NewBuffer(body))
			w := httptest.NewRecorder()

			handler := &Handler{
				svc: mockSvc,
				cfg: &config.Config{},
			}
			handler.RefreshToken(w, req)

			assert.Equal(t, c.expectedStatus, w.Code)
			c.checkResponse(t, w)

			mockSvc.AssertExpectations(t)
		})
	}
}
