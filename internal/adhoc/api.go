package adhoc

import (
	"io"
	"net/http"
	"os"

	routing "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/qomarullah/go-rest-api/internal/errors"
	"github.com/qomarullah/go-rest-api/pkg/log"
	"github.com/qomarullah/go-rest-api/pkg/pagination"
)

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r *routing.RouteGroup, service Service, authHandler routing.Handler, logger log.Logger) {
	res := resource{service, logger}

	r.Get("/adhoc/<id>", res.get)
	r.Get("/adhoc", res.query)

	r.Use(authHandler)
	// the following endpoints require a valid JWT
	r.Post("/adhoc", res.create)
	r.Put("/adhoc/<id>", res.update)
	r.Delete("/adhoc/<id>", res.delete)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) get(c *routing.Context) error {
	adhoc, err := r.service.Get(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}

	return errors.Success(adhoc)
}

func (r resource) query(c *routing.Context) error {
	ctx := c.Request.Context()
	count, err := r.service.Count(ctx)
	if err != nil {
		return err
	}
	pages := pagination.NewFromRequest(c.Request, count)
	adhocs, err := r.service.Query(ctx, pages.Offset(), pages.Limit())
	if err != nil {
		return err
	}
	pages.Items = adhocs
	return errors.Success(pages)
}

func (r resource) create(c *routing.Context) error {

	var input CreateRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}

	//handle file
	c.Request.ParseMultipartForm(32 << 20)
	file, handler, err := c.Request.FormFile("filename")
	if err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return nil
	}
	defer file.Close()
	f, err := os.OpenFile("upload/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0664)
	if err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return nil
	}
	defer f.Close()
	io.Copy(f, file)
	//return c.Write(handler.Filename)

	//continue
	input.Filename = handler.Filename
	input.Name = c.Request.FormValue("name")
	input.ScheduledAt = c.Request.FormValue("scheduled_at")
	err = r.service.Create(c.Request.Context(), input)
	if err != nil {
		return err
	}

	return errors.SuccessWithStatus(nil, http.StatusCreated)

}

func (r resource) update(c *routing.Context) error {
	var input UpdateRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}

	//handle file
	c.Request.ParseMultipartForm(32 << 20)
	file, handler, err := c.Request.FormFile("filename")
	if err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return nil
	}
	defer file.Close()
	f, err := os.OpenFile("upload/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0664)
	if err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return nil
	}
	defer f.Close()
	io.Copy(f, file)
	//return c.Write(handler.Filename)

	//continue
	input.Filename = handler.Filename
	input.Name = c.Request.FormValue("name")
	input.ScheduledAt = c.Request.FormValue("scheduled_at")
	input.Status = c.Request.FormValue("status")

	adhoc, err := r.service.Update(c.Request.Context(), c.Param("id"), input)
	if err != nil {
		return err
	}

	return errors.Success(adhoc)

}

func (r resource) delete(c *routing.Context) error {
	adhoc, err := r.service.Delete(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}

	return errors.Success(adhoc)
}
