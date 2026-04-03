package config

import (
	"errors"
	"fmt"
	"finance-backend/models"
	"log/slog"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() error {
	gormCfg := buildGORMConfig()

	db, err := gorm.Open(sqlite.Open("finance.db"), gormCfg)
	if err != nil {
		return fmt.Errorf("open database: %w", err)
	}

	if err := configureSQLitePool(db); err != nil {
		return fmt.Errorf("configure connection pool: %w", err)
	}

	if err := runMigrations(db); err != nil {
		return fmt.Errorf("run migrations: %w", err)
	}

	DB = db
	slog.Info("database connected and migrated successfully")
	return nil
}

// buildGORMConfig returns a GORM config with a structured logger.
func buildGORMConfig() *gorm.Config {
	return &gorm.Config{
		Logger: logger.New(
			slogWriter{},
			logger.Config{
				SlowThreshold:             200 * time.Millisecond,
				LogLevel:                  logger.Warn, // only log warnings + slow queries
				IgnoreRecordNotFoundError: true,        // don't log expected "not found" lookups
				Colorful:                  false,
			},
		),
		// Protect against accidental mass updates/deletes without a WHERE clause
		AllowGlobalUpdate: false,

		// Automatically use createdAt/updatedAt fields if present on models
		NowFunc: time.Now,
	}
}

// configureSQLitePool sets up the underlying sql.DB connection pool.
// SQLite doesn't support concurrent writes, so max open connections is 1.
func configureSQLitePool(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("get underlying sql.DB: %w", err)
	}

	sqlDB.SetMaxOpenConns(1)               // SQLite only supports one writer at a time
	sqlDB.SetMaxIdleConns(1)
	sqlDB.SetConnMaxLifetime(time.Hour)    // recycle connections to avoid stale handles

	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("ping database: %w", err)
	}

	return nil
}

// runMigrations auto-migrates all known models in dependency order.
func runMigrations(db *gorm.DB) error {
	models := []any{
		&models.User{},
		&models.Record{},
	}

	for _, model := range models {
		if err := db.AutoMigrate(model); err != nil {
			return fmt.Errorf("migrate %T: %w", model, err)
		}
	}

	return nil
}

// MustGetDB returns the DB instance or panics — useful in tests and CLI tools
// where there's no meaningful way to recover from a missing DB.
func MustGetDB() *gorm.DB {
	if DB == nil {
		panic("database not initialized: call ConnectDB() first")
	}
	return DB
}

// IsNotFound returns true if the error is a GORM "record not found" error.
// Use this in service/repository layers instead of importing gorm directly.
func IsNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

// slogWriter bridges GORM's logger interface to slog.
type slogWriter struct{}

func (slogWriter) Printf(format string, args ...any) {
	slog.Debug(fmt.Sprintf(format, args...))
}