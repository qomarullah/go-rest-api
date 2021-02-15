package history

import (
	"context"

	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/qomarullah/go-rest-api/internal/entity"
	"github.com/qomarullah/go-rest-api/pkg/dbcontext"
	"github.com/qomarullah/go-rest-api/pkg/log"
	"github.com/spf13/cast"
)

// Repository encapsulates the logic to access historys from the data source.
type Repository interface {
	// Get returns the history with the specified history ID.
	Get(ctx context.Context, id string) (entity.History, error)
	// Get returns the history with the specified history ID.
	GetByTrxId(ctx context.Context, id string) (entity.History, error)
	// Count returns the number of historys.
	Count(ctx context.Context) (int, error)
	// Query returns the list of historys with the given offset and limit.
	Query(ctx context.Context, offset, limit int) ([]entity.History, error)
	// Create saves a new history in the storage.
	Create(ctx context.Context, history entity.History) error
	// Update updates the history with given ID in the storage.
	Update(ctx context.Context, history entity.History) error
	// Delete removes the history with given ID from the storage.
	Delete(ctx context.Context, id string) error
}

// repository persists historys in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new history repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the history with the MSISDN from the database.
func (r repository) Get(ctx context.Context, id string) (entity.History, error) {
	shard := cast.ToInt(id) % 10
	r.db.DB().TableMapper = func(a interface{}) string {
		return dbx.GetTableName(a) + cast.ToString(shard)
	}

	var history entity.History
	err := r.db.With(ctx).Select().Where(dbx.HashExp{"msisdn": id}).One(&history)
	return history, err
}

// Get reads the history with the specified ID from the database.
func (r repository) GetByTrxId(ctx context.Context, id string) (entity.History, error) {
	//TODO find table transactionid store
	shard := cast.ToInt(id) % 10
	r.db.DB().TableMapper = func(a interface{}) string {
		return dbx.GetTableName(a) + cast.ToString(shard)
	}

	var history entity.History
	err := r.db.With(ctx).Select().Model(id, &history)
	return history, err
}

// Create saves a new history record in the database.
// It returns the ID of the newly inserted history record.
func (r repository) Create(ctx context.Context, history entity.History) error {
	primary := history.Msisdn
	lastId := primary[len(primary)-1:]
	shard := cast.ToInt(lastId) % 10
	r.db.DB().TableMapper = func(a interface{}) string {
		return dbx.GetTableName(a) + cast.ToString(shard)
	}

	return r.db.With(ctx).Model(&history).Insert()
}

// Update saves the changes to an history in the database.
func (r repository) Update(ctx context.Context, history entity.History) error {
	primary := history.Msisdn
	lastId := primary[len(primary)-1:]
	shard := cast.ToInt(lastId) % 10
	r.db.DB().TableMapper = func(a interface{}) string {
		return dbx.GetTableName(a) + cast.ToString(shard)
	}

	return r.db.With(ctx).Model(&history).Update()
}

// Delete deletes an history with the specified ID from the database.
func (r repository) Delete(ctx context.Context, id string) error {

	shard := cast.ToInt(id) % 10
	q := r.db.With(ctx).NewQuery("DELETE  from history" + cast.ToString(shard) + " where msisdn={:msisdn}")
	q.Bind(dbx.Params{"msisdn": id})
	_, err := q.Execute()

	return err
}

// Count returns the number of the history records in the database.
func (r repository) Count(ctx context.Context) (int, error) {
	var count int
	//table := entity.History{}
	//err := r.db.With(ctx).Select("COUNT(*)").From("history0").Row(&count)
	sql := "SELECT sum(count) as count FROM (" +
		"SELECT count(*) as count FROM `history0` WHERE 1 UNION ALL " +
		"SELECT count(*) as count FROM `history1` WHERE 1 UNION ALL " +
		"SELECT count(*) as count FROM `history2` WHERE 1 UNION ALL " +
		"SELECT count(*) as count FROM `history3` WHERE 1 UNION ALL " +
		"SELECT count(*) as count FROM `history4` WHERE 1 UNION ALL " +
		"SELECT count(*) as count FROM `history5` WHERE 1 UNION ALL " +
		"SELECT count(*) as count FROM `history6` WHERE 1 UNION ALL " +
		"SELECT count(*) as count FROM `history7` WHERE 1 UNION ALL " +
		"SELECT count(*) as count FROM `history8` WHERE 1 UNION ALL " +
		"SELECT count(*) as count FROM `history9` WHERE 1)x"
	q := r.db.With(ctx).NewQuery(sql)
	err := q.Row(&count)
	return count, err
}

// Query retrieves the history records with the specified offset and limit from the database.
func (r repository) Query(ctx context.Context, offset, limit int) ([]entity.History, error) {
	var historys []entity.History
	/*err := r.db.With(ctx).
	Select().
	OrderBy("id").
	Offset(int64(offset)).
	Limit(int64(limit)).
	All(&historys)
	*/
	sql := "SELECT * FROM (" +
		"SELECT * FROM `history0` WHERE 1 UNION ALL " +
		"SELECT * FROM `history1` WHERE 1 UNION ALL " +
		"SELECT * FROM `history2` WHERE 1 UNION ALL " +
		"SELECT * FROM `history3` WHERE 1 UNION ALL " +
		"SELECT * FROM `history4` WHERE 1 UNION ALL " +
		"SELECT * FROM `history5` WHERE 1 UNION ALL " +
		"SELECT * FROM `history6` WHERE 1 UNION ALL " +
		"SELECT * FROM `history7` WHERE 1 UNION ALL " +
		"SELECT * FROM `history8` WHERE 1 UNION ALL " +
		"SELECT * FROM `history9` WHERE 1)x  LIMIT {:limit} OFFSET {:offset}"
	q := r.db.With(ctx).NewQuery(sql)
	q.Bind(dbx.Params{"offset": int64(offset), "limit": int64(limit)})
	err := q.All(&historys)

	return historys, err
}
