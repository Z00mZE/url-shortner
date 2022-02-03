package postgres

import (
	"database/sql"
	"github.com/Z00mZE/url-shortner/ent/service"
	"time"
)

type Option func(db *sql.DB)

// MaxPoolSize set max pool connections size
func MaxPoolSize(size int) service.DBOption {
	return func(c *sql.DB) {
		c.SetMaxOpenConns(size)
		c.SetMaxIdleConns(size)
	}
}

func SetConnMaxIdleTime(idleDuration time.Duration) service.DBOption {
	return func(db *sql.DB) {
		db.SetConnMaxIdleTime(idleDuration)
	}
}
func SetConnMaxLifetime(lifetime time.Duration) service.DBOption {
	return func(db *sql.DB) {
		db.SetConnMaxLifetime(lifetime)
	}
}
