// Code generated by ent, DO NOT EDIT.

package transcript

import (
	"genomedb/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Seqname applies equality check predicate on the "seqname" field. It's identical to SeqnameEQ.
func Seqname(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSeqname), v))
	})
}

// Strand applies equality check predicate on the "strand" field. It's identical to StrandEQ.
func Strand(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStrand), v))
	})
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// Start applies equality check predicate on the "start" field. It's identical to StartEQ.
func Start(v int32) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStart), v))
	})
}

// End applies equality check predicate on the "end" field. It's identical to EndEQ.
func End(v int32) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEnd), v))
	})
}

// GenomicSequence applies equality check predicate on the "genomic_sequence" field. It's identical to GenomicSequenceEQ.
func GenomicSequence(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGenomicSequence), v))
	})
}

// ExonSequence applies equality check predicate on the "exon_sequence" field. It's identical to ExonSequenceEQ.
func ExonSequence(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExonSequence), v))
	})
}

// CdsSequence applies equality check predicate on the "cds_sequence" field. It's identical to CdsSequenceEQ.
func CdsSequence(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCdsSequence), v))
	})
}

// ProteinSequence applies equality check predicate on the "protein_sequence" field. It's identical to ProteinSequenceEQ.
func ProteinSequence(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProteinSequence), v))
	})
}

// SeqnameEQ applies the EQ predicate on the "seqname" field.
func SeqnameEQ(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSeqname), v))
	})
}

// SeqnameNEQ applies the NEQ predicate on the "seqname" field.
func SeqnameNEQ(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSeqname), v))
	})
}

// SeqnameIn applies the In predicate on the "seqname" field.
func SeqnameIn(vs ...string) predicate.Transcript {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSeqname), v...))
	})
}

// SeqnameNotIn applies the NotIn predicate on the "seqname" field.
func SeqnameNotIn(vs ...string) predicate.Transcript {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSeqname), v...))
	})
}

// SeqnameGT applies the GT predicate on the "seqname" field.
func SeqnameGT(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSeqname), v))
	})
}

// SeqnameGTE applies the GTE predicate on the "seqname" field.
func SeqnameGTE(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSeqname), v))
	})
}

// SeqnameLT applies the LT predicate on the "seqname" field.
func SeqnameLT(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSeqname), v))
	})
}

// SeqnameLTE applies the LTE predicate on the "seqname" field.
func SeqnameLTE(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSeqname), v))
	})
}

// SeqnameContains applies the Contains predicate on the "seqname" field.
func SeqnameContains(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSeqname), v))
	})
}

// SeqnameHasPrefix applies the HasPrefix predicate on the "seqname" field.
func SeqnameHasPrefix(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSeqname), v))
	})
}

// SeqnameHasSuffix applies the HasSuffix predicate on the "seqname" field.
func SeqnameHasSuffix(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSeqname), v))
	})
}

// SeqnameEqualFold applies the EqualFold predicate on the "seqname" field.
func SeqnameEqualFold(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSeqname), v))
	})
}

// SeqnameContainsFold applies the ContainsFold predicate on the "seqname" field.
func SeqnameContainsFold(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSeqname), v))
	})
}

// StrandEQ applies the EQ predicate on the "strand" field.
func StrandEQ(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStrand), v))
	})
}

// StrandNEQ applies the NEQ predicate on the "strand" field.
func StrandNEQ(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStrand), v))
	})
}

// StrandIn applies the In predicate on the "strand" field.
func StrandIn(vs ...string) predicate.Transcript {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStrand), v...))
	})
}

// StrandNotIn applies the NotIn predicate on the "strand" field.
func StrandNotIn(vs ...string) predicate.Transcript {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStrand), v...))
	})
}

// StrandGT applies the GT predicate on the "strand" field.
func StrandGT(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStrand), v))
	})
}

// StrandGTE applies the GTE predicate on the "strand" field.
func StrandGTE(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStrand), v))
	})
}

// StrandLT applies the LT predicate on the "strand" field.
func StrandLT(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStrand), v))
	})
}

// StrandLTE applies the LTE predicate on the "strand" field.
func StrandLTE(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStrand), v))
	})
}

