// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/d-exclaimation/exclaimation-api/ent/predicate"
	"github.com/d-exclaimation/exclaimation-api/ent/profile"
)

// ProfileUpdate is the builder for updating Profile entities.
type ProfileUpdate struct {
	config
	hooks    []Hook
	mutation *ProfileMutation
}

// Where adds a new predicate for the ProfileUpdate builder.
func (pu *ProfileUpdate) Where(ps ...predicate.Profile) *ProfileUpdate {
	pu.mutation.predicates = append(pu.mutation.predicates, ps...)
	return pu
}

// SetName sets the "name" field.
func (pu *ProfileUpdate) SetName(s string) *ProfileUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetAvatarURL sets the "avatar_url" field.
func (pu *ProfileUpdate) SetAvatarURL(s string) *ProfileUpdate {
	pu.mutation.SetAvatarURL(s)
	return pu
}

// SetGithubURL sets the "github_url" field.
func (pu *ProfileUpdate) SetGithubURL(s string) *ProfileUpdate {
	pu.mutation.SetGithubURL(s)
	return pu
}

// SetLocation sets the "location" field.
func (pu *ProfileUpdate) SetLocation(s string) *ProfileUpdate {
	pu.mutation.SetLocation(s)
	return pu
}

// SetBio sets the "bio" field.
func (pu *ProfileUpdate) SetBio(s string) *ProfileUpdate {
	pu.mutation.SetBio(s)
	return pu
}

// SetTwitterUsername sets the "twitter_username" field.
func (pu *ProfileUpdate) SetTwitterUsername(s string) *ProfileUpdate {
	pu.mutation.SetTwitterUsername(s)
	return pu
}

// SetPublicRepo sets the "public_repo" field.
func (pu *ProfileUpdate) SetPublicRepo(i int) *ProfileUpdate {
	pu.mutation.ResetPublicRepo()
	pu.mutation.SetPublicRepo(i)
	return pu
}

// AddPublicRepo adds i to the "public_repo" field.
func (pu *ProfileUpdate) AddPublicRepo(i int) *ProfileUpdate {
	pu.mutation.AddPublicRepo(i)
	return pu
}

// SetFollowers sets the "followers" field.
func (pu *ProfileUpdate) SetFollowers(i int) *ProfileUpdate {
	pu.mutation.ResetFollowers()
	pu.mutation.SetFollowers(i)
	return pu
}

// AddFollowers adds i to the "followers" field.
func (pu *ProfileUpdate) AddFollowers(i int) *ProfileUpdate {
	pu.mutation.AddFollowers(i)
	return pu
}

// SetFollowing sets the "following" field.
func (pu *ProfileUpdate) SetFollowing(i int) *ProfileUpdate {
	pu.mutation.ResetFollowing()
	pu.mutation.SetFollowing(i)
	return pu
}

// AddFollowing adds i to the "following" field.
func (pu *ProfileUpdate) AddFollowing(i int) *ProfileUpdate {
	pu.mutation.AddFollowing(i)
	return pu
}

// SetLastUpdated sets the "last_updated" field.
func (pu *ProfileUpdate) SetLastUpdated(t time.Time) *ProfileUpdate {
	pu.mutation.SetLastUpdated(t)
	return pu
}

