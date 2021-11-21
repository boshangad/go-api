// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/boshangad/v1/ent/app"
	"github.com/boshangad/v1/ent/emaillog"
)

// EmailLogCreate is the builder for creating a EmailLog entity.
type EmailLogCreate struct {
	config
	mutation *EmailLogMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (elc *EmailLogCreate) SetCreateTime(i int64) *EmailLogCreate {
	elc.mutation.SetCreateTime(i)
	return elc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (elc *EmailLogCreate) SetNillableCreateTime(i *int64) *EmailLogCreate {
	if i != nil {
		elc.SetCreateTime(*i)
	}
	return elc
}

// SetCreateBy sets the "create_by" field.
func (elc *EmailLogCreate) SetCreateBy(u uint64) *EmailLogCreate {
	elc.mutation.SetCreateBy(u)
	return elc
}

// SetNillableCreateBy sets the "create_by" field if the given value is not nil.
func (elc *EmailLogCreate) SetNillableCreateBy(u *uint64) *EmailLogCreate {
	if u != nil {
		elc.SetCreateBy(*u)
	}
	return elc
}

// SetUpdateTime sets the "update_time" field.
func (elc *EmailLogCreate) SetUpdateTime(i int64) *EmailLogCreate {
	elc.mutation.SetUpdateTime(i)
	return elc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (elc *EmailLogCreate) SetNillableUpdateTime(i *int64) *EmailLogCreate {
	if i != nil {
		elc.SetUpdateTime(*i)
	}
	return elc
}

// SetUpdateBy sets the "update_by" field.
func (elc *EmailLogCreate) SetUpdateBy(u uint64) *EmailLogCreate {
	elc.mutation.SetUpdateBy(u)
	return elc
}

// SetNillableUpdateBy sets the "update_by" field if the given value is not nil.
func (elc *EmailLogCreate) SetNillableUpdateBy(u *uint64) *EmailLogCreate {
	if u != nil {
		elc.SetUpdateBy(*u)
	}
	return elc
}

// SetAppID sets the "app_id" field.
func (elc *EmailLogCreate) SetAppID(u uint64) *EmailLogCreate {
	elc.mutation.SetAppID(u)
	return elc
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (elc *EmailLogCreate) SetNillableAppID(u *uint64) *EmailLogCreate {
	if u != nil {
		elc.SetAppID(*u)
	}
	return elc
}

// SetEmail sets the "email" field.
func (elc *EmailLogCreate) SetEmail(s string) *EmailLogCreate {
	elc.mutation.SetEmail(s)
	return elc
}

// SetScope sets the "scope" field.
func (elc *EmailLogCreate) SetScope(s string) *EmailLogCreate {
	elc.mutation.SetScope(s)
	return elc
}

// SetNillableScope sets the "scope" field if the given value is not nil.
func (elc *EmailLogCreate) SetNillableScope(s *string) *EmailLogCreate {
	if s != nil {
		elc.SetScope(*s)
	}
	return elc
}

// SetTypeID sets the "type_id" field.
func (elc *EmailLogCreate) SetTypeID(u uint64) *EmailLogCreate {
	elc.mutation.SetTypeID(u)
	return elc
}

// SetNillableTypeID sets the "type_id" field if the given value is not nil.
func (elc *EmailLogCreate) SetNillableTypeID(u *uint64) *EmailLogCreate {
	if u != nil {
		elc.SetTypeID(*u)
	}
	return elc
}

// SetGateway sets the "gateway" field.
func (elc *EmailLogCreate) SetGateway(s string) *EmailLogCreate {
	elc.mutation.SetGateway(s)
	return elc
}

// SetNillableGateway sets the "gateway" field if the given value is not nil.
func (elc *EmailLogCreate) SetNillableGateway(s *string) *EmailLogCreate {
	if s != nil {
		elc.SetGateway(*s)
	}
	return elc
}

// SetIP sets the "ip" field.
func (elc *EmailLogCreate) SetIP(s string) *EmailLogCreate {
	elc.mutation.SetIP(s)
	return elc
}

// SetNillableIP sets the "ip" field if the given value is not nil.
func (elc *EmailLogCreate) SetNillableIP(s *string) *EmailLogCreate {
	if s != nil {
		elc.SetIP(*s)
	}
	return elc
}

// SetFromName sets the "from_name" field.
func (elc *EmailLogCreate) SetFromName(s string) *EmailLogCreate {
	elc.mutation.SetFromName(s)
	return elc
}

// SetNillableFromName sets the "from_name" field if the given value is not nil.
func (elc *EmailLogCreate) SetNillableFromName(s *string) *EmailLogCreate {
	if s != nil {
		elc.SetFromName(*s)
	}
	return elc
}

// SetFromAddress sets the "from_address" field.
func (elc *EmailLogCreate) SetFromAddress(s string) *EmailLogCreate {
	elc.mutation.SetFromAddress(s)
	return elc
}

// SetNillableFromAddress sets the "from_address" field if the given value is not nil.
func (elc *EmailLogCreate) SetNillableFromAddress(s *string) *EmailLogCreate {
	if s != nil {
		elc.SetFromAddress(*s)
	}
	return elc
}

// SetTitle sets the "title" field.
func (elc *EmailLogCreate) SetTitle(s string) *EmailLogCreate {
	elc.mutation.SetTitle(s)
	return elc
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (elc *EmailLogCreate) SetNillableTitle(s *string) *EmailLogCreate {
	if s != nil {
		elc.SetTitle(*s)
	}
	return elc
}

// SetContent sets the "content" field.
func (elc *EmailLogCreate) SetContent(s string) *EmailLogCreate {
	elc.mutation.SetContent(s)
	return elc
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (elc *EmailLogCreate) SetNillableContent(s *string) *EmailLogCreate {
	if s != nil {
		elc.SetContent(*s)
	}
	return elc
}

// SetData sets the "data" field.
func (elc *EmailLogCreate) SetData(s string) *EmailLogCreate {
	elc.mutation.SetData(s)
	return elc
}

// SetNillableData sets the "data" field if the given value is not nil.
func (elc *EmailLogCreate) SetNillableData(s *string) *EmailLogCreate {
	if s != nil {
		elc.SetData(*s)
	}
	return elc
}

// SetCheckCount sets the "check_count" field.
func (elc *EmailLogCreate) SetCheckCount(u uint8) *EmailLogCreate {
	elc.mutation.SetCheckCount(u)
	return elc
}

// SetNillableCheckCount sets the "check_count" field if the given value is not nil.
func (elc *EmailLogCreate) SetNillableCheckCount(u *uint8) *EmailLogCreate {
	if u != nil {
		elc.SetCheckCount(*u)
	}
	return elc
}

// SetStatus sets the "status" field.
func (elc *EmailLogCreate) SetStatus(u uint) *EmailLogCreate {
	elc.mutation.SetStatus(u)
	return elc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (elc *EmailLogCreate) SetNillableStatus(u *uint) *EmailLogCreate {
	if u != nil {
		elc.SetStatus(*u)
	}
	return elc
}

// SetReturnMsg sets the "return_msg" field.
func (elc *EmailLogCreate) SetReturnMsg(s string) *EmailLogCreate {
	elc.mutation.SetReturnMsg(s)
	return elc
}

// SetNillableReturnMsg sets the "return_msg" field if the given value is not nil.
func (elc *EmailLogCreate) SetNillableReturnMsg(s *string) *EmailLogCreate {
	if s != nil {
		elc.SetReturnMsg(*s)
	}
	return elc
}

// SetID sets the "id" field.
func (elc *EmailLogCreate) SetID(u uint64) *EmailLogCreate {
	elc.mutation.SetID(u)
	return elc
}

// SetApp sets the "app" edge to the App entity.
func (elc *EmailLogCreate) SetApp(a *App) *EmailLogCreate {
	return elc.SetAppID(a.ID)
}

// Mutation returns the EmailLogMutation object of the builder.
func (elc *EmailLogCreate) Mutation() *EmailLogMutation {
	return elc.mutation
}

// Save creates the EmailLog in the database.
func (elc *EmailLogCreate) Save(ctx context.Context) (*EmailLog, error) {
	var (
		err  error
		node *EmailLog
	)
	elc.defaults()
	if len(elc.hooks) == 0 {
		if err = elc.check(); err != nil {
			return nil, err
		}
		node, err = elc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EmailLogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = elc.check(); err != nil {
				return nil, err
			}
			elc.mutation = mutation
			if node, err = elc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(elc.hooks) - 1; i >= 0; i-- {
			if elc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = elc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, elc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (elc *EmailLogCreate) SaveX(ctx context.Context) *EmailLog {
	v, err := elc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (elc *EmailLogCreate) Exec(ctx context.Context) error {
	_, err := elc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (elc *EmailLogCreate) ExecX(ctx context.Context) {
	if err := elc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (elc *EmailLogCreate) defaults() {
	if _, ok := elc.mutation.CreateTime(); !ok {
		v := emaillog.DefaultCreateTime
		elc.mutation.SetCreateTime(v)
	}
	if _, ok := elc.mutation.CreateBy(); !ok {
		v := emaillog.DefaultCreateBy
		elc.mutation.SetCreateBy(v)
	}
	if _, ok := elc.mutation.UpdateTime(); !ok {
		v := emaillog.DefaultUpdateTime
		elc.mutation.SetUpdateTime(v)
	}
	if _, ok := elc.mutation.UpdateBy(); !ok {
		v := emaillog.DefaultUpdateBy
		elc.mutation.SetUpdateBy(v)
	}
	if _, ok := elc.mutation.AppID(); !ok {
		v := emaillog.DefaultAppID
		elc.mutation.SetAppID(v)
	}
	if _, ok := elc.mutation.Scope(); !ok {
		v := emaillog.DefaultScope
		elc.mutation.SetScope(v)
	}
	if _, ok := elc.mutation.TypeID(); !ok {
		v := emaillog.DefaultTypeID
		elc.mutation.SetTypeID(v)
	}
	if _, ok := elc.mutation.Gateway(); !ok {
		v := emaillog.DefaultGateway
		elc.mutation.SetGateway(v)
	}
	if _, ok := elc.mutation.IP(); !ok {
		v := emaillog.DefaultIP
		elc.mutation.SetIP(v)
	}
	if _, ok := elc.mutation.FromName(); !ok {
		v := emaillog.DefaultFromName
		elc.mutation.SetFromName(v)
	}
	if _, ok := elc.mutation.FromAddress(); !ok {
		v := emaillog.DefaultFromAddress
		elc.mutation.SetFromAddress(v)
	}
	if _, ok := elc.mutation.Title(); !ok {
		v := emaillog.DefaultTitle
		elc.mutation.SetTitle(v)
	}
	if _, ok := elc.mutation.Content(); !ok {
		v := emaillog.DefaultContent
		elc.mutation.SetContent(v)
	}
	if _, ok := elc.mutation.Data(); !ok {
		v := emaillog.DefaultData
		elc.mutation.SetData(v)
	}
	if _, ok := elc.mutation.CheckCount(); !ok {
		v := emaillog.DefaultCheckCount
		elc.mutation.SetCheckCount(v)
	}
	if _, ok := elc.mutation.Status(); !ok {
		v := emaillog.DefaultStatus
		elc.mutation.SetStatus(v)
	}
	if _, ok := elc.mutation.ReturnMsg(); !ok {
		v := emaillog.DefaultReturnMsg
		elc.mutation.SetReturnMsg(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (elc *EmailLogCreate) check() error {
	if _, ok := elc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "create_time"`)}
	}
	if _, ok := elc.mutation.CreateBy(); !ok {
		return &ValidationError{Name: "create_by", err: errors.New(`ent: missing required field "create_by"`)}
	}
	if _, ok := elc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "update_time"`)}
	}
	if _, ok := elc.mutation.UpdateBy(); !ok {
		return &ValidationError{Name: "update_by", err: errors.New(`ent: missing required field "update_by"`)}
	}
	if _, ok := elc.mutation.AppID(); !ok {
		return &ValidationError{Name: "app_id", err: errors.New(`ent: missing required field "app_id"`)}
	}
	if _, ok := elc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "email"`)}
	}
	if v, ok := elc.mutation.Email(); ok {
		if err := emaillog.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "email": %w`, err)}
		}
	}
	if _, ok := elc.mutation.Scope(); !ok {
		return &ValidationError{Name: "scope", err: errors.New(`ent: missing required field "scope"`)}
	}
	if v, ok := elc.mutation.Scope(); ok {
		if err := emaillog.ScopeValidator(v); err != nil {
			return &ValidationError{Name: "scope", err: fmt.Errorf(`ent: validator failed for field "scope": %w`, err)}
		}
	}
	if _, ok := elc.mutation.TypeID(); !ok {
		return &ValidationError{Name: "type_id", err: errors.New(`ent: missing required field "type_id"`)}
	}
	if _, ok := elc.mutation.Gateway(); !ok {
		return &ValidationError{Name: "gateway", err: errors.New(`ent: missing required field "gateway"`)}
	}
	if v, ok := elc.mutation.Gateway(); ok {
		if err := emaillog.GatewayValidator(v); err != nil {
			return &ValidationError{Name: "gateway", err: fmt.Errorf(`ent: validator failed for field "gateway": %w`, err)}
		}
	}
	if _, ok := elc.mutation.IP(); !ok {
		return &ValidationError{Name: "ip", err: errors.New(`ent: missing required field "ip"`)}
	}
	if v, ok := elc.mutation.IP(); ok {
		if err := emaillog.IPValidator(v); err != nil {
			return &ValidationError{Name: "ip", err: fmt.Errorf(`ent: validator failed for field "ip": %w`, err)}
		}
	}
	if _, ok := elc.mutation.FromName(); !ok {
		return &ValidationError{Name: "from_name", err: errors.New(`ent: missing required field "from_name"`)}
	}
	if v, ok := elc.mutation.FromName(); ok {
		if err := emaillog.FromNameValidator(v); err != nil {
			return &ValidationError{Name: "from_name", err: fmt.Errorf(`ent: validator failed for field "from_name": %w`, err)}
		}
	}
	if _, ok := elc.mutation.FromAddress(); !ok {
		return &ValidationError{Name: "from_address", err: errors.New(`ent: missing required field "from_address"`)}
	}
	if v, ok := elc.mutation.FromAddress(); ok {
		if err := emaillog.FromAddressValidator(v); err != nil {
			return &ValidationError{Name: "from_address", err: fmt.Errorf(`ent: validator failed for field "from_address": %w`, err)}
		}
	}
	if _, ok := elc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "title"`)}
	}
	if v, ok := elc.mutation.Title(); ok {
		if err := emaillog.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "title": %w`, err)}
		}
	}
	if _, ok := elc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "content"`)}
	}
	if _, ok := elc.mutation.Data(); !ok {
		return &ValidationError{Name: "data", err: errors.New(`ent: missing required field "data"`)}
	}
	if v, ok := elc.mutation.Data(); ok {
		if err := emaillog.DataValidator(v); err != nil {
			return &ValidationError{Name: "data", err: fmt.Errorf(`ent: validator failed for field "data": %w`, err)}
		}
	}
	if _, ok := elc.mutation.CheckCount(); !ok {
		return &ValidationError{Name: "check_count", err: errors.New(`ent: missing required field "check_count"`)}
	}
	if _, ok := elc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "status"`)}
	}
	if _, ok := elc.mutation.ReturnMsg(); !ok {
		return &ValidationError{Name: "return_msg", err: errors.New(`ent: missing required field "return_msg"`)}
	}
	if v, ok := elc.mutation.ReturnMsg(); ok {
		if err := emaillog.ReturnMsgValidator(v); err != nil {
			return &ValidationError{Name: "return_msg", err: fmt.Errorf(`ent: validator failed for field "return_msg": %w`, err)}
		}
	}
	if _, ok := elc.mutation.AppID(); !ok {
		return &ValidationError{Name: "app", err: errors.New("ent: missing required edge \"app\"")}
	}
	return nil
}

func (elc *EmailLogCreate) sqlSave(ctx context.Context) (*EmailLog, error) {
	_node, _spec := elc.createSpec()
	if err := sqlgraph.CreateNode(ctx, elc.driver, _spec); err != nil {
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

func (elc *EmailLogCreate) createSpec() (*EmailLog, *sqlgraph.CreateSpec) {
	var (
		_node = &EmailLog{config: elc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: emaillog.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: emaillog.FieldID,
			},
		}
	)
	if id, ok := elc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := elc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: emaillog.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := elc.mutation.CreateBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: emaillog.FieldCreateBy,
		})
		_node.CreateBy = value
	}
	if value, ok := elc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: emaillog.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := elc.mutation.UpdateBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: emaillog.FieldUpdateBy,
		})
		_node.UpdateBy = value
	}
	if value, ok := elc.mutation.Email(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: emaillog.FieldEmail,
		})
		_node.Email = value
	}
	if value, ok := elc.mutation.Scope(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: emaillog.FieldScope,
		})
		_node.Scope = value
	}
	if value, ok := elc.mutation.TypeID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: emaillog.FieldTypeID,
		})
		_node.TypeID = value
	}
	if value, ok := elc.mutation.Gateway(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: emaillog.FieldGateway,
		})
		_node.Gateway = value
	}
	if value, ok := elc.mutation.IP(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: emaillog.FieldIP,
		})
		_node.IP = value
	}
	if value, ok := elc.mutation.FromName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: emaillog.FieldFromName,
		})
		_node.FromName = value
	}
	if value, ok := elc.mutation.FromAddress(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: emaillog.FieldFromAddress,
		})
		_node.FromAddress = value
	}
	if value, ok := elc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: emaillog.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := elc.mutation.Content(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: emaillog.FieldContent,
		})
		_node.Content = value
	}
	if value, ok := elc.mutation.Data(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: emaillog.FieldData,
		})
		_node.Data = value
	}
	if value, ok := elc.mutation.CheckCount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: emaillog.FieldCheckCount,
		})
		_node.CheckCount = value
	}
	if value, ok := elc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: emaillog.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := elc.mutation.ReturnMsg(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: emaillog.FieldReturnMsg,
		})
		_node.ReturnMsg = value
	}
	if nodes := elc.mutation.AppIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   emaillog.AppTable,
			Columns: []string{emaillog.AppColumn},
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

// EmailLogCreateBulk is the builder for creating many EmailLog entities in bulk.
type EmailLogCreateBulk struct {
	config
	builders []*EmailLogCreate
}

// Save creates the EmailLog entities in the database.
func (elcb *EmailLogCreateBulk) Save(ctx context.Context) ([]*EmailLog, error) {
	specs := make([]*sqlgraph.CreateSpec, len(elcb.builders))
	nodes := make([]*EmailLog, len(elcb.builders))
	mutators := make([]Mutator, len(elcb.builders))
	for i := range elcb.builders {
		func(i int, root context.Context) {
			builder := elcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EmailLogMutation)
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
					_, err = mutators[i+1].Mutate(root, elcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, elcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, elcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (elcb *EmailLogCreateBulk) SaveX(ctx context.Context) []*EmailLog {
	v, err := elcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (elcb *EmailLogCreateBulk) Exec(ctx context.Context) error {
	_, err := elcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (elcb *EmailLogCreateBulk) ExecX(ctx context.Context) {
	if err := elcb.Exec(ctx); err != nil {
		panic(err)
	}
}
