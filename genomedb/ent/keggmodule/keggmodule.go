// Code generated by ent, DO NOT EDIT.

package keggmodule

const (
	// Label holds the string label denoting the keggmodule type in the database.
	Label = "kegg_module"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// Table holds the table name of the keggmodule in the database.
	Table = "kegg_modules"
)

// Columns holds all SQL columns for keggmodule fields.
var Columns = []string{
	FieldID,
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