package auth

import (
	routing "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/qomarullah/go-rest-api/internal/errors"
	"github.com/qomarullah/go-rest-api/pkg/log"
)

// RegisterHandlers registers handlers for different HTTP requests.
func RegisterHandlers(r *routing.RouteGroup, service Service, logger log.Logger) {
	res := resource{service, logger}

	r.Post("/login", res.login)

}

type resource struct {
	service Service
	logger  log.Logger
}

// login returns a handler that handles user login request.
func (r resource) login(c *routing.Context) error {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.Read(&req); err != nil {
		r.logger.With(c.Request.Context()).Errorf("invalid request: %v", err)
		return errors.BadRequest("")
	}

	user, err := r.service.Login(c.Request.Context(), req.Username, req.Password)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	//return c.Write(user)

	return errors.Success(user)

}
