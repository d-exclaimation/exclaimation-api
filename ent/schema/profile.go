package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Profile holds the schema definition for the Profile entity.
type Profile struct {
	ent.Schema
}

// Fields of the Profile.
func (Profile) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),
		field.String("avatar_url"),
		field.String("github_url").
			NotEmpty(),
		field.String("location").
			NotEmpty(),
		field.Text("bio"),
		field.String("twitter_username"),
		field.Int("public_repo"),
		field.Int("followers").
			NonNegative(),
		field.Int("following").
			NonNegative(),
		field.Time("last_updated"),
	}
}

// Edges of the Profile.
func (Profile) Edges() []ent.Edge {
	return nil
}
