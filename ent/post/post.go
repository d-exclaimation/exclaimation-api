// Code generated by entc, DO NOT EDIT.

package post

const (
	// Label holds the string label denoting the post type in the database.
	Label = "post"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldBody holds the string denoting the body field in the database.
	FieldBody = "body"
	// FieldCrabrave holds the string denoting the crabrave field in the database.
	FieldCrabrave = "crabrave"
	// Table holds the table name of the post in the database.
	Table = "posts"
)

// Columns holds all SQL columns for post fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldBody,
	FieldCrabrave,
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

var (
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// BodyValidator is a validator for the "body" field. It is called by the builders before save.
	BodyValidator func(string) error
	// DefaultCrabrave holds the default value on creation for the "crabrave" field.
	DefaultCrabrave int
)
