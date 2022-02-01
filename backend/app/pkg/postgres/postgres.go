package postgres

import (
	"context"
	"github.com/Z00mZE/url-shortner/ent"
)

// NewEntPostgres -
func NewEntPostgres(url string, opts ...ent.DBOption) (*ent.Client, error) {
	orm, ormError := ent.Open(ent.PgPgx, url, nil, opts)
	if ormError != nil {
		return nil, ormError
	}
	if migrateError := orm.Schema.Create(context.Background()); migrateError != nil {
		return nil, migrateError
	}
	return orm, nil
}
