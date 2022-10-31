package schema

import "entgo.io/ent"

// MessageFeed holds the schema definition for the MessageFeed entity.
type MessageFeed struct {
	ent.Schema
}

// Fields of the MessageFeed.
func (MessageFeed) Fields() []ent.Field {
	return nil
}

// Edges of the MessageFeed.
func (MessageFeed) Edges() []ent.Edge {
	return nil
}
