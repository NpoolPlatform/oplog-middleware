// Code generated by ent, DO NOT EDIT.

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
	"github.com/NpoolPlatform/oplog-middleware/pkg/db/ent/oplog"
	"github.com/NpoolPlatform/oplog-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// OpLogQuery is the builder for querying OpLog entities.
type OpLogQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.OpLog
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OpLogQuery builder.
func (olq *OpLogQuery) Where(ps ...predicate.OpLog) *OpLogQuery {
	olq.predicates = append(olq.predicates, ps...)
	return olq
}

// Limit adds a limit step to the query.
func (olq *OpLogQuery) Limit(limit int) *OpLogQuery {
	olq.limit = &limit
	return olq
}

// Offset adds an offset step to the query.
func (olq *OpLogQuery) Offset(offset int) *OpLogQuery {
	olq.offset = &offset
	return olq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (olq *OpLogQuery) Unique(unique bool) *OpLogQuery {
	olq.unique = &unique
	return olq
}

// Order adds an order step to the query.
func (olq *OpLogQuery) Order(o ...OrderFunc) *OpLogQuery {
	olq.order = append(olq.order, o...)
	return olq
}

// First returns the first OpLog entity from the query.
// Returns a *NotFoundError when no OpLog was found.
func (olq *OpLogQuery) First(ctx context.Context) (*OpLog, error) {
	nodes, err := olq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{oplog.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (olq *OpLogQuery) FirstX(ctx context.Context) *OpLog {
	node, err := olq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OpLog ID from the query.
// Returns a *NotFoundError when no OpLog ID was found.
func (olq *OpLogQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = olq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{oplog.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (olq *OpLogQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := olq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OpLog entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OpLog entity is found.
// Returns a *NotFoundError when no OpLog entities are found.
func (olq *OpLogQuery) Only(ctx context.Context) (*OpLog, error) {
	nodes, err := olq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{oplog.Label}
	default:
		return nil, &NotSingularError{oplog.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (olq *OpLogQuery) OnlyX(ctx context.Context) *OpLog {
	node, err := olq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OpLog ID in the query.
// Returns a *NotSingularError when more than one OpLog ID is found.
// Returns a *NotFoundError when no entities are found.
func (olq *OpLogQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = olq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{oplog.Label}
	default:
		err = &NotSingularError{oplog.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (olq *OpLogQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := olq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OpLogs.
func (olq *OpLogQuery) All(ctx context.Context) ([]*OpLog, error) {
	if err := olq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return olq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (olq *OpLogQuery) AllX(ctx context.Context) []*OpLog {
	nodes, err := olq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OpLog IDs.
func (olq *OpLogQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := olq.Select(oplog.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (olq *OpLogQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := olq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (olq *OpLogQuery) Count(ctx context.Context) (int, error) {
	if err := olq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return olq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (olq *OpLogQuery) CountX(ctx context.Context) int {
	count, err := olq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (olq *OpLogQuery) Exist(ctx context.Context) (bool, error) {
	if err := olq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return olq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (olq *OpLogQuery) ExistX(ctx context.Context) bool {
	exist, err := olq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OpLogQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (olq *OpLogQuery) Clone() *OpLogQuery {
	if olq == nil {
		return nil
	}
	return &OpLogQuery{
		config:     olq.config,
		limit:      olq.limit,
		offset:     olq.offset,
		order:      append([]OrderFunc{}, olq.order...),
		predicates: append([]predicate.OpLog{}, olq.predicates...),
		// clone intermediate query.
		sql:    olq.sql.Clone(),
		path:   olq.path,
		unique: olq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OpLog.Query().
//		GroupBy(oplog.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (olq *OpLogQuery) GroupBy(field string, fields ...string) *OpLogGroupBy {
	grbuild := &OpLogGroupBy{config: olq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := olq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return olq.sqlQuery(ctx), nil
	}
	grbuild.label = oplog.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//	}
//
//	client.OpLog.Query().
//		Select(oplog.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (olq *OpLogQuery) Select(fields ...string) *OpLogSelect {
	olq.fields = append(olq.fields, fields...)
	selbuild := &OpLogSelect{OpLogQuery: olq}
	selbuild.label = oplog.Label
	selbuild.flds, selbuild.scan = &olq.fields, selbuild.Scan
	return selbuild
}

func (olq *OpLogQuery) prepareQuery(ctx context.Context) error {
	for _, f := range olq.fields {
		if !oplog.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if olq.path != nil {
		prev, err := olq.path(ctx)
		if err != nil {
			return err
		}
		olq.sql = prev
	}
	if oplog.Policy == nil {
		return errors.New("ent: uninitialized oplog.Policy (forgotten import ent/runtime?)")
	}
	if err := oplog.Policy.EvalQuery(ctx, olq); err != nil {
		return err
	}
	return nil
}

func (olq *OpLogQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OpLog, error) {
	var (
		nodes = []*OpLog{}
		_spec = olq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*OpLog).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &OpLog{config: olq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(olq.modifiers) > 0 {
		_spec.Modifiers = olq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, olq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (olq *OpLogQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := olq.querySpec()
	if len(olq.modifiers) > 0 {
		_spec.Modifiers = olq.modifiers
	}
	_spec.Node.Columns = olq.fields
	if len(olq.fields) > 0 {
		_spec.Unique = olq.unique != nil && *olq.unique
	}
	return sqlgraph.CountNodes(ctx, olq.driver, _spec)
}

func (olq *OpLogQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := olq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (olq *OpLogQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   oplog.Table,
			Columns: oplog.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: oplog.FieldID,
			},
		},
		From:   olq.sql,
		Unique: true,
	}
	if unique := olq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := olq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, oplog.FieldID)
		for i := range fields {
			if fields[i] != oplog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := olq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := olq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := olq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := olq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (olq *OpLogQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(olq.driver.Dialect())
	t1 := builder.Table(oplog.Table)
	columns := olq.fields
	if len(columns) == 0 {
		columns = oplog.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if olq.sql != nil {
		selector = olq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if olq.unique != nil && *olq.unique {
		selector.Distinct()
	}
	for _, m := range olq.modifiers {
		m(selector)
	}
	for _, p := range olq.predicates {
		p(selector)
	}
	for _, p := range olq.order {
		p(selector)
	}
	if offset := olq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := olq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (olq *OpLogQuery) ForUpdate(opts ...sql.LockOption) *OpLogQuery {
	if olq.driver.Dialect() == dialect.Postgres {
		olq.Unique(false)
	}
	olq.modifiers = append(olq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return olq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (olq *OpLogQuery) ForShare(opts ...sql.LockOption) *OpLogQuery {
	if olq.driver.Dialect() == dialect.Postgres {
		olq.Unique(false)
	}
	olq.modifiers = append(olq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return olq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (olq *OpLogQuery) Modify(modifiers ...func(s *sql.Selector)) *OpLogSelect {
	olq.modifiers = append(olq.modifiers, modifiers...)
	return olq.Select()
}

// OpLogGroupBy is the group-by builder for OpLog entities.
type OpLogGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (olgb *OpLogGroupBy) Aggregate(fns ...AggregateFunc) *OpLogGroupBy {
	olgb.fns = append(olgb.fns, fns...)
	return olgb
}

// Scan applies the group-by query and scans the result into the given value.
func (olgb *OpLogGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := olgb.path(ctx)
	if err != nil {
		return err
	}
	olgb.sql = query
	return olgb.sqlScan(ctx, v)
}

func (olgb *OpLogGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range olgb.fields {
		if !oplog.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := olgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := olgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (olgb *OpLogGroupBy) sqlQuery() *sql.Selector {
	selector := olgb.sql.Select()
	aggregation := make([]string, 0, len(olgb.fns))
	for _, fn := range olgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(olgb.fields)+len(olgb.fns))
		for _, f := range olgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(olgb.fields...)...)
}

// OpLogSelect is the builder for selecting fields of OpLog entities.
type OpLogSelect struct {
	*OpLogQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ols *OpLogSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ols.prepareQuery(ctx); err != nil {
		return err
	}
	ols.sql = ols.OpLogQuery.sqlQuery(ctx)
	return ols.sqlScan(ctx, v)
}

func (ols *OpLogSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ols.sql.Query()
	if err := ols.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ols *OpLogSelect) Modify(modifiers ...func(s *sql.Selector)) *OpLogSelect {
	ols.modifiers = append(ols.modifiers, modifiers...)
	return ols
}
