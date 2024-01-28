package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Unique().
			Annotations(
				entproto.Field(2),
			),
		field.String("email_address").
			Unique().
			Annotations(
				entproto.Field(3),
			),
		field.String("alias").
			Optional().
			Annotations(entproto.Field(4)),
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(
		//entproto.Methods(entproto.MethodCreate | entproto.MethodGet | entproto.MethodList | entproto.MethodBatchCreate),
		// age mikhay baraye mesal delete va update ro as auto ganaerate bardari, code bala ro bezan
		), // <-- add this
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("administered", Category.Type).
			Ref("admin").
			Annotations(entproto.Field(5)),
	}
}
