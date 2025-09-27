package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sqlx.DB
}

func (p PostgresRepository) New(sdn string) (*PostgresRepository, error) {
	db, err := sqlx.Connect("postgres", sdn)
	if err != nil {
		return nil, err
	}

	return &PostgresRepository{db: db}, nil
}
