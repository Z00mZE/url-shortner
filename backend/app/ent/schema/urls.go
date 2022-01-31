package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Urls holds the schema definition for the Urls entity.
type Urls struct {
	ent.Schema
}

// Fields of the Urls.
func (Urls) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("url").
			Unique().
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

// Edges of the Urls.
func (Urls) Edges() []ent.Edge {
	return nil
}
