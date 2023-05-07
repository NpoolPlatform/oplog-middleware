// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/NpoolPlatform/oplog-middleware/pkg/db/ent/oplog"
	"github.com/NpoolPlatform/oplog-middleware/pkg/db/ent/pubsubmessage"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 2)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   oplog.Table,
			Columns: oplog.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: oplog.FieldID,
			},
		},
		Type: "OpLog",
		Fields: map[string]*sqlgraph.FieldSpec{
			oplog.FieldCreatedAt:        {Type: field.TypeUint32, Column: oplog.FieldCreatedAt},
			oplog.FieldUpdatedAt:        {Type: field.TypeUint32, Column: oplog.FieldUpdatedAt},
			oplog.FieldDeletedAt:        {Type: field.TypeUint32, Column: oplog.FieldDeletedAt},
			oplog.FieldAutoID:           {Type: field.TypeUint32, Column: oplog.FieldAutoID},
			oplog.FieldAppID:            {Type: field.TypeUUID, Column: oplog.FieldAppID},
			oplog.FieldUserID:           {Type: field.TypeUUID, Column: oplog.FieldUserID},
			oplog.FieldMethod:           {Type: field.TypeString, Column: oplog.FieldMethod},
			oplog.FieldArguments:        {Type: field.TypeString, Column: oplog.FieldArguments},
			oplog.FieldCurValue:         {Type: field.TypeString, Column: oplog.FieldCurValue},
			oplog.FieldHumanReadable:    {Type: field.TypeString, Column: oplog.FieldHumanReadable},
			oplog.FieldResult:           {Type: field.TypeString, Column: oplog.FieldResult},
			oplog.FieldFailReason:       {Type: field.TypeString, Column: oplog.FieldFailReason},
			oplog.FieldElapsedMillisecs: {Type: field.TypeUint32, Column: oplog.FieldElapsedMillisecs},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   pubsubmessage.Table,
			Columns: pubsubmessage.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: pubsubmessage.FieldID,
			},
		},
		Type: "PubsubMessage",
		Fields: map[string]*sqlgraph.FieldSpec{
			pubsubmessage.FieldCreatedAt: {Type: field.TypeUint32, Column: pubsubmessage.FieldCreatedAt},
			pubsubmessage.FieldUpdatedAt: {Type: field.TypeUint32, Column: pubsubmessage.FieldUpdatedAt},
			pubsubmessage.FieldDeletedAt: {Type: field.TypeUint32, Column: pubsubmessage.FieldDeletedAt},
			pubsubmessage.FieldAutoID:    {Type: field.TypeUint32, Column: pubsubmessage.FieldAutoID},
			pubsubmessage.FieldMessageID: {Type: field.TypeString, Column: pubsubmessage.FieldMessageID},
			pubsubmessage.FieldState:     {Type: field.TypeString, Column: pubsubmessage.FieldState},
			pubsubmessage.FieldRespToID:  {Type: field.TypeUUID, Column: pubsubmessage.FieldRespToID},
			pubsubmessage.FieldUndoID:    {Type: field.TypeUUID, Column: pubsubmessage.FieldUndoID},
			pubsubmessage.FieldArguments: {Type: field.TypeString, Column: pubsubmessage.FieldArguments},
		},
	}
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (olq *OpLogQuery) addPredicate(pred func(s *sql.Selector)) {
	olq.predicates = append(olq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the OpLogQuery builder.
func (olq *OpLogQuery) Filter() *OpLogFilter {
	return &OpLogFilter{config: olq.config, predicateAdder: olq}
}

// addPredicate implements the predicateAdder interface.
func (m *OpLogMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the OpLogMutation builder.
func (m *OpLogMutation) Filter() *OpLogFilter {
	return &OpLogFilter{config: m.config, predicateAdder: m}
}

// OpLogFilter provides a generic filtering capability at runtime for OpLogQuery.
type OpLogFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *OpLogFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *OpLogFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(oplog.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *OpLogFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(oplog.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *OpLogFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(oplog.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *OpLogFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(oplog.FieldDeletedAt))
}

// WhereAutoID applies the entql uint32 predicate on the auto_id field.
func (f *OpLogFilter) WhereAutoID(p entql.Uint32P) {
	f.Where(p.Field(oplog.FieldAutoID))
}

// WhereAppID applies the entql [16]byte predicate on the app_id field.
func (f *OpLogFilter) WhereAppID(p entql.ValueP) {
	f.Where(p.Field(oplog.FieldAppID))
}

// WhereUserID applies the entql [16]byte predicate on the user_id field.
func (f *OpLogFilter) WhereUserID(p entql.ValueP) {
	f.Where(p.Field(oplog.FieldUserID))
}

// WhereMethod applies the entql string predicate on the method field.
func (f *OpLogFilter) WhereMethod(p entql.StringP) {
	f.Where(p.Field(oplog.FieldMethod))
}

// WhereArguments applies the entql string predicate on the arguments field.
func (f *OpLogFilter) WhereArguments(p entql.StringP) {
	f.Where(p.Field(oplog.FieldArguments))
}

// WhereCurValue applies the entql string predicate on the cur_value field.
func (f *OpLogFilter) WhereCurValue(p entql.StringP) {
	f.Where(p.Field(oplog.FieldCurValue))
}

// WhereHumanReadable applies the entql string predicate on the human_readable field.
func (f *OpLogFilter) WhereHumanReadable(p entql.StringP) {
	f.Where(p.Field(oplog.FieldHumanReadable))
}

// WhereResult applies the entql string predicate on the result field.
func (f *OpLogFilter) WhereResult(p entql.StringP) {
	f.Where(p.Field(oplog.FieldResult))
}

// WhereFailReason applies the entql string predicate on the fail_reason field.
func (f *OpLogFilter) WhereFailReason(p entql.StringP) {
	f.Where(p.Field(oplog.FieldFailReason))
}

// WhereElapsedMillisecs applies the entql uint32 predicate on the elapsed_millisecs field.
func (f *OpLogFilter) WhereElapsedMillisecs(p entql.Uint32P) {
	f.Where(p.Field(oplog.FieldElapsedMillisecs))
}

// addPredicate implements the predicateAdder interface.
func (pmq *PubsubMessageQuery) addPredicate(pred func(s *sql.Selector)) {
	pmq.predicates = append(pmq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the PubsubMessageQuery builder.
func (pmq *PubsubMessageQuery) Filter() *PubsubMessageFilter {
	return &PubsubMessageFilter{config: pmq.config, predicateAdder: pmq}
}

// addPredicate implements the predicateAdder interface.
func (m *PubsubMessageMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the PubsubMessageMutation builder.
func (m *PubsubMessageMutation) Filter() *PubsubMessageFilter {
	return &PubsubMessageFilter{config: m.config, predicateAdder: m}
}

// PubsubMessageFilter provides a generic filtering capability at runtime for PubsubMessageQuery.
type PubsubMessageFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *PubsubMessageFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *PubsubMessageFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(pubsubmessage.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *PubsubMessageFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(pubsubmessage.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *PubsubMessageFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(pubsubmessage.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *PubsubMessageFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(pubsubmessage.FieldDeletedAt))
}

// WhereAutoID applies the entql uint32 predicate on the auto_id field.
func (f *PubsubMessageFilter) WhereAutoID(p entql.Uint32P) {
	f.Where(p.Field(pubsubmessage.FieldAutoID))
}

// WhereMessageID applies the entql string predicate on the message_id field.
func (f *PubsubMessageFilter) WhereMessageID(p entql.StringP) {
	f.Where(p.Field(pubsubmessage.FieldMessageID))
}

// WhereState applies the entql string predicate on the state field.
func (f *PubsubMessageFilter) WhereState(p entql.StringP) {
	f.Where(p.Field(pubsubmessage.FieldState))
}

// WhereRespToID applies the entql [16]byte predicate on the resp_to_id field.
func (f *PubsubMessageFilter) WhereRespToID(p entql.ValueP) {
	f.Where(p.Field(pubsubmessage.FieldRespToID))
}

// WhereUndoID applies the entql [16]byte predicate on the undo_id field.
func (f *PubsubMessageFilter) WhereUndoID(p entql.ValueP) {
	f.Where(p.Field(pubsubmessage.FieldUndoID))
}

// WhereArguments applies the entql string predicate on the arguments field.
func (f *PubsubMessageFilter) WhereArguments(p entql.StringP) {
	f.Where(p.Field(pubsubmessage.FieldArguments))
}
