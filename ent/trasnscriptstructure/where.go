// Code generated by ent, DO NOT EDIT.

package trasnscriptstructure

import (
	"genomedb/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// TranscriptId applies equality check predicate on the "transcriptId" field. It's identical to TranscriptIdEQ.
func TranscriptId(v string) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTranscriptId), v))
	})
}

// Feature applies equality check predicate on the "feature" field. It's identical to FeatureEQ.
func Feature(v string) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFeature), v))
	})
}

// Seqname applies equality check predicate on the "seqname" field. It's identical to SeqnameEQ.
func Seqname(v string) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSeqname), v))
	})
}

// Start applies equality check predicate on the "start" field. It's identical to StartEQ.
func Start(v int32) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStart), v))
	})
}

// End applies equality check predicate on the "end" field. It's identical to EndEQ.
func End(v int32) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEnd), v))
	})
}

// Strand applies equality check predicate on the "strand" field. It's identical to StrandEQ.
func Strand(v string) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStrand), v))
	})
}

// TranscriptIdEQ applies the EQ predicate on the "transcriptId" field.
func TranscriptIdEQ(v string) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTranscriptId), v))
	})
}

// TranscriptIdNEQ applies the NEQ predicate on the "transcriptId" field.
func TranscriptIdNEQ(v string) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTranscriptId), v))
	})
}

// TranscriptIdIn applies the In predicate on the "transcriptId" field.
func TranscriptIdIn(vs ...string) predicate.TrasnscriptStructure {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTranscriptId), v...))
	})
}

// TranscriptIdNotIn applies the NotIn predicate on the "transcriptId" field.
func TranscriptIdNotIn(vs ...string) predicate.TrasnscriptStructure {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTranscriptId), v...))
	})
}

// TranscriptIdGT applies the GT predicate on the "transcriptId" field.
func TranscriptIdGT(v string) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTranscriptId), v))
	})
}

// TranscriptIdGTE applies the GTE predicate on the "transcriptId" field.
func TranscriptIdGTE(v string) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTranscriptId), v))
	})
}

// TranscriptIdLT applies the LT predicate on the "transcriptId" field.
func TranscriptIdLT(v string) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTranscriptId), v))
	})
}

// TranscriptIdLTE applies the LTE predicate on the "transcriptId" field.
func TranscriptIdLTE(v string) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTranscriptId), v))
	})
}

// TranscriptIdContains applies the Contains predicate on the "transcriptId" field.
func TranscriptIdContains(v string) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTranscriptId), v))
	})
}

// TranscriptIdHasPrefix applies the HasPrefix predicate on the "transcriptId" field.
func TranscriptIdHasPrefix(v string) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTranscriptId), v))
	})
}

// TranscriptIdHasSuffix applies the HasSuffix predicate on the "transcriptId" field.
func TranscriptIdHasSuffix(v string) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTranscriptId), v))
	})
}

// TranscriptIdEqualFold applies the EqualFold predicate on the "transcriptId" field.
func TranscriptIdEqualFold(v string) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTranscriptId), v))
	})
}

// TranscriptIdContainsFold applies the ContainsFold predicate on the "transcriptId" field.
func TranscriptIdContainsFold(v string) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTranscriptId), v))
	})
}

// FeatureEQ applies the EQ predicate on the "feature" field.
func FeatureEQ(v string) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFeature), v))
	})
}

// FeatureNEQ applies the NEQ predicate on the "feature" field.
func FeatureNEQ(v string) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFeature), v))
	})
}

// FeatureIn applies the In predicate on the "feature" field.
func FeatureIn(vs ...string) predicate.TrasnscriptStructure {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldFeature), v...))
	})
}

// FeatureNotIn applies the NotIn predicate on the "feature" field.
func FeatureNotIn(vs ...string) predicate.TrasnscriptStructure {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldFeature), v...))
	})
}

// FeatureGT applies the GT predicate on the "feature" field.
func FeatureGT(v string) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFeature), v))
	})
}

// FeatureGTE applies the GTE predicate on the "feature" field.
func FeatureGTE(v string) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFeature), v))
	})
}

// FeatureLT applies the LT predicate on the "feature" field.
func FeatureLT(v string) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFeature), v))
	})
}

// FeatureLTE applies the LTE predicate on the "feature" field.
func FeatureLTE(v string) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFeature), v))
	})
}

// FeatureContains applies the Contains predicate on the "feature" field.
func FeatureContains(v string) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFeature), v))
	})
}

// FeatureHasPrefix applies the HasPrefix predicate on the "feature" field.
func FeatureHasPrefix(v string) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFeature), v))
	})
}

// FeatureHasSuffix applies the HasSuffix predicate on the "feature" field.
func FeatureHasSuffix(v string) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFeature), v))
	})
}

// FeatureEqualFold applies the EqualFold predicate on the "feature" field.
func FeatureEqualFold(v string) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFeature), v))
	})
}

// FeatureContainsFold applies the ContainsFold predicate on the "feature" field.
func FeatureContainsFold(v string) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFeature), v))
	})
}

// SeqnameEQ applies the EQ predicate on the "seqname" field.
func SeqnameEQ(v string) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSeqname), v))
	})
}

// SeqnameNEQ applies the NEQ predicate on the "seqname" field.
func SeqnameNEQ(v string) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSeqname), v))
	})
}

// SeqnameIn applies the In predicate on the "seqname" field.
func SeqnameIn(vs ...string) predicate.TrasnscriptStructure {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSeqname), v...))
	})
}

