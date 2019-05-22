package data

import (
	"fmt"
	"time"

	"github.com/axiomzen/beanstalks-api/config"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/sirupsen/logrus"
)

// DAL represents the data abstraction layer and provides an interface to the
// database. This is just a wrapper around the PG database object.
type DAL struct {
	db orm.DB
}

type dbLogger struct{}

func (d dbLogger) BeforeQuery(q *pg.QueryEvent) {}

func (d dbLogger) AfterQuery(q *pg.QueryEvent) {
	fmt.Println(q.FormattedQuery())
}

// New returns a new DAL instance based on a configuration object.
func New(c *config.Config) *DAL {
	opts := &pg.Options{
		Addr:            c.PostgresHost + ":" + c.PostgresPort,
		User:            c.PostgresUser,
		Password:        c.PostgresPass,
		Database:        c.PostgresDatabase,
		MaxRetries:      10,
		MinRetryBackoff: time.Second,
		MaxRetryBackoff: time.Second * 10,
	}

	db := pg.Connect(opts)
	// For debugging
	db.AddQueryHook(dbLogger{})
	dal := &DAL{db}

	if err := dal.Ping(); err != nil {
		logrus.WithError(err).Fatal("Error initializing the database")
	} else {
		logrus.Info("Successfully connected to the database")
	}

	return dal
}

// Ping checks that we can reach the database.
func (dal *DAL) Ping() error {
	i := 0
	_, err := dal.db.QueryOne(pg.Scan(&i), "SELECT 1")
	return err
}
