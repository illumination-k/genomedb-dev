// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"genomedb/ent/keggorthlogy"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// KeggOrthlogy is the model entity for the KeggOrthlogy schema.
type KeggOrthlogy struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the KeggOrthlogyQuery when eager-loading is set.
	Edges KeggOrthlogyEdges `json:"edges"`
}

// KeggOrthlogyEdges holds the relations/edges for other nodes in the graph.
type KeggOrthlogyEdges struct {
	// Pathways holds the value of the pathways edge.
	Pathways []*KeggPathway `json:"pathways,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// PathwaysOrErr returns the Pathways value or an error if the edge
// was not loaded in eager-loading.
func (e KeggOrthlogyEdges) PathwaysOrErr() ([]*KeggPathway, error) {
	if e.loadedTypes[0] {
		return e.Pathways, nil
	}
	return nil, &NotLoadedError{edge: "pathways"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*KeggOrthlogy) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case keggorthlogy.FieldID, keggorthlogy.FieldName:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type KeggOrthlogy", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the KeggOrthlogy fields.
func (ko *KeggOrthlogy) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case keggorthlogy.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				ko.ID = value.String
			}
		case keggorthlogy.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ko.Name = value.String
			}
		}
	}
	return nil
}

// QueryPathways queries the "pathways" edge of the KeggOrthlogy entity.
func (ko *KeggOrthlogy) QueryPathways() *KeggPathwayQuery {
	return (&KeggOrthlogyClient{config: ko.config}).QueryPathways(ko)
}

// Update returns a builder for updating this KeggOrthlogy.
// Note that you need to call KeggOrthlogy.Unwrap() before calling this method if this KeggOrthlogy
// was returned from a transaction, and the transaction was committed or rolled back.
func (ko *KeggOrthlogy) Update() *KeggOrthlogyUpdateOne {
	return (&KeggOrthlogyClient{config: ko.config}).UpdateOne(ko)
}

// Unwrap unwraps the KeggOrthlogy entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ko *KeggOrthlogy) Unwrap() *KeggOrthlogy {
	_tx, ok := ko.config.driver.(*txDriver)
	if !ok {
		panic("ent: KeggOrthlogy is not a transactional entity")
	}
	ko.config.driver = _tx.drv
	return ko
}

// String implements the fmt.Stringer.
func (ko *KeggOrthlogy) String() string {
	var builder strings.Builder
	builder.WriteString("KeggOrthlogy(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ko.ID))
	builder.WriteString("name=")
	builder.WriteString(ko.Name)
	builder.WriteByte(')')
	return builder.String()
}

// KeggOrthlogies is a parsable slice of KeggOrthlogy.
type KeggOrthlogies []*KeggOrthlogy

func (ko KeggOrthlogies) config(cfg config) {
	for _i := range ko {
		ko[_i].config = cfg
	}
}