// StrandContains applies the Contains predicate on the "strand" field.
func StrandContains(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStrand), v))
	})
}

// StrandHasPrefix applies the HasPrefix predicate on the "strand" field.
func StrandHasPrefix(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStrand), v))
	})
}

// StrandHasSuffix applies the HasSuffix predicate on the "strand" field.
func StrandHasSuffix(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStrand), v))
	})
}

// StrandEqualFold applies the EqualFold predicate on the "strand" field.
func StrandEqualFold(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStrand), v))
	})
}

// StrandContainsFold applies the ContainsFold predicate on the "strand" field.
func StrandContainsFold(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStrand), v))
	})
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldType), v))
	})
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.Transcript {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldType), v...))
	})
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.Transcript {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldType), v...))
	})
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldType), v))
	})
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldType), v))
	})
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldType), v))
	})
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldType), v))
	})
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldType), v))
	})
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldType), v))
	})
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldType), v))
	})
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldType), v))
	})
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldType), v))
	})
}

// StartEQ applies the EQ predicate on the "start" field.
func StartEQ(v int32) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStart), v))
	})
}

// StartNEQ applies the NEQ predicate on the "start" field.
func StartNEQ(v int32) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStart), v))
	})
}

// StartIn applies the In predicate on the "start" field.
func StartIn(vs ...int32) predicate.Transcript {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStart), v...))
	})
}

// StartNotIn applies the NotIn predicate on the "start" field.
func StartNotIn(vs ...int32) predicate.Transcript {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStart), v...))
	})
}

// StartGT applies the GT predicate on the "start" field.
func StartGT(v int32) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStart), v))
	})
}

// StartGTE applies the GTE predicate on the "start" field.
func StartGTE(v int32) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStart), v))
	})
}

// StartLT applies the LT predicate on the "start" field.
func StartLT(v int32) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStart), v))
	})
}

// StartLTE applies the LTE predicate on the "start" field.
func StartLTE(v int32) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStart), v))
	})
}

// EndEQ applies the EQ predicate on the "end" field.
func EndEQ(v int32) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEnd), v))
	})
}

// EndNEQ applies the NEQ predicate on the "end" field.
func EndNEQ(v int32) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEnd), v))
	})
}

// EndIn applies the In predicate on the "end" field.
func EndIn(vs ...int32) predicate.Transcript {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEnd), v...))
	})
}

// EndNotIn applies the NotIn predicate on the "end" field.
func EndNotIn(vs ...int32) predicate.Transcript {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEnd), v...))
	})
}

// EndGT applies the GT predicate on the "end" field.
func EndGT(v int32) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEnd), v))
	})
}

// EndGTE applies the GTE predicate on the "end" field.
func EndGTE(v int32) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEnd), v))
	})
}

// EndLT applies the LT predicate on the "end" field.
func EndLT(v int32) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEnd), v))
	})
}

// EndLTE applies the LTE predicate on the "end" field.
func EndLTE(v int32) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEnd), v))
	})
}

// GenomicSequenceEQ applies the EQ predicate on the "genomic_sequence" field.
func GenomicSequenceEQ(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGenomicSequence), v))
	})
}

// GenomicSequenceNEQ applies the NEQ predicate on the "genomic_sequence" field.
func GenomicSequenceNEQ(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGenomicSequence), v))
	})
}

// GenomicSequenceIn applies the In predicate on the "genomic_sequence" field.
func GenomicSequenceIn(vs ...string) predicate.Transcript {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldGenomicSequence), v...))
	})
}

// GenomicSequenceNotIn applies the NotIn predicate on the "genomic_sequence" field.
func GenomicSequenceNotIn(vs ...string) predicate.Transcript {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldGenomicSequence), v...))
	})
}

// GenomicSequenceGT applies the GT predicate on the "genomic_sequence" field.
func GenomicSequenceGT(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGenomicSequence), v))
	})
}

// GenomicSequenceGTE applies the GTE predicate on the "genomic_sequence" field.
func GenomicSequenceGTE(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGenomicSequence), v))
	})
}

