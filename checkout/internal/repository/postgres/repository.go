package postgres

import "github.com/jackc/pgx/v4/pgxpool"

type CheckoutRepository struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *CheckoutRepository {
	return &CheckoutRepository{
		pool: pool,
	}
}
