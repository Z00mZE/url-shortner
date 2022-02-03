package postgres

import (
	"context"
	"github.com/Z00mZE/url-shortner/ent/service"
)

// NewEntPostgres -
func NewEntPostgres(url string, opts ...service.DBOption) (*service.Client, error) {
	orm, ormError := service.Open(service.PgPgx, url, nil, opts)
	if ormError != nil {
		return nil, ormError
	}
	if migrateError := orm.Schema.Create(context.Background()); migrateError != nil {
		return nil, migrateError
	}
	return orm, nil
}
