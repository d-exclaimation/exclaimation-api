// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/d-exclaimation/exclaimation-api/ent/repo"
)

// Repo is the model entity for the Repo schema.
type Repo struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// RepoName holds the value of the "repo_name" field.
	RepoName string `json:"repo_name,omitempty"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Language holds the value of the "language" field.
	Language string `json:"language,omitempty"`
	// LastUpdated holds the value of the "last_updated" field.
	LastUpdated time.Time `json:"last_updated,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Repo) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case repo.FieldID:
			values[i] = new(sql.NullInt64)
		case repo.FieldName, repo.FieldRepoName, repo.FieldURL, repo.FieldDescription, repo.FieldLanguage:
			values[i] = new(sql.NullString)
		case repo.FieldLastUpdated:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Repo", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Repo fields.
func (r *Repo) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case repo.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int(value.Int64)
		case repo.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				r.Name = value.String
			}
		case repo.FieldRepoName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field repo_name", values[i])
			} else if value.Valid {
				r.RepoName = value.String
			}
		case repo.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				r.URL = value.String
			}
		case repo.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				r.Description = value.String
			}
		case repo.FieldLanguage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field language", values[i])
			} else if value.Valid {
				r.Language = value.String
			}
		case repo.FieldLastUpdated:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_updated", values[i])
			} else if value.Valid {
				r.LastUpdated = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Repo.
// Note that you need to call Repo.Unwrap() before calling this method if this Repo
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Repo) Update() *RepoUpdateOne {
	return (&RepoClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the Repo entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Repo) Unwrap() *Repo {
	tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Repo is not a transactional entity")
	}
	r.config.driver = tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Repo) String() string {
	var builder strings.Builder
	builder.WriteString("Repo(")
	builder.WriteString(fmt.Sprintf("id=%v", r.ID))
	builder.WriteString(", name=")
	builder.WriteString(r.Name)
	builder.WriteString(", repo_name=")
	builder.WriteString(r.RepoName)
	builder.WriteString(", url=")
	builder.WriteString(r.URL)
	builder.WriteString(", description=")
	builder.WriteString(r.Description)
	builder.WriteString(", language=")
	builder.WriteString(r.Language)
	builder.WriteString(", last_updated=")
	builder.WriteString(r.LastUpdated.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Repos is a parsable slice of Repo.
type Repos []*Repo

func (r Repos) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
