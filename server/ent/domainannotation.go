// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"genomedb/ent/domainannotation"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// DomainAnnotation is the model entity for the DomainAnnotation schema.
type DomainAnnotation struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Analysis holds the value of the "Analysis" field.
	Analysis domainannotation.Analysis `json:"Analysis,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DomainAnnotationQuery when eager-loading is set.
	Edges DomainAnnotationEdges `json:"edges"`
}

// DomainAnnotationEdges holds the relations/edges for other nodes in the graph.
type DomainAnnotationEdges struct {
	// Transcripts holds the value of the transcripts edge.
	Transcripts []*Transcript `json:"transcripts,omitempty"`
	// DomainTranscript holds the value of the domain_transcript edge.
	DomainTranscript []*DomainAnnotationToTranscript `json:"domain_transcript,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// TranscriptsOrErr returns the Transcripts value or an error if the edge
// was not loaded in eager-loading.
func (e DomainAnnotationEdges) TranscriptsOrErr() ([]*Transcript, error) {
	if e.loadedTypes[0] {
		return e.Transcripts, nil
	}
	return nil, &NotLoadedError{edge: "transcripts"}
}

// DomainTranscriptOrErr returns the DomainTranscript value or an error if the edge
// was not loaded in eager-loading.
func (e DomainAnnotationEdges) DomainTranscriptOrErr() ([]*DomainAnnotationToTranscript, error) {
	if e.loadedTypes[1] {
		return e.DomainTranscript, nil
	}
	return nil, &NotLoadedError{edge: "domain_transcript"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DomainAnnotation) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case domainannotation.FieldID, domainannotation.FieldDescription, domainannotation.FieldAnalysis:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type DomainAnnotation", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DomainAnnotation fields.
func (da *DomainAnnotation) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case domainannotation.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				da.ID = value.String
			}
		case domainannotation.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				da.Description = value.String
			}
		case domainannotation.FieldAnalysis:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Analysis", values[i])
			} else if value.Valid {
				da.Analysis = domainannotation.Analysis(value.String)
			}
		}
	}
	return nil
}

// QueryTranscripts queries the "transcripts" edge of the DomainAnnotation entity.
func (da *DomainAnnotation) QueryTranscripts() *TranscriptQuery {
	return (&DomainAnnotationClient{config: da.config}).QueryTranscripts(da)
}

// QueryDomainTranscript queries the "domain_transcript" edge of the DomainAnnotation entity.
func (da *DomainAnnotation) QueryDomainTranscript() *DomainAnnotationToTranscriptQuery {
	return (&DomainAnnotationClient{config: da.config}).QueryDomainTranscript(da)
}

// Update returns a builder for updating this DomainAnnotation.
// Note that you need to call DomainAnnotation.Unwrap() before calling this method if this DomainAnnotation
// was returned from a transaction, and the transaction was committed or rolled back.
func (da *DomainAnnotation) Update() *DomainAnnotationUpdateOne {
	return (&DomainAnnotationClient{config: da.config}).UpdateOne(da)
}

// Unwrap unwraps the DomainAnnotation entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (da *DomainAnnotation) Unwrap() *DomainAnnotation {
	_tx, ok := da.config.driver.(*txDriver)
	if !ok {
		panic("ent: DomainAnnotation is not a transactional entity")
	}
	da.config.driver = _tx.drv
	return da
}

// String implements the fmt.Stringer.
func (da *DomainAnnotation) String() string {
	var builder strings.Builder
	builder.WriteString("DomainAnnotation(")
	builder.WriteString(fmt.Sprintf("id=%v, ", da.ID))
	builder.WriteString("description=")
	builder.WriteString(da.Description)
	builder.WriteString(", ")
	builder.WriteString("Analysis=")
	builder.WriteString(fmt.Sprintf("%v", da.Analysis))
	builder.WriteByte(')')
	return builder.String()
}

// DomainAnnotations is a parsable slice of DomainAnnotation.
type DomainAnnotations []*DomainAnnotation

func (da DomainAnnotations) config(cfg config) {
	for _i := range da {
		da[_i].config = cfg
	}
}
