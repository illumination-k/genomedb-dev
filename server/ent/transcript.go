// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"genomedb/bio/gffio"
	"genomedb/ent/locus"
	"genomedb/ent/transcript"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Transcript is the model entity for the Transcript schema.
type Transcript struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Seqname holds the value of the "seqname" field.
	Seqname string `json:"seqname,omitempty"`
	// Strand holds the value of the "strand" field.
	Strand string `json:"strand,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// Source holds the value of the "source" field.
	Source string `json:"source,omitempty"`
	// Start holds the value of the "start" field.
	Start int32 `json:"start,omitempty"`
	// End holds the value of the "end" field.
	End int32 `json:"end,omitempty"`
	// Exon holds the value of the "exon" field.
	Exon []gffio.GffRecord `json:"exon,omitempty"`
	// FivePrimeUtr holds the value of the "five_prime_utr" field.
	FivePrimeUtr []gffio.GffRecord `json:"five_prime_utr,omitempty"`
	// ThreePrimeUtr holds the value of the "three_prime_utr" field.
	ThreePrimeUtr []gffio.GffRecord `json:"three_prime_utr,omitempty"`
	// Cds holds the value of the "cds" field.
	Cds []gffio.GffRecord `json:"cds,omitempty"`
	// GenomicSequence holds the value of the "genomic_sequence" field.
	GenomicSequence string `json:"genomic_sequence,omitempty"`
	// ExonSequence holds the value of the "exon_sequence" field.
	ExonSequence string `json:"exon_sequence,omitempty"`
	// CdsSequence holds the value of the "cds_sequence" field.
	CdsSequence string `json:"cds_sequence,omitempty"`
	// ProteinSequence holds the value of the "protein_sequence" field.
	ProteinSequence string `json:"protein_sequence,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TranscriptQuery when eager-loading is set.
	Edges             TranscriptEdges `json:"edges"`
	locus_transcripts *string
}

