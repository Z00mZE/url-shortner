// Code generated by entc, DO NOT EDIT.

package service

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/Z00mZE/url-shortner/ent/service/shorturl"
)

// ShortUrl is the model entity for the ShortUrl schema.
type ShortUrl struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
	// ExpiredAt holds the value of the "expired_at" field.
	ExpiredAt time.Time `json:"expired_at,omitempty"`
	HashID    string    `json:"-"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ShortUrl) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case shorturl.FieldID:
			values[i] = new(sql.NullInt64)
		case shorturl.FieldURL:
			values[i] = new(sql.NullString)
		case shorturl.FieldExpiredAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ShortUrl", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ShortUrl fields.
func (su *ShortUrl) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case shorturl.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			su.ID = int(value.Int64)
		case shorturl.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				su.URL = value.String
			}
		case shorturl.FieldExpiredAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expired_at", values[i])
			} else if value.Valid {
				su.ExpiredAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this ShortUrl.
// Note that you need to call ShortUrl.Unwrap() before calling this method if this ShortUrl
// was returned from a transaction, and the transaction was committed or rolled back.
func (su *ShortUrl) Update() *ShortUrlUpdateOne {
	return (&ShortUrlClient{config: su.config}).UpdateOne(su)
}

// Unwrap unwraps the ShortUrl entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (su *ShortUrl) Unwrap() *ShortUrl {
	tx, ok := su.config.driver.(*txDriver)
	if !ok {
		panic("service: ShortUrl is not a transactional entity")
	}
	su.config.driver = tx.drv
	return su
}

// String implements the fmt.Stringer.
func (su *ShortUrl) String() string {
	var builder strings.Builder
	builder.WriteString("ShortUrl(")
	builder.WriteString(fmt.Sprintf("id=%v", su.ID))
	builder.WriteString(", url=")
	builder.WriteString(su.URL)
	builder.WriteString(", expired_at=")
	builder.WriteString(su.ExpiredAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// ShortUrls is a parsable slice of ShortUrl.
type ShortUrls []*ShortUrl

func (su ShortUrls) config(cfg config) {
	for _i := range su {
		su[_i].config = cfg
	}
}
