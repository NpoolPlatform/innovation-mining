package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

// TechniqueAnalysis holds the schema definition for the TechniqueAnalysis entity.
type TechniqueAnalysis struct {
	ent.Schema
}

// Fields of the TechniqueAnalysis.
func (TechniqueAnalysis) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("author_id", uuid.UUID{}),
		field.String("title"),
		field.String("content"),
		field.String("abstract"),
		field.UUID("project_id", uuid.UUID{}),
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

// Edges of the TechniqueAnalysis.
func (TechniqueAnalysis) Edges() []ent.Edge {
	return nil
}
