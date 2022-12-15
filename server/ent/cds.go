// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"genomedb/ent/cds"
	"genomedb/ent/transcript"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Cds is the model entity for the Cds schema.
type Cds struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Seqname holds the value of the "seqname" field.
	Seqname string `json:"seqname,omitempty"`
	// Start holds the value of the "start" field.
	Start int32 `json:"start,omitempty"`
	// End holds the value of the "end" field.
	End int32 `json:"end,omitempty"`
	// Frame holds the value of the "frame" field.
	Frame int8 `json:"frame,omitempty"`
	// Strand holds the value of the "strand" field.
	Strand string `json:"strand,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CdsQuery when eager-loading is set.
	Edges          CdsEdges `json:"edges"`
	transcript_cds *string
}

// CdsEdges holds the relations/edges for other nodes in the graph.
type CdsEdges struct {
	// Transcript holds the value of the transcript edge.
	Transcript *Transcript `json:"transcript,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// TranscriptOrErr returns the Transcript value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CdsEdges) TranscriptOrErr() (*Transcript, error) {
	if e.loadedTypes[0] {
		if e.Transcript == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: transcript.Label}
		}
		return e.Transcript, nil
	}
	return nil, &NotLoadedError{edge: "transcript"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Cds) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case cds.FieldID, cds.FieldStart, cds.FieldEnd, cds.FieldFrame:
			values[i] = new(sql.NullInt64)
		case cds.FieldSeqname, cds.FieldStrand:
			values[i] = new(sql.NullString)
		case cds.ForeignKeys[0]: // transcript_cds
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Cds", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Cds fields.
func (c *Cds) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case cds.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case cds.FieldSeqname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field seqname", values[i])
			} else if value.Valid {
				c.Seqname = value.String
			}
		case cds.FieldStart:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field start", values[i])
			} else if value.Valid {
				c.Start = int32(value.Int64)
			}
		case cds.FieldEnd:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field end", values[i])
			} else if value.Valid {
				c.End = int32(value.Int64)
			}
		case cds.FieldFrame:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field frame", values[i])
			} else if value.Valid {
				c.Frame = int8(value.Int64)
			}
		case cds.FieldStrand:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field strand", values[i])
			} else if value.Valid {
				c.Strand = value.String
			}
		case cds.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field transcript_cds", values[i])
			} else if value.Valid {
				c.transcript_cds = new(string)
				*c.transcript_cds = value.String
			}
		}
	}
	return nil
}

// QueryTranscript queries the "transcript" edge of the Cds entity.
func (c *Cds) QueryTranscript() *TranscriptQuery {
	return (&CdsClient{config: c.config}).QueryTranscript(c)
}

// Update returns a builder for updating this Cds.
// Note that you need to call Cds.Unwrap() before calling this method if this Cds
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Cds) Update() *CdsUpdateOne {
	return (&CdsClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Cds entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Cds) Unwrap() *Cds {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Cds is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Cds) String() string {
	var builder strings.Builder
	builder.WriteString("Cds(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("seqname=")
	builder.WriteString(c.Seqname)
	builder.WriteString(", ")
	builder.WriteString("start=")
	builder.WriteString(fmt.Sprintf("%v", c.Start))
	builder.WriteString(", ")
	builder.WriteString("end=")
	builder.WriteString(fmt.Sprintf("%v", c.End))
	builder.WriteString(", ")
	builder.WriteString("frame=")
	builder.WriteString(fmt.Sprintf("%v", c.Frame))
	builder.WriteString(", ")
	builder.WriteString("strand=")
	builder.WriteString(c.Strand)
	builder.WriteByte(')')
	return builder.String()
}

// CdsSlice is a parsable slice of Cds.
type CdsSlice []*Cds

func (c CdsSlice) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
