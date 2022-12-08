// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"genomedb/ent/trasnscriptstructure"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// TrasnscriptStructure is the model entity for the TrasnscriptStructure schema.
type TrasnscriptStructure struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// TranscriptId holds the value of the "transcriptId" field.
	TranscriptId string `json:"transcriptId,omitempty"`
	// Feature holds the value of the "feature" field.
	Feature string `json:"feature,omitempty"`
	// Seqname holds the value of the "seqname" field.
	Seqname string `json:"seqname,omitempty"`
	// Start holds the value of the "start" field.
	Start int32 `json:"start,omitempty"`
	// End holds the value of the "end" field.
	End int32 `json:"end,omitempty"`
	// Strand holds the value of the "strand" field.
	Strand string `json:"strand,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TrasnscriptStructure) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case trasnscriptstructure.FieldID, trasnscriptstructure.FieldStart, trasnscriptstructure.FieldEnd:
			values[i] = new(sql.NullInt64)
		case trasnscriptstructure.FieldTranscriptId, trasnscriptstructure.FieldFeature, trasnscriptstructure.FieldSeqname, trasnscriptstructure.FieldStrand:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type TrasnscriptStructure", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TrasnscriptStructure fields.
func (ts *TrasnscriptStructure) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case trasnscriptstructure.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ts.ID = int(value.Int64)
		case trasnscriptstructure.FieldTranscriptId:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field transcriptId", values[i])
			} else if value.Valid {
				ts.TranscriptId = value.String
			}
		case trasnscriptstructure.FieldFeature:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field feature", values[i])
			} else if value.Valid {
				ts.Feature = value.String
			}
		case trasnscriptstructure.FieldSeqname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field seqname", values[i])
			} else if value.Valid {
				ts.Seqname = value.String
			}
		case trasnscriptstructure.FieldStart:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field start", values[i])
			} else if value.Valid {
				ts.Start = int32(value.Int64)
			}
		case trasnscriptstructure.FieldEnd:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field end", values[i])
			} else if value.Valid {
				ts.End = int32(value.Int64)
			}
		case trasnscriptstructure.FieldStrand:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field strand", values[i])
			} else if value.Valid {
				ts.Strand = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this TrasnscriptStructure.
// Note that you need to call TrasnscriptStructure.Unwrap() before calling this method if this TrasnscriptStructure
// was returned from a transaction, and the transaction was committed or rolled back.
func (ts *TrasnscriptStructure) Update() *TrasnscriptStructureUpdateOne {
	return (&TrasnscriptStructureClient{config: ts.config}).UpdateOne(ts)
}

// Unwrap unwraps the TrasnscriptStructure entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ts *TrasnscriptStructure) Unwrap() *TrasnscriptStructure {
	_tx, ok := ts.config.driver.(*txDriver)
	if !ok {
		panic("ent: TrasnscriptStructure is not a transactional entity")
	}
	ts.config.driver = _tx.drv
	return ts
}

// String implements the fmt.Stringer.
func (ts *TrasnscriptStructure) String() string {
	var builder strings.Builder
	builder.WriteString("TrasnscriptStructure(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ts.ID))
	builder.WriteString("transcriptId=")
	builder.WriteString(ts.TranscriptId)
	builder.WriteString(", ")
	builder.WriteString("feature=")
	builder.WriteString(ts.Feature)
	builder.WriteString(", ")
	builder.WriteString("seqname=")
	builder.WriteString(ts.Seqname)
	builder.WriteString(", ")
	builder.WriteString("start=")
	builder.WriteString(fmt.Sprintf("%v", ts.Start))
	builder.WriteString(", ")
	builder.WriteString("end=")
	builder.WriteString(fmt.Sprintf("%v", ts.End))
	builder.WriteString(", ")
	builder.WriteString("strand=")
	builder.WriteString(ts.Strand)
	builder.WriteByte(')')
	return builder.String()
}

// TrasnscriptStructures is a parsable slice of TrasnscriptStructure.
type TrasnscriptStructures []*TrasnscriptStructure

func (ts TrasnscriptStructures) config(cfg config) {
	for _i := range ts {
		ts[_i].config = cfg
	}
}