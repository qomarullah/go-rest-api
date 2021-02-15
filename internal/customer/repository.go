package customer

import (
	"context"

	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/qomarullah/go-rest-api/internal/entity"
	"github.com/qomarullah/go-rest-api/pkg/dbcontext"
	"github.com/qomarullah/go-rest-api/pkg/log"
	"github.com/spf13/cast"
)

// Repository encapsulates the logic to access customers from the data source.
type Repository interface {
	// Get returns the customer with the specified customer ID.
	Get(ctx context.Context, id string) (entity.Customer, error)
	// Count returns the number of customers.
	Count(ctx context.Context) (int, error)
	// Query returns the list of customers with the given offset and limit.
	Query(ctx context.Context, offset, limit int) ([]entity.Customer, error)
	// Create saves a new customer in the storage.
	Create(ctx context.Context, customer entity.Customer) error
	// Update updates the customer with given ID in the storage.
	Update(ctx context.Context, customer entity.Customer) error
	// Delete removes the customer with given ID from the storage.
	Delete(ctx context.Context, id string) error
}

// repository persists customers in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new customer repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the customer with the specified ID from the database.
func (r repository) Get(ctx context.Context, id string) (entity.Customer, error) {
	shard := cast.ToInt(id) % 10
	r.db.DB().TableMapper = func(a interface{}) string {
		return dbx.GetTableName(a) + cast.ToString(shard)
	}

	var customer entity.Customer
	err := r.db.With(ctx).Select().Model(id, &customer)
	return customer, err
}

// Create saves a new customer record in the database.
// It returns the ID of the newly inserted customer record.
func (r repository) Create(ctx context.Context, customer entity.Customer) error {
	primary := customer.Msisdn
	lastId := primary[len(primary)-1:]
	shard := cast.ToInt(lastId) % 10
	r.db.DB().TableMapper = func(a interface{}) string {
		return dbx.GetTableName(a) + cast.ToString(shard)
	}

	return r.db.With(ctx).Model(&customer).Insert()
}

// Update saves the changes to an customer in the database.
func (r repository) Update(ctx context.Context, customer entity.Customer) error {
	primary := customer.Msisdn
	lastId := primary[len(primary)-1:]
	shard := cast.ToInt(lastId) % 10
	r.db.DB().TableMapper = func(a interface{}) string {
		return dbx.GetTableName(a) + cast.ToString(shard)
	}

	return r.db.With(ctx).Model(&customer).Update()
}

// Delete deletes an customer with the specified ID from the database.
func (r repository) Delete(ctx context.Context, id string) error {

	shard := cast.ToInt(id) % 10
	q := r.db.With(ctx).NewQuery("DELETE  from profile" + cast.ToString(shard) + " where msisdn={:msisdn}")
	q.Bind(dbx.Params{"msisdn": id})
	_, err := q.Execute()

	return err
}

// Count returns the number of the customer records in the database.
func (r repository) Count(ctx context.Context) (int, error) {
	var count int
	//table := entity.Customer{}
	//err := r.db.With(ctx).Select("COUNT(*)").From("profile0").Row(&count)
	sql := "SELECT sum(count) as count FROM (" +
		"SELECT count(*) as count FROM `profile0` WHERE 1 UNION ALL " +
		"SELECT count(*) as count FROM `profile1` WHERE 1 UNION ALL " +
		"SELECT count(*) as count FROM `profile2` WHERE 1 UNION ALL " +
		"SELECT count(*) as count FROM `profile3` WHERE 1 UNION ALL " +
		"SELECT count(*) as count FROM `profile4` WHERE 1 UNION ALL " +
		"SELECT count(*) as count FROM `profile5` WHERE 1 UNION ALL " +
		"SELECT count(*) as count FROM `profile6` WHERE 1 UNION ALL " +
		"SELECT count(*) as count FROM `profile7` WHERE 1 UNION ALL " +
		"SELECT count(*) as count FROM `profile8` WHERE 1 UNION ALL " +
		"SELECT count(*) as count FROM `profile9` WHERE 1)x"
	q := r.db.With(ctx).NewQuery(sql)
	err := q.Row(&count)
	return count, err
}

// Query retrieves the customer records with the specified offset and limit from the database.
func (r repository) Query(ctx context.Context, offset, limit int) ([]entity.Customer, error) {
	var customers []entity.Customer
	/*err := r.db.With(ctx).
	Select().
	OrderBy("id").
	Offset(int64(offset)).
	Limit(int64(limit)).
	All(&customers)
	*/
	sql := "SELECT * FROM (" +
		"SELECT * FROM `profile0` WHERE 1 UNION ALL " +
		"SELECT * FROM `profile1` WHERE 1 UNION ALL " +
		"SELECT * FROM `profile2` WHERE 1 UNION ALL " +
		"SELECT * FROM `profile3` WHERE 1 UNION ALL " +
		"SELECT * FROM `profile4` WHERE 1 UNION ALL " +
		"SELECT * FROM `profile5` WHERE 1 UNION ALL " +
		"SELECT * FROM `profile6` WHERE 1 UNION ALL " +
		"SELECT * FROM `profile7` WHERE 1 UNION ALL " +
		"SELECT * FROM `profile8` WHERE 1 UNION ALL " +
		"SELECT * FROM `profile9` WHERE 1)x  LIMIT {:limit} OFFSET {:offset}"
	q := r.db.With(ctx).NewQuery(sql)
	q.Bind(dbx.Params{"offset": int64(offset), "limit": int64(limit)})
	err := q.All(&customers)

	return customers, err
}
