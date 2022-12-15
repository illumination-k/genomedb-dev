// Code generated by ent, DO NOT EDIT.

package fiveprimeutr

import (
	"genomedb/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.FivePrimeUtr {
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.FivePrimeUtr {
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.FivePrimeUtr {
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.FivePrimeUtr {
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.FivePrimeUtr {
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.FivePrimeUtr {
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.FivePrimeUtr {
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.FivePrimeUtr {
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.FivePrimeUtr {
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Seqname applies equality check predicate on the "seqname" field. It's identical to SeqnameEQ.
func Seqname(v string) predicate.FivePrimeUtr {
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSeqname), v))
	})
}

// Start applies equality check predicate on the "start" field. It's identical to StartEQ.
func Start(v int32) predicate.FivePrimeUtr {
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStart), v))
	})
}

// End applies equality check predicate on the "end" field. It's identical to EndEQ.
func End(v int32) predicate.FivePrimeUtr {
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEnd), v))
	})
}

// Strand applies equality check predicate on the "strand" field. It's identical to StrandEQ.
func Strand(v string) predicate.FivePrimeUtr {
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStrand), v))
	})
}

// SeqnameEQ applies the EQ predicate on the "seqname" field.
func SeqnameEQ(v string) predicate.FivePrimeUtr {
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSeqname), v))
	})
}

// SeqnameNEQ applies the NEQ predicate on the "seqname" field.
func SeqnameNEQ(v string) predicate.FivePrimeUtr {
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSeqname), v))
	})
}

// SeqnameIn applies the In predicate on the "seqname" field.
func SeqnameIn(vs ...string) predicate.FivePrimeUtr {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSeqname), v...))
	})
}

// SeqnameNotIn applies the NotIn predicate on the "seqname" field.
func SeqnameNotIn(vs ...string) predicate.FivePrimeUtr {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSeqname), v...))
	})
}

// SeqnameGT applies the GT predicate on the "seqname" field.
func SeqnameGT(v string) predicate.FivePrimeUtr {
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSeqname), v))
	})
}

// SeqnameGTE applies the GTE predicate on the "seqname" field.
func SeqnameGTE(v string) predicate.FivePrimeUtr {
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSeqname), v))
	})
}

// SeqnameLT applies the LT predicate on the "seqname" field.
func SeqnameLT(v string) predicate.FivePrimeUtr {
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSeqname), v))
	})
}

// SeqnameLTE applies the LTE predicate on the "seqname" field.
func SeqnameLTE(v string) predicate.FivePrimeUtr {
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSeqname), v))
	})
}

// SeqnameContains applies the Contains predicate on the "seqname" field.
func SeqnameContains(v string) predicate.FivePrimeUtr {
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSeqname), v))
	})
}

// SeqnameHasPrefix applies the HasPrefix predicate on the "seqname" field.
func SeqnameHasPrefix(v string) predicate.FivePrimeUtr {
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSeqname), v))
	})
}

// SeqnameHasSuffix applies the HasSuffix predicate on the "seqname" field.
func SeqnameHasSuffix(v string) predicate.FivePrimeUtr {
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSeqname), v))
	})
}

// SeqnameEqualFold applies the EqualFold predicate on the "seqname" field.
func SeqnameEqualFold(v string) predicate.FivePrimeUtr {
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSeqname), v))
	})
}

// SeqnameContainsFold applies the ContainsFold predicate on the "seqname" field.
func SeqnameContainsFold(v string) predicate.FivePrimeUtr {
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSeqname), v))
	})
}

// StartEQ applies the EQ predicate on the "start" field.
func StartEQ(v int32) predicate.FivePrimeUtr {
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStart), v))
	})
}

// StartNEQ applies the NEQ predicate on the "start" field.
func StartNEQ(v int32) predicate.FivePrimeUtr {
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStart), v))
	})
}

// StartIn applies the In predicate on the "start" field.
func StartIn(vs ...int32) predicate.FivePrimeUtr {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStart), v...))
	})
}

// StartNotIn applies the NotIn predicate on the "start" field.
func StartNotIn(vs ...int32) predicate.FivePrimeUtr {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStart), v...))
	})
}

// StartGT applies the GT predicate on the "start" field.
func StartGT(v int32) predicate.FivePrimeUtr {
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStart), v))
	})
}

// StartGTE applies the GTE predicate on the "start" field.
func StartGTE(v int32) predicate.FivePrimeUtr {
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStart), v))
	})
}

