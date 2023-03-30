package postgres

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type LomsRepository struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *LomsRepository {
	return &LomsRepository{
		pool: pool,
	}
}
