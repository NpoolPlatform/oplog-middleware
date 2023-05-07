// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/oplog-middleware/pkg/db/ent/oplog"
	"github.com/google/uuid"
)

// OpLog is the model entity for the OpLog schema.
type OpLog struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
	// AutoID holds the value of the "auto_id" field.
	AutoID uint32 `json:"auto_id,omitempty"`
	// AppID holds the value of the "app_id" field.
	AppID uuid.UUID `json:"app_id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// Path holds the value of the "path" field.
	Path string `json:"path,omitempty"`
	// Method holds the value of the "method" field.
	Method string `json:"method,omitempty"`
	// Arguments holds the value of the "arguments" field.
	Arguments string `json:"arguments,omitempty"`
	// CurValue holds the value of the "cur_value" field.
	CurValue string `json:"cur_value,omitempty"`
	// HumanReadable holds the value of the "human_readable" field.
	HumanReadable string `json:"human_readable,omitempty"`
	// Result holds the value of the "result" field.
	Result string `json:"result,omitempty"`
	// FailReason holds the value of the "fail_reason" field.
	FailReason string `json:"fail_reason,omitempty"`
	// ElapsedMillisecs holds the value of the "elapsed_millisecs" field.
	ElapsedMillisecs uint32 `json:"elapsed_millisecs,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OpLog) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case oplog.FieldCreatedAt, oplog.FieldUpdatedAt, oplog.FieldDeletedAt, oplog.FieldAutoID, oplog.FieldElapsedMillisecs:
			values[i] = new(sql.NullInt64)
		case oplog.FieldPath, oplog.FieldMethod, oplog.FieldArguments, oplog.FieldCurValue, oplog.FieldHumanReadable, oplog.FieldResult, oplog.FieldFailReason:
			values[i] = new(sql.NullString)
		case oplog.FieldID, oplog.FieldAppID, oplog.FieldUserID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type OpLog", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OpLog fields.
func (ol *OpLog) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case oplog.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ol.ID = *value
			}
		case oplog.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ol.CreatedAt = uint32(value.Int64)
			}
		case oplog.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ol.UpdatedAt = uint32(value.Int64)
			}
		case oplog.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ol.DeletedAt = uint32(value.Int64)
			}
		case oplog.FieldAutoID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field auto_id", values[i])
			} else if value.Valid {
				ol.AutoID = uint32(value.Int64)
			}
		case oplog.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				ol.AppID = *value
			}
		case oplog.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				ol.UserID = *value
			}
		case oplog.FieldPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field path", values[i])
			} else if value.Valid {
				ol.Path = value.String
			}
		case oplog.FieldMethod:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field method", values[i])
			} else if value.Valid {
				ol.Method = value.String
			}
		case oplog.FieldArguments:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field arguments", values[i])
			} else if value.Valid {
				ol.Arguments = value.String
			}
		case oplog.FieldCurValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cur_value", values[i])
			} else if value.Valid {
				ol.CurValue = value.String
			}
		case oplog.FieldHumanReadable:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field human_readable", values[i])
			} else if value.Valid {
				ol.HumanReadable = value.String
			}
		case oplog.FieldResult:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field result", values[i])
			} else if value.Valid {
				ol.Result = value.String
			}
		case oplog.FieldFailReason:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field fail_reason", values[i])
			} else if value.Valid {
				ol.FailReason = value.String
			}
		case oplog.FieldElapsedMillisecs:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field elapsed_millisecs", values[i])
			} else if value.Valid {
				ol.ElapsedMillisecs = uint32(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this OpLog.
// Note that you need to call OpLog.Unwrap() before calling this method if this OpLog
// was returned from a transaction, and the transaction was committed or rolled back.
func (ol *OpLog) Update() *OpLogUpdateOne {
	return (&OpLogClient{config: ol.config}).UpdateOne(ol)
}

// Unwrap unwraps the OpLog entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ol *OpLog) Unwrap() *OpLog {
	_tx, ok := ol.config.driver.(*txDriver)
	if !ok {
		panic("ent: OpLog is not a transactional entity")
	}
	ol.config.driver = _tx.drv
	return ol
}

// String implements the fmt.Stringer.
func (ol *OpLog) String() string {
	var builder strings.Builder
	builder.WriteString("OpLog(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ol.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", ol.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", ol.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", ol.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("auto_id=")
	builder.WriteString(fmt.Sprintf("%v", ol.AutoID))
	builder.WriteString(", ")
	builder.WriteString("app_id=")
	builder.WriteString(fmt.Sprintf("%v", ol.AppID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", ol.UserID))
	builder.WriteString(", ")
	builder.WriteString("path=")
	builder.WriteString(ol.Path)
	builder.WriteString(", ")
	builder.WriteString("method=")
	builder.WriteString(ol.Method)
	builder.WriteString(", ")
	builder.WriteString("arguments=")
	builder.WriteString(ol.Arguments)
	builder.WriteString(", ")
	builder.WriteString("cur_value=")
	builder.WriteString(ol.CurValue)
	builder.WriteString(", ")
	builder.WriteString("human_readable=")
	builder.WriteString(ol.HumanReadable)
	builder.WriteString(", ")
	builder.WriteString("result=")
	builder.WriteString(ol.Result)
	builder.WriteString(", ")
	builder.WriteString("fail_reason=")
	builder.WriteString(ol.FailReason)
	builder.WriteString(", ")
	builder.WriteString("elapsed_millisecs=")
	builder.WriteString(fmt.Sprintf("%v", ol.ElapsedMillisecs))
	builder.WriteByte(')')
	return builder.String()
}

// OpLogs is a parsable slice of OpLog.
type OpLogs []*OpLog

func (ol OpLogs) config(cfg config) {
	for _i := range ol {
		ol[_i].config = cfg
	}
}
