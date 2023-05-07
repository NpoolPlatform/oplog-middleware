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
	"github.com/NpoolPlatform/oplog-middleware/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/oplog-middleware/pkg/db/ent/pubsubmessage"
)

// PubsubMessageQuery is the builder for querying PubsubMessage entities.
type PubsubMessageQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.PubsubMessage
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PubsubMessageQuery builder.
func (pmq *PubsubMessageQuery) Where(ps ...predicate.PubsubMessage) *PubsubMessageQuery {
	pmq.predicates = append(pmq.predicates, ps...)
	return pmq
}

// Limit adds a limit step to the query.
func (pmq *PubsubMessageQuery) Limit(limit int) *PubsubMessageQuery {
	pmq.limit = &limit
	return pmq
}

// Offset adds an offset step to the query.
func (pmq *PubsubMessageQuery) Offset(offset int) *PubsubMessageQuery {
	pmq.offset = &offset
	return pmq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pmq *PubsubMessageQuery) Unique(unique bool) *PubsubMessageQuery {
	pmq.unique = &unique
	return pmq
}

// Order adds an order step to the query.
func (pmq *PubsubMessageQuery) Order(o ...OrderFunc) *PubsubMessageQuery {
	pmq.order = append(pmq.order, o...)
	return pmq
}

