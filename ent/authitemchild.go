// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/boshangad/v1/ent/authitemchild"
)

// AuthItemChild is the model entity for the AuthItemChild schema.
type AuthItemChild struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AuthItemChild) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case authitemchild.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AuthItemChild", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AuthItemChild fields.
func (aic *AuthItemChild) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case authitemchild.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			aic.ID = int(value.Int64)
		}
	}
	return nil
}

// Update returns a builder for updating this AuthItemChild.
// Note that you need to call AuthItemChild.Unwrap() before calling this method if this AuthItemChild
// was returned from a transaction, and the transaction was committed or rolled back.
func (aic *AuthItemChild) Update() *AuthItemChildUpdateOne {
	return (&AuthItemChildClient{config: aic.config}).UpdateOne(aic)
}

// Unwrap unwraps the AuthItemChild entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (aic *AuthItemChild) Unwrap() *AuthItemChild {
	tx, ok := aic.config.driver.(*txDriver)
	if !ok {
		panic("ent: AuthItemChild is not a transactional entity")
	}
	aic.config.driver = tx.drv
	return aic
}

// String implements the fmt.Stringer.
func (aic *AuthItemChild) String() string {
	var builder strings.Builder
	builder.WriteString("AuthItemChild(")
	builder.WriteString(fmt.Sprintf("id=%v", aic.ID))
	builder.WriteByte(')')
	return builder.String()
}

// AuthItemChilds is a parsable slice of AuthItemChild.
type AuthItemChilds []*AuthItemChild

func (aic AuthItemChilds) config(cfg config) {
	for _i := range aic {
		aic[_i].config = cfg
	}
}
