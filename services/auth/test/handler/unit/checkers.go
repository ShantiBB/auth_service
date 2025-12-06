package unit

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"auth/internal/http/dto/response"
	"fukuro-reserve/pkg/utils/consts"
)

type ResponseChecker func(*testing.T, *httptest.ResponseRecorder)

var (
	CheckSuccessUserResponse = func() ResponseChecker {
		return func(t *testing.T, w *httptest.ResponseRecorder) {
			var resp response.User
			err := json.Unmarshal(w.Body.Bytes(), &resp)
			assert.NoError(t, err)
			assert.Equal(t, resp.ID, UserMock.ID)
			assert.Equal(t, resp.Email, UserMock.Email)
			assert.Equal(t, resp.Username, UserMock.Username)
			assert.Equal(t, resp.Role, UserMock.Role)
			assert.Equal(t, resp.IsActive, UserMock.IsActive)
			assert.NotEmpty(t, resp.CreatedAt)
			assert.NotEmpty(t, resp.UpdatedAt)
		}
	}

	CheckSuccessUserGetAllResponse = func() ResponseChecker {
		return func(t *testing.T, w *httptest.ResponseRecorder) {
			var resp []response.User
			err := json.Unmarshal(w.Body.Bytes(), &resp)
			assert.NoError(t, err)
			assert.Equal(t, resp[0].ID, UserMock.ID)
			assert.Equal(t, resp[0].Email, UserMock.Email)
			assert.Equal(t, resp[0].Username, UserMock.Username)
			assert.Equal(t, resp[0].Role, UserMock.Role)
			assert.Equal(t, resp[0].IsActive, UserMock.IsActive)
			assert.NotEmpty(t, resp[0].CreatedAt)
			assert.NotEmpty(t, resp[0].UpdatedAt)
		}
	}

	//checkSuccessUserGetAllResponse = func() ResponseChecker {
	//	return func(t *testing.T, w *httptest.ResponseRecorder) {
	//		var resp response.UserList
	//		err := json.Unmarshal(w.Body.Bytes(), &resp)
	//		assert.NoError(t, err)
	//		assert.Equal(t, resp.Users[0].ID, userMock.ID)
	//		assert.Equal(t, resp.Users[0].Email, userMock.Email)
	//		assert.Equal(t, resp.Users[0].Username, userMock.Username)
	//		assert.Equal(t, resp.Users[0].Role, userMock.Role)
	//		assert.Equal(t, resp.Users[0].IsActive, userMock.IsActive)
	//		assert.NotEmpty(t, resp.Users[0].CreatedAt)
	//		assert.NotEmpty(t, resp.Users[0].UpdatedAt)
	//	}
	//}

	CheckSuccessTokenResponse = func() ResponseChecker {
		return func(t *testing.T, w *httptest.ResponseRecorder) {
			var resp response.Token
			assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
			assert.Equal(t, TokensMock.Access, resp.Access)
			assert.Equal(t, TokensMock.Refresh, resp.Refresh)
		}
	}

	CheckMessageError = func(expectedErr error) ResponseChecker {
		return func(t *testing.T, w *httptest.ResponseRecorder) {
			var resp response.ErrorSchema
			err := json.Unmarshal(w.Body.Bytes(), &resp)
			assert.NoError(t, err)

			assert.Equal(t, expectedErr.Error(), resp.Message)
		}
	}

	CheckFieldsRequired = func(fields ...string) ResponseChecker {
		return func(t *testing.T, w *httptest.ResponseRecorder) {
			var resp response.ValidateError
			err := json.Unmarshal(w.Body.Bytes(), &resp)
			assert.NoError(t, err)

			for _, field := range fields {
				assert.Equal(t, consts.FieldRequired.Error(), resp.Errors[field])
			}
		}
	}

	CheckFieldsInvalid = func(fields map[string]error) ResponseChecker {
		return func(t *testing.T, w *httptest.ResponseRecorder) {
			var resp response.ValidateError
			err := json.Unmarshal(w.Body.Bytes(), &resp)
			assert.NoError(t, err)

			for field, expectedErr := range fields {
				assert.Equal(t, expectedErr.Error(), resp.Errors[field])
			}
		}
	}
)
