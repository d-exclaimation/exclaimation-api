package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Repo holds the schema definition for the Repo, entity.
type Repo struct {
	ent.Schema
}

// Fields of the Repo.
func (Repo) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),
		field.String("repo_name").
			NotEmpty(),
		field.String("url").
			NotEmpty(),
		field.Text("description"),
		field.String("language"),
		field.Time("last_updated"),
	}
}

// Edges of the Repo,.
func (Repo) Edges() []ent.Edge {
	return nil
}
