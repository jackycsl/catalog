package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/jackycsl/catalog/catalog-service/internal/biz"
)

// Game holds the schema definition for the Game entity.
type Game struct {
	ent.Schema
}

// Fields of the Game.
func (Game) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("name"),
		field.String("description"),
		field.Int64("count"),
		field.Int64("price"),
		field.JSON("images", []biz.Image{}),
		field.Time("created_at").
			Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
		field.Time("updated_at").
			Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
	}
}

// Edges of the Beer.
func (Game) Edges() []ent.Edge {
	return nil
}
