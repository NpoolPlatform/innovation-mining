// Code generated by entc, DO NOT EDIT.

package techniqueanalysis

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the techniqueanalysis type in the database.
	Label = "technique_analysis"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAuthorID holds the string denoting the author_id field in the database.
	FieldAuthorID = "author_id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldAbstract holds the string denoting the abstract field in the database.
	FieldAbstract = "abstract"
	// FieldProjectID holds the string denoting the project_id field in the database.
	FieldProjectID = "project_id"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"
	// FieldDeleteAt holds the string denoting the delete_at field in the database.
	FieldDeleteAt = "delete_at"
	// Table holds the table name of the techniqueanalysis in the database.
	Table = "technique_analyses"
)

// Columns holds all SQL columns for techniqueanalysis fields.
var Columns = []string{
	FieldID,
	FieldAuthorID,
	FieldTitle,
	FieldContent,
	FieldAbstract,
	FieldProjectID,
	FieldCreateAt,
	FieldUpdateAt,
	FieldDeleteAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateAt holds the default value on creation for the "create_at" field.
	DefaultCreateAt func() uint32
	// DefaultUpdateAt holds the default value on creation for the "update_at" field.
	DefaultUpdateAt func() uint32
	// UpdateDefaultUpdateAt holds the default value on update for the "update_at" field.
	UpdateDefaultUpdateAt func() uint32
	// DefaultDeleteAt holds the default value on creation for the "delete_at" field.
	DefaultDeleteAt func() uint32
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