// SeqnameNotIn applies the NotIn predicate on the "seqname" field.
func SeqnameNotIn(vs ...string) predicate.TrasnscriptStructure {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSeqname), v...))
	})
}

// SeqnameGT applies the GT predicate on the "seqname" field.
func SeqnameGT(v string) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSeqname), v))
	})
}

// SeqnameGTE applies the GTE predicate on the "seqname" field.
func SeqnameGTE(v string) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSeqname), v))
	})
}

// SeqnameLT applies the LT predicate on the "seqname" field.
func SeqnameLT(v string) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSeqname), v))
	})
}

// SeqnameLTE applies the LTE predicate on the "seqname" field.
func SeqnameLTE(v string) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSeqname), v))
	})
}

// SeqnameContains applies the Contains predicate on the "seqname" field.
func SeqnameContains(v string) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSeqname), v))
	})
}

// SeqnameHasPrefix applies the HasPrefix predicate on the "seqname" field.
func SeqnameHasPrefix(v string) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSeqname), v))
	})
}

// SeqnameHasSuffix applies the HasSuffix predicate on the "seqname" field.
func SeqnameHasSuffix(v string) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSeqname), v))
	})
}

// SeqnameEqualFold applies the EqualFold predicate on the "seqname" field.
func SeqnameEqualFold(v string) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSeqname), v))
	})
}

// SeqnameContainsFold applies the ContainsFold predicate on the "seqname" field.
func SeqnameContainsFold(v string) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSeqname), v))
	})
}

// StartEQ applies the EQ predicate on the "start" field.
func StartEQ(v int32) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStart), v))
	})
}

// StartNEQ applies the NEQ predicate on the "start" field.
func StartNEQ(v int32) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStart), v))
	})
}

// StartIn applies the In predicate on the "start" field.
func StartIn(vs ...int32) predicate.TrasnscriptStructure {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStart), v...))
	})
}

// StartNotIn applies the NotIn predicate on the "start" field.
func StartNotIn(vs ...int32) predicate.TrasnscriptStructure {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStart), v...))
	})
}

// StartGT applies the GT predicate on the "start" field.
func StartGT(v int32) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStart), v))
	})
}

// StartGTE applies the GTE predicate on the "start" field.
func StartGTE(v int32) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStart), v))
	})
}

// StartLT applies the LT predicate on the "start" field.
func StartLT(v int32) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStart), v))
	})
}

// StartLTE applies the LTE predicate on the "start" field.
func StartLTE(v int32) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStart), v))
	})
}

// EndEQ applies the EQ predicate on the "end" field.
func EndEQ(v int32) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEnd), v))
	})
}

// EndNEQ applies the NEQ predicate on the "end" field.
func EndNEQ(v int32) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEnd), v))
	})
}

// EndIn applies the In predicate on the "end" field.
func EndIn(vs ...int32) predicate.TrasnscriptStructure {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEnd), v...))
	})
}

// EndNotIn applies the NotIn predicate on the "end" field.
func EndNotIn(vs ...int32) predicate.TrasnscriptStructure {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEnd), v...))
	})
}

// EndGT applies the GT predicate on the "end" field.
func EndGT(v int32) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEnd), v))
	})
}

// EndGTE applies the GTE predicate on the "end" field.
func EndGTE(v int32) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEnd), v))
	})
}

// EndLT applies the LT predicate on the "end" field.
func EndLT(v int32) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEnd), v))
	})
}

// EndLTE applies the LTE predicate on the "end" field.
func EndLTE(v int32) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEnd), v))
	})
}

// StrandEQ applies the EQ predicate on the "strand" field.
func StrandEQ(v string) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStrand), v))
	})
}

// StrandNEQ applies the NEQ predicate on the "strand" field.
func StrandNEQ(v string) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStrand), v))
	})
}

// StrandIn applies the In predicate on the "strand" field.
func StrandIn(vs ...string) predicate.TrasnscriptStructure {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStrand), v...))
	})
}

// StrandNotIn applies the NotIn predicate on the "strand" field.
func StrandNotIn(vs ...string) predicate.TrasnscriptStructure {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStrand), v...))
	})
}

// StrandGT applies the GT predicate on the "strand" field.
func StrandGT(v string) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStrand), v))
	})
}

// StrandGTE applies the GTE predicate on the "strand" field.
func StrandGTE(v string) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStrand), v))
	})
}

// StrandLT applies the LT predicate on the "strand" field.
func StrandLT(v string) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStrand), v))
	})
}

// StrandLTE applies the LTE predicate on the "strand" field.
func StrandLTE(v string) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStrand), v))
	})
}

// StrandContains applies the Contains predicate on the "strand" field.
func StrandContains(v string) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStrand), v))
	})
}

// StrandHasPrefix applies the HasPrefix predicate on the "strand" field.
func StrandHasPrefix(v string) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStrand), v))
	})
}

// StrandHasSuffix applies the HasSuffix predicate on the "strand" field.
func StrandHasSuffix(v string) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStrand), v))
	})
}

// StrandEqualFold applies the EqualFold predicate on the "strand" field.
func StrandEqualFold(v string) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStrand), v))
	})
}

// StrandContainsFold applies the ContainsFold predicate on the "strand" field.
func StrandContainsFold(v string) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStrand), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TrasnscriptStructure) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TrasnscriptStructure) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
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
func Not(p predicate.TrasnscriptStructure) predicate.TrasnscriptStructure {
	return predicate.TrasnscriptStructure(func(s *sql.Selector) {
		p(s.Not())
	})
}
