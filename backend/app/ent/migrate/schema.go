// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// UrlsColumns holds the columns for the "urls" table.
	UrlsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "url", Type: field.TypeString, Unique: true},
		{Name: "expired_at", Type: field.TypeTime},
	}
	// UrlsTable holds the schema information for the "urls" table.
	UrlsTable = &schema.Table{
		Name:       "urls",
		Columns:    UrlsColumns,
		PrimaryKey: []*schema.Column{UrlsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		UrlsTable,
	}
)

func init() {
}