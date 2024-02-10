package data

import (
	"context"

	"github.com/Astak/otus-docker-basics-homework/web-service-gin/config"
	pgxuuid "github.com/jackc/pgx-gofrs-uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

type Database struct {
	Conn *pgxpool.Pool
}

func NewDatabase(c config.Config) *Database {
	pool, _ := NewPgConnectionPool(c)
	return &Database{Conn: pool}
}

func NewPgConnectionPool(c config.Config) (*pgxpool.Pool, error) {
	dbConfig, err := pgxpool.ParseConfig(c.GetDbUrl())
	if err != nil {
		return nil, err
	}
	dbConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		pgxuuid.Register(conn.TypeMap())
		return nil
	}
	conn, err := pgxpool.NewWithConfig(context.Background(), dbConfig)
	if err != nil {
		return nil, err
	}
	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatal().Err(err).Msgf("Error connecting to the database.")
	}
	return conn, nil
}
