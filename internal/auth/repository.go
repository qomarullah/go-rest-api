package auth

import (
	"context"

	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/qomarullah/go-rest-api/internal/entity"
	"github.com/qomarullah/go-rest-api/pkg/dbcontext"
	"github.com/qomarullah/go-rest-api/pkg/helpers"
	"github.com/qomarullah/go-rest-api/pkg/log"
)

// Repository encapsulates the logic to access auth from the data source.
type Repository interface {
	// Get returns the user
	Get(ctx context.Context, username, password string) (entity.User, error)
}

// repository persists  in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new  repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the album with the specified ID from the database.
func (r repository) Get(ctx context.Context, username, password string) (entity.User, error) {
	var auth entity.User
	q := r.db.With(ctx).NewQuery("SELECT * FROM users where email={:username} and password={:password} LIMIT 1")
	q.Bind(dbx.Params{"username": username, "password": helpers.MD5Hash(password)})
	err := q.One(&auth)
	return auth, err
}
