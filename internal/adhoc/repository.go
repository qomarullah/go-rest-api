package adhoc

import (
	"context"

	"github.com/qomarullah/go-rest-api/internal/entity"
	"github.com/qomarullah/go-rest-api/pkg/dbcontext"
	"github.com/qomarullah/go-rest-api/pkg/log"
)

// Repository encapsulates the logic to access adhocs from the data source.
type Repository interface {
	// Get returns the adhoc with the specified adhoc ID.
	Get(ctx context.Context, id string) (entity.Adhoc, error)

	// Count returns the number of adhocs.
	Count(ctx context.Context) (int, error)
	// Query returns the list of adhocs with the given offset and limit.
	Query(ctx context.Context, offset, limit int) ([]entity.Adhoc, error)
	// Create saves a new adhoc in the storage.
	Create(ctx context.Context, adhoc entity.Adhoc) error
	// Update updates the adhoc with given ID in the storage.
	Update(ctx context.Context, adhoc entity.Adhoc) error
	// Delete removes the adhoc with given ID from the storage.
	Delete(ctx context.Context, id string) error
}

// repository persists adhocs in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new adhoc repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the adhoc with the specified ID from the database.
func (r repository) Get(ctx context.Context, id string) (entity.Adhoc, error) {
	var adhoc entity.Adhoc
	err := r.db.With(ctx).Select().Model(id, &adhoc)
	return adhoc, err
}

// Create saves a new adhoc record in the database.
// It returns the ID of the newly inserted adhoc record.
func (r repository) Create(ctx context.Context, adhoc entity.Adhoc) error {
	return r.db.With(ctx).Model(&adhoc).Insert()
}

// Update saves the changes to an adhoc in the database.
func (r repository) Update(ctx context.Context, adhoc entity.Adhoc) error {
	return r.db.With(ctx).Model(&adhoc).Update()
}

// Delete deletes an adhoc with the specified ID from the database.
func (r repository) Delete(ctx context.Context, id string) error {
	adhoc, err := r.Get(ctx, id)
	if err != nil {
		return err
	}
	return r.db.With(ctx).Model(&adhoc).Delete()
}

// Count returns the number of the adhoc records in the database.
func (r repository) Count(ctx context.Context) (int, error) {
	var count int
	table := entity.Adhoc{}
	err := r.db.With(ctx).Select("COUNT(*)").From(table.TableName()).Row(&count)
	return count, err
}

// Query retrieves the adhoc records with the specified offset and limit from the database.
func (r repository) Query(ctx context.Context, offset, limit int) ([]entity.Adhoc, error) {
	var adhocs []entity.Adhoc
	err := r.db.With(ctx).
		Select().
		OrderBy("id").
		Offset(int64(offset)).
		Limit(int64(limit)).
		All(&adhocs)
	return adhocs, err
}