// First returns the first PubsubMessage entity from the query.
// Returns a *NotFoundError when no PubsubMessage was found.
func (pmq *PubsubMessageQuery) First(ctx context.Context) (*PubsubMessage, error) {
	nodes, err := pmq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{pubsubmessage.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pmq *PubsubMessageQuery) FirstX(ctx context.Context) *PubsubMessage {
	node, err := pmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PubsubMessage ID from the query.
// Returns a *NotFoundError when no PubsubMessage ID was found.
func (pmq *PubsubMessageQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = pmq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{pubsubmessage.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pmq *PubsubMessageQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := pmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PubsubMessage entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PubsubMessage entity is found.
// Returns a *NotFoundError when no PubsubMessage entities are found.
func (pmq *PubsubMessageQuery) Only(ctx context.Context) (*PubsubMessage, error) {
	nodes, err := pmq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{pubsubmessage.Label}
	default:
		return nil, &NotSingularError{pubsubmessage.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pmq *PubsubMessageQuery) OnlyX(ctx context.Context) *PubsubMessage {
	node, err := pmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PubsubMessage ID in the query.
// Returns a *NotSingularError when more than one PubsubMessage ID is found.
// Returns a *NotFoundError when no entities are found.
func (pmq *PubsubMessageQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = pmq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{pubsubmessage.Label}
	default:
		err = &NotSingularError{pubsubmessage.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pmq *PubsubMessageQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := pmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PubsubMessages.
func (pmq *PubsubMessageQuery) All(ctx context.Context) ([]*PubsubMessage, error) {
	if err := pmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return pmq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (pmq *PubsubMessageQuery) AllX(ctx context.Context) []*PubsubMessage {
	nodes, err := pmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PubsubMessage IDs.
func (pmq *PubsubMessageQuery) IDs(ctx context.Context) ([]uint32, error) {
	var ids []uint32
	if err := pmq.Select(pubsubmessage.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pmq *PubsubMessageQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := pmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pmq *PubsubMessageQuery) Count(ctx context.Context) (int, error) {
	if err := pmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return pmq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (pmq *PubsubMessageQuery) CountX(ctx context.Context) int {
	count, err := pmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pmq *PubsubMessageQuery) Exist(ctx context.Context) (bool, error) {
	if err := pmq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return pmq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (pmq *PubsubMessageQuery) ExistX(ctx context.Context) bool {
	exist, err := pmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PubsubMessageQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pmq *PubsubMessageQuery) Clone() *PubsubMessageQuery {
	if pmq == nil {
		return nil
	}
	return &PubsubMessageQuery{
		config:     pmq.config,
		limit:      pmq.limit,
		offset:     pmq.offset,
		order:      append([]OrderFunc{}, pmq.order...),
		predicates: append([]predicate.PubsubMessage{}, pmq.predicates...),
		// clone intermediate query.
		sql:    pmq.sql.Clone(),
		path:   pmq.path,
		unique: pmq.unique,
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
//	client.PubsubMessage.Query().
//		GroupBy(pubsubmessage.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (pmq *PubsubMessageQuery) GroupBy(field string, fields ...string) *PubsubMessageGroupBy {
	grbuild := &PubsubMessageGroupBy{config: pmq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := pmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return pmq.sqlQuery(ctx), nil
	}
	grbuild.label = pubsubmessage.Label
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
//	client.PubsubMessage.Query().
//		Select(pubsubmessage.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (pmq *PubsubMessageQuery) Select(fields ...string) *PubsubMessageSelect {
	pmq.fields = append(pmq.fields, fields...)
	selbuild := &PubsubMessageSelect{PubsubMessageQuery: pmq}
	selbuild.label = pubsubmessage.Label
	selbuild.flds, selbuild.scan = &pmq.fields, selbuild.Scan
	return selbuild
}

func (pmq *PubsubMessageQuery) prepareQuery(ctx context.Context) error {
	for _, f := range pmq.fields {
		if !pubsubmessage.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pmq.path != nil {
		prev, err := pmq.path(ctx)
		if err != nil {
			return err
		}
		pmq.sql = prev
	}
	if pubsubmessage.Policy == nil {
		return errors.New("ent: uninitialized pubsubmessage.Policy (forgotten import ent/runtime?)")
	}
	if err := pubsubmessage.Policy.EvalQuery(ctx, pmq); err != nil {
		return err
	}
	return nil
}

func (pmq *PubsubMessageQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PubsubMessage, error) {
	var (
		nodes = []*PubsubMessage{}
		_spec = pmq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*PubsubMessage).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &PubsubMessage{config: pmq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(pmq.modifiers) > 0 {
		_spec.Modifiers = pmq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (pmq *PubsubMessageQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pmq.querySpec()
	if len(pmq.modifiers) > 0 {
		_spec.Modifiers = pmq.modifiers
	}
	_spec.Node.Columns = pmq.fields
	if len(pmq.fields) > 0 {
		_spec.Unique = pmq.unique != nil && *pmq.unique
	}
	return sqlgraph.CountNodes(ctx, pmq.driver, _spec)
}

func (pmq *PubsubMessageQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := pmq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (pmq *PubsubMessageQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   pubsubmessage.Table,
			Columns: pubsubmessage.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: pubsubmessage.FieldID,
			},
		},
		From:   pmq.sql,
		Unique: true,
	}
	if unique := pmq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := pmq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pubsubmessage.FieldID)
		for i := range fields {
			if fields[i] != pubsubmessage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pmq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pmq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pmq *PubsubMessageQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pmq.driver.Dialect())
	t1 := builder.Table(pubsubmessage.Table)
	columns := pmq.fields
	if len(columns) == 0 {
		columns = pubsubmessage.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pmq.sql != nil {
		selector = pmq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pmq.unique != nil && *pmq.unique {
		selector.Distinct()
	}
	for _, m := range pmq.modifiers {
		m(selector)
	}
	for _, p := range pmq.predicates {
		p(selector)
	}
	for _, p := range pmq.order {
		p(selector)
	}
	if offset := pmq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pmq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (pmq *PubsubMessageQuery) ForUpdate(opts ...sql.LockOption) *PubsubMessageQuery {
	if pmq.driver.Dialect() == dialect.Postgres {
		pmq.Unique(false)
	}
	pmq.modifiers = append(pmq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return pmq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (pmq *PubsubMessageQuery) ForShare(opts ...sql.LockOption) *PubsubMessageQuery {
	if pmq.driver.Dialect() == dialect.Postgres {
		pmq.Unique(false)
	}
	pmq.modifiers = append(pmq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return pmq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (pmq *PubsubMessageQuery) Modify(modifiers ...func(s *sql.Selector)) *PubsubMessageSelect {
	pmq.modifiers = append(pmq.modifiers, modifiers...)
	return pmq.Select()
}

// PubsubMessageGroupBy is the group-by builder for PubsubMessage entities.
type PubsubMessageGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pmgb *PubsubMessageGroupBy) Aggregate(fns ...AggregateFunc) *PubsubMessageGroupBy {
	pmgb.fns = append(pmgb.fns, fns...)
	return pmgb
}

// Scan applies the group-by query and scans the result into the given value.
func (pmgb *PubsubMessageGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := pmgb.path(ctx)
	if err != nil {
		return err
	}
	pmgb.sql = query
	return pmgb.sqlScan(ctx, v)
}

func (pmgb *PubsubMessageGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range pmgb.fields {
		if !pubsubmessage.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := pmgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pmgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pmgb *PubsubMessageGroupBy) sqlQuery() *sql.Selector {
	selector := pmgb.sql.Select()
	aggregation := make([]string, 0, len(pmgb.fns))
	for _, fn := range pmgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(pmgb.fields)+len(pmgb.fns))
		for _, f := range pmgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(pmgb.fields...)...)
}

// PubsubMessageSelect is the builder for selecting fields of PubsubMessage entities.
type PubsubMessageSelect struct {
	*PubsubMessageQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (pms *PubsubMessageSelect) Scan(ctx context.Context, v interface{}) error {
	if err := pms.prepareQuery(ctx); err != nil {
		return err
	}
	pms.sql = pms.PubsubMessageQuery.sqlQuery(ctx)
	return pms.sqlScan(ctx, v)
}

func (pms *PubsubMessageSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := pms.sql.Query()
	if err := pms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (pms *PubsubMessageSelect) Modify(modifiers ...func(s *sql.Selector)) *PubsubMessageSelect {
	pms.modifiers = append(pms.modifiers, modifiers...)
	return pms
}
