package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/qrave1/lamoda_test/config"
)

func NewConnect(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.DSN())
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	return db, err
}
