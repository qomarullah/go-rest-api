package reminder

import (
	"context"

	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/qomarullah/go-rest-api/internal/entity"
	"github.com/qomarullah/go-rest-api/pkg/dbcontext"
	"github.com/qomarullah/go-rest-api/pkg/log"
	"github.com/spf13/cast"
)

// Repository encapsulates the logic to access reminders from the data source.
type Repository interface {
	// Get returns the reminder with the specified reminder ID.
	Get(ctx context.Context, id string) (entity.Reminder, error)
	// Count returns the number of reminders.
	Count(ctx context.Context) (int, error)
	// Query returns the list of reminders with the given offset and limit.
	Query(ctx context.Context, offset, limit int) ([]entity.Reminder, error)
	// Create saves a new reminder in the storage.
	Create(ctx context.Context, reminder entity.Reminder) error
	// Update updates the reminder with given ID in the storage.
	Update(ctx context.Context, reminder entity.Reminder) error
	// Delete removes the reminder with given ID from the storage.
	Delete(ctx context.Context, id string) error
}

// repository persists reminders in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new reminder repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the reminder with the MSISDN from the database.
func (r repository) Get(ctx context.Context, id string) (entity.Reminder, error) {
	shard := cast.ToInt(id) % 10
	r.db.DB().TableMapper = func(a interface{}) string {
		return dbx.GetTableName(a) + cast.ToString(shard)
	}

	var reminder entity.Reminder
	err := r.db.With(ctx).Select().Where(dbx.HashExp{"msisdn": id}).One(&reminder)
	return reminder, err
}

// Create saves a new reminder record in the database.
// It returns the ID of the newly inserted reminder record.
func (r repository) Create(ctx context.Context, reminder entity.Reminder) error {
	primary := reminder.Msisdn
	lastId := primary[len(primary)-1:]
	shard := cast.ToInt(lastId) % 10
	r.db.DB().TableMapper = func(a interface{}) string {
		return dbx.GetTableName(a) + cast.ToString(shard)
	}

	return r.db.With(ctx).Model(&reminder).Insert()
}

// Update saves the changes to an reminder in the database.
func (r repository) Update(ctx context.Context, reminder entity.Reminder) error {
	primary := reminder.Msisdn
	lastId := primary[len(primary)-1:]
	shard := cast.ToInt(lastId) % 10
	r.db.DB().TableMapper = func(a interface{}) string {
		return dbx.GetTableName(a) + cast.ToString(shard)
	}

	return r.db.With(ctx).Model(&reminder).Update()
}

// Delete deletes an reminder with the specified ID from the database.
func (r repository) Delete(ctx context.Context, id string) error {

	shard := cast.ToInt(id) % 10
	q := r.db.With(ctx).NewQuery("DELETE  from queue" + cast.ToString(shard) + " where msisdn={:msisdn}")
	q.Bind(dbx.Params{"msisdn": id})
	_, err := q.Execute()

	return err
}

// Count returns the number of the reminder records in the database.
func (r repository) Count(ctx context.Context) (int, error) {
	var count int
	//table := entity.Reminder{}
	//err := r.db.With(ctx).Select("COUNT(*)").From("reminder0").Row(&count)
	sql := "SELECT sum(count) as count FROM (" +
		"SELECT count(*) as count FROM `queue0` WHERE 1 UNION ALL " +
		"SELECT count(*) as count FROM `queue1` WHERE 1 UNION ALL " +
		"SELECT count(*) as count FROM `queue2` WHERE 1 UNION ALL " +
		"SELECT count(*) as count FROM `queue3` WHERE 1 UNION ALL " +
		"SELECT count(*) as count FROM `queue4` WHERE 1 UNION ALL " +
		"SELECT count(*) as count FROM `queue5` WHERE 1 UNION ALL " +
		"SELECT count(*) as count FROM `queue6` WHERE 1 UNION ALL " +
		"SELECT count(*) as count FROM `queue7` WHERE 1 UNION ALL " +
		"SELECT count(*) as count FROM `queue8` WHERE 1 UNION ALL " +
		"SELECT count(*) as count FROM `queue9` WHERE 1)x"
	q := r.db.With(ctx).NewQuery(sql)
	err := q.Row(&count)
	return count, err
}

// Query retrieves the reminder records with the specified offset and limit from the database.
func (r repository) Query(ctx context.Context, offset, limit int) ([]entity.Reminder, error) {
	var reminders []entity.Reminder
	/*err := r.db.With(ctx).
	Select().
	OrderBy("id").
	Offset(int64(offset)).
	Limit(int64(limit)).
	All(&reminders)
	*/
	sql := "SELECT * FROM (" +
		"SELECT * FROM `queue0` WHERE 1 UNION ALL " +
		"SELECT * FROM `queue1` WHERE 1 UNION ALL " +
		"SELECT * FROM `queue2` WHERE 1 UNION ALL " +
		"SELECT * FROM `queue3` WHERE 1 UNION ALL " +
		"SELECT * FROM `queue4` WHERE 1 UNION ALL " +
		"SELECT * FROM `queue5` WHERE 1 UNION ALL " +
		"SELECT * FROM `queue6` WHERE 1 UNION ALL " +
		"SELECT * FROM `queue7` WHERE 1 UNION ALL " +
		"SELECT * FROM `queue8` WHERE 1 UNION ALL " +
		"SELECT * FROM `queue9` WHERE 1)x  LIMIT {:limit} OFFSET {:offset}"
	q := r.db.With(ctx).NewQuery(sql)
	q.Bind(dbx.Params{"offset": int64(offset), "limit": int64(limit)})
	err := q.All(&reminders)

	return reminders, err
}
