package handler

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"auth_service/package/utils/errs"
)

var (
	checkTokenResponse = func(t *testing.T, w *httptest.ResponseRecorder) {
		var resp map[string]interface{}
		assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
		assert.Equal(t, "access-token", resp["access_token"])
		assert.Equal(t, "refresh-token", resp["refresh_token"])
		assert.Equal(t, "Bearer", resp["token_type"])
	}

	checkUnauthorizedResponse = func(t *testing.T, w *httptest.ResponseRecorder) {
		var resp map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NotEmpty(t, resp["message"])
	}

	checkRegisterConflictResponse = func(t *testing.T, w *httptest.ResponseRecorder) {
		var resp map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &resp)
		assert.Contains(t, resp["message"], errs.UniqueUserField.Error())
	}

	checkRegisterEmailAndPasswordRequired = func(t *testing.T, w *httptest.ResponseRecorder) {
		var resp struct {
			Errors map[string]string `json:"errors"`
		}

		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)

		assert.Equal(t, errs.FieldRequired.Error(), resp.Errors["Email"])
		assert.Equal(t, errs.FieldRequired.Error(), resp.Errors["Password"])
	}

	checkRegisterInvalidEmailAndPassword = func(t *testing.T, w *httptest.ResponseRecorder) {
		var resp struct {
			Errors map[string]string `json:"errors"`
		}

		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)

		assert.Equal(t, errs.InvalidEmail.Error(), resp.Errors["Email"])
		assert.Equal(t, errs.InvalidPassword.Error(), resp.Errors["Password"])
	}

	checkRegisterServerErrorResponse = func(t *testing.T, w *httptest.ResponseRecorder) {
		var resp map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &resp)
		assert.Contains(t, resp["message"], errs.InternalServer.Error())
	}

	checkLoginServerErrorResponse = func(t *testing.T, w *httptest.ResponseRecorder) {
		var resp map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &resp)
		assert.Contains(t, resp["message"], errs.InternalServer.Error())
	}

	checkRefreshServerErrorResponse = func(t *testing.T, w *httptest.ResponseRecorder) {
		var resp map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &resp)
		assert.Contains(t, resp["message"], errs.InternalServer.Error())
	}
)
