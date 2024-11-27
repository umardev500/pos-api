package pkg

import (
	"context"
	"fmt"
	nativeLogger "log"
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type GormDB struct {
	db *gorm.DB
}

var (
	gormInstance *GormDB
	gormOnce     sync.Once
)

func NewGormDB() *GormDB {
	gormOnce.Do(func() {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"),
		)
		newLoger := logger.New(
			nativeLogger.New(os.Stdout, "\r\n", nativeLogger.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             200 * time.Millisecond, // Slow SQL threshold
				LogLevel:                  logger.Info,            // Log level
				IgnoreRecordNotFoundError: false,                  // Ignore ErrRecordNotFound error for logger
				Colorful:                  true,                   // Disable color
			},
		)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: newLoger,
		})
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to connect to database")
		}

		gormInstance = &GormDB{db: db}
	})

	return gormInstance
}

func (g *GormDB) WithTransaction(ctx context.Context, f func(ctx context.Context) error) error {
	err := g.db.Transaction(func(tx *gorm.DB) error {
		ctx := context.WithValue(ctx, TrxKey, tx)
		err := f(ctx)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}

func (g *GormDB) GetDB() *gorm.DB {
	return g.db
}

func (g *GormDB) GetConn(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(TrxKey).(*gorm.DB); ok {
		return tx.WithContext(ctx)
	}

	return g.db.WithContext(ctx)
}
