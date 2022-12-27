// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"genomedb/ent/keggmodule"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// KeggModule is the model entity for the KeggModule schema.
type KeggModule struct {
	config
	// ID of the ent.
	ID string `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*KeggModule) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case keggmodule.FieldID:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type KeggModule", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the KeggModule fields.
func (km *KeggModule) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case keggmodule.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				km.ID = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this KeggModule.
// Note that you need to call KeggModule.Unwrap() before calling this method if this KeggModule
// was returned from a transaction, and the transaction was committed or rolled back.
func (km *KeggModule) Update() *KeggModuleUpdateOne {
	return (&KeggModuleClient{config: km.config}).UpdateOne(km)
}

// Unwrap unwraps the KeggModule entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (km *KeggModule) Unwrap() *KeggModule {
	_tx, ok := km.config.driver.(*txDriver)
	if !ok {
		panic("ent: KeggModule is not a transactional entity")
	}
	km.config.driver = _tx.drv
	return km
}

// String implements the fmt.Stringer.
func (km *KeggModule) String() string {
	var builder strings.Builder
	builder.WriteString("KeggModule(")
	builder.WriteString(fmt.Sprintf("id=%v", km.ID))
	builder.WriteByte(')')
	return builder.String()
}

// KeggModules is a parsable slice of KeggModule.
type KeggModules []*KeggModule

func (km KeggModules) config(cfg config) {
	for _i := range km {
		km[_i].config = cfg
	}
}