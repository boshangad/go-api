// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/boshangad/v1/ent/authrole"
	"github.com/boshangad/v1/ent/internal"
	"github.com/boshangad/v1/ent/predicate"
)

// AuthRoleQuery is the builder for querying AuthRole entities.
type AuthRoleQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.AuthRole
	modifiers  []func(s *sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AuthRoleQuery builder.
func (arq *AuthRoleQuery) Where(ps ...predicate.AuthRole) *AuthRoleQuery {
	arq.predicates = append(arq.predicates, ps...)
	return arq
}

// Limit adds a limit step to the query.
func (arq *AuthRoleQuery) Limit(limit int) *AuthRoleQuery {
	arq.limit = &limit
	return arq
}

// Offset adds an offset step to the query.
func (arq *AuthRoleQuery) Offset(offset int) *AuthRoleQuery {
	arq.offset = &offset
	return arq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (arq *AuthRoleQuery) Unique(unique bool) *AuthRoleQuery {
	arq.unique = &unique
	return arq
}

// Order adds an order step to the query.
func (arq *AuthRoleQuery) Order(o ...OrderFunc) *AuthRoleQuery {
	arq.order = append(arq.order, o...)
	return arq
}

// First returns the first AuthRole entity from the query.
// Returns a *NotFoundError when no AuthRole was found.
func (arq *AuthRoleQuery) First(ctx context.Context) (*AuthRole, error) {
	nodes, err := arq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{authrole.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (arq *AuthRoleQuery) FirstX(ctx context.Context) *AuthRole {
	node, err := arq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AuthRole ID from the query.
// Returns a *NotFoundError when no AuthRole ID was found.
func (arq *AuthRoleQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = arq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{authrole.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (arq *AuthRoleQuery) FirstIDX(ctx context.Context) int {
	id, err := arq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AuthRole entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one AuthRole entity is not found.
// Returns a *NotFoundError when no AuthRole entities are found.
func (arq *AuthRoleQuery) Only(ctx context.Context) (*AuthRole, error) {
	nodes, err := arq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{authrole.Label}
	default:
		return nil, &NotSingularError{authrole.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (arq *AuthRoleQuery) OnlyX(ctx context.Context) *AuthRole {
	node, err := arq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AuthRole ID in the query.
// Returns a *NotSingularError when exactly one AuthRole ID is not found.
// Returns a *NotFoundError when no entities are found.
func (arq *AuthRoleQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = arq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{authrole.Label}
	default:
		err = &NotSingularError{authrole.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (arq *AuthRoleQuery) OnlyIDX(ctx context.Context) int {
	id, err := arq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AuthRoles.
func (arq *AuthRoleQuery) All(ctx context.Context) ([]*AuthRole, error) {
	if err := arq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return arq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (arq *AuthRoleQuery) AllX(ctx context.Context) []*AuthRole {
	nodes, err := arq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AuthRole IDs.
func (arq *AuthRoleQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := arq.Select(authrole.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (arq *AuthRoleQuery) IDsX(ctx context.Context) []int {
	ids, err := arq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (arq *AuthRoleQuery) Count(ctx context.Context) (int, error) {
	if err := arq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return arq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (arq *AuthRoleQuery) CountX(ctx context.Context) int {
	count, err := arq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (arq *AuthRoleQuery) Exist(ctx context.Context) (bool, error) {
	if err := arq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return arq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (arq *AuthRoleQuery) ExistX(ctx context.Context) bool {
	exist, err := arq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AuthRoleQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (arq *AuthRoleQuery) Clone() *AuthRoleQuery {
	if arq == nil {
		return nil
	}
	return &AuthRoleQuery{
		config:     arq.config,
		limit:      arq.limit,
		offset:     arq.offset,
		order:      append([]OrderFunc{}, arq.order...),
		predicates: append([]predicate.AuthRole{}, arq.predicates...),
		// clone intermediate query.
		sql:  arq.sql.Clone(),
		path: arq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (arq *AuthRoleQuery) GroupBy(field string, fields ...string) *AuthRoleGroupBy {
	group := &AuthRoleGroupBy{config: arq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := arq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return arq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (arq *AuthRoleQuery) Select(fields ...string) *AuthRoleSelect {
	arq.fields = append(arq.fields, fields...)
	return &AuthRoleSelect{AuthRoleQuery: arq}
}

func (arq *AuthRoleQuery) prepareQuery(ctx context.Context) error {
	for _, f := range arq.fields {
		if !authrole.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if arq.path != nil {
		prev, err := arq.path(ctx)
		if err != nil {
			return err
		}
		arq.sql = prev
	}
	return nil
}

func (arq *AuthRoleQuery) sqlAll(ctx context.Context) ([]*AuthRole, error) {
	var (
		nodes = []*AuthRole{}
		_spec = arq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &AuthRole{config: arq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if len(arq.modifiers) > 0 {
		_spec.Modifiers = arq.modifiers
	}
	_spec.Node.Schema = arq.schemaConfig.AuthRole
	ctx = internal.NewSchemaConfigContext(ctx, arq.schemaConfig)
	if err := sqlgraph.QueryNodes(ctx, arq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (arq *AuthRoleQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := arq.querySpec()
	if len(arq.modifiers) > 0 {
		_spec.Modifiers = arq.modifiers
	}
	_spec.Node.Schema = arq.schemaConfig.AuthRole
	ctx = internal.NewSchemaConfigContext(ctx, arq.schemaConfig)
	return sqlgraph.CountNodes(ctx, arq.driver, _spec)
}

func (arq *AuthRoleQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := arq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (arq *AuthRoleQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   authrole.Table,
			Columns: authrole.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: authrole.FieldID,
			},
		},
		From:   arq.sql,
		Unique: true,
	}
	if unique := arq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := arq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, authrole.FieldID)
		for i := range fields {
			if fields[i] != authrole.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := arq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := arq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := arq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := arq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (arq *AuthRoleQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(arq.driver.Dialect())
	t1 := builder.Table(authrole.Table)
	columns := arq.fields
	if len(columns) == 0 {
		columns = authrole.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if arq.sql != nil {
		selector = arq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, m := range arq.modifiers {
		m(selector)
	}
	t1.Schema(arq.schemaConfig.AuthRole)
	ctx = internal.NewSchemaConfigContext(ctx, arq.schemaConfig)
	selector.WithContext(ctx)
	for _, p := range arq.predicates {
		p(selector)
	}
	for _, p := range arq.order {
		p(selector)
	}
	if offset := arq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := arq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (arq *AuthRoleQuery) ForUpdate(opts ...sql.LockOption) *AuthRoleQuery {
	if arq.driver.Dialect() == dialect.Postgres {
		arq.Unique(false)
	}
	arq.modifiers = append(arq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return arq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (arq *AuthRoleQuery) ForShare(opts ...sql.LockOption) *AuthRoleQuery {
	if arq.driver.Dialect() == dialect.Postgres {
		arq.Unique(false)
	}
	arq.modifiers = append(arq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return arq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (arq *AuthRoleQuery) Modify(modifiers ...func(s *sql.Selector)) *AuthRoleSelect {
	arq.modifiers = append(arq.modifiers, modifiers...)
	return arq.Select()
}

// AuthRoleGroupBy is the group-by builder for AuthRole entities.
type AuthRoleGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (argb *AuthRoleGroupBy) Aggregate(fns ...AggregateFunc) *AuthRoleGroupBy {
	argb.fns = append(argb.fns, fns...)
	return argb
}

// Scan applies the group-by query and scans the result into the given value.
func (argb *AuthRoleGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := argb.path(ctx)
	if err != nil {
		return err
	}
	argb.sql = query
	return argb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (argb *AuthRoleGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := argb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (argb *AuthRoleGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(argb.fields) > 1 {
		return nil, errors.New("ent: AuthRoleGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := argb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (argb *AuthRoleGroupBy) StringsX(ctx context.Context) []string {
	v, err := argb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (argb *AuthRoleGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = argb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authrole.Label}
	default:
		err = fmt.Errorf("ent: AuthRoleGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (argb *AuthRoleGroupBy) StringX(ctx context.Context) string {
	v, err := argb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (argb *AuthRoleGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(argb.fields) > 1 {
		return nil, errors.New("ent: AuthRoleGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := argb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (argb *AuthRoleGroupBy) IntsX(ctx context.Context) []int {
	v, err := argb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (argb *AuthRoleGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = argb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authrole.Label}
	default:
		err = fmt.Errorf("ent: AuthRoleGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (argb *AuthRoleGroupBy) IntX(ctx context.Context) int {
	v, err := argb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (argb *AuthRoleGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(argb.fields) > 1 {
		return nil, errors.New("ent: AuthRoleGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := argb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (argb *AuthRoleGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := argb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (argb *AuthRoleGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = argb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authrole.Label}
	default:
		err = fmt.Errorf("ent: AuthRoleGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (argb *AuthRoleGroupBy) Float64X(ctx context.Context) float64 {
	v, err := argb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (argb *AuthRoleGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(argb.fields) > 1 {
		return nil, errors.New("ent: AuthRoleGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := argb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (argb *AuthRoleGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := argb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (argb *AuthRoleGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = argb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authrole.Label}
	default:
		err = fmt.Errorf("ent: AuthRoleGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (argb *AuthRoleGroupBy) BoolX(ctx context.Context) bool {
	v, err := argb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (argb *AuthRoleGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range argb.fields {
		if !authrole.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := argb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := argb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (argb *AuthRoleGroupBy) sqlQuery() *sql.Selector {
	selector := argb.sql.Select()
	aggregation := make([]string, 0, len(argb.fns))
	for _, fn := range argb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(argb.fields)+len(argb.fns))
		for _, f := range argb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(argb.fields...)...)
}

// AuthRoleSelect is the builder for selecting fields of AuthRole entities.
type AuthRoleSelect struct {
	*AuthRoleQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ars *AuthRoleSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ars.prepareQuery(ctx); err != nil {
		return err
	}
	ars.sql = ars.AuthRoleQuery.sqlQuery(ctx)
	return ars.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ars *AuthRoleSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ars.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ars *AuthRoleSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ars.fields) > 1 {
		return nil, errors.New("ent: AuthRoleSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ars.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ars *AuthRoleSelect) StringsX(ctx context.Context) []string {
	v, err := ars.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ars *AuthRoleSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ars.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authrole.Label}
	default:
		err = fmt.Errorf("ent: AuthRoleSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ars *AuthRoleSelect) StringX(ctx context.Context) string {
	v, err := ars.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ars *AuthRoleSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ars.fields) > 1 {
		return nil, errors.New("ent: AuthRoleSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ars.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ars *AuthRoleSelect) IntsX(ctx context.Context) []int {
	v, err := ars.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ars *AuthRoleSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ars.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authrole.Label}
	default:
		err = fmt.Errorf("ent: AuthRoleSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ars *AuthRoleSelect) IntX(ctx context.Context) int {
	v, err := ars.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ars *AuthRoleSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ars.fields) > 1 {
		return nil, errors.New("ent: AuthRoleSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ars.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ars *AuthRoleSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ars.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ars *AuthRoleSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ars.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authrole.Label}
	default:
		err = fmt.Errorf("ent: AuthRoleSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ars *AuthRoleSelect) Float64X(ctx context.Context) float64 {
	v, err := ars.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ars *AuthRoleSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ars.fields) > 1 {
		return nil, errors.New("ent: AuthRoleSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ars.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ars *AuthRoleSelect) BoolsX(ctx context.Context) []bool {
	v, err := ars.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ars *AuthRoleSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ars.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authrole.Label}
	default:
		err = fmt.Errorf("ent: AuthRoleSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ars *AuthRoleSelect) BoolX(ctx context.Context) bool {
	v, err := ars.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ars *AuthRoleSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ars.sql.Query()
	if err := ars.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ars *AuthRoleSelect) Modify(modifiers ...func(s *sql.Selector)) *AuthRoleSelect {
	ars.modifiers = append(ars.modifiers, modifiers...)
	return ars
}