// GenomicSequenceLT applies the LT predicate on the "genomic_sequence" field.
func GenomicSequenceLT(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGenomicSequence), v))
	})
}

// GenomicSequenceLTE applies the LTE predicate on the "genomic_sequence" field.
func GenomicSequenceLTE(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGenomicSequence), v))
	})
}

// GenomicSequenceContains applies the Contains predicate on the "genomic_sequence" field.
func GenomicSequenceContains(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldGenomicSequence), v))
	})
}

// GenomicSequenceHasPrefix applies the HasPrefix predicate on the "genomic_sequence" field.
func GenomicSequenceHasPrefix(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldGenomicSequence), v))
	})
}

// GenomicSequenceHasSuffix applies the HasSuffix predicate on the "genomic_sequence" field.
func GenomicSequenceHasSuffix(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldGenomicSequence), v))
	})
}

// GenomicSequenceEqualFold applies the EqualFold predicate on the "genomic_sequence" field.
func GenomicSequenceEqualFold(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldGenomicSequence), v))
	})
}

// GenomicSequenceContainsFold applies the ContainsFold predicate on the "genomic_sequence" field.
func GenomicSequenceContainsFold(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldGenomicSequence), v))
	})
}

// ExonSequenceEQ applies the EQ predicate on the "exon_sequence" field.
func ExonSequenceEQ(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExonSequence), v))
	})
}

// ExonSequenceNEQ applies the NEQ predicate on the "exon_sequence" field.
func ExonSequenceNEQ(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldExonSequence), v))
	})
}

// ExonSequenceIn applies the In predicate on the "exon_sequence" field.
func ExonSequenceIn(vs ...string) predicate.Transcript {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldExonSequence), v...))
	})
}

// ExonSequenceNotIn applies the NotIn predicate on the "exon_sequence" field.
func ExonSequenceNotIn(vs ...string) predicate.Transcript {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldExonSequence), v...))
	})
}

// ExonSequenceGT applies the GT predicate on the "exon_sequence" field.
func ExonSequenceGT(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldExonSequence), v))
	})
}

// ExonSequenceGTE applies the GTE predicate on the "exon_sequence" field.
func ExonSequenceGTE(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldExonSequence), v))
	})
}

// ExonSequenceLT applies the LT predicate on the "exon_sequence" field.
func ExonSequenceLT(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldExonSequence), v))
	})
}

// ExonSequenceLTE applies the LTE predicate on the "exon_sequence" field.
func ExonSequenceLTE(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldExonSequence), v))
	})
}

// ExonSequenceContains applies the Contains predicate on the "exon_sequence" field.
func ExonSequenceContains(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldExonSequence), v))
	})
}

// ExonSequenceHasPrefix applies the HasPrefix predicate on the "exon_sequence" field.
func ExonSequenceHasPrefix(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldExonSequence), v))
	})
}

// ExonSequenceHasSuffix applies the HasSuffix predicate on the "exon_sequence" field.
func ExonSequenceHasSuffix(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldExonSequence), v))
	})
}

// ExonSequenceEqualFold applies the EqualFold predicate on the "exon_sequence" field.
func ExonSequenceEqualFold(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldExonSequence), v))
	})
}

// ExonSequenceContainsFold applies the ContainsFold predicate on the "exon_sequence" field.
func ExonSequenceContainsFold(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldExonSequence), v))
	})
}

// CdsSequenceEQ applies the EQ predicate on the "cds_sequence" field.
func CdsSequenceEQ(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCdsSequence), v))
	})
}

// CdsSequenceNEQ applies the NEQ predicate on the "cds_sequence" field.
func CdsSequenceNEQ(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCdsSequence), v))
	})
}

// CdsSequenceIn applies the In predicate on the "cds_sequence" field.
func CdsSequenceIn(vs ...string) predicate.Transcript {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCdsSequence), v...))
	})
}

// CdsSequenceNotIn applies the NotIn predicate on the "cds_sequence" field.
func CdsSequenceNotIn(vs ...string) predicate.Transcript {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCdsSequence), v...))
	})
}

// CdsSequenceGT applies the GT predicate on the "cds_sequence" field.
func CdsSequenceGT(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCdsSequence), v))
	})
}

