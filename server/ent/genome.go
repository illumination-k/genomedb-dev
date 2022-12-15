// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"genomedb/ent/genome"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Genome is the model entity for the Genome schema.
type Genome struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// See https://www.ncbi.nlm.nih.gov/Taxonomy/taxonomyhome.html/index.cgi?chapter=tgencodes
	CodonTable int32 `json:"codon_table,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GenomeQuery when eager-loading is set.
	Edges GenomeEdges `json:"edges"`
}

// GenomeEdges holds the relations/edges for other nodes in the graph.
type GenomeEdges struct {
	// Genes holds the value of the genes edge.
	Genes []*Gene `json:"genes,omitempty"`
	// Scaffolds holds the value of the scaffolds edge.
	Scaffolds []*Scaffold `json:"scaffolds,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// GenesOrErr returns the Genes value or an error if the edge
// was not loaded in eager-loading.
func (e GenomeEdges) GenesOrErr() ([]*Gene, error) {
	if e.loadedTypes[0] {
		return e.Genes, nil
	}
	return nil, &NotLoadedError{edge: "genes"}
}

// ScaffoldsOrErr returns the Scaffolds value or an error if the edge
// was not loaded in eager-loading.
func (e GenomeEdges) ScaffoldsOrErr() ([]*Scaffold, error) {
	if e.loadedTypes[1] {
		return e.Scaffolds, nil
	}
	return nil, &NotLoadedError{edge: "scaffolds"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Genome) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case genome.FieldCodonTable:
			values[i] = new(sql.NullInt64)
		case genome.FieldID:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Genome", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Genome fields.
func (ge *Genome) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case genome.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				ge.ID = value.String
			}
		case genome.FieldCodonTable:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field codon_table", values[i])
			} else if value.Valid {
				ge.CodonTable = int32(value.Int64)
			}
		}
	}
	return nil
}

// QueryGenes queries the "genes" edge of the Genome entity.
func (ge *Genome) QueryGenes() *GeneQuery {
	return (&GenomeClient{config: ge.config}).QueryGenes(ge)
}

// QueryScaffolds queries the "scaffolds" edge of the Genome entity.
func (ge *Genome) QueryScaffolds() *ScaffoldQuery {
	return (&GenomeClient{config: ge.config}).QueryScaffolds(ge)
}

// Update returns a builder for updating this Genome.
// Note that you need to call Genome.Unwrap() before calling this method if this Genome
// was returned from a transaction, and the transaction was committed or rolled back.
func (ge *Genome) Update() *GenomeUpdateOne {
	return (&GenomeClient{config: ge.config}).UpdateOne(ge)
}

// Unwrap unwraps the Genome entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ge *Genome) Unwrap() *Genome {
	_tx, ok := ge.config.driver.(*txDriver)
	if !ok {
		panic("ent: Genome is not a transactional entity")
	}
	ge.config.driver = _tx.drv
	return ge
}

// String implements the fmt.Stringer.
func (ge *Genome) String() string {
	var builder strings.Builder
	builder.WriteString("Genome(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ge.ID))
	builder.WriteString("codon_table=")
	builder.WriteString(fmt.Sprintf("%v", ge.CodonTable))
	builder.WriteByte(')')
	return builder.String()
}

// Genomes is a parsable slice of Genome.
type Genomes []*Genome

func (ge Genomes) config(cfg config) {
	for _i := range ge {
		ge[_i].config = cfg
	}
}