// Mutation returns the ProfileMutation object of the builder.
func (pu *ProfileUpdate) Mutation() *ProfileMutation {
	return pu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *ProfileUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		if err = pu.check(); err != nil {
			return 0, err
		}
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProfileMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pu.check(); err != nil {
				return 0, err
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProfileUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProfileUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProfileUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *ProfileUpdate) check() error {
	if v, ok := pu.mutation.Name(); ok {
		if err := profile.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := pu.mutation.AvatarURL(); ok {
		if err := profile.AvatarURLValidator(v); err != nil {
			return &ValidationError{Name: "avatar_url", err: fmt.Errorf("ent: validator failed for field \"avatar_url\": %w", err)}
		}
	}
	if v, ok := pu.mutation.GithubURL(); ok {
		if err := profile.GithubURLValidator(v); err != nil {
			return &ValidationError{Name: "github_url", err: fmt.Errorf("ent: validator failed for field \"github_url\": %w", err)}
		}
	}
	if v, ok := pu.mutation.Location(); ok {
		if err := profile.LocationValidator(v); err != nil {
			return &ValidationError{Name: "location", err: fmt.Errorf("ent: validator failed for field \"location\": %w", err)}
		}
	}
	if v, ok := pu.mutation.TwitterUsername(); ok {
		if err := profile.TwitterUsernameValidator(v); err != nil {
			return &ValidationError{Name: "twitter_username", err: fmt.Errorf("ent: validator failed for field \"twitter_username\": %w", err)}
		}
	}
	if v, ok := pu.mutation.PublicRepo(); ok {
		if err := profile.PublicRepoValidator(v); err != nil {
			return &ValidationError{Name: "public_repo", err: fmt.Errorf("ent: validator failed for field \"public_repo\": %w", err)}
		}
	}
	if v, ok := pu.mutation.Followers(); ok {
		if err := profile.FollowersValidator(v); err != nil {
			return &ValidationError{Name: "followers", err: fmt.Errorf("ent: validator failed for field \"followers\": %w", err)}
		}
	}
	if v, ok := pu.mutation.Following(); ok {
		if err := profile.FollowingValidator(v); err != nil {
			return &ValidationError{Name: "following", err: fmt.Errorf("ent: validator failed for field \"following\": %w", err)}
		}
	}
	return nil
}

func (pu *ProfileUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   profile.Table,
			Columns: profile.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: profile.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldName,
		})
	}
	if value, ok := pu.mutation.AvatarURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldAvatarURL,
		})
	}
	if value, ok := pu.mutation.GithubURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldGithubURL,
		})
	}
	if value, ok := pu.mutation.Location(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldLocation,
		})
	}
	if value, ok := pu.mutation.Bio(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldBio,
		})
	}
	if value, ok := pu.mutation.TwitterUsername(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldTwitterUsername,
		})
	}
	if value, ok := pu.mutation.PublicRepo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: profile.FieldPublicRepo,
		})
	}
	if value, ok := pu.mutation.AddedPublicRepo(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: profile.FieldPublicRepo,
		})
	}
	if value, ok := pu.mutation.Followers(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: profile.FieldFollowers,
		})
	}
	if value, ok := pu.mutation.AddedFollowers(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: profile.FieldFollowers,
		})
	}
	if value, ok := pu.mutation.Following(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: profile.FieldFollowing,
		})
	}
	if value, ok := pu.mutation.AddedFollowing(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: profile.FieldFollowing,
		})
	}
	if value, ok := pu.mutation.LastUpdated(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: profile.FieldLastUpdated,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{profile.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ProfileUpdateOne is the builder for updating a single Profile entity.
type ProfileUpdateOne struct {
	config
	hooks    []Hook
	mutation *ProfileMutation
}

// SetName sets the "name" field.
func (puo *ProfileUpdateOne) SetName(s string) *ProfileUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetAvatarURL sets the "avatar_url" field.
func (puo *ProfileUpdateOne) SetAvatarURL(s string) *ProfileUpdateOne {
	puo.mutation.SetAvatarURL(s)
	return puo
}

// SetGithubURL sets the "github_url" field.
func (puo *ProfileUpdateOne) SetGithubURL(s string) *ProfileUpdateOne {
	puo.mutation.SetGithubURL(s)
	return puo
}

// SetLocation sets the "location" field.
func (puo *ProfileUpdateOne) SetLocation(s string) *ProfileUpdateOne {
	puo.mutation.SetLocation(s)
	return puo
}

// SetBio sets the "bio" field.
func (puo *ProfileUpdateOne) SetBio(s string) *ProfileUpdateOne {
	puo.mutation.SetBio(s)
	return puo
}

// SetTwitterUsername sets the "twitter_username" field.
func (puo *ProfileUpdateOne) SetTwitterUsername(s string) *ProfileUpdateOne {
	puo.mutation.SetTwitterUsername(s)
	return puo
}

// SetPublicRepo sets the "public_repo" field.
func (puo *ProfileUpdateOne) SetPublicRepo(i int) *ProfileUpdateOne {
	puo.mutation.ResetPublicRepo()
	puo.mutation.SetPublicRepo(i)
	return puo
}

// AddPublicRepo adds i to the "public_repo" field.
func (puo *ProfileUpdateOne) AddPublicRepo(i int) *ProfileUpdateOne {
	puo.mutation.AddPublicRepo(i)
	return puo
}

// SetFollowers sets the "followers" field.
func (puo *ProfileUpdateOne) SetFollowers(i int) *ProfileUpdateOne {
	puo.mutation.ResetFollowers()
	puo.mutation.SetFollowers(i)
	return puo
}

// AddFollowers adds i to the "followers" field.
func (puo *ProfileUpdateOne) AddFollowers(i int) *ProfileUpdateOne {
	puo.mutation.AddFollowers(i)
	return puo
}

// SetFollowing sets the "following" field.
func (puo *ProfileUpdateOne) SetFollowing(i int) *ProfileUpdateOne {
	puo.mutation.ResetFollowing()
	puo.mutation.SetFollowing(i)
	return puo
}

// AddFollowing adds i to the "following" field.
func (puo *ProfileUpdateOne) AddFollowing(i int) *ProfileUpdateOne {
	puo.mutation.AddFollowing(i)
	return puo
}

// SetLastUpdated sets the "last_updated" field.
func (puo *ProfileUpdateOne) SetLastUpdated(t time.Time) *ProfileUpdateOne {
	puo.mutation.SetLastUpdated(t)
	return puo
}

// Mutation returns the ProfileMutation object of the builder.
func (puo *ProfileUpdateOne) Mutation() *ProfileMutation {
	return puo.mutation
}

// Save executes the query and returns the updated Profile entity.
func (puo *ProfileUpdateOne) Save(ctx context.Context) (*Profile, error) {
	var (
		err  error
		node *Profile
	)
	if len(puo.hooks) == 0 {
		if err = puo.check(); err != nil {
			return nil, err
		}
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProfileMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = puo.check(); err != nil {
				return nil, err
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProfileUpdateOne) SaveX(ctx context.Context) *Profile {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *ProfileUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProfileUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *ProfileUpdateOne) check() error {
	if v, ok := puo.mutation.Name(); ok {
		if err := profile.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := puo.mutation.AvatarURL(); ok {
		if err := profile.AvatarURLValidator(v); err != nil {
			return &ValidationError{Name: "avatar_url", err: fmt.Errorf("ent: validator failed for field \"avatar_url\": %w", err)}
		}
	}
	if v, ok := puo.mutation.GithubURL(); ok {
		if err := profile.GithubURLValidator(v); err != nil {
			return &ValidationError{Name: "github_url", err: fmt.Errorf("ent: validator failed for field \"github_url\": %w", err)}
		}
	}
	if v, ok := puo.mutation.Location(); ok {
		if err := profile.LocationValidator(v); err != nil {
			return &ValidationError{Name: "location", err: fmt.Errorf("ent: validator failed for field \"location\": %w", err)}
		}
	}
	if v, ok := puo.mutation.TwitterUsername(); ok {
		if err := profile.TwitterUsernameValidator(v); err != nil {
			return &ValidationError{Name: "twitter_username", err: fmt.Errorf("ent: validator failed for field \"twitter_username\": %w", err)}
		}
	}
	if v, ok := puo.mutation.PublicRepo(); ok {
		if err := profile.PublicRepoValidator(v); err != nil {
			return &ValidationError{Name: "public_repo", err: fmt.Errorf("ent: validator failed for field \"public_repo\": %w", err)}
		}
	}
	if v, ok := puo.mutation.Followers(); ok {
		if err := profile.FollowersValidator(v); err != nil {
			return &ValidationError{Name: "followers", err: fmt.Errorf("ent: validator failed for field \"followers\": %w", err)}
		}
	}
	if v, ok := puo.mutation.Following(); ok {
		if err := profile.FollowingValidator(v); err != nil {
			return &ValidationError{Name: "following", err: fmt.Errorf("ent: validator failed for field \"following\": %w", err)}
		}
	}
	return nil
}

func (puo *ProfileUpdateOne) sqlSave(ctx context.Context) (_node *Profile, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   profile.Table,
			Columns: profile.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: profile.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Profile.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldName,
		})
	}
	if value, ok := puo.mutation.AvatarURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldAvatarURL,
		})
	}
	if value, ok := puo.mutation.GithubURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldGithubURL,
		})
	}
	if value, ok := puo.mutation.Location(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldLocation,
		})
	}
	if value, ok := puo.mutation.Bio(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldBio,
		})
	}
	if value, ok := puo.mutation.TwitterUsername(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldTwitterUsername,
		})
	}
	if value, ok := puo.mutation.PublicRepo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: profile.FieldPublicRepo,
		})
	}
	if value, ok := puo.mutation.AddedPublicRepo(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: profile.FieldPublicRepo,
		})
	}
	if value, ok := puo.mutation.Followers(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: profile.FieldFollowers,
		})
	}
	if value, ok := puo.mutation.AddedFollowers(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: profile.FieldFollowers,
		})
	}
	if value, ok := puo.mutation.Following(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: profile.FieldFollowing,
		})
	}
	if value, ok := puo.mutation.AddedFollowing(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: profile.FieldFollowing,
		})
	}
	if value, ok := puo.mutation.LastUpdated(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: profile.FieldLastUpdated,
		})
	}
	_node = &Profile{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{profile.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
