package menu

import (
	"context"

	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/qomarullah/go-rest-api/internal/entity"
	"github.com/qomarullah/go-rest-api/pkg/dbcontext"
	"github.com/qomarullah/go-rest-api/pkg/log"
)

// Repository encapsulates the logic to access menus from the data source.
type Repository interface {
	// Get returns the menu with the specified menu ID.
	Get(ctx context.Context, id string) (entity.Menu, error)

	// Get returns the menu with the specified menu ID by User.
	GetByRoles(ctx context.Context, id string) ([]entity.MenuRoles, error)

	// Count returns the number of menus.
	Count(ctx context.Context) (int, error)
	// Query returns the list of menus with the given offset and limit.
	Query(ctx context.Context, offset, limit int) ([]entity.Menu, error)
	// Create saves a new menu in the storage.
	Create(ctx context.Context, menu entity.Menu) error
	// Update updates the menu with given ID in the storage.
	Update(ctx context.Context, menu entity.Menu) error
	// Delete removes the menu with given ID from the storage.
	Delete(ctx context.Context, id string) error
}

// repository persists menus in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new menu repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the menu with the specified ID from the database.
func (r repository) Get(ctx context.Context, id string) (entity.Menu, error) {
	var menu entity.Menu
	err := r.db.With(ctx).Select().Model(id, &menu)
	return menu, err
}

// Create saves a new menu record in the database.
// It returns the ID of the newly inserted menu record.
func (r repository) Create(ctx context.Context, menu entity.Menu) error {
	return r.db.With(ctx).Model(&menu).Insert()
}

// Update saves the changes to an menu in the database.
func (r repository) Update(ctx context.Context, menu entity.Menu) error {
	return r.db.With(ctx).Model(&menu).Update()
}

// Delete deletes an menu with the specified ID from the database.
func (r repository) Delete(ctx context.Context, id string) error {
	menu, err := r.Get(ctx, id)
	if err != nil {
		return err
	}
	return r.db.With(ctx).Model(&menu).Delete()
}

// Count returns the number of the menu records in the database.
func (r repository) Count(ctx context.Context) (int, error) {
	var count int
	table := entity.Menu{}
	err := r.db.With(ctx).Select("COUNT(*)").From(table.TableName()).Row(&count)
	return count, err
}

// Query retrieves the menu records with the specified offset and limit from the database.
func (r repository) Query(ctx context.Context, offset, limit int) ([]entity.Menu, error) {
	var menus []entity.Menu
	err := r.db.With(ctx).
		Select().
		OrderBy("id").
		Offset(int64(offset)).
		Limit(int64(limit)).
		All(&menus)
	return menus, err
}

// Query retrieves the menu records with the specified offset and limit from the database.
func (r repository) GetByRoles(ctx context.Context, id string) ([]entity.MenuRoles, error) {
	var menus []entity.MenuRoles
	/*err := r.db.With(ctx).
	Select().
	OrderBy("id").
	Offset(int64(offset)).
	Limit(int64(limit)).
	All(&menus)
	*/

	sql := "SELECT * FROM menus m LEFT join menus_roles mp on m.id=mp.menus_id where mp.roles_id={:id}"
	//sql := "SELECT * FROM menus where roles_id={:id}"
	q := r.db.With(ctx).NewQuery(sql)
	q.Bind(dbx.Params{"id": id})
	err := q.All(&menus)
	return menus, err
}
