// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/boshangad/v1/ent/app"
	"github.com/boshangad/v1/ent/appoption"
)

// AppOptionCreate is the builder for creating a AppOption entity.
type AppOptionCreate struct {
	config
	mutation *AppOptionMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (aoc *AppOptionCreate) SetCreateTime(i int64) *AppOptionCreate {
	aoc.mutation.SetCreateTime(i)
	return aoc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (aoc *AppOptionCreate) SetNillableCreateTime(i *int64) *AppOptionCreate {
	if i != nil {
		aoc.SetCreateTime(*i)
	}
	return aoc
}

// SetCreateBy sets the "create_by" field.
func (aoc *AppOptionCreate) SetCreateBy(u uint64) *AppOptionCreate {
	aoc.mutation.SetCreateBy(u)
	return aoc
}

// SetNillableCreateBy sets the "create_by" field if the given value is not nil.
func (aoc *AppOptionCreate) SetNillableCreateBy(u *uint64) *AppOptionCreate {
	if u != nil {
		aoc.SetCreateBy(*u)
	}
	return aoc
}

// SetUpdateTime sets the "update_time" field.
func (aoc *AppOptionCreate) SetUpdateTime(i int64) *AppOptionCreate {
	aoc.mutation.SetUpdateTime(i)
	return aoc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (aoc *AppOptionCreate) SetNillableUpdateTime(i *int64) *AppOptionCreate {
	if i != nil {
		aoc.SetUpdateTime(*i)
	}
	return aoc
}

// SetUpdateBy sets the "update_by" field.
func (aoc *AppOptionCreate) SetUpdateBy(u uint64) *AppOptionCreate {
	aoc.mutation.SetUpdateBy(u)
	return aoc
}

// SetNillableUpdateBy sets the "update_by" field if the given value is not nil.
func (aoc *AppOptionCreate) SetNillableUpdateBy(u *uint64) *AppOptionCreate {
	if u != nil {
		aoc.SetUpdateBy(*u)
	}
	return aoc
}

// SetAppID sets the "app_id" field.
func (aoc *AppOptionCreate) SetAppID(u uint64) *AppOptionCreate {
	aoc.mutation.SetAppID(u)
	return aoc
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (aoc *AppOptionCreate) SetNillableAppID(u *uint64) *AppOptionCreate {
	if u != nil {
		aoc.SetAppID(*u)
	}
	return aoc
}

// SetTitle sets the "title" field.
func (aoc *AppOptionCreate) SetTitle(s string) *AppOptionCreate {
	aoc.mutation.SetTitle(s)
	return aoc
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (aoc *AppOptionCreate) SetNillableTitle(s *string) *AppOptionCreate {
	if s != nil {
		aoc.SetTitle(*s)
	}
	return aoc
}

// SetDescription sets the "description" field.
func (aoc *AppOptionCreate) SetDescription(s string) *AppOptionCreate {
	aoc.mutation.SetDescription(s)
	return aoc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (aoc *AppOptionCreate) SetNillableDescription(s *string) *AppOptionCreate {
	if s != nil {
		aoc.SetDescription(*s)
	}
	return aoc
}

// SetName sets the "name" field.
func (aoc *AppOptionCreate) SetName(s string) *AppOptionCreate {
	aoc.mutation.SetName(s)
	return aoc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (aoc *AppOptionCreate) SetNillableName(s *string) *AppOptionCreate {
	if s != nil {
		aoc.SetName(*s)
	}
	return aoc
}

// SetValue sets the "value" field.
func (aoc *AppOptionCreate) SetValue(s string) *AppOptionCreate {
	aoc.mutation.SetValue(s)
	return aoc
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (aoc *AppOptionCreate) SetNillableValue(s *string) *AppOptionCreate {
	if s != nil {
		aoc.SetValue(*s)
	}
	return aoc
}

// SetExpireTime sets the "expire_time" field.
func (aoc *AppOptionCreate) SetExpireTime(u uint64) *AppOptionCreate {
	aoc.mutation.SetExpireTime(u)
	return aoc
}

// SetNillableExpireTime sets the "expire_time" field if the given value is not nil.
func (aoc *AppOptionCreate) SetNillableExpireTime(u *uint64) *AppOptionCreate {
	if u != nil {
		aoc.SetExpireTime(*u)
	}
	return aoc
}

// SetEditType sets the "edit_type" field.
func (aoc *AppOptionCreate) SetEditType(u uint) *AppOptionCreate {
	aoc.mutation.SetEditType(u)
	return aoc
}

// SetNillableEditType sets the "edit_type" field if the given value is not nil.
func (aoc *AppOptionCreate) SetNillableEditType(u *uint) *AppOptionCreate {
	if u != nil {
		aoc.SetEditType(*u)
	}
	return aoc
}

// SetID sets the "id" field.
func (aoc *AppOptionCreate) SetID(u uint64) *AppOptionCreate {
	aoc.mutation.SetID(u)
	return aoc
}

// SetApp sets the "app" edge to the App entity.
func (aoc *AppOptionCreate) SetApp(a *App) *AppOptionCreate {
	return aoc.SetAppID(a.ID)
}

// Mutation returns the AppOptionMutation object of the builder.
func (aoc *AppOptionCreate) Mutation() *AppOptionMutation {
	return aoc.mutation
}

// Save creates the AppOption in the database.
func (aoc *AppOptionCreate) Save(ctx context.Context) (*AppOption, error) {
	var (
		err  error
		node *AppOption
	)
	aoc.defaults()
	if len(aoc.hooks) == 0 {
		if err = aoc.check(); err != nil {
			return nil, err
		}
		node, err = aoc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppOptionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = aoc.check(); err != nil {
				return nil, err
			}
			aoc.mutation = mutation
			if node, err = aoc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(aoc.hooks) - 1; i >= 0; i-- {
			if aoc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = aoc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aoc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (aoc *AppOptionCreate) SaveX(ctx context.Context) *AppOption {
	v, err := aoc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (aoc *AppOptionCreate) Exec(ctx context.Context) error {
	_, err := aoc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aoc *AppOptionCreate) ExecX(ctx context.Context) {
	if err := aoc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aoc *AppOptionCreate) defaults() {
	if _, ok := aoc.mutation.CreateTime(); !ok {
		v := appoption.DefaultCreateTime
		aoc.mutation.SetCreateTime(v)
	}
	if _, ok := aoc.mutation.CreateBy(); !ok {
		v := appoption.DefaultCreateBy
		aoc.mutation.SetCreateBy(v)
	}
	if _, ok := aoc.mutation.UpdateTime(); !ok {
		v := appoption.DefaultUpdateTime
		aoc.mutation.SetUpdateTime(v)
	}
	if _, ok := aoc.mutation.UpdateBy(); !ok {
		v := appoption.DefaultUpdateBy
		aoc.mutation.SetUpdateBy(v)
	}
	if _, ok := aoc.mutation.AppID(); !ok {
		v := appoption.DefaultAppID
		aoc.mutation.SetAppID(v)
	}
	if _, ok := aoc.mutation.Title(); !ok {
		v := appoption.DefaultTitle
		aoc.mutation.SetTitle(v)
	}
	if _, ok := aoc.mutation.Description(); !ok {
		v := appoption.DefaultDescription
		aoc.mutation.SetDescription(v)
	}
	if _, ok := aoc.mutation.Name(); !ok {
		v := appoption.DefaultName
		aoc.mutation.SetName(v)
	}
	if _, ok := aoc.mutation.Value(); !ok {
		v := appoption.DefaultValue
		aoc.mutation.SetValue(v)
	}
	if _, ok := aoc.mutation.ExpireTime(); !ok {
		v := appoption.DefaultExpireTime
		aoc.mutation.SetExpireTime(v)
	}
	if _, ok := aoc.mutation.EditType(); !ok {
		v := appoption.DefaultEditType
		aoc.mutation.SetEditType(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (aoc *AppOptionCreate) check() error {
	if _, ok := aoc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "create_time"`)}
	}
	if _, ok := aoc.mutation.CreateBy(); !ok {
		return &ValidationError{Name: "create_by", err: errors.New(`ent: missing required field "create_by"`)}
	}
	if _, ok := aoc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "update_time"`)}
	}
	if _, ok := aoc.mutation.UpdateBy(); !ok {
		return &ValidationError{Name: "update_by", err: errors.New(`ent: missing required field "update_by"`)}
	}
	if _, ok := aoc.mutation.AppID(); !ok {
		return &ValidationError{Name: "app_id", err: errors.New(`ent: missing required field "app_id"`)}
	}
	if _, ok := aoc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "title"`)}
	}
	if v, ok := aoc.mutation.Title(); ok {
		if err := appoption.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "title": %w`, err)}
		}
	}
	if _, ok := aoc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "description"`)}
	}
	if v, ok := aoc.mutation.Description(); ok {
		if err := appoption.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "description": %w`, err)}
		}
	}
	if _, ok := aoc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	if v, ok := aoc.mutation.Name(); ok {
		if err := appoption.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "name": %w`, err)}
		}
	}
	if _, ok := aoc.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`ent: missing required field "value"`)}
	}
	if v, ok := aoc.mutation.Value(); ok {
		if err := appoption.ValueValidator(v); err != nil {
			return &ValidationError{Name: "value", err: fmt.Errorf(`ent: validator failed for field "value": %w`, err)}
		}
	}
	if _, ok := aoc.mutation.ExpireTime(); !ok {
		return &ValidationError{Name: "expire_time", err: errors.New(`ent: missing required field "expire_time"`)}
	}
	if _, ok := aoc.mutation.EditType(); !ok {
		return &ValidationError{Name: "edit_type", err: errors.New(`ent: missing required field "edit_type"`)}
	}
	if _, ok := aoc.mutation.AppID(); !ok {
		return &ValidationError{Name: "app", err: errors.New("ent: missing required edge \"app\"")}
	}
	return nil
}

func (aoc *AppOptionCreate) sqlSave(ctx context.Context) (*AppOption, error) {
	_node, _spec := aoc.createSpec()
	if err := sqlgraph.CreateNode(ctx, aoc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	return _node, nil
}

func (aoc *AppOptionCreate) createSpec() (*AppOption, *sqlgraph.CreateSpec) {
	var (
		_node = &AppOption{config: aoc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: appoption.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: appoption.FieldID,
			},
		}
	)
	if id, ok := aoc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := aoc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: appoption.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := aoc.mutation.CreateBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: appoption.FieldCreateBy,
		})
		_node.CreateBy = value
	}
	if value, ok := aoc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: appoption.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := aoc.mutation.UpdateBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: appoption.FieldUpdateBy,
		})
		_node.UpdateBy = value
	}
	if value, ok := aoc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appoption.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := aoc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appoption.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := aoc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appoption.FieldName,
		})
		_node.Name = value
	}
	if value, ok := aoc.mutation.Value(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appoption.FieldValue,
		})
		_node.Value = value
	}
	if value, ok := aoc.mutation.ExpireTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: appoption.FieldExpireTime,
		})
		_node.ExpireTime = value
	}
	if value, ok := aoc.mutation.EditType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: appoption.FieldEditType,
		})
		_node.EditType = value
	}
	if nodes := aoc.mutation.AppIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   appoption.AppTable,
			Columns: []string{appoption.AppColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: app.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.AppID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AppOptionCreateBulk is the builder for creating many AppOption entities in bulk.
type AppOptionCreateBulk struct {
	config
	builders []*AppOptionCreate
}

// Save creates the AppOption entities in the database.
func (aocb *AppOptionCreateBulk) Save(ctx context.Context) ([]*AppOption, error) {
	specs := make([]*sqlgraph.CreateSpec, len(aocb.builders))
	nodes := make([]*AppOption, len(aocb.builders))
	mutators := make([]Mutator, len(aocb.builders))
	for i := range aocb.builders {
		func(i int, root context.Context) {
			builder := aocb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AppOptionMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, aocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, aocb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, aocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (aocb *AppOptionCreateBulk) SaveX(ctx context.Context) []*AppOption {
	v, err := aocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (aocb *AppOptionCreateBulk) Exec(ctx context.Context) error {
	_, err := aocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aocb *AppOptionCreateBulk) ExecX(ctx context.Context) {
	if err := aocb.Exec(ctx); err != nil {
		panic(err)
	}
}
