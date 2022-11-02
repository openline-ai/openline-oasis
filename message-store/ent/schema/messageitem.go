package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
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
		field.Enum("channel").Values("CHAT", "MAIL", "WHATSAPP", "FACEBOOK", "TWITTER").
			Annotations(
				entproto.Field(5),
				entproto.Enum(map[string]int32{
					"CHAT":     1,
					"MAIL":     2,
					"WHATSAPP": 3,
					"FACEBOOK": 4,
					"TWITTER":  5,
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
	}
}

// Edges of the MessageItem.
func (MessageItem) Edges() []ent.Edge {
	return nil
}

func (MessageItem) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}
