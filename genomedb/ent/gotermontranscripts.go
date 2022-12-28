// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"genomedb/ent/goterm"
	"genomedb/ent/gotermontranscripts"
	"genomedb/ent/transcript"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// GoTermOnTranscripts is the model entity for the GoTermOnTranscripts schema.
type GoTermOnTranscripts struct {
	config `json:"-"`
	// EvidenceCode holds the value of the "evidence_code" field.
	EvidenceCode string `json:"evidence_code,omitempty"`
	// GoTermID holds the value of the "go_term_id" field.
	GoTermID string `json:"go_term_id,omitempty"`
	// TranscriptID holds the value of the "transcript_id" field.
	TranscriptID string `json:"transcript_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GoTermOnTranscriptsQuery when eager-loading is set.
	Edges GoTermOnTranscriptsEdges `json:"edges"`
}

// GoTermOnTranscriptsEdges holds the relations/edges for other nodes in the graph.
type GoTermOnTranscriptsEdges struct {
	// GoTerm holds the value of the go_term edge.
	GoTerm *GoTerm `json:"go_term,omitempty"`
	// Transcript holds the value of the transcript edge.
	Transcript *Transcript `json:"transcript,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// GoTermOrErr returns the GoTerm value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GoTermOnTranscriptsEdges) GoTermOrErr() (*GoTerm, error) {
	if e.loadedTypes[0] {
		if e.GoTerm == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: goterm.Label}
		}
		return e.GoTerm, nil
	}
	return nil, &NotLoadedError{edge: "go_term"}
}

// TranscriptOrErr returns the Transcript value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GoTermOnTranscriptsEdges) TranscriptOrErr() (*Transcript, error) {
	if e.loadedTypes[1] {
		if e.Transcript == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: transcript.Label}
		}
		return e.Transcript, nil
	}
	return nil, &NotLoadedError{edge: "transcript"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GoTermOnTranscripts) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case gotermontranscripts.FieldEvidenceCode, gotermontranscripts.FieldGoTermID, gotermontranscripts.FieldTranscriptID:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type GoTermOnTranscripts", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GoTermOnTranscripts fields.
func (gtot *GoTermOnTranscripts) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case gotermontranscripts.FieldEvidenceCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field evidence_code", values[i])
			} else if value.Valid {
				gtot.EvidenceCode = value.String
			}
		case gotermontranscripts.FieldGoTermID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field go_term_id", values[i])
			} else if value.Valid {
				gtot.GoTermID = value.String
			}
		case gotermontranscripts.FieldTranscriptID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field transcript_id", values[i])
			} else if value.Valid {
				gtot.TranscriptID = value.String
			}
		}
	}
	return nil
}

// QueryGoTerm queries the "go_term" edge of the GoTermOnTranscripts entity.
func (gtot *GoTermOnTranscripts) QueryGoTerm() *GoTermQuery {
	return (&GoTermOnTranscriptsClient{config: gtot.config}).QueryGoTerm(gtot)
}

// QueryTranscript queries the "transcript" edge of the GoTermOnTranscripts entity.
func (gtot *GoTermOnTranscripts) QueryTranscript() *TranscriptQuery {
	return (&GoTermOnTranscriptsClient{config: gtot.config}).QueryTranscript(gtot)
}

// Update returns a builder for updating this GoTermOnTranscripts.
// Note that you need to call GoTermOnTranscripts.Unwrap() before calling this method if this GoTermOnTranscripts
// was returned from a transaction, and the transaction was committed or rolled back.
func (gtot *GoTermOnTranscripts) Update() *GoTermOnTranscriptsUpdateOne {
	return (&GoTermOnTranscriptsClient{config: gtot.config}).UpdateOne(gtot)
}

// Unwrap unwraps the GoTermOnTranscripts entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gtot *GoTermOnTranscripts) Unwrap() *GoTermOnTranscripts {
	_tx, ok := gtot.config.driver.(*txDriver)
	if !ok {
		panic("ent: GoTermOnTranscripts is not a transactional entity")
	}
	gtot.config.driver = _tx.drv
	return gtot
}

// String implements the fmt.Stringer.
func (gtot *GoTermOnTranscripts) String() string {
	var builder strings.Builder
	builder.WriteString("GoTermOnTranscripts(")
	builder.WriteString("evidence_code=")
	builder.WriteString(gtot.EvidenceCode)
	builder.WriteString(", ")
	builder.WriteString("go_term_id=")
	builder.WriteString(gtot.GoTermID)
	builder.WriteString(", ")
	builder.WriteString("transcript_id=")
	builder.WriteString(gtot.TranscriptID)
	builder.WriteByte(')')
	return builder.String()
}

// GoTermOnTranscriptsSlice is a parsable slice of GoTermOnTranscripts.
type GoTermOnTranscriptsSlice []*GoTermOnTranscripts

func (gtot GoTermOnTranscriptsSlice) config(cfg config) {
	for _i := range gtot {
		gtot[_i].config = cfg
	}
}
