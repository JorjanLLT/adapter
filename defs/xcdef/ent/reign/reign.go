// Code generated by ent, DO NOT EDIT.

package reign

import (
	"time"

	"github.com/auroraride/adapter/defs/batdef"
)

const (
	// Label holds the string label denoting the reign type in the database.
	Label = "reign"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAction holds the string denoting the action field in the database.
	FieldAction = "action"
	// FieldSn holds the string denoting the sn field in the database.
	FieldSn = "sn"
	// FieldBatteryID holds the string denoting the battery_id field in the database.
	FieldBatteryID = "battery_id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldSerial holds the string denoting the serial field in the database.
	FieldSerial = "serial"
	// FieldOrdinal holds the string denoting the ordinal field in the database.
	FieldOrdinal = "ordinal"
	// FieldCabinetName holds the string denoting the cabinet_name field in the database.
	FieldCabinetName = "cabinet_name"
	// FieldRemark holds the string denoting the remark field in the database.
	FieldRemark = "remark"
	// FieldGeom holds the string denoting the geom field in the database.
	FieldGeom = "geom"
	// EdgeBattery holds the string denoting the battery edge name in mutations.
	EdgeBattery = "battery"
	// Table holds the table name of the reign in the database.
	Table = "reign"
	// BatteryTable is the table that holds the battery relation/edge.
	BatteryTable = "reign"
	// BatteryInverseTable is the table name for the Battery entity.
	// It exists in this package in order to avoid circular dependency with the "battery" package.
	BatteryInverseTable = "battery"
	// BatteryColumn is the table column denoting the battery relation/edge.
	BatteryColumn = "battery_id"
)

// Columns holds all SQL columns for reign fields.
var Columns = []string{
	FieldID,
	FieldAction,
	FieldSn,
	FieldBatteryID,
	FieldCreatedAt,
	FieldSerial,
	FieldOrdinal,
	FieldCabinetName,
	FieldRemark,
	FieldGeom,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultAction holds the default value on creation for the "action" field.
	DefaultAction batdef.ReignAction
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)
