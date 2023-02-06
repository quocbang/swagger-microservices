package impl

import (
	"context"
	"fmt"
	pgimpl "quocbang/swagger-microservices/pg-impl"
	"quocbang/swagger-microservices/pg-impl/impl/orm/models"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PGConfig struct {
	Address  string
	Port     int
	UserName string
	Password string
	Database string

	// default 1 sec.
	LockTimeout time.Duration
}

// Option definition.
type Option func(*options)

type options struct {
	schema string
	// pdaServiceEndpoint string
	autoMigrateTables bool
	// migrateCloudTables bool
	// adAuth bool
}

// DBOptions are option settings for gorm database.
type DBOptions struct {
	Schema string
	Logger logger.Interface
}

// DataManager definition.
type DataManager struct {
	db *gorm.DB

	// pdaService pda.WebService

	// agent *commonsAccount.ADAgent

	lockTimeout time.Duration
}

func pareOptions(opts []Option) options {
	var o options

	for _, opt := range opts {
		opt(&o)
	}
	return o
}

func NewDB(pg PGConfig, opts DBOptions) (*gorm.DB, error) {
	connectionString := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		pg.Address,
		pg.Port,
		pg.UserName,
		pg.Database,
		pg.Password,
	)
	if opts.Schema != "" {
		connectionString += fmt.Sprintf(" search_path=%s", opts.Schema)
	}
	db, err := gorm.Open(postgres.New(
		postgres.Config{
			// DriverName: "postgres",
			DSN: connectionString,
		},
	), &gorm.Config{
		Logger: opts.Logger,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgreSQL database: %v", err)
	}
	return db, nil
}

func newDataManager(pg PGConfig, o options) (*DataManager, error) {

	db, err := NewDB(pg, DBOptions{
		Schema: o.schema,
		// Logger: logger.NewLogger(), // TODO: complicated
	})
	if err != nil {
		return nil, err
	}

	return &DataManager{db: db, lockTimeout: pg.LockTimeout}, nil
}

func New(ctx context.Context, pg PGConfig, opts ...Option) (pgimpl.DataManager, error) {
	o := pareOptions(opts)

	if pg.LockTimeout == 0 {
		pg.LockTimeout = time.Second
	}

	dm, err := newDataManager(pg, o)
	if err != nil {
		return nil, err
	}

	if o.autoMigrateTables {
		if err := dm.maybeMigrate(o.autoMigrateTables); err != nil {
			return nil, err
		}
	}
	return dm, nil
}

func (dm *DataManager) maybeMigrate(migrateCloudTable bool) error {
	// ms := models.GetModelList()
	// if migrateCloudTable {
	// 	cms := models.GetCloudModelList()
	// 	ms = append(ms, cms...)
	// }
	ms := models.GetModelList()

	if err := maybeMigrateTables(dm.db, ms...); err != nil {
		return err
	}
	// if err := maybeMigrateFunctions(dm.db, ms...); err != nil {
	// 	return err
	// }
	return maybeMigrateTriggers(dm.db, ms...)
}

// maybeMigrateTables attempts to create tables automatically if implement
// models.Model interface.
func maybeMigrateTables(db *gorm.DB, ms ...models.Model) error {
	dst := []interface{}{}
	for _, m := range ms {
		dst = append(dst, m)
	}
	return db.AutoMigrate(dst...)
}

// maybeMigrateTriggers attempts to create triggers automatically if implement
// models.Trigger interface.
func maybeMigrateTriggers(db *gorm.DB, ms ...models.Model) error {
	for _, m := range ms {
		t, ok := m.(models.Trigger)
		if !ok {
			continue
		}
		if err := t.MigrateTrigger(db); err != nil {
			return err
		}
	}
	return nil
}

// WithPostgreSQLSchema with given schema name for postgreSQL.
func WithPostgreSQLSchema(schema string) Option {
	return func(o *options) {
		o.schema = schema
	}
}

type session struct {
	DataManager
	ctx context.Context
}

// newSession returns a new session.
func (dm *DataManager) newSession(ctx context.Context) session {
	var res session
	res.db = dm.db.WithContext(ctx)
	res.lockTimeout = dm.lockTimeout
	res.ctx = ctx
	return res
}
