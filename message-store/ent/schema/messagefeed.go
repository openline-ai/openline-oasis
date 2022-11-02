package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// MessageFeed holds the schema definition for the MessageFeed entity.
type MessageFeed struct {
	ent.Schema
}

// Fields of the MessageFeed.
func (MessageFeed) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").
			Unique().
			Annotations(
				entproto.Field(2),
			),
	}
}

// Edges of the MessageFeed.
func (MessageFeed) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("message_item", MessageItem.Type).
			Annotations(entproto.Field(3)),
	}
}

func (MessageFeed) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("username").Unique(),
	}
}
func (MessageFeed) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}
