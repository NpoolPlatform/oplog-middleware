// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// OpLogsColumns holds the columns for the "op_logs" table.
	OpLogsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "path", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "method", Type: field.TypeString, Nullable: true, Default: "DefaultMethod"},
		{Name: "arguments", Type: field.TypeString, Nullable: true, Size: 2147483647, Default: ""},
		{Name: "cur_value", Type: field.TypeString, Nullable: true, Size: 2147483647, Default: ""},
		{Name: "new_value", Type: field.TypeString, Nullable: true, Size: 2147483647, Default: ""},
		{Name: "human_readable", Type: field.TypeString, Nullable: true, Size: 2147483647, Default: ""},
		{Name: "result", Type: field.TypeString, Nullable: true, Default: "DefaultResult"},
		{Name: "fail_reason", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "elapsed_millisecs", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "status_code", Type: field.TypeInt32, Nullable: true, Default: 0},
		{Name: "req_headers", Type: field.TypeString, Nullable: true, Size: 2147483647, Default: "{}"},
		{Name: "resp_headers", Type: field.TypeString, Nullable: true, Size: 2147483647, Default: "{}"},
	}
	// OpLogsTable holds the schema information for the "op_logs" table.
	OpLogsTable = &schema.Table{
		Name:       "op_logs",
		Columns:    OpLogsColumns,
		PrimaryKey: []*schema.Column{OpLogsColumns[0]},
	}
	// PubsubMessagesColumns holds the columns for the "pubsub_messages" table.
	PubsubMessagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "message_id", Type: field.TypeString, Nullable: true, Default: "DefaultMsgID"},
		{Name: "state", Type: field.TypeString, Nullable: true, Default: "DefaultMsgState"},
		{Name: "resp_to_id", Type: field.TypeUUID, Nullable: true},
		{Name: "undo_id", Type: field.TypeUUID, Nullable: true},
		{Name: "arguments", Type: field.TypeString, Nullable: true, Default: ""},
	}
	// PubsubMessagesTable holds the schema information for the "pubsub_messages" table.
	PubsubMessagesTable = &schema.Table{
		Name:       "pubsub_messages",
		Columns:    PubsubMessagesColumns,
		PrimaryKey: []*schema.Column{PubsubMessagesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "pubsubmessage_state_resp_to_id",
				Unique:  false,
				Columns: []*schema.Column{PubsubMessagesColumns[6], PubsubMessagesColumns[7]},
			},
			{
				Name:    "pubsubmessage_state_undo_id",
				Unique:  false,
				Columns: []*schema.Column{PubsubMessagesColumns[6], PubsubMessagesColumns[8]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		OpLogsTable,
		PubsubMessagesTable,
	}
)

func init() {
}
