package models

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetDatabaseConnection(dbConn string) (*pgxpool.Pool, error) {
	dbpool, err := pgxpool.New(context.Background(), dbConn)
	return dbpool, err
}
