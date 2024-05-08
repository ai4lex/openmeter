package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/oklog/ulid/v2"

	"github.com/openmeterio/openmeter/internal/credit/postgres_connector/ent/pgulid"
)

// IDMixin adds the ID field to the schema
type IDMixin struct {
	mixin.Schema
}

// Fields of the IDMixin.
func (IDMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Other("id", pgulid.ULID{}).
			Default(func() pgulid.ULID {
				return pgulid.ULID{ULID: ulid.Make()}
			}).
			Unique().
			Immutable().
			SchemaType(map[string]string{
				dialect.Postgres: "char(26)",
			}),
	}
}

// TimeMixin adds the created_at and updated_at fields to the schema
type TimeMixin struct {
	mixin.Schema
}

// Fields of the TimeMixin.
func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}