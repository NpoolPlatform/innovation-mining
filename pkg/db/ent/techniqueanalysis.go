// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/innovation-minning/pkg/db/ent/techniqueanalysis"
	"github.com/google/uuid"
)

// TechniqueAnalysis is the model entity for the TechniqueAnalysis schema.
type TechniqueAnalysis struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// AuthorID holds the value of the "author_id" field.
	AuthorID uuid.UUID `json:"author_id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Content holds the value of the "content" field.
	Content string `json:"content,omitempty"`
	// Abstract holds the value of the "abstract" field.
	Abstract string `json:"abstract,omitempty"`
	// ProjectID holds the value of the "project_id" field.
	ProjectID uuid.UUID `json:"project_id,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt uint32 `json:"create_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt uint32 `json:"update_at,omitempty"`
	// DeleteAt holds the value of the "delete_at" field.
	DeleteAt uint32 `json:"delete_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TechniqueAnalysis) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case techniqueanalysis.FieldCreateAt, techniqueanalysis.FieldUpdateAt, techniqueanalysis.FieldDeleteAt:
			values[i] = new(sql.NullInt64)
		case techniqueanalysis.FieldTitle, techniqueanalysis.FieldContent, techniqueanalysis.FieldAbstract:
			values[i] = new(sql.NullString)
		case techniqueanalysis.FieldID, techniqueanalysis.FieldAuthorID, techniqueanalysis.FieldProjectID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type TechniqueAnalysis", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TechniqueAnalysis fields.
func (ta *TechniqueAnalysis) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case techniqueanalysis.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ta.ID = *value
			}
		case techniqueanalysis.FieldAuthorID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field author_id", values[i])
			} else if value != nil {
				ta.AuthorID = *value
			}
		case techniqueanalysis.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				ta.Title = value.String
			}
		case techniqueanalysis.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				ta.Content = value.String
			}
		case techniqueanalysis.FieldAbstract:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field abstract", values[i])
			} else if value.Valid {
				ta.Abstract = value.String
			}
		case techniqueanalysis.FieldProjectID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field project_id", values[i])
			} else if value != nil {
				ta.ProjectID = *value
			}
		case techniqueanalysis.FieldCreateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				ta.CreateAt = uint32(value.Int64)
			}
		case techniqueanalysis.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				ta.UpdateAt = uint32(value.Int64)
			}
		case techniqueanalysis.FieldDeleteAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete_at", values[i])
			} else if value.Valid {
				ta.DeleteAt = uint32(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this TechniqueAnalysis.
// Note that you need to call TechniqueAnalysis.Unwrap() before calling this method if this TechniqueAnalysis
// was returned from a transaction, and the transaction was committed or rolled back.
func (ta *TechniqueAnalysis) Update() *TechniqueAnalysisUpdateOne {
	return (&TechniqueAnalysisClient{config: ta.config}).UpdateOne(ta)
}

// Unwrap unwraps the TechniqueAnalysis entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ta *TechniqueAnalysis) Unwrap() *TechniqueAnalysis {
	tx, ok := ta.config.driver.(*txDriver)
	if !ok {
		panic("ent: TechniqueAnalysis is not a transactional entity")
	}
	ta.config.driver = tx.drv
	return ta
}

// String implements the fmt.Stringer.
func (ta *TechniqueAnalysis) String() string {
	var builder strings.Builder
	builder.WriteString("TechniqueAnalysis(")
	builder.WriteString(fmt.Sprintf("id=%v", ta.ID))
	builder.WriteString(", author_id=")
	builder.WriteString(fmt.Sprintf("%v", ta.AuthorID))
	builder.WriteString(", title=")
	builder.WriteString(ta.Title)
	builder.WriteString(", content=")
	builder.WriteString(ta.Content)
	builder.WriteString(", abstract=")
	builder.WriteString(ta.Abstract)
	builder.WriteString(", project_id=")
	builder.WriteString(fmt.Sprintf("%v", ta.ProjectID))
	builder.WriteString(", create_at=")
	builder.WriteString(fmt.Sprintf("%v", ta.CreateAt))
	builder.WriteString(", update_at=")
	builder.WriteString(fmt.Sprintf("%v", ta.UpdateAt))
	builder.WriteString(", delete_at=")
	builder.WriteString(fmt.Sprintf("%v", ta.DeleteAt))
	builder.WriteByte(')')
	return builder.String()
}

// TechniqueAnalyses is a parsable slice of TechniqueAnalysis.
type TechniqueAnalyses []*TechniqueAnalysis

func (ta TechniqueAnalyses) config(cfg config) {
	for _i := range ta {
		ta[_i].config = cfg
	}
}
