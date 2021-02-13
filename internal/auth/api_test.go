package auth

import (
	"context"
	"net/http"
	"testing"

	"github.com/qiangxue/go-rest-api/internal/entity"
	"github.com/qiangxue/go-rest-api/internal/errors"
	"github.com/qiangxue/go-rest-api/internal/test"
	"github.com/qiangxue/go-rest-api/pkg/log"
)

type mockService struct {
}
type mockRepository struct {
}

func (m mockService) Login(ctx context.Context, username, password string) (string, error) {
	if username == "test" && password == "pass" {
		return "token-100", nil
	}
	return "", errors.Unauthorized("")
}

func (r mockRepository) Get(ctx context.Context, username, password string) (entity.User, error) {
	var auth entity.User
	if username == "test" && password == "pass" {
		auth.Token = "1234"
		return auth, nil
	}
	return auth, nil
}

func TestAPI(t *testing.T) {
	logger, _ := log.NewForTest()
	router := test.MockRouter(logger)
	repo := &mockRepository{}
	RegisterHandlers(router.Group(""), NewService(repo, "", 0, logger), logger)

	tests := []test.APITestCase{
		{"success", "POST", "/login", `{"username":"test","password":"pass"}`, nil, http.StatusOK, `{"token":"token-100"}`},
		{"bad credential", "POST", "/login", `{"username":"test","password":"wrong pass"}`, nil, http.StatusUnauthorized, ""},
		{"bad json", "POST", "/login", `"username":"test","password":"wrong pass"}`, nil, http.StatusBadRequest, ""},
	}
	for _, tc := range tests {
		test.Endpoint(t, router, tc)
	}
}
