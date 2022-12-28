// Code generated by ent, DO NOT EDIT.

package transcript

const (
	// Label holds the string label denoting the transcript type in the database.
	Label = "transcript"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSeqname holds the string denoting the seqname field in the database.
	FieldSeqname = "seqname"
	// FieldStrand holds the string denoting the strand field in the database.
	FieldStrand = "strand"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldSource holds the string denoting the source field in the database.
	FieldSource = "source"
	// FieldStart holds the string denoting the start field in the database.
	FieldStart = "start"
	// FieldEnd holds the string denoting the end field in the database.
	FieldEnd = "end"
	// FieldExon holds the string denoting the exon field in the database.
	FieldExon = "exon"
	// FieldFivePrimeUtr holds the string denoting the five_prime_utr field in the database.
	FieldFivePrimeUtr = "five_prime_utr"
	// FieldThreePrimeUtr holds the string denoting the three_prime_utr field in the database.
	FieldThreePrimeUtr = "three_prime_utr"
	// FieldCds holds the string denoting the cds field in the database.
	FieldCds = "cds"
	// FieldGenomicSequence holds the string denoting the genomic_sequence field in the database.
	FieldGenomicSequence = "genomic_sequence"
	// FieldExonSequence holds the string denoting the exon_sequence field in the database.
	FieldExonSequence = "exon_sequence"
	// FieldCdsSequence holds the string denoting the cds_sequence field in the database.
	FieldCdsSequence = "cds_sequence"
	// FieldProteinSequence holds the string denoting the protein_sequence field in the database.
	FieldProteinSequence = "protein_sequence"
	// EdgeGene holds the string denoting the gene edge name in mutations.
	EdgeGene = "gene"
	// EdgeGoterms holds the string denoting the goterms edge name in mutations.
	EdgeGoterms = "goterms"
	// EdgeDomains holds the string denoting the domains edge name in mutations.
	EdgeDomains = "domains"
	// EdgeGotermTranscript holds the string denoting the goterm_transcript edge name in mutations.
	EdgeGotermTranscript = "goterm_transcript"
	// EdgeDomainTranscript holds the string denoting the domain_transcript edge name in mutations.
	EdgeDomainTranscript = "domain_transcript"
	// Table holds the table name of the transcript in the database.
	Table = "transcripts"
	// GeneTable is the table that holds the gene relation/edge.
	GeneTable = "transcripts"
	// GeneInverseTable is the table name for the Gene entity.
	// It exists in this package in order to avoid circular dependency with the "gene" package.
	GeneInverseTable = "genes"
	// GeneColumn is the table column denoting the gene relation/edge.
	GeneColumn = "gene_transcripts"
	// GotermsTable is the table that holds the goterms relation/edge. The primary key declared below.
	GotermsTable = "go_term_on_transcripts"
	// GotermsInverseTable is the table name for the GoTerm entity.
	// It exists in this package in order to avoid circular dependency with the "goterm" package.
	GotermsInverseTable = "go_terms"
	// DomainsTable is the table that holds the domains relation/edge. The primary key declared below.
	DomainsTable = "domain_annotation_to_transcripts"
	// DomainsInverseTable is the table name for the DomainAnnotation entity.
	// It exists in this package in order to avoid circular dependency with the "domainannotation" package.
	DomainsInverseTable = "domain_annotations"
	// GotermTranscriptTable is the table that holds the goterm_transcript relation/edge.
	GotermTranscriptTable = "go_term_on_transcripts"
	// GotermTranscriptInverseTable is the table name for the GoTermOnTranscripts entity.
	// It exists in this package in order to avoid circular dependency with the "gotermontranscripts" package.
	GotermTranscriptInverseTable = "go_term_on_transcripts"
	// GotermTranscriptColumn is the table column denoting the goterm_transcript relation/edge.
	GotermTranscriptColumn = "transcript_id"
	// DomainTranscriptTable is the table that holds the domain_transcript relation/edge.
	DomainTranscriptTable = "domain_annotation_to_transcripts"
	// DomainTranscriptInverseTable is the table name for the DomainAnnotationToTranscript entity.
	// It exists in this package in order to avoid circular dependency with the "domainannotationtotranscript" package.
	DomainTranscriptInverseTable = "domain_annotation_to_transcripts"
	// DomainTranscriptColumn is the table column denoting the domain_transcript relation/edge.
	DomainTranscriptColumn = "transcript_id"
)

// Columns holds all SQL columns for transcript fields.
var Columns = []string{
	FieldID,
	FieldSeqname,
	FieldStrand,
	FieldType,
	FieldSource,
	FieldStart,
	FieldEnd,
	FieldExon,
	FieldFivePrimeUtr,
	FieldThreePrimeUtr,
	FieldCds,
	FieldGenomicSequence,
	FieldExonSequence,
	FieldCdsSequence,
	FieldProteinSequence,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "transcripts"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"gene_transcripts",
}

var (
	// GotermsPrimaryKey and GotermsColumn2 are the table columns denoting the
	// primary key for the goterms relation (M2M).
	GotermsPrimaryKey = []string{"go_term_id", "transcript_id"}
	// DomainsPrimaryKey and DomainsColumn2 are the table columns denoting the
	// primary key for the domains relation (M2M).
	DomainsPrimaryKey = []string{"domain_annotation_id", "transcript_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// StartValidator is a validator for the "start" field. It is called by the builders before save.
	StartValidator func(int32) error
	// EndValidator is a validator for the "end" field. It is called by the builders before save.
	EndValidator func(int32) error
)
