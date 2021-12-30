// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AuthorsColumns holds the columns for the "authors" table.
	AuthorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// AuthorsTable holds the schema information for the "authors" table.
	AuthorsTable = &schema.Table{
		Name:       "authors",
		Columns:    AuthorsColumns,
		PrimaryKey: []*schema.Column{AuthorsColumns[0]},
	}
	// CapitalsColumns holds the columns for the "capitals" table.
	CapitalsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// CapitalsTable holds the schema information for the "capitals" table.
	CapitalsTable = &schema.Table{
		Name:       "capitals",
		Columns:    CapitalsColumns,
		PrimaryKey: []*schema.Column{CapitalsColumns[0]},
	}
	// ProjectsColumns holds the columns for the "projects" table.
	ProjectsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// ProjectsTable holds the schema information for the "projects" table.
	ProjectsTable = &schema.Table{
		Name:       "projects",
		Columns:    ProjectsColumns,
		PrimaryKey: []*schema.Column{ProjectsColumns[0]},
	}
	// TeamsColumns holds the columns for the "teams" table.
	TeamsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// TeamsTable holds the schema information for the "teams" table.
	TeamsTable = &schema.Table{
		Name:       "teams",
		Columns:    TeamsColumns,
		PrimaryKey: []*schema.Column{TeamsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AuthorsTable,
		CapitalsTable,
		ProjectsTable,
		TeamsTable,
	}
)

func init() {
}
