package schema

import (
	"time"
	"github.com/google/uuid"

	"entgo.io/ent"
    "entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
        field.UUID("id", uuid.UUID{}).
            Default(uuid.New).
            Immutable(),
        field.String("email"),
		field.Time("created_at").Default(time.Now),
    }
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