// StartLT applies the LT predicate on the "start" field.
func StartLT(v int32) predicate.FivePrimeUtr {
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStart), v))
	})
}

// StartLTE applies the LTE predicate on the "start" field.
func StartLTE(v int32) predicate.FivePrimeUtr {
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStart), v))
	})
}

// EndEQ applies the EQ predicate on the "end" field.
func EndEQ(v int32) predicate.FivePrimeUtr {
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEnd), v))
	})
}

// EndNEQ applies the NEQ predicate on the "end" field.
func EndNEQ(v int32) predicate.FivePrimeUtr {
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEnd), v))
	})
}

// EndIn applies the In predicate on the "end" field.
func EndIn(vs ...int32) predicate.FivePrimeUtr {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEnd), v...))
	})
}

// EndNotIn applies the NotIn predicate on the "end" field.
func EndNotIn(vs ...int32) predicate.FivePrimeUtr {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEnd), v...))
	})
}

// EndGT applies the GT predicate on the "end" field.
func EndGT(v int32) predicate.FivePrimeUtr {
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEnd), v))
	})
}

// EndGTE applies the GTE predicate on the "end" field.
func EndGTE(v int32) predicate.FivePrimeUtr {
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEnd), v))
	})
}

// EndLT applies the LT predicate on the "end" field.
func EndLT(v int32) predicate.FivePrimeUtr {
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEnd), v))
	})
}

// EndLTE applies the LTE predicate on the "end" field.
func EndLTE(v int32) predicate.FivePrimeUtr {
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEnd), v))
	})
}

// StrandEQ applies the EQ predicate on the "strand" field.
func StrandEQ(v string) predicate.FivePrimeUtr {
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStrand), v))
	})
}

// StrandNEQ applies the NEQ predicate on the "strand" field.
func StrandNEQ(v string) predicate.FivePrimeUtr {
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStrand), v))
	})
}

// StrandIn applies the In predicate on the "strand" field.
func StrandIn(vs ...string) predicate.FivePrimeUtr {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStrand), v...))
	})
}

// StrandNotIn applies the NotIn predicate on the "strand" field.
func StrandNotIn(vs ...string) predicate.FivePrimeUtr {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStrand), v...))
	})
}

// StrandGT applies the GT predicate on the "strand" field.
func StrandGT(v string) predicate.FivePrimeUtr {
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStrand), v))
	})
}

// StrandGTE applies the GTE predicate on the "strand" field.
func StrandGTE(v string) predicate.FivePrimeUtr {
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStrand), v))
	})
}

// StrandLT applies the LT predicate on the "strand" field.
func StrandLT(v string) predicate.FivePrimeUtr {
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStrand), v))
	})
}

// StrandLTE applies the LTE predicate on the "strand" field.
func StrandLTE(v string) predicate.FivePrimeUtr {
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStrand), v))
	})
}

// StrandContains applies the Contains predicate on the "strand" field.
func StrandContains(v string) predicate.FivePrimeUtr {
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStrand), v))
	})
}

// StrandHasPrefix applies the HasPrefix predicate on the "strand" field.
func StrandHasPrefix(v string) predicate.FivePrimeUtr {
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStrand), v))
	})
}

// StrandHasSuffix applies the HasSuffix predicate on the "strand" field.
func StrandHasSuffix(v string) predicate.FivePrimeUtr {
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStrand), v))
	})
}

// StrandEqualFold applies the EqualFold predicate on the "strand" field.
func StrandEqualFold(v string) predicate.FivePrimeUtr {
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStrand), v))
	})
}

// StrandContainsFold applies the ContainsFold predicate on the "strand" field.
func StrandContainsFold(v string) predicate.FivePrimeUtr {
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStrand), v))
	})
}

// HasTranscript applies the HasEdge predicate on the "transcript" edge.
func HasTranscript() predicate.FivePrimeUtr {
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TranscriptTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TranscriptTable, TranscriptColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTranscriptWith applies the HasEdge predicate on the "transcript" edge with a given conditions (other predicates).
func HasTranscriptWith(preds ...predicate.Transcript) predicate.FivePrimeUtr {
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TranscriptInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TranscriptTable, TranscriptColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.FivePrimeUtr) predicate.FivePrimeUtr {
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.FivePrimeUtr) predicate.FivePrimeUtr {
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
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
func Not(p predicate.FivePrimeUtr) predicate.FivePrimeUtr {
	return predicate.FivePrimeUtr(func(s *sql.Selector) {
		p(s.Not())
	})
}
