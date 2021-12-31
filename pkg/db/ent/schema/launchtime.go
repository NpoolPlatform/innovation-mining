package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

// LaunchTime holds the schema definition for the LaunchTime entity.
type LaunchTime struct {
	ent.Schema
}

// Fields of the LaunchTime.
func (LaunchTime) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("project_id", uuid.UUID{}).
			Unique(),
		field.String("network_name"),
		field.String("network_version"),
		field.String("introduction"),
		field.Uint32("launch_time"),
		field.Bool("incentive"),
		field.Uint32("incentive_total"),
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

// Edges of the LaunchTime.
func (LaunchTime) Edges() []ent.Edge {
	return nil
}
