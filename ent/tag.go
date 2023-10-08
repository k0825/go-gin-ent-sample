// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/k0825/go-gin-ent-sample/ent/tag"
	"github.com/k0825/go-gin-ent-sample/ent/todo"
)

// Tag is the model entity for the Tag schema.
type Tag struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// TodoID holds the value of the "todo_id" field.
	TodoID uuid.UUID `json:"todo_id,omitempty"`
	// Keyword holds the value of the "keyword" field.
	Keyword string `json:"keyword,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TagQuery when eager-loading is set.
	Edges        TagEdges `json:"edges"`
	selectValues sql.SelectValues
}

// TagEdges holds the relations/edges for other nodes in the graph.
type TagEdges struct {
	// Todo holds the value of the todo edge.
	Todo *Todo `json:"todo,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// TodoOrErr returns the Todo value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TagEdges) TodoOrErr() (*Todo, error) {
	if e.loadedTypes[0] {
		if e.Todo == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: todo.Label}
		}
		return e.Todo, nil
	}
	return nil, &NotLoadedError{edge: "todo"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Tag) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case tag.FieldKeyword:
			values[i] = new(sql.NullString)
		case tag.FieldID, tag.FieldTodoID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Tag fields.
func (t *Tag) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tag.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				t.ID = *value
			}
		case tag.FieldTodoID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field todo_id", values[i])
			} else if value != nil {
				t.TodoID = *value
			}
		case tag.FieldKeyword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field keyword", values[i])
			} else if value.Valid {
				t.Keyword = value.String
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Tag.
// This includes values selected through modifiers, order, etc.
func (t *Tag) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// QueryTodo queries the "todo" edge of the Tag entity.
func (t *Tag) QueryTodo() *TodoQuery {
	return NewTagClient(t.config).QueryTodo(t)
}

// Update returns a builder for updating this Tag.
// Note that you need to call Tag.Unwrap() before calling this method if this Tag
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Tag) Update() *TagUpdateOne {
	return NewTagClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Tag entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Tag) Unwrap() *Tag {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Tag is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Tag) String() string {
	var builder strings.Builder
	builder.WriteString("Tag(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("todo_id=")
	builder.WriteString(fmt.Sprintf("%v", t.TodoID))
	builder.WriteString(", ")
	builder.WriteString("keyword=")
	builder.WriteString(t.Keyword)
	builder.WriteByte(')')
	return builder.String()
}

// Tags is a parsable slice of Tag.
type Tags []*Tag
