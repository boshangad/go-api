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
	"github.com/boshangad/v1/ent/app"
	"github.com/boshangad/v1/ent/emaillog"
	"github.com/boshangad/v1/ent/internal"
	"github.com/boshangad/v1/ent/predicate"
)

// EmailLogQuery is the builder for querying EmailLog entities.
type EmailLogQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.EmailLog
	// eager-loading edges.
	withApp   *AppQuery
	modifiers []func(s *sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EmailLogQuery builder.
func (elq *EmailLogQuery) Where(ps ...predicate.EmailLog) *EmailLogQuery {
	elq.predicates = append(elq.predicates, ps...)
	return elq
}

// Limit adds a limit step to the query.
func (elq *EmailLogQuery) Limit(limit int) *EmailLogQuery {
	elq.limit = &limit
	return elq
}

// Offset adds an offset step to the query.
func (elq *EmailLogQuery) Offset(offset int) *EmailLogQuery {
	elq.offset = &offset
	return elq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (elq *EmailLogQuery) Unique(unique bool) *EmailLogQuery {
	elq.unique = &unique
	return elq
}

// Order adds an order step to the query.
func (elq *EmailLogQuery) Order(o ...OrderFunc) *EmailLogQuery {
	elq.order = append(elq.order, o...)
	return elq
}

// QueryApp chains the current query on the "app" edge.
func (elq *EmailLogQuery) QueryApp() *AppQuery {
	query := &AppQuery{config: elq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := elq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := elq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(emaillog.Table, emaillog.FieldID, selector),
			sqlgraph.To(app.Table, app.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, emaillog.AppTable, emaillog.AppColumn),
		)
		schemaConfig := elq.schemaConfig
		step.To.Schema = schemaConfig.App
		step.Edge.Schema = schemaConfig.EmailLog
		fromU = sqlgraph.SetNeighbors(elq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first EmailLog entity from the query.
// Returns a *NotFoundError when no EmailLog was found.
func (elq *EmailLogQuery) First(ctx context.Context) (*EmailLog, error) {
	nodes, err := elq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{emaillog.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (elq *EmailLogQuery) FirstX(ctx context.Context) *EmailLog {
	node, err := elq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first EmailLog ID from the query.
// Returns a *NotFoundError when no EmailLog ID was found.
func (elq *EmailLogQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = elq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{emaillog.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (elq *EmailLogQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := elq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single EmailLog entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one EmailLog entity is not found.
// Returns a *NotFoundError when no EmailLog entities are found.
func (elq *EmailLogQuery) Only(ctx context.Context) (*EmailLog, error) {
	nodes, err := elq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{emaillog.Label}
	default:
		return nil, &NotSingularError{emaillog.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (elq *EmailLogQuery) OnlyX(ctx context.Context) *EmailLog {
	node, err := elq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only EmailLog ID in the query.
// Returns a *NotSingularError when exactly one EmailLog ID is not found.
// Returns a *NotFoundError when no entities are found.
func (elq *EmailLogQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = elq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{emaillog.Label}
	default:
		err = &NotSingularError{emaillog.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (elq *EmailLogQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := elq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of EmailLogs.
func (elq *EmailLogQuery) All(ctx context.Context) ([]*EmailLog, error) {
	if err := elq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return elq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (elq *EmailLogQuery) AllX(ctx context.Context) []*EmailLog {
	nodes, err := elq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of EmailLog IDs.
func (elq *EmailLogQuery) IDs(ctx context.Context) ([]uint64, error) {
	var ids []uint64
	if err := elq.Select(emaillog.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (elq *EmailLogQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := elq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (elq *EmailLogQuery) Count(ctx context.Context) (int, error) {
	if err := elq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return elq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (elq *EmailLogQuery) CountX(ctx context.Context) int {
	count, err := elq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (elq *EmailLogQuery) Exist(ctx context.Context) (bool, error) {
	if err := elq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return elq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (elq *EmailLogQuery) ExistX(ctx context.Context) bool {
	exist, err := elq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EmailLogQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (elq *EmailLogQuery) Clone() *EmailLogQuery {
	if elq == nil {
		return nil
	}
	return &EmailLogQuery{
		config:     elq.config,
		limit:      elq.limit,
		offset:     elq.offset,
		order:      append([]OrderFunc{}, elq.order...),
		predicates: append([]predicate.EmailLog{}, elq.predicates...),
		withApp:    elq.withApp.Clone(),
		// clone intermediate query.
		sql:  elq.sql.Clone(),
		path: elq.path,
	}
}

// WithApp tells the query-builder to eager-load the nodes that are connected to
// the "app" edge. The optional arguments are used to configure the query builder of the edge.
func (elq *EmailLogQuery) WithApp(opts ...func(*AppQuery)) *EmailLogQuery {
	query := &AppQuery{config: elq.config}
	for _, opt := range opts {
		opt(query)
	}
	elq.withApp = query
	return elq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime int64 `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.EmailLog.Query().
//		GroupBy(emaillog.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (elq *EmailLogQuery) GroupBy(field string, fields ...string) *EmailLogGroupBy {
	group := &EmailLogGroupBy{config: elq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := elq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return elq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime int64 `json:"create_time,omitempty"`
//	}
//
//	client.EmailLog.Query().
//		Select(emaillog.FieldCreateTime).
//		Scan(ctx, &v)
//
func (elq *EmailLogQuery) Select(fields ...string) *EmailLogSelect {
	elq.fields = append(elq.fields, fields...)
	return &EmailLogSelect{EmailLogQuery: elq}
}

func (elq *EmailLogQuery) prepareQuery(ctx context.Context) error {
	for _, f := range elq.fields {
		if !emaillog.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if elq.path != nil {
		prev, err := elq.path(ctx)
		if err != nil {
			return err
		}
		elq.sql = prev
	}
	return nil
}

func (elq *EmailLogQuery) sqlAll(ctx context.Context) ([]*EmailLog, error) {
	var (
		nodes       = []*EmailLog{}
		_spec       = elq.querySpec()
		loadedTypes = [1]bool{
			elq.withApp != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &EmailLog{config: elq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(elq.modifiers) > 0 {
		_spec.Modifiers = elq.modifiers
	}
	_spec.Node.Schema = elq.schemaConfig.EmailLog
	ctx = internal.NewSchemaConfigContext(ctx, elq.schemaConfig)
	if err := sqlgraph.QueryNodes(ctx, elq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := elq.withApp; query != nil {
		ids := make([]uint64, 0, len(nodes))
		nodeids := make(map[uint64][]*EmailLog)
		for i := range nodes {
			fk := nodes[i].AppID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(app.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "app_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.App = n
			}
		}
	}

	return nodes, nil
}

func (elq *EmailLogQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := elq.querySpec()
	if len(elq.modifiers) > 0 {
		_spec.Modifiers = elq.modifiers
	}
	_spec.Node.Schema = elq.schemaConfig.EmailLog
	ctx = internal.NewSchemaConfigContext(ctx, elq.schemaConfig)
	return sqlgraph.CountNodes(ctx, elq.driver, _spec)
}

func (elq *EmailLogQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := elq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (elq *EmailLogQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   emaillog.Table,
			Columns: emaillog.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: emaillog.FieldID,
			},
		},
		From:   elq.sql,
		Unique: true,
	}
	if unique := elq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := elq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, emaillog.FieldID)
		for i := range fields {
			if fields[i] != emaillog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := elq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := elq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := elq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := elq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (elq *EmailLogQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(elq.driver.Dialect())
	t1 := builder.Table(emaillog.Table)
	columns := elq.fields
	if len(columns) == 0 {
		columns = emaillog.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if elq.sql != nil {
		selector = elq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, m := range elq.modifiers {
		m(selector)
	}
	t1.Schema(elq.schemaConfig.EmailLog)
	ctx = internal.NewSchemaConfigContext(ctx, elq.schemaConfig)
	selector.WithContext(ctx)
	for _, p := range elq.predicates {
		p(selector)
	}
	for _, p := range elq.order {
		p(selector)
	}
	if offset := elq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := elq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (elq *EmailLogQuery) ForUpdate(opts ...sql.LockOption) *EmailLogQuery {
	if elq.driver.Dialect() == dialect.Postgres {
		elq.Unique(false)
	}
	elq.modifiers = append(elq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return elq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (elq *EmailLogQuery) ForShare(opts ...sql.LockOption) *EmailLogQuery {
	if elq.driver.Dialect() == dialect.Postgres {
		elq.Unique(false)
	}
	elq.modifiers = append(elq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return elq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (elq *EmailLogQuery) Modify(modifiers ...func(s *sql.Selector)) *EmailLogSelect {
	elq.modifiers = append(elq.modifiers, modifiers...)
	return elq.Select()
}

// EmailLogGroupBy is the group-by builder for EmailLog entities.
type EmailLogGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (elgb *EmailLogGroupBy) Aggregate(fns ...AggregateFunc) *EmailLogGroupBy {
	elgb.fns = append(elgb.fns, fns...)
	return elgb
}

// Scan applies the group-by query and scans the result into the given value.
func (elgb *EmailLogGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := elgb.path(ctx)
	if err != nil {
		return err
	}
	elgb.sql = query
	return elgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (elgb *EmailLogGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := elgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (elgb *EmailLogGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(elgb.fields) > 1 {
		return nil, errors.New("ent: EmailLogGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := elgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (elgb *EmailLogGroupBy) StringsX(ctx context.Context) []string {
	v, err := elgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (elgb *EmailLogGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = elgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{emaillog.Label}
	default:
		err = fmt.Errorf("ent: EmailLogGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (elgb *EmailLogGroupBy) StringX(ctx context.Context) string {
	v, err := elgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (elgb *EmailLogGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(elgb.fields) > 1 {
		return nil, errors.New("ent: EmailLogGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := elgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (elgb *EmailLogGroupBy) IntsX(ctx context.Context) []int {
	v, err := elgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (elgb *EmailLogGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = elgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{emaillog.Label}
	default:
		err = fmt.Errorf("ent: EmailLogGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (elgb *EmailLogGroupBy) IntX(ctx context.Context) int {
	v, err := elgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (elgb *EmailLogGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(elgb.fields) > 1 {
		return nil, errors.New("ent: EmailLogGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := elgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (elgb *EmailLogGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := elgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (elgb *EmailLogGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = elgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{emaillog.Label}
	default:
		err = fmt.Errorf("ent: EmailLogGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (elgb *EmailLogGroupBy) Float64X(ctx context.Context) float64 {
	v, err := elgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (elgb *EmailLogGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(elgb.fields) > 1 {
		return nil, errors.New("ent: EmailLogGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := elgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (elgb *EmailLogGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := elgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (elgb *EmailLogGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = elgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{emaillog.Label}
	default:
		err = fmt.Errorf("ent: EmailLogGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (elgb *EmailLogGroupBy) BoolX(ctx context.Context) bool {
	v, err := elgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (elgb *EmailLogGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range elgb.fields {
		if !emaillog.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := elgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := elgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (elgb *EmailLogGroupBy) sqlQuery() *sql.Selector {
	selector := elgb.sql.Select()
	aggregation := make([]string, 0, len(elgb.fns))
	for _, fn := range elgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(elgb.fields)+len(elgb.fns))
		for _, f := range elgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(elgb.fields...)...)
}

// EmailLogSelect is the builder for selecting fields of EmailLog entities.
type EmailLogSelect struct {
	*EmailLogQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (els *EmailLogSelect) Scan(ctx context.Context, v interface{}) error {
	if err := els.prepareQuery(ctx); err != nil {
		return err
	}
	els.sql = els.EmailLogQuery.sqlQuery(ctx)
	return els.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (els *EmailLogSelect) ScanX(ctx context.Context, v interface{}) {
	if err := els.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (els *EmailLogSelect) Strings(ctx context.Context) ([]string, error) {
	if len(els.fields) > 1 {
		return nil, errors.New("ent: EmailLogSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := els.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (els *EmailLogSelect) StringsX(ctx context.Context) []string {
	v, err := els.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (els *EmailLogSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = els.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{emaillog.Label}
	default:
		err = fmt.Errorf("ent: EmailLogSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (els *EmailLogSelect) StringX(ctx context.Context) string {
	v, err := els.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (els *EmailLogSelect) Ints(ctx context.Context) ([]int, error) {
	if len(els.fields) > 1 {
		return nil, errors.New("ent: EmailLogSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := els.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (els *EmailLogSelect) IntsX(ctx context.Context) []int {
	v, err := els.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (els *EmailLogSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = els.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{emaillog.Label}
	default:
		err = fmt.Errorf("ent: EmailLogSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (els *EmailLogSelect) IntX(ctx context.Context) int {
	v, err := els.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (els *EmailLogSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(els.fields) > 1 {
		return nil, errors.New("ent: EmailLogSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := els.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (els *EmailLogSelect) Float64sX(ctx context.Context) []float64 {
	v, err := els.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (els *EmailLogSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = els.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{emaillog.Label}
	default:
		err = fmt.Errorf("ent: EmailLogSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (els *EmailLogSelect) Float64X(ctx context.Context) float64 {
	v, err := els.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (els *EmailLogSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(els.fields) > 1 {
		return nil, errors.New("ent: EmailLogSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := els.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (els *EmailLogSelect) BoolsX(ctx context.Context) []bool {
	v, err := els.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (els *EmailLogSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = els.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{emaillog.Label}
	default:
		err = fmt.Errorf("ent: EmailLogSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (els *EmailLogSelect) BoolX(ctx context.Context) bool {
	v, err := els.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (els *EmailLogSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := els.sql.Query()
	if err := els.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (els *EmailLogSelect) Modify(modifiers ...func(s *sql.Selector)) *EmailLogSelect {
	els.modifiers = append(els.modifiers, modifiers...)
	return els
}
