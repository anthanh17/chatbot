package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func ConnectDB(ctx context.Context, dbSource string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, dbSource)
	if err != nil {
		panic(err)
	}

	// enable pgvector extension
	_, err = conn.Exec(ctx, "CREATE EXTENSION IF NOT EXISTS vector")
	if err != nil {
		panic(err)
	}

	return conn, nil
}
