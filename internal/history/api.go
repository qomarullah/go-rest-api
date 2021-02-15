package history

import (
	"net/http"

	routing "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/qomarullah/go-rest-api/internal/errors"
	"github.com/qomarullah/go-rest-api/pkg/log"
	"github.com/qomarullah/go-rest-api/pkg/pagination"
)

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r *routing.RouteGroup, service Service, authHandler routing.Handler, logger log.Logger) {
	res := resource{service, logger}

	r.Get("/history/<id>", res.get)
	r.Get("/history/trxid/<id>", res.getByTrxId)
	r.Get("/history", res.query)

	r.Use(authHandler)
	// the following endpoints require a valid JWT
	r.Post("/history", res.create)
	r.Put("/history/<id>", res.update)
	r.Delete("/history/<id>", res.delete)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) get(c *routing.Context) error {
	menu, err := r.service.Get(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}

	return errors.Success(menu)
}

func (r resource) getByTrxId(c *routing.Context) error {
	menu, err := r.service.GetByTrxId(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}

	return errors.Success(menu)
}

func (r resource) query(c *routing.Context) error {
	ctx := c.Request.Context()
	count, err := r.service.Count(ctx)
	if err != nil {
		return err
	}
	pages := pagination.NewFromRequest(c.Request, count)
	menus, err := r.service.Query(ctx, pages.Offset(), pages.Limit())
	if err != nil {
		return err
	}
	pages.Items = menus
	return errors.Success(pages)
}

func (r resource) create(c *routing.Context) error {
	var input CreateRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	menu, err := r.service.Create(c.Request.Context(), input)
	if err != nil {
		return err
	}

	return errors.SuccessWithStatus(menu, http.StatusCreated)
}

func (r resource) update(c *routing.Context) error {
	var input UpdateRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}

	menu, err := r.service.Update(c.Request.Context(), c.Param("id"), input)
	if err != nil {
		return err
	}

	return errors.Success(menu)

}

func (r resource) delete(c *routing.Context) error {
	menu, err := r.service.Delete(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}

	return errors.Success(menu)
}
