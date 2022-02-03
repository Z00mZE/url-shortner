package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"net/url"
	"time"
)

// ShortUrl holds the schema definition for the ShortUrl entity.
type ShortUrl struct {
	ent.Schema
}

// Fields of the ShortUrl.
func (ShortUrl) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("url").
			Validate(func(s string) error {
				if _, err := url.ParseRequestURI(s); err != nil {
					return err
				}
				return nil
			}).
			NotEmpty(),
		field.
			Time("expired_at").
			Default(
				func() time.Time {
					return time.Now().AddDate(0, 0, 7)
				},
			),
	}
}

// Edges of the ShortUrl.
func (ShortUrl) Edges() []ent.Edge {
	return nil
}
