// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/innovation-mining/pkg/db/ent/capital"
)

// Capital is the model entity for the Capital schema.
type Capital struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Capital) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case capital.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Capital", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Capital fields.
func (c *Capital) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case capital.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		}
	}
	return nil
}

// Update returns a builder for updating this Capital.
// Note that you need to call Capital.Unwrap() before calling this method if this Capital
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Capital) Update() *CapitalUpdateOne {
	return (&CapitalClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Capital entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Capital) Unwrap() *Capital {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Capital is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Capital) String() string {
	var builder strings.Builder
	builder.WriteString("Capital(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Capitals is a parsable slice of Capital.
type Capitals []*Capital

func (c Capitals) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
