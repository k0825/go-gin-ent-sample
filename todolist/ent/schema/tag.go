package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Tag struct {
	ent.Schema
}

func (Tag) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("todo_id", uuid.UUID{}),
		field.Text("keyword").NotEmpty(),
	}
}

func (Tag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("todo", Todo.Type).
			Ref("tags").
			Unique().
			Field("todo_id").
			Required(),
	}
}
