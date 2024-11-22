package pkg

import (
	"context"
	"os"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

type PGXQuery interface {
	CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)
	SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults

	Exec(ctx context.Context, sql string, arguments ...any) (commandTag pgconn.CommandTag, err error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

type PGX struct {
	pool *pgxpool.Pool
}

var (
	instance *PGX
	once     sync.Once
)

func NewPGX(ctx context.Context) *PGX {
	once.Do(func() {
		dsn := os.Getenv("DATABASE_URL")
		pool, err := pgxpool.New(ctx, dsn)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to connect to database")
		}

		err = pool.Ping(ctx)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to ping database")
		}
		instance = &PGX{pool: pool}
		log.Info().Msg("Connected to database")
	})

	return instance
}

func (p *PGX) WithTransaction(ctx context.Context, f func(ctx context.Context) error) error {
	tx, err := p.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Commit(ctx)

	ctx = context.WithValue(ctx, TrxKey, tx)
	err = f(ctx)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}
	return nil
}

func (p *PGX) GetConn(ctx context.Context) PGXQuery {
	if tx, ok := ctx.Value(TrxKey).(PGXQuery); ok {
		return tx
	}
	return p.pool
}
