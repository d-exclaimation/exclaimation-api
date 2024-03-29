// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/d-exclaimation/exclaimation-api/ent/profile"
)

// Profile is the model entity for the Profile schema.
type Profile struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// AvatarURL holds the value of the "avatar_url" field.
	AvatarURL string `json:"avatar_url,omitempty"`
	// GithubURL holds the value of the "github_url" field.
	GithubURL string `json:"github_url,omitempty"`
	// Location holds the value of the "location" field.
	Location string `json:"location,omitempty"`
	// Bio holds the value of the "bio" field.
	Bio string `json:"bio,omitempty"`
	// TwitterUsername holds the value of the "twitter_username" field.
	TwitterUsername string `json:"twitter_username,omitempty"`
	// PublicRepo holds the value of the "public_repo" field.
	PublicRepo int `json:"public_repo,omitempty"`
	// Followers holds the value of the "followers" field.
	Followers int `json:"followers,omitempty"`
	// Following holds the value of the "following" field.
	Following int `json:"following,omitempty"`
	// LastUpdated holds the value of the "last_updated" field.
	LastUpdated time.Time `json:"last_updated,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Profile) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case profile.FieldID, profile.FieldPublicRepo, profile.FieldFollowers, profile.FieldFollowing:
			values[i] = new(sql.NullInt64)
		case profile.FieldName, profile.FieldAvatarURL, profile.FieldGithubURL, profile.FieldLocation, profile.FieldBio, profile.FieldTwitterUsername:
			values[i] = new(sql.NullString)
		case profile.FieldLastUpdated:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Profile", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Profile fields.
func (pr *Profile) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case profile.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pr.ID = int(value.Int64)
		case profile.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pr.Name = value.String
			}
		case profile.FieldAvatarURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar_url", values[i])
			} else if value.Valid {
				pr.AvatarURL = value.String
			}
		case profile.FieldGithubURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field github_url", values[i])
			} else if value.Valid {
				pr.GithubURL = value.String
			}
		case profile.FieldLocation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field location", values[i])
			} else if value.Valid {
				pr.Location = value.String
			}
		case profile.FieldBio:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field bio", values[i])
			} else if value.Valid {
				pr.Bio = value.String
			}
		case profile.FieldTwitterUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field twitter_username", values[i])
			} else if value.Valid {
				pr.TwitterUsername = value.String
			}
		case profile.FieldPublicRepo:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field public_repo", values[i])
			} else if value.Valid {
				pr.PublicRepo = int(value.Int64)
			}
		case profile.FieldFollowers:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field followers", values[i])
			} else if value.Valid {
				pr.Followers = int(value.Int64)
			}
		case profile.FieldFollowing:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field following", values[i])
			} else if value.Valid {
				pr.Following = int(value.Int64)
			}
		case profile.FieldLastUpdated:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_updated", values[i])
			} else if value.Valid {
				pr.LastUpdated = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Profile.
// Note that you need to call Profile.Unwrap() before calling this method if this Profile
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Profile) Update() *ProfileUpdateOne {
	return (&ProfileClient{config: pr.config}).UpdateOne(pr)
}

// Unwrap unwraps the Profile entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Profile) Unwrap() *Profile {
	tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Profile is not a transactional entity")
	}
	pr.config.driver = tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Profile) String() string {
	var builder strings.Builder
	builder.WriteString("Profile(")
	builder.WriteString(fmt.Sprintf("id=%v", pr.ID))
	builder.WriteString(", name=")
	builder.WriteString(pr.Name)
	builder.WriteString(", avatar_url=")
	builder.WriteString(pr.AvatarURL)
	builder.WriteString(", github_url=")
	builder.WriteString(pr.GithubURL)
	builder.WriteString(", location=")
	builder.WriteString(pr.Location)
	builder.WriteString(", bio=")
	builder.WriteString(pr.Bio)
	builder.WriteString(", twitter_username=")
	builder.WriteString(pr.TwitterUsername)
	builder.WriteString(", public_repo=")
	builder.WriteString(fmt.Sprintf("%v", pr.PublicRepo))
	builder.WriteString(", followers=")
	builder.WriteString(fmt.Sprintf("%v", pr.Followers))
	builder.WriteString(", following=")
	builder.WriteString(fmt.Sprintf("%v", pr.Following))
	builder.WriteString(", last_updated=")
	builder.WriteString(pr.LastUpdated.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Profiles is a parsable slice of Profile.
type Profiles []*Profile

func (pr Profiles) config(cfg config) {
	for _i := range pr {
		pr[_i].config = cfg
	}
}
