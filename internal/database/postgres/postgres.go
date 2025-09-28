package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"auth_service/internal/config"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(config *config.Config) (*Repository, error) {
	sdn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Postgres.Host,
		config.Postgres.Port,
		config.Postgres.User,
		config.Postgres.Password,
		config.Postgres.DB,
	)

	db, err := sqlx.Connect("postgres", sdn)
	if err != nil {
		return nil, err
	}

	return &Repository{db: db}, nil
}
