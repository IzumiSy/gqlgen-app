// Code generated by ent, DO NOT EDIT.

package todo

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the todo type in the database.
	Label = "todo"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldText holds the string denoting the text field in the database.
	FieldText = "text"
	// FieldDone holds the string denoting the done field in the database.
	FieldDone = "done"
	// EdgeAssignee holds the string denoting the assignee edge name in mutations.
	EdgeAssignee = "assignee"
	// EdgeCategories holds the string denoting the categories edge name in mutations.
	EdgeCategories = "categories"
	// Table holds the table name of the todo in the database.
	Table = "todos"
	// AssigneeTable is the table that holds the assignee relation/edge.
	AssigneeTable = "todos"
	// AssigneeInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	AssigneeInverseTable = "users"
	// AssigneeColumn is the table column denoting the assignee relation/edge.
	AssigneeColumn = "user_todos"
	// CategoriesTable is the table that holds the categories relation/edge. The primary key declared below.
	CategoriesTable = "category_todos"
	// CategoriesInverseTable is the table name for the Category entity.
	// It exists in this package in order to avoid circular dependency with the "category" package.
	CategoriesInverseTable = "categories"
)

// Columns holds all SQL columns for todo fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldText,
	FieldDone,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "todos"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_todos",
}

var (
	// CategoriesPrimaryKey and CategoriesColumn2 are the table columns denoting the
	// primary key for the categories relation (M2M).
	CategoriesPrimaryKey = []string{"category_id", "todo_id"}
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
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// TextValidator is a validator for the "text" field. It is called by the builders before save.
	TextValidator func(string) error
	// DefaultDone holds the default value on creation for the "done" field.
	DefaultDone bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