// TranscriptEdges holds the relations/edges for other nodes in the graph.
type TranscriptEdges struct {
	// Locus holds the value of the locus edge.
	Locus *Locus `json:"locus,omitempty"`
	// Goterms holds the value of the goterms edge.
	Goterms []*GoTerm `json:"goterms,omitempty"`
	// Domains holds the value of the domains edge.
	Domains []*DomainAnnotation `json:"domains,omitempty"`
	// GotermTranscript holds the value of the goterm_transcript edge.
	GotermTranscript []*GoTermOnTranscripts `json:"goterm_transcript,omitempty"`
	// DomainTranscript holds the value of the domain_transcript edge.
	DomainTranscript []*DomainAnnotationToTranscript `json:"domain_transcript,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// LocusOrErr returns the Locus value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TranscriptEdges) LocusOrErr() (*Locus, error) {
	if e.loadedTypes[0] {
		if e.Locus == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: locus.Label}
		}
		return e.Locus, nil
	}
	return nil, &NotLoadedError{edge: "locus"}
}

// GotermsOrErr returns the Goterms value or an error if the edge
// was not loaded in eager-loading.
func (e TranscriptEdges) GotermsOrErr() ([]*GoTerm, error) {
	if e.loadedTypes[1] {
		return e.Goterms, nil
	}
	return nil, &NotLoadedError{edge: "goterms"}
}

// DomainsOrErr returns the Domains value or an error if the edge
// was not loaded in eager-loading.
func (e TranscriptEdges) DomainsOrErr() ([]*DomainAnnotation, error) {
	if e.loadedTypes[2] {
		return e.Domains, nil
	}
	return nil, &NotLoadedError{edge: "domains"}
}

// GotermTranscriptOrErr returns the GotermTranscript value or an error if the edge
// was not loaded in eager-loading.
func (e TranscriptEdges) GotermTranscriptOrErr() ([]*GoTermOnTranscripts, error) {
	if e.loadedTypes[3] {
		return e.GotermTranscript, nil
	}
	return nil, &NotLoadedError{edge: "goterm_transcript"}
}

// DomainTranscriptOrErr returns the DomainTranscript value or an error if the edge
// was not loaded in eager-loading.
func (e TranscriptEdges) DomainTranscriptOrErr() ([]*DomainAnnotationToTranscript, error) {
	if e.loadedTypes[4] {
		return e.DomainTranscript, nil
	}
	return nil, &NotLoadedError{edge: "domain_transcript"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Transcript) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case transcript.FieldExon, transcript.FieldFivePrimeUtr, transcript.FieldThreePrimeUtr, transcript.FieldCds:
			values[i] = new([]byte)
		case transcript.FieldStart, transcript.FieldEnd:
			values[i] = new(sql.NullInt64)
		case transcript.FieldID, transcript.FieldSeqname, transcript.FieldStrand, transcript.FieldType, transcript.FieldSource, transcript.FieldGenomicSequence, transcript.FieldExonSequence, transcript.FieldCdsSequence, transcript.FieldProteinSequence:
			values[i] = new(sql.NullString)
		case transcript.ForeignKeys[0]: // locus_transcripts
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Transcript", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Transcript fields.
func (t *Transcript) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case transcript.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				t.ID = value.String
			}
		case transcript.FieldSeqname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field seqname", values[i])
			} else if value.Valid {
				t.Seqname = value.String
			}
		case transcript.FieldStrand:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field strand", values[i])
			} else if value.Valid {
				t.Strand = value.String
			}
		case transcript.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				t.Type = value.String
			}
		case transcript.FieldSource:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field source", values[i])
			} else if value.Valid {
				t.Source = value.String
			}
		case transcript.FieldStart:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field start", values[i])
			} else if value.Valid {
				t.Start = int32(value.Int64)
			}
		case transcript.FieldEnd:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field end", values[i])
			} else if value.Valid {
				t.End = int32(value.Int64)
			}
		case transcript.FieldExon:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field exon", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &t.Exon); err != nil {
					return fmt.Errorf("unmarshal field exon: %w", err)
				}
			}
		case transcript.FieldFivePrimeUtr:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field five_prime_utr", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &t.FivePrimeUtr); err != nil {
					return fmt.Errorf("unmarshal field five_prime_utr: %w", err)
				}
			}
		case transcript.FieldThreePrimeUtr:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field three_prime_utr", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &t.ThreePrimeUtr); err != nil {
					return fmt.Errorf("unmarshal field three_prime_utr: %w", err)
				}
			}
		case transcript.FieldCds:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field cds", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &t.Cds); err != nil {
					return fmt.Errorf("unmarshal field cds: %w", err)
				}
			}
		case transcript.FieldGenomicSequence:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field genomic_sequence", values[i])
			} else if value.Valid {
				t.GenomicSequence = value.String
			}
		case transcript.FieldExonSequence:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field exon_sequence", values[i])
			} else if value.Valid {
				t.ExonSequence = value.String
			}
		case transcript.FieldCdsSequence:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cds_sequence", values[i])
			} else if value.Valid {
				t.CdsSequence = value.String
			}
		case transcript.FieldProteinSequence:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field protein_sequence", values[i])
			} else if value.Valid {
				t.ProteinSequence = value.String
			}
		case transcript.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field locus_transcripts", values[i])
			} else if value.Valid {
				t.locus_transcripts = new(string)
				*t.locus_transcripts = value.String
			}
		}
	}
	return nil
}

// QueryLocus queries the "locus" edge of the Transcript entity.
func (t *Transcript) QueryLocus() *LocusQuery {
	return (&TranscriptClient{config: t.config}).QueryLocus(t)
}

// QueryGoterms queries the "goterms" edge of the Transcript entity.
func (t *Transcript) QueryGoterms() *GoTermQuery {
	return (&TranscriptClient{config: t.config}).QueryGoterms(t)
}

// QueryDomains queries the "domains" edge of the Transcript entity.
func (t *Transcript) QueryDomains() *DomainAnnotationQuery {
	return (&TranscriptClient{config: t.config}).QueryDomains(t)
}

// QueryGotermTranscript queries the "goterm_transcript" edge of the Transcript entity.
func (t *Transcript) QueryGotermTranscript() *GoTermOnTranscriptsQuery {
	return (&TranscriptClient{config: t.config}).QueryGotermTranscript(t)
}

// QueryDomainTranscript queries the "domain_transcript" edge of the Transcript entity.
func (t *Transcript) QueryDomainTranscript() *DomainAnnotationToTranscriptQuery {
	return (&TranscriptClient{config: t.config}).QueryDomainTranscript(t)
}

// Update returns a builder for updating this Transcript.
// Note that you need to call Transcript.Unwrap() before calling this method if this Transcript
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Transcript) Update() *TranscriptUpdateOne {
	return (&TranscriptClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Transcript entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Transcript) Unwrap() *Transcript {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Transcript is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Transcript) String() string {
	var builder strings.Builder
	builder.WriteString("Transcript(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("seqname=")
	builder.WriteString(t.Seqname)
	builder.WriteString(", ")
	builder.WriteString("strand=")
	builder.WriteString(t.Strand)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(t.Type)
	builder.WriteString(", ")
	builder.WriteString("source=")
	builder.WriteString(t.Source)
	builder.WriteString(", ")
	builder.WriteString("start=")
	builder.WriteString(fmt.Sprintf("%v", t.Start))
	builder.WriteString(", ")
	builder.WriteString("end=")
	builder.WriteString(fmt.Sprintf("%v", t.End))
	builder.WriteString(", ")
	builder.WriteString("exon=")
	builder.WriteString(fmt.Sprintf("%v", t.Exon))
	builder.WriteString(", ")
	builder.WriteString("five_prime_utr=")
	builder.WriteString(fmt.Sprintf("%v", t.FivePrimeUtr))
	builder.WriteString(", ")
	builder.WriteString("three_prime_utr=")
	builder.WriteString(fmt.Sprintf("%v", t.ThreePrimeUtr))
	builder.WriteString(", ")
	builder.WriteString("cds=")
	builder.WriteString(fmt.Sprintf("%v", t.Cds))
	builder.WriteString(", ")
	builder.WriteString("genomic_sequence=")
	builder.WriteString(t.GenomicSequence)
	builder.WriteString(", ")
	builder.WriteString("exon_sequence=")
	builder.WriteString(t.ExonSequence)
	builder.WriteString(", ")
	builder.WriteString("cds_sequence=")
	builder.WriteString(t.CdsSequence)
	builder.WriteString(", ")
	builder.WriteString("protein_sequence=")
	builder.WriteString(t.ProteinSequence)
	builder.WriteByte(')')
	return builder.String()
}

// Transcripts is a parsable slice of Transcript.
type Transcripts []*Transcript

func (t Transcripts) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
