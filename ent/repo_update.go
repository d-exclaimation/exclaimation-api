// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/d-exclaimation/exclaimation-api/ent/predicate"
	"github.com/d-exclaimation/exclaimation-api/ent/repo"
)

// RepoUpdate is the builder for updating Repo entities.
type RepoUpdate struct {
	config
	hooks    []Hook
	mutation *RepoMutation
}

// Where appends a list predicates to the RepoUpdate builder.
func (ru *RepoUpdate) Where(ps ...predicate.Repo) *RepoUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetName sets the "name" field.
func (ru *RepoUpdate) SetName(s string) *RepoUpdate {
	ru.mutation.SetName(s)
	return ru
}

// SetRepoName sets the "repo_name" field.
func (ru *RepoUpdate) SetRepoName(s string) *RepoUpdate {
	ru.mutation.SetRepoName(s)
	return ru
}

// SetURL sets the "url" field.
func (ru *RepoUpdate) SetURL(s string) *RepoUpdate {
	ru.mutation.SetURL(s)
	return ru
}

// SetDescription sets the "description" field.
func (ru *RepoUpdate) SetDescription(s string) *RepoUpdate {
	ru.mutation.SetDescription(s)
	return ru
}

// SetLanguage sets the "language" field.
func (ru *RepoUpdate) SetLanguage(s string) *RepoUpdate {
	ru.mutation.SetLanguage(s)
	return ru
}

// SetLastUpdated sets the "last_updated" field.
func (ru *RepoUpdate) SetLastUpdated(t time.Time) *RepoUpdate {
	ru.mutation.SetLastUpdated(t)
	return ru
}

// Mutation returns the RepoMutation object of the builder.
func (ru *RepoUpdate) Mutation() *RepoMutation {
	return ru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RepoUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ru.hooks) == 0 {
		if err = ru.check(); err != nil {
			return 0, err
		}
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RepoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ru.check(); err != nil {
				return 0, err
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			if ru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RepoUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RepoUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RepoUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *RepoUpdate) check() error {
	if v, ok := ru.mutation.Name(); ok {
		if err := repo.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Repo.name": %w`, err)}
		}
	}
	if v, ok := ru.mutation.RepoName(); ok {
		if err := repo.RepoNameValidator(v); err != nil {
			return &ValidationError{Name: "repo_name", err: fmt.Errorf(`ent: validator failed for field "Repo.repo_name": %w`, err)}
		}
	}
	return nil
}

func (ru *RepoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   repo.Table,
			Columns: repo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: repo.FieldID,
			},
		},
	}
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repo.FieldName,
		})
	}
	if value, ok := ru.mutation.RepoName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repo.FieldRepoName,
		})
	}
	if value, ok := ru.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repo.FieldURL,
		})
	}
	if value, ok := ru.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repo.FieldDescription,
		})
	}
	if value, ok := ru.mutation.Language(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repo.FieldLanguage,
		})
	}
	if value, ok := ru.mutation.LastUpdated(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: repo.FieldLastUpdated,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{repo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// RepoUpdateOne is the builder for updating a single Repo entity.
type RepoUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RepoMutation
}

// SetName sets the "name" field.
func (ruo *RepoUpdateOne) SetName(s string) *RepoUpdateOne {
	ruo.mutation.SetName(s)
	return ruo
}

// SetRepoName sets the "repo_name" field.
func (ruo *RepoUpdateOne) SetRepoName(s string) *RepoUpdateOne {
	ruo.mutation.SetRepoName(s)
	return ruo
}

// SetURL sets the "url" field.
func (ruo *RepoUpdateOne) SetURL(s string) *RepoUpdateOne {
	ruo.mutation.SetURL(s)
	return ruo
}

// SetDescription sets the "description" field.
func (ruo *RepoUpdateOne) SetDescription(s string) *RepoUpdateOne {
	ruo.mutation.SetDescription(s)
	return ruo
}

// SetLanguage sets the "language" field.
func (ruo *RepoUpdateOne) SetLanguage(s string) *RepoUpdateOne {
	ruo.mutation.SetLanguage(s)
	return ruo
}

// SetLastUpdated sets the "last_updated" field.
func (ruo *RepoUpdateOne) SetLastUpdated(t time.Time) *RepoUpdateOne {
	ruo.mutation.SetLastUpdated(t)
	return ruo
}

// Mutation returns the RepoMutation object of the builder.
func (ruo *RepoUpdateOne) Mutation() *RepoMutation {
	return ruo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RepoUpdateOne) Select(field string, fields ...string) *RepoUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Repo entity.
func (ruo *RepoUpdateOne) Save(ctx context.Context) (*Repo, error) {
	var (
		err  error
		node *Repo
	)
	if len(ruo.hooks) == 0 {
		if err = ruo.check(); err != nil {
			return nil, err
		}
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RepoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ruo.check(); err != nil {
				return nil, err
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			if ruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RepoUpdateOne) SaveX(ctx context.Context) *Repo {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RepoUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RepoUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *RepoUpdateOne) check() error {
	if v, ok := ruo.mutation.Name(); ok {
		if err := repo.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Repo.name": %w`, err)}
		}
	}
	if v, ok := ruo.mutation.RepoName(); ok {
		if err := repo.RepoNameValidator(v); err != nil {
			return &ValidationError{Name: "repo_name", err: fmt.Errorf(`ent: validator failed for field "Repo.repo_name": %w`, err)}
		}
	}
	return nil
}

func (ruo *RepoUpdateOne) sqlSave(ctx context.Context) (_node *Repo, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   repo.Table,
			Columns: repo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: repo.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Repo.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, repo.FieldID)
		for _, f := range fields {
			if !repo.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != repo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repo.FieldName,
		})
	}
	if value, ok := ruo.mutation.RepoName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repo.FieldRepoName,
		})
	}
	if value, ok := ruo.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repo.FieldURL,
		})
	}
	if value, ok := ruo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repo.FieldDescription,
		})
	}
	if value, ok := ruo.mutation.Language(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: repo.FieldLanguage,
		})
	}
	if value, ok := ruo.mutation.LastUpdated(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: repo.FieldLastUpdated,
		})
	}
	_node = &Repo{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{repo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
