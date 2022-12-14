// Code generated by ent, DO NOT EDIT.

package scaffold

import (
	"genomedb/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Scaffold {
	return predicate.Scaffold(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Scaffold {
	return predicate.Scaffold(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Scaffold {
	return predicate.Scaffold(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Scaffold {
	return predicate.Scaffold(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Scaffold {
	return predicate.Scaffold(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Scaffold {
	return predicate.Scaffold(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Scaffold {
	return predicate.Scaffold(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Scaffold {
	return predicate.Scaffold(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Scaffold {
	return predicate.Scaffold(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Seqname applies equality check predicate on the "seqname" field. It's identical to SeqnameEQ.
func Seqname(v string) predicate.Scaffold {
	return predicate.Scaffold(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSeqname), v))
	})
}

// Seq applies equality check predicate on the "seq" field. It's identical to SeqEQ.
func Seq(v string) predicate.Scaffold {
	return predicate.Scaffold(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSeq), v))
	})
}

// SeqnameEQ applies the EQ predicate on the "seqname" field.
func SeqnameEQ(v string) predicate.Scaffold {
	return predicate.Scaffold(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSeqname), v))
	})
}

// SeqnameNEQ applies the NEQ predicate on the "seqname" field.
func SeqnameNEQ(v string) predicate.Scaffold {
	return predicate.Scaffold(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSeqname), v))
	})
}

// SeqnameIn applies the In predicate on the "seqname" field.
func SeqnameIn(vs ...string) predicate.Scaffold {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Scaffold(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSeqname), v...))
	})
}

// SeqnameNotIn applies the NotIn predicate on the "seqname" field.
func SeqnameNotIn(vs ...string) predicate.Scaffold {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Scaffold(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSeqname), v...))
	})
}

// SeqnameGT applies the GT predicate on the "seqname" field.
func SeqnameGT(v string) predicate.Scaffold {
	return predicate.Scaffold(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSeqname), v))
	})
}

// SeqnameGTE applies the GTE predicate on the "seqname" field.
func SeqnameGTE(v string) predicate.Scaffold {
	return predicate.Scaffold(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSeqname), v))
	})
}

// SeqnameLT applies the LT predicate on the "seqname" field.
func SeqnameLT(v string) predicate.Scaffold {
	return predicate.Scaffold(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSeqname), v))
	})
}

// SeqnameLTE applies the LTE predicate on the "seqname" field.
func SeqnameLTE(v string) predicate.Scaffold {
	return predicate.Scaffold(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSeqname), v))
	})
}

// SeqnameContains applies the Contains predicate on the "seqname" field.
func SeqnameContains(v string) predicate.Scaffold {
	return predicate.Scaffold(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSeqname), v))
	})
}

// SeqnameHasPrefix applies the HasPrefix predicate on the "seqname" field.
func SeqnameHasPrefix(v string) predicate.Scaffold {
	return predicate.Scaffold(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSeqname), v))
	})
}

// SeqnameHasSuffix applies the HasSuffix predicate on the "seqname" field.
func SeqnameHasSuffix(v string) predicate.Scaffold {
	return predicate.Scaffold(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSeqname), v))
	})
}

// SeqnameEqualFold applies the EqualFold predicate on the "seqname" field.
func SeqnameEqualFold(v string) predicate.Scaffold {
	return predicate.Scaffold(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSeqname), v))
	})
}

// SeqnameContainsFold applies the ContainsFold predicate on the "seqname" field.
func SeqnameContainsFold(v string) predicate.Scaffold {
	return predicate.Scaffold(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSeqname), v))
	})
}

// SeqEQ applies the EQ predicate on the "seq" field.
func SeqEQ(v string) predicate.Scaffold {
	return predicate.Scaffold(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSeq), v))
	})
}

// SeqNEQ applies the NEQ predicate on the "seq" field.
func SeqNEQ(v string) predicate.Scaffold {
	return predicate.Scaffold(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSeq), v))
	})
}

// SeqIn applies the In predicate on the "seq" field.
func SeqIn(vs ...string) predicate.Scaffold {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Scaffold(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSeq), v...))
	})
}

// SeqNotIn applies the NotIn predicate on the "seq" field.
func SeqNotIn(vs ...string) predicate.Scaffold {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Scaffold(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSeq), v...))
	})
}

// SeqGT applies the GT predicate on the "seq" field.
func SeqGT(v string) predicate.Scaffold {
	return predicate.Scaffold(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSeq), v))
	})
}

// SeqGTE applies the GTE predicate on the "seq" field.
func SeqGTE(v string) predicate.Scaffold {
	return predicate.Scaffold(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSeq), v))
	})
}

// SeqLT applies the LT predicate on the "seq" field.
func SeqLT(v string) predicate.Scaffold {
	return predicate.Scaffold(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSeq), v))
	})
}

// SeqLTE applies the LTE predicate on the "seq" field.
func SeqLTE(v string) predicate.Scaffold {
	return predicate.Scaffold(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSeq), v))
	})
}

// SeqContains applies the Contains predicate on the "seq" field.
func SeqContains(v string) predicate.Scaffold {
	return predicate.Scaffold(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSeq), v))
	})
}

// SeqHasPrefix applies the HasPrefix predicate on the "seq" field.
func SeqHasPrefix(v string) predicate.Scaffold {
	return predicate.Scaffold(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSeq), v))
	})
}

// SeqHasSuffix applies the HasSuffix predicate on the "seq" field.
func SeqHasSuffix(v string) predicate.Scaffold {
	return predicate.Scaffold(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSeq), v))
	})
}

// SeqEqualFold applies the EqualFold predicate on the "seq" field.
func SeqEqualFold(v string) predicate.Scaffold {
	return predicate.Scaffold(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSeq), v))
	})
}

// SeqContainsFold applies the ContainsFold predicate on the "seq" field.
func SeqContainsFold(v string) predicate.Scaffold {
	return predicate.Scaffold(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSeq), v))
	})
}

// HasGenome applies the HasEdge predicate on the "genome" edge.
func HasGenome() predicate.Scaffold {
	return predicate.Scaffold(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GenomeTable, GenomeFieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, GenomeTable, GenomeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGenomeWith applies the HasEdge predicate on the "genome" edge with a given conditions (other predicates).
func HasGenomeWith(preds ...predicate.Genome) predicate.Scaffold {
	return predicate.Scaffold(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GenomeInverseTable, GenomeFieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, GenomeTable, GenomeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Scaffold) predicate.Scaffold {
	return predicate.Scaffold(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Scaffold) predicate.Scaffold {
	return predicate.Scaffold(func(s *sql.Selector) {
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
func Not(p predicate.Scaffold) predicate.Scaffold {
	return predicate.Scaffold(func(s *sql.Selector) {
		p(s.Not())
	})
}
