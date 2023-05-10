package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	"github.com/NpoolPlatform/oplog-middleware/pkg/db/mixin"
	"github.com/google/uuid"
)

// OpLog holds the schema definition for the OpLog entity.
type OpLog struct {
	ent.Schema
}

func (OpLog) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the OpLog.
func (OpLog) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			UUID("user_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			String("path").
			Optional().
			Default(""),
		field.
			String("method").
			Optional().
			Default(basetypes.HTTPMethod_DefaultMethod.String()),
		field.
			Text("arguments").
			Optional().
			Default(""),
		field.
			Text("cur_value").
			Optional().
			Default(""),
		field.
			Text("new_value").
			Optional().
			Default(""),
		field.
			Text("human_readable").
			Optional().
			Default(""),
		field.
			String("result").
			Optional().
			Default(basetypes.Result_DefaultResult.String()),
		field.
			String("fail_reason").
			Optional().
			Default(""),
		field.
			Uint32("elapsed_millisecs").
			Optional().
			Default(0),
		field.
			Int32("status_code").
			Optional().
			Default(0),
		field.
			Text("req_headers").
			Optional().
			Default("{}"),
		field.
			Text("resp_headers").
			Optional().
			Default("{}"),
	}
}

// Edges of the OpLog.
func (OpLog) Edges() []ent.Edge {
	return nil
}
