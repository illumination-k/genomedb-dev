// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// DomainAnnotationsColumns holds the columns for the "domain_annotations" table.
	DomainAnnotationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "analysis", Type: field.TypeString},
	}
	// DomainAnnotationsTable holds the schema information for the "domain_annotations" table.
	DomainAnnotationsTable = &schema.Table{
		Name:       "domain_annotations",
		Columns:    DomainAnnotationsColumns,
		PrimaryKey: []*schema.Column{DomainAnnotationsColumns[0]},
	}
	// DomainAnnotationToTranscriptsColumns holds the columns for the "domain_annotation_to_transcripts" table.
	DomainAnnotationToTranscriptsColumns = []*schema.Column{
		{Name: "start", Type: field.TypeInt32},
		{Name: "stop", Type: field.TypeInt32},
		{Name: "score", Type: field.TypeFloat64},
		{Name: "domain_annotation_id", Type: field.TypeString},
		{Name: "transcript_id", Type: field.TypeString},
	}
	// DomainAnnotationToTranscriptsTable holds the schema information for the "domain_annotation_to_transcripts" table.
	DomainAnnotationToTranscriptsTable = &schema.Table{
		Name:       "domain_annotation_to_transcripts",
		Columns:    DomainAnnotationToTranscriptsColumns,
		PrimaryKey: []*schema.Column{DomainAnnotationToTranscriptsColumns[3], DomainAnnotationToTranscriptsColumns[4]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "domain_annotation_to_transcripts_domain_annotations_domain",
				Columns:    []*schema.Column{DomainAnnotationToTranscriptsColumns[3]},
				RefColumns: []*schema.Column{DomainAnnotationsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "domain_annotation_to_transcripts_transcripts_transcript",
				Columns:    []*schema.Column{DomainAnnotationToTranscriptsColumns[4]},
				RefColumns: []*schema.Column{TranscriptsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// GenesColumns holds the columns for the "genes" table.
	GenesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "genome_genes", Type: field.TypeString, Nullable: true},
	}
	// GenesTable holds the schema information for the "genes" table.
	GenesTable = &schema.Table{
		Name:       "genes",
		Columns:    GenesColumns,
		PrimaryKey: []*schema.Column{GenesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "genes_genomes_genes",
				Columns:    []*schema.Column{GenesColumns[1]},
				RefColumns: []*schema.Column{GenomesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// GenomesColumns holds the columns for the "genomes" table.
	GenomesColumns = []*schema.Column{
		{Name: "name", Type: field.TypeString},
		{Name: "codon_table", Type: field.TypeInt32},
	}
	// GenomesTable holds the schema information for the "genomes" table.
	GenomesTable = &schema.Table{
		Name:       "genomes",
		Columns:    GenomesColumns,
		PrimaryKey: []*schema.Column{GenomesColumns[0]},
	}
	// GoTermsColumns holds the columns for the "go_terms" table.
	GoTermsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "namespace", Type: field.TypeEnum, Enums: []string{"BP", "MF", "CC"}},
		{Name: "name", Type: field.TypeString},
		{Name: "def", Type: field.TypeString},
		{Name: "level", Type: field.TypeInt32},
		{Name: "depth", Type: field.TypeInt32},
		{Name: "go_term_children", Type: field.TypeString, Nullable: true},
	}
	// GoTermsTable holds the schema information for the "go_terms" table.
	GoTermsTable = &schema.Table{
		Name:       "go_terms",
		Columns:    GoTermsColumns,
		PrimaryKey: []*schema.Column{GoTermsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "go_terms_go_terms_children",
				Columns:    []*schema.Column{GoTermsColumns[6]},
				RefColumns: []*schema.Column{GoTermsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// GoTermOnTranscriptsColumns holds the columns for the "go_term_on_transcripts" table.
	GoTermOnTranscriptsColumns = []*schema.Column{
		{Name: "evidence_code", Type: field.TypeString},
		{Name: "go_term_id", Type: field.TypeString},
		{Name: "transcript_id", Type: field.TypeString},
	}
	// GoTermOnTranscriptsTable holds the schema information for the "go_term_on_transcripts" table.
	GoTermOnTranscriptsTable = &schema.Table{
		Name:       "go_term_on_transcripts",
		Columns:    GoTermOnTranscriptsColumns,
		PrimaryKey: []*schema.Column{GoTermOnTranscriptsColumns[1], GoTermOnTranscriptsColumns[2]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "go_term_on_transcripts_go_terms_go_term",
				Columns:    []*schema.Column{GoTermOnTranscriptsColumns[1]},
				RefColumns: []*schema.Column{GoTermsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "go_term_on_transcripts_transcripts_transcript",
				Columns:    []*schema.Column{GoTermOnTranscriptsColumns[2]},
				RefColumns: []*schema.Column{TranscriptsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// KeggCompoundsColumns holds the columns for the "kegg_compounds" table.
	KeggCompoundsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
	}
	// KeggCompoundsTable holds the schema information for the "kegg_compounds" table.
	KeggCompoundsTable = &schema.Table{
		Name:       "kegg_compounds",
		Columns:    KeggCompoundsColumns,
		PrimaryKey: []*schema.Column{KeggCompoundsColumns[0]},
	}
	// KeggModulesColumns holds the columns for the "kegg_modules" table.
	KeggModulesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
	}
	// KeggModulesTable holds the schema information for the "kegg_modules" table.
	KeggModulesTable = &schema.Table{
		Name:       "kegg_modules",
		Columns:    KeggModulesColumns,
		PrimaryKey: []*schema.Column{KeggModulesColumns[0]},
	}
	// KeggOntologiesColumns holds the columns for the "kegg_ontologies" table.
	KeggOntologiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "symbol", Type: field.TypeString},
	}
	// KeggOntologiesTable holds the schema information for the "kegg_ontologies" table.
	KeggOntologiesTable = &schema.Table{
		Name:       "kegg_ontologies",
		Columns:    KeggOntologiesColumns,
		PrimaryKey: []*schema.Column{KeggOntologiesColumns[0]},
	}
	// KeggPathwaysColumns holds the columns for the "kegg_pathways" table.
	KeggPathwaysColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
	}
	// KeggPathwaysTable holds the schema information for the "kegg_pathways" table.
	KeggPathwaysTable = &schema.Table{
		Name:       "kegg_pathways",
		Columns:    KeggPathwaysColumns,
		PrimaryKey: []*schema.Column{KeggPathwaysColumns[0]},
	}
	// KeggReactionsColumns holds the columns for the "kegg_reactions" table.
	KeggReactionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// KeggReactionsTable holds the schema information for the "kegg_reactions" table.
	KeggReactionsTable = &schema.Table{
		Name:       "kegg_reactions",
		Columns:    KeggReactionsColumns,
		PrimaryKey: []*schema.Column{KeggReactionsColumns[0]},
	}
	// ScaffoldsColumns holds the columns for the "scaffolds" table.
	ScaffoldsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "seqname", Type: field.TypeString},
		{Name: "seq", Type: field.TypeString, Size: 2147483647},
		{Name: "genome_scaffolds", Type: field.TypeString, Nullable: true},
	}
	// ScaffoldsTable holds the schema information for the "scaffolds" table.
	ScaffoldsTable = &schema.Table{
		Name:       "scaffolds",
		Columns:    ScaffoldsColumns,
		PrimaryKey: []*schema.Column{ScaffoldsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "scaffolds_genomes_scaffolds",
				Columns:    []*schema.Column{ScaffoldsColumns[3]},
				RefColumns: []*schema.Column{GenomesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TranscriptsColumns holds the columns for the "transcripts" table.
	TranscriptsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "seqname", Type: field.TypeString},
		{Name: "strand", Type: field.TypeString},
		{Name: "type", Type: field.TypeString},
		{Name: "source", Type: field.TypeString},
		{Name: "start", Type: field.TypeInt32},
		{Name: "end", Type: field.TypeInt32},
		{Name: "exon", Type: field.TypeJSON},
		{Name: "five_prime_utr", Type: field.TypeJSON},
		{Name: "three_prime_utr", Type: field.TypeJSON},
		{Name: "cds", Type: field.TypeJSON},
		{Name: "genomic_sequence", Type: field.TypeString, Size: 2147483647},
		{Name: "exon_sequence", Type: field.TypeString, Size: 2147483647},
		{Name: "cds_sequence", Type: field.TypeString, Size: 2147483647},
		{Name: "protein_sequence", Type: field.TypeString, Size: 2147483647},
		{Name: "gene_transcripts", Type: field.TypeString, Nullable: true},
	}
	// TranscriptsTable holds the schema information for the "transcripts" table.
	TranscriptsTable = &schema.Table{
		Name:       "transcripts",
		Columns:    TranscriptsColumns,
		PrimaryKey: []*schema.Column{TranscriptsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "transcripts_genes_transcripts",
				Columns:    []*schema.Column{TranscriptsColumns[15]},
				RefColumns: []*schema.Column{GenesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DomainAnnotationsTable,
		DomainAnnotationToTranscriptsTable,
		GenesTable,
		GenomesTable,
		GoTermsTable,
		GoTermOnTranscriptsTable,
		KeggCompoundsTable,
		KeggModulesTable,
		KeggOntologiesTable,
		KeggPathwaysTable,
		KeggReactionsTable,
		ScaffoldsTable,
		TranscriptsTable,
	}
)

func init() {
	DomainAnnotationToTranscriptsTable.ForeignKeys[0].RefTable = DomainAnnotationsTable
	DomainAnnotationToTranscriptsTable.ForeignKeys[1].RefTable = TranscriptsTable
	GenesTable.ForeignKeys[0].RefTable = GenomesTable
	GoTermsTable.ForeignKeys[0].RefTable = GoTermsTable
	GoTermOnTranscriptsTable.ForeignKeys[0].RefTable = GoTermsTable
	GoTermOnTranscriptsTable.ForeignKeys[1].RefTable = TranscriptsTable
	ScaffoldsTable.ForeignKeys[0].RefTable = GenomesTable
	TranscriptsTable.ForeignKeys[0].RefTable = GenesTable
}
