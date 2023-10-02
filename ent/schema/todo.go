package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// ToDo holds the schema definition for the ToDo entity.
type ToDo struct {
	ent.Schema
}

// Fields of the ToDo.
func (ToDo) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.Enum("status").
			Values("todo", "doing", "done").
			Default("todo"),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}
