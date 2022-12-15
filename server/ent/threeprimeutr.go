// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"genomedb/ent/threeprimeutr"
	"genomedb/ent/transcript"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// ThreePrimeUtr is the model entity for the ThreePrimeUtr schema.
type ThreePrimeUtr struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Seqname holds the value of the "seqname" field.
	Seqname string `json:"seqname,omitempty"`
	// Start holds the value of the "start" field.
	Start int32 `json:"start,omitempty"`
	// End holds the value of the "end" field.
	End int32 `json:"end,omitempty"`
	// Strand holds the value of the "strand" field.
	Strand string `json:"strand,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ThreePrimeUtrQuery when eager-loading is set.
	Edges                      ThreePrimeUtrEdges `json:"edges"`
	transcript_three_prime_utr *string
}

// ThreePrimeUtrEdges holds the relations/edges for other nodes in the graph.
type ThreePrimeUtrEdges struct {
	// Transcript holds the value of the transcript edge.
	Transcript *Transcript `json:"transcript,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// TranscriptOrErr returns the Transcript value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ThreePrimeUtrEdges) TranscriptOrErr() (*Transcript, error) {
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
func (*ThreePrimeUtr) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case threeprimeutr.FieldID, threeprimeutr.FieldStart, threeprimeutr.FieldEnd:
			values[i] = new(sql.NullInt64)
		case threeprimeutr.FieldSeqname, threeprimeutr.FieldStrand:
			values[i] = new(sql.NullString)
		case threeprimeutr.ForeignKeys[0]: // transcript_three_prime_utr
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ThreePrimeUtr", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ThreePrimeUtr fields.
func (tpu *ThreePrimeUtr) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case threeprimeutr.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			tpu.ID = int(value.Int64)
		case threeprimeutr.FieldSeqname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field seqname", values[i])
			} else if value.Valid {
				tpu.Seqname = value.String
			}
		case threeprimeutr.FieldStart:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field start", values[i])
			} else if value.Valid {
				tpu.Start = int32(value.Int64)
			}
		case threeprimeutr.FieldEnd:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field end", values[i])
			} else if value.Valid {
				tpu.End = int32(value.Int64)
			}
		case threeprimeutr.FieldStrand:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field strand", values[i])
			} else if value.Valid {
				tpu.Strand = value.String
			}
		case threeprimeutr.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field transcript_three_prime_utr", values[i])
			} else if value.Valid {
				tpu.transcript_three_prime_utr = new(string)
				*tpu.transcript_three_prime_utr = value.String
			}
		}
	}
	return nil
}

// QueryTranscript queries the "transcript" edge of the ThreePrimeUtr entity.
func (tpu *ThreePrimeUtr) QueryTranscript() *TranscriptQuery {
	return (&ThreePrimeUtrClient{config: tpu.config}).QueryTranscript(tpu)
}

// Update returns a builder for updating this ThreePrimeUtr.
// Note that you need to call ThreePrimeUtr.Unwrap() before calling this method if this ThreePrimeUtr
// was returned from a transaction, and the transaction was committed or rolled back.
func (tpu *ThreePrimeUtr) Update() *ThreePrimeUtrUpdateOne {
	return (&ThreePrimeUtrClient{config: tpu.config}).UpdateOne(tpu)
}

// Unwrap unwraps the ThreePrimeUtr entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tpu *ThreePrimeUtr) Unwrap() *ThreePrimeUtr {
	_tx, ok := tpu.config.driver.(*txDriver)
	if !ok {
		panic("ent: ThreePrimeUtr is not a transactional entity")
	}
	tpu.config.driver = _tx.drv
	return tpu
}

// String implements the fmt.Stringer.
func (tpu *ThreePrimeUtr) String() string {
	var builder strings.Builder
	builder.WriteString("ThreePrimeUtr(")
	builder.WriteString(fmt.Sprintf("id=%v, ", tpu.ID))
	builder.WriteString("seqname=")
	builder.WriteString(tpu.Seqname)
	builder.WriteString(", ")
	builder.WriteString("start=")
	builder.WriteString(fmt.Sprintf("%v", tpu.Start))
	builder.WriteString(", ")
	builder.WriteString("end=")
	builder.WriteString(fmt.Sprintf("%v", tpu.End))
	builder.WriteString(", ")
	builder.WriteString("strand=")
	builder.WriteString(tpu.Strand)
	builder.WriteByte(')')
	return builder.String()
}

// ThreePrimeUtrs is a parsable slice of ThreePrimeUtr.
type ThreePrimeUtrs []*ThreePrimeUtr

func (tpu ThreePrimeUtrs) config(cfg config) {
	for _i := range tpu {
		tpu[_i].config = cfg
	}
}
