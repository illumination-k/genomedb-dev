// Code generated by ent, DO NOT EDIT.

package domainannotation

import (
	"fmt"
)

const (
	// Label holds the string label denoting the domainannotation type in the database.
	Label = "domain_annotation"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldAnalysis holds the string denoting the analysis field in the database.
	FieldAnalysis = "analysis"
	// EdgeTranscripts holds the string denoting the transcripts edge name in mutations.
	EdgeTranscripts = "transcripts"
	// EdgeDomainTranscript holds the string denoting the domain_transcript edge name in mutations.
	EdgeDomainTranscript = "domain_transcript"
	// Table holds the table name of the domainannotation in the database.
	Table = "domain_annotations"
	// TranscriptsTable is the table that holds the transcripts relation/edge. The primary key declared below.
	TranscriptsTable = "domain_annotation_to_transcripts"
	// TranscriptsInverseTable is the table name for the Transcript entity.
	// It exists in this package in order to avoid circular dependency with the "transcript" package.
	TranscriptsInverseTable = "transcripts"
	// DomainTranscriptTable is the table that holds the domain_transcript relation/edge.
	DomainTranscriptTable = "domain_annotation_to_transcripts"
	// DomainTranscriptInverseTable is the table name for the DomainAnnotationToTranscript entity.
	// It exists in this package in order to avoid circular dependency with the "domainannotationtotranscript" package.
	DomainTranscriptInverseTable = "domain_annotation_to_transcripts"
	// DomainTranscriptColumn is the table column denoting the domain_transcript relation/edge.
	DomainTranscriptColumn = "domain_annotation_id"
)

// Columns holds all SQL columns for domainannotation fields.
var Columns = []string{
	FieldID,
	FieldDescription,
	FieldAnalysis,
}

var (
	// TranscriptsPrimaryKey and TranscriptsColumn2 are the table columns denoting the
	// primary key for the transcripts relation (M2M).
	TranscriptsPrimaryKey = []string{"domain_annotation_id", "transcript_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Analysis defines the type for the "Analysis" enum field.
type Analysis string

// Analysis values.
const (
	AnalysisCDD         Analysis = "CDD"
	AnalysisCOILS       Analysis = "COILS"
	AnalysisGene3D      Analysis = "Gene3D"
	AnalysisHAMAP       Analysis = "HAMAP"
	AnalysisMOBIDB      Analysis = "MOBIDB"
	AnalysisPANTHER     Analysis = "PANTHER"
	AnalysisPfam        Analysis = "Pfam"
	AnalysisPIRSF       Analysis = "PIRSF"
	AnalysisPRINTS      Analysis = "PRINTS"
	AnalysisPROSITE     Analysis = "PROSITE"
	AnalysisSFLD        Analysis = "SFLD"
	AnalysisSMART       Analysis = "SMART"
	AnalysisSUPERFAMILY Analysis = "SUPERFAMILY"
	AnalysisTIGRFAMs    Analysis = "TIGRFAMs"
)

func (_analysis Analysis) String() string {
	return string(_analysis)
}

// AnalysisValidator is a validator for the "Analysis" field enum values. It is called by the builders before save.
func AnalysisValidator(_analysis Analysis) error {
	switch _analysis {
	case AnalysisCDD, AnalysisCOILS, AnalysisGene3D, AnalysisHAMAP, AnalysisMOBIDB, AnalysisPANTHER, AnalysisPfam, AnalysisPIRSF, AnalysisPRINTS, AnalysisPROSITE, AnalysisSFLD, AnalysisSMART, AnalysisSUPERFAMILY, AnalysisTIGRFAMs:
		return nil
	default:
		return fmt.Errorf("domainannotation: invalid enum value for Analysis field: %q", _analysis)
	}
}