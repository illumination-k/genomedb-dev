// Code generated by ent, DO NOT EDIT.

package genome

const (
	// Label holds the string label denoting the genome type in the database.
	Label = "genome"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCodonTable holds the string denoting the codon_table field in the database.
	FieldCodonTable = "codon_table"
	// FieldSeq holds the string denoting the seq field in the database.
	FieldSeq = "seq"
	// Table holds the table name of the genome in the database.
	Table = "genomes"
)

// Columns holds all SQL columns for genome fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldCodonTable,
	FieldSeq,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