// CdsSequenceGTE applies the GTE predicate on the "cds_sequence" field.
func CdsSequenceGTE(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCdsSequence), v))
	})
}

// CdsSequenceLT applies the LT predicate on the "cds_sequence" field.
func CdsSequenceLT(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCdsSequence), v))
	})
}

// CdsSequenceLTE applies the LTE predicate on the "cds_sequence" field.
func CdsSequenceLTE(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCdsSequence), v))
	})
}

// CdsSequenceContains applies the Contains predicate on the "cds_sequence" field.
func CdsSequenceContains(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCdsSequence), v))
	})
}

// CdsSequenceHasPrefix applies the HasPrefix predicate on the "cds_sequence" field.
func CdsSequenceHasPrefix(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCdsSequence), v))
	})
}

// CdsSequenceHasSuffix applies the HasSuffix predicate on the "cds_sequence" field.
func CdsSequenceHasSuffix(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCdsSequence), v))
	})
}

// CdsSequenceEqualFold applies the EqualFold predicate on the "cds_sequence" field.
func CdsSequenceEqualFold(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCdsSequence), v))
	})
}

// CdsSequenceContainsFold applies the ContainsFold predicate on the "cds_sequence" field.
func CdsSequenceContainsFold(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCdsSequence), v))
	})
}

// ProteinSequenceEQ applies the EQ predicate on the "protein_sequence" field.
func ProteinSequenceEQ(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProteinSequence), v))
	})
}

// ProteinSequenceNEQ applies the NEQ predicate on the "protein_sequence" field.
func ProteinSequenceNEQ(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldProteinSequence), v))
	})
}

// ProteinSequenceIn applies the In predicate on the "protein_sequence" field.
func ProteinSequenceIn(vs ...string) predicate.Transcript {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldProteinSequence), v...))
	})
}

// ProteinSequenceNotIn applies the NotIn predicate on the "protein_sequence" field.
func ProteinSequenceNotIn(vs ...string) predicate.Transcript {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldProteinSequence), v...))
	})
}

// ProteinSequenceGT applies the GT predicate on the "protein_sequence" field.
func ProteinSequenceGT(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldProteinSequence), v))
	})
}

// ProteinSequenceGTE applies the GTE predicate on the "protein_sequence" field.
func ProteinSequenceGTE(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldProteinSequence), v))
	})
}

// ProteinSequenceLT applies the LT predicate on the "protein_sequence" field.
func ProteinSequenceLT(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldProteinSequence), v))
	})
}

// ProteinSequenceLTE applies the LTE predicate on the "protein_sequence" field.
func ProteinSequenceLTE(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldProteinSequence), v))
	})
}

// ProteinSequenceContains applies the Contains predicate on the "protein_sequence" field.
func ProteinSequenceContains(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldProteinSequence), v))
	})
}

// ProteinSequenceHasPrefix applies the HasPrefix predicate on the "protein_sequence" field.
func ProteinSequenceHasPrefix(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldProteinSequence), v))
	})
}

// ProteinSequenceHasSuffix applies the HasSuffix predicate on the "protein_sequence" field.
func ProteinSequenceHasSuffix(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldProteinSequence), v))
	})
}

// ProteinSequenceEqualFold applies the EqualFold predicate on the "protein_sequence" field.
func ProteinSequenceEqualFold(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldProteinSequence), v))
	})
}

// ProteinSequenceContainsFold applies the ContainsFold predicate on the "protein_sequence" field.
func ProteinSequenceContainsFold(v string) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldProteinSequence), v))
	})
}

// HasLocus applies the HasEdge predicate on the "locus" edge.
func HasLocus() predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(LocusTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, LocusTable, LocusColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLocusWith applies the HasEdge predicate on the "locus" edge with a given conditions (other predicates).
func HasLocusWith(preds ...predicate.Locus) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(LocusInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, LocusTable, LocusColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Transcript) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Transcript) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Transcript) predicate.Transcript {
	return predicate.Transcript(func(s *sql.Selector) {
		p(s.Not())
	})
}
