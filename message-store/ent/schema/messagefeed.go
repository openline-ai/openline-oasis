package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
)

// MessageFeed holds the schema definition for the MessageFeed entity.
type MessageFeed struct {
	ent.Schema
}

// Fields of the MessageFeed.
func (MessageFeed) Fields() []ent.Field {
	return []ent.Field{}
}

// Edges of the MessageFeed.
func (MessageFeed) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("message_item", MessageItem.Type),
	}
}

func (MessageFeed) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}
