// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"gqlgen-app/ent/todo"
	"gqlgen-app/ent/user"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Todo is the model entity for the Todo schema.
type Todo struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Text holds the value of the "text" field.
	Text string `json:"text,omitempty"`
	// Done holds the value of the "done" field.
	Done bool `json:"done,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TodoQuery when eager-loading is set.
	Edges      TodoEdges `json:"edges"`
	user_todos *uuid.UUID
}

// TodoEdges holds the relations/edges for other nodes in the graph.
type TodoEdges struct {
	// Assignee holds the value of the assignee edge.
	Assignee *User `json:"assignee,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// AssigneeOrErr returns the Assignee value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TodoEdges) AssigneeOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Assignee == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Assignee, nil
	}
	return nil, &NotLoadedError{edge: "assignee"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Todo) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case todo.FieldDone:
			values[i] = new(sql.NullBool)
		case todo.FieldText:
			values[i] = new(sql.NullString)
		case todo.FieldID:
			values[i] = new(uuid.UUID)
		case todo.ForeignKeys[0]: // user_todos
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Todo", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Todo fields.
func (t *Todo) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case todo.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				t.ID = *value
			}
		case todo.FieldText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field text", values[i])
			} else if value.Valid {
				t.Text = value.String
			}
		case todo.FieldDone:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field done", values[i])
			} else if value.Valid {
				t.Done = value.Bool
			}
		case todo.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_todos", values[i])
			} else if value.Valid {
				t.user_todos = new(uuid.UUID)
				*t.user_todos = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryAssignee queries the "assignee" edge of the Todo entity.
func (t *Todo) QueryAssignee() *UserQuery {
	return (&TodoClient{config: t.config}).QueryAssignee(t)
}

// Update returns a builder for updating this Todo.
// Note that you need to call Todo.Unwrap() before calling this method if this Todo
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Todo) Update() *TodoUpdateOne {
	return (&TodoClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Todo entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Todo) Unwrap() *Todo {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Todo is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Todo) String() string {
	var builder strings.Builder
	builder.WriteString("Todo(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("text=")
	builder.WriteString(t.Text)
	builder.WriteString(", ")
	builder.WriteString("done=")
	builder.WriteString(fmt.Sprintf("%v", t.Done))
	builder.WriteByte(')')
	return builder.String()
}

// Todos is a parsable slice of Todo.
type Todos []*Todo

func (t Todos) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
