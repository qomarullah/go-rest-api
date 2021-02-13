package auth

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/qiangxue/go-rest-api/internal/entity"
	"github.com/qiangxue/go-rest-api/internal/errors"
	"github.com/qiangxue/go-rest-api/pkg/log"
)

// Service encapsulates the authentication logic.
type Service interface {
	// authenticate authenticates a user using username and password.
	// It returns a JWT token if authentication succeeds. Otherwise, an error is returned.
	Login(ctx context.Context, username, password string) (entity.User, error)
}

type service struct {
	repo            Repository
	signingKey      string
	tokenExpiration int
	logger          log.Logger
}

// NewService creates a new authentication service.
func NewService(repo Repository, signingKey string, tokenExpiration int, logger log.Logger) Service {
	return service{repo, signingKey, tokenExpiration, logger}
}

// Login authenticates a user and generates a JWT token if authentication succeeds.
// Otherwise, an error is returned.
func (s service) Login(ctx context.Context, username, password string) (entity.User, error) {
	user, err := s.repo.Get(ctx, username, password)
	if err != nil {
		return user, err
	}
	if user.ID != "" {
		token, _ := s.generateJWT(user)
		user.Token = token
		return user, nil

	}

	return user, errors.Unauthorized("")
}

// generateJWT generates a JWT that encodes an identity.
func (s service) generateJWT(user entity.User) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
		"exp":   time.Now().Add(time.Duration(s.tokenExpiration) * time.Hour).Unix(),
	}).SignedString([]byte(s.signingKey))
}
