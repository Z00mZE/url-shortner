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
	"github.com/Z00mZE/url-shortner/ent/predicate"
	"github.com/Z00mZE/url-shortner/ent/urls"
)

// UrlsUpdate is the builder for updating Urls entities.
type UrlsUpdate struct {
	config
	hooks    []Hook
	mutation *UrlsMutation
}

// Where appends a list predicates to the UrlsUpdate builder.
func (uu *UrlsUpdate) Where(ps ...predicate.Urls) *UrlsUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetURL sets the "url" field.
func (uu *UrlsUpdate) SetURL(s string) *UrlsUpdate {
	uu.mutation.SetURL(s)
	return uu
}

// SetExpiredAt sets the "expired_at" field.
func (uu *UrlsUpdate) SetExpiredAt(t time.Time) *UrlsUpdate {
	uu.mutation.SetExpiredAt(t)
	return uu
}

// SetNillableExpiredAt sets the "expired_at" field if the given value is not nil.
func (uu *UrlsUpdate) SetNillableExpiredAt(t *time.Time) *UrlsUpdate {
	if t != nil {
		uu.SetExpiredAt(*t)
	}
	return uu
}

// Mutation returns the UrlsMutation object of the builder.
func (uu *UrlsUpdate) Mutation() *UrlsMutation {
	return uu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UrlsUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(uu.hooks) == 0 {
		if err = uu.check(); err != nil {
			return 0, err
		}
		affected, err = uu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UrlsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uu.check(); err != nil {
				return 0, err
			}
			uu.mutation = mutation
			affected, err = uu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uu.hooks) - 1; i >= 0; i-- {
			if uu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UrlsUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UrlsUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UrlsUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UrlsUpdate) check() error {
	if v, ok := uu.mutation.URL(); ok {
		if err := urls.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "Urls.url": %w`, err)}
		}
	}
	return nil
}

func (uu *UrlsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   urls.Table,
			Columns: urls.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: urls.FieldID,
			},
		},
	}
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: urls.FieldURL,
		})
	}
	if value, ok := uu.mutation.ExpiredAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: urls.FieldExpiredAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{urls.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// UrlsUpdateOne is the builder for updating a single Urls entity.
type UrlsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UrlsMutation
}

// SetURL sets the "url" field.
func (uuo *UrlsUpdateOne) SetURL(s string) *UrlsUpdateOne {
	uuo.mutation.SetURL(s)
	return uuo
}

// SetExpiredAt sets the "expired_at" field.
func (uuo *UrlsUpdateOne) SetExpiredAt(t time.Time) *UrlsUpdateOne {
	uuo.mutation.SetExpiredAt(t)
	return uuo
}

// SetNillableExpiredAt sets the "expired_at" field if the given value is not nil.
func (uuo *UrlsUpdateOne) SetNillableExpiredAt(t *time.Time) *UrlsUpdateOne {
	if t != nil {
		uuo.SetExpiredAt(*t)
	}
	return uuo
}

// Mutation returns the UrlsMutation object of the builder.
func (uuo *UrlsUpdateOne) Mutation() *UrlsMutation {
	return uuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UrlsUpdateOne) Select(field string, fields ...string) *UrlsUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated Urls entity.
func (uuo *UrlsUpdateOne) Save(ctx context.Context) (*Urls, error) {
	var (
		err  error
		node *Urls
	)
	if len(uuo.hooks) == 0 {
		if err = uuo.check(); err != nil {
			return nil, err
		}
		node, err = uuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UrlsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uuo.check(); err != nil {
				return nil, err
			}
			uuo.mutation = mutation
			node, err = uuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uuo.hooks) - 1; i >= 0; i-- {
			if uuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UrlsUpdateOne) SaveX(ctx context.Context) *Urls {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UrlsUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UrlsUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UrlsUpdateOne) check() error {
	if v, ok := uuo.mutation.URL(); ok {
		if err := urls.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "Urls.url": %w`, err)}
		}
	}
	return nil
}

func (uuo *UrlsUpdateOne) sqlSave(ctx context.Context) (_node *Urls, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   urls.Table,
			Columns: urls.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: urls.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Urls.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, urls.FieldID)
		for _, f := range fields {
			if !urls.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != urls.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: urls.FieldURL,
		})
	}
	if value, ok := uuo.mutation.ExpiredAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: urls.FieldExpiredAt,
		})
	}
	_node = &Urls{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{urls.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}