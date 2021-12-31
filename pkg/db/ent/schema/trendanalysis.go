package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

// TrendAnalysis holds the schema definition for the TrendAnalysis entity.
type TrendAnalysis struct {
	ent.Schema
}

// Fields of the TrendAnalysis.
func (TrendAnalysis) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("project_id", uuid.UUID{}),
		field.UUID("author_id", uuid.UUID{}),
		field.String("title"),
		field.String("content"),
		field.String("abstract"),
		field.Uint32("create_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("update_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}).
			UpdateDefault(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("delete_at").
			DefaultFunc(func() uint32 {
				return 0
			}),
	}
}

// Edges of the TrendAnalysis.
func (TrendAnalysis) Edges() []ent.Edge {
	return nil
}
