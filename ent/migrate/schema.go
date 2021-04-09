// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// PostsColumns holds the columns for the "posts" table.
	PostsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "body", Type: field.TypeString, Size: 2147483647},
		{Name: "crabrave", Type: field.TypeInt, Default: 0},
	}
	// PostsTable holds the schema information for the "posts" table.
	PostsTable = &schema.Table{
		Name:        "posts",
		Columns:     PostsColumns,
		PrimaryKey:  []*schema.Column{PostsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// ProfilesColumns holds the columns for the "profiles" table.
	ProfilesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "avatar_url", Type: field.TypeString},
		{Name: "github_url", Type: field.TypeString},
		{Name: "location", Type: field.TypeString},
		{Name: "bio", Type: field.TypeString, Size: 2147483647},
		{Name: "twitter_username", Type: field.TypeString},
		{Name: "public_repo", Type: field.TypeInt},
		{Name: "followers", Type: field.TypeInt},
		{Name: "following", Type: field.TypeInt},
		{Name: "last_updated", Type: field.TypeTime},
	}
	// ProfilesTable holds the schema information for the "profiles" table.
	ProfilesTable = &schema.Table{
		Name:        "profiles",
		Columns:     ProfilesColumns,
		PrimaryKey:  []*schema.Column{ProfilesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// ReposColumns holds the columns for the "repos" table.
	ReposColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "repo_name", Type: field.TypeString},
		{Name: "url", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Size: 2147483647},
		{Name: "language", Type: field.TypeString},
		{Name: "last_updated", Type: field.TypeTime},
	}
	// ReposTable holds the schema information for the "repos" table.
	ReposTable = &schema.Table{
		Name:        "repos",
		Columns:     ReposColumns,
		PrimaryKey:  []*schema.Column{ReposColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		PostsTable,
		ProfilesTable,
		ReposTable,
	}
)

func init() {
}
