package handler

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"auth_service/internal/http/lib/schemas/response"
	"auth_service/package/utils/errs"
)

var (
	checkSuccessUserCreateResponse = func(t *testing.T, w *httptest.ResponseRecorder) {
		var resp response.User
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Equal(t, resp.ID, userMock.ID)
		assert.Equal(t, resp.Email, userMock.Email)
		assert.Equal(t, resp.Username, userMock.Username)
		assert.Equal(t, resp.Role, userMock.Role)
		assert.Equal(t, resp.IsActive, userMock.IsActive)
		assert.NotEmpty(t, resp.CreatedAt)
		assert.NotEmpty(t, resp.UpdatedAt)
	}

	checkSuccessUserListResponse = func(t *testing.T, w *httptest.ResponseRecorder) {
		var resp []response.User
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Equal(t, resp[0].ID, userMock.ID)
		assert.Equal(t, resp[0].Email, userMock.Email)
		assert.Equal(t, resp[0].Username, userMock.Username)
		assert.Equal(t, resp[0].Role, userMock.Role)
		assert.Equal(t, resp[0].IsActive, userMock.IsActive)
		assert.NotEmpty(t, resp[0].CreatedAt)
		assert.NotEmpty(t, resp[0].UpdatedAt)
	}

	checkTokenResponse = func(t *testing.T, w *httptest.ResponseRecorder) {
		var resp map[string]interface{}
		assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
		assert.Equal(t, "access-token", resp["access_token"])
		assert.Equal(t, "refresh-token", resp["refresh_token"])
		assert.Equal(t, "Bearer", resp["token_type"])
	}

	checkEmailAndPasswordRequired = func(t *testing.T, w *httptest.ResponseRecorder) {
		var resp struct {
			Errors map[string]string `json:"errors"`
		}

		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)

		assert.Equal(t, errs.FieldRequired.Error(), resp.Errors["Email"])
		assert.Equal(t, errs.FieldRequired.Error(), resp.Errors["Password"])
	}

	checkInvalidEmailAndPassword = func(t *testing.T, w *httptest.ResponseRecorder) {
		var resp struct {
			Errors map[string]string `json:"errors"`
		}

		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)

		assert.Equal(t, errs.InvalidEmail.Error(), resp.Errors["Email"])
		assert.Equal(t, errs.InvalidPassword.Error(), resp.Errors["Password"])
	}

	checkLoginInvalidEmail = func(t *testing.T, w *httptest.ResponseRecorder) {
		var resp struct {
			Errors map[string]string `json:"errors"`
		}

		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)

		assert.Equal(t, errs.InvalidEmail.Error(), resp.Errors["Email"])
	}

	checkRefreshTokenRequired = func(t *testing.T, w *httptest.ResponseRecorder) {
		var resp struct {
			Errors map[string]string `json:"errors"`
		}

		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)

		assert.Equal(t, errs.FieldRequired.Error(), resp.Errors["RefreshToken"])
	}

	checkInvalidJSONResponse = func(t *testing.T, w *httptest.ResponseRecorder) {
		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Contains(t, response["message"], errs.InvalidJSON.Error())
	}

	checkConflictResponse = func(t *testing.T, w *httptest.ResponseRecorder) {
		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Contains(t, response["message"], errs.UniqueUserField.Error())
	}

	checkUnauthorizedResponse = func(t *testing.T, w *httptest.ResponseRecorder) {
		var resp map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NotEmpty(t, resp["message"])
	}

	checkServerErrorResponse = func(t *testing.T, w *httptest.ResponseRecorder) {
		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Contains(t, response["message"], errs.InternalServer.Error())
	}
)
