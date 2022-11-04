package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// MessageItem holds the schema definition for the MessageItem entity.
type MessageItem struct {
	ent.Schema
}

// Fields of the MessageItem.
func (MessageItem) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").Values("MESSAGE", "FILE").
			Annotations(
				entproto.Field(2),
				entproto.Enum(map[string]int32{
					"MESSAGE": 1,
					"FILE":    2,
				}),
			),
		field.String("username").
			Annotations(
				entproto.Field(3),
			),
		field.String("message").
			Annotations(
				entproto.Field(4),
			).
			SchemaType(map[string]string{
				dialect.Postgres: "text", // Override Postgres.
			}),
		field.Enum("channel").Values("CHAT", "MAIL", "WHATSAPP", "FACEBOOK", "TWITTER", "VOICE").
			Annotations(
				entproto.Field(5),
				entproto.Enum(map[string]int32{
					"CHAT":     1,
					"MAIL":     2,
					"WHATSAPP": 3,
					"FACEBOOK": 4,
					"TWITTER":  5,
					"VOICE":    6,
				}),
			),
		field.Enum("direction").Values("INBOUND", "OUTBOUND").
			Annotations(
				entproto.Field(6),
				entproto.Enum(map[string]int32{
					"INBOUND":  1,
					"OUTBOUND": 2,
				}),
			),
		field.Time("time").Default(time.Now()).
			Annotations(
				entproto.Field(8),
				&entsql.Annotation{
					Default: "CURRENT_TIMESTAMP",
				},
			),
	}
}

// Edges of the MessageItem.
func (MessageItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("message_feed", MessageFeed.Type).
			Ref("message_item").
			Unique().
			Required().
			Immutable().
			Annotations(entproto.Field(7)),
	}
}
