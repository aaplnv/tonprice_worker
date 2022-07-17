// Code generated by entc, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTelegramId holds the string denoting the telegramid field in the database.
	FieldTelegramId = "telegram_id"
	// FieldAllStables holds the string denoting the allstables field in the database.
	FieldAllStables = "all_stables"
	// FieldRegTime holds the string denoting the regtime field in the database.
	FieldRegTime = "reg_time"
	// Table holds the table name of the user in the database.
	Table = "users"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldTelegramId,
	FieldAllStables,
	FieldRegTime,
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