//nolint:dupl
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/g11n-manager/pkg/db/mixin"
	"github.com/google/uuid"
)

// Lang holds the schema definition for the Lang entity.
type Lang struct {
	ent.Schema
}

func (Lang) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the Lang.
func (Lang) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.
			String("lang").
			Optional().
			Default(""),
		field.
			String("logo").
			Optional().
			Default(""),
		field.
			String("name").
			Optional().
			Default(""),
		field.
			String("short").
			Optional().
			Default(""),
	}
}

// Edges of the Lang.
func (Lang) Edges() []ent.Edge {
	return nil
}
