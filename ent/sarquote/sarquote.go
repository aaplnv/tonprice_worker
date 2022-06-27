// Code generated by entc, DO NOT EDIT.

package sarquote

const (
	// Label holds the string label denoting the sarquote type in the database.
	Label = "sar_quote"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// FieldTimestamp holds the string denoting the timestamp field in the database.
	FieldTimestamp = "timestamp"
	// Table holds the table name of the sarquote in the database.
	Table = "SAR"
)

// Columns holds all SQL columns for sarquote fields.
var Columns = []string{
	FieldID,
	FieldPrice,
	FieldTimestamp,
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