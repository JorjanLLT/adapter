// Code generated by ent, DO NOT EDIT.

package fault

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/auroraride/adapter/defs/xcdef"
	"github.com/auroraride/adapter/defs/xcdef/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Fault {
	return predicate.Fault(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Fault {
	return predicate.Fault(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Fault {
	return predicate.Fault(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Fault {
	return predicate.Fault(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Fault {
	return predicate.Fault(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Fault {
	return predicate.Fault(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Fault {
	return predicate.Fault(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Fault {
	return predicate.Fault(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Fault {
	return predicate.Fault(sql.FieldLTE(FieldID, id))
}

// Sn applies equality check predicate on the "sn" field. It's identical to SnEQ.
func Sn(v string) predicate.Fault {
	return predicate.Fault(sql.FieldEQ(FieldSn, v))
}

// BatteryID applies equality check predicate on the "battery_id" field. It's identical to BatteryIDEQ.
func BatteryID(v int) predicate.Fault {
	return predicate.Fault(sql.FieldEQ(FieldBatteryID, v))
}

// Fault applies equality check predicate on the "fault" field. It's identical to FaultEQ.
func Fault(v xcdef.Fault) predicate.Fault {
	return predicate.Fault(sql.FieldEQ(FieldFault, v))
}

// BeginAt applies equality check predicate on the "begin_at" field. It's identical to BeginAtEQ.
func BeginAt(v time.Time) predicate.Fault {
	return predicate.Fault(sql.FieldEQ(FieldBeginAt, v))
}

// EndAt applies equality check predicate on the "end_at" field. It's identical to EndAtEQ.
func EndAt(v time.Time) predicate.Fault {
	return predicate.Fault(sql.FieldEQ(FieldEndAt, v))
}

// SnEQ applies the EQ predicate on the "sn" field.
func SnEQ(v string) predicate.Fault {
	return predicate.Fault(sql.FieldEQ(FieldSn, v))
}

// SnNEQ applies the NEQ predicate on the "sn" field.
func SnNEQ(v string) predicate.Fault {
	return predicate.Fault(sql.FieldNEQ(FieldSn, v))
}

// SnIn applies the In predicate on the "sn" field.
func SnIn(vs ...string) predicate.Fault {
	return predicate.Fault(sql.FieldIn(FieldSn, vs...))
}

// SnNotIn applies the NotIn predicate on the "sn" field.
func SnNotIn(vs ...string) predicate.Fault {
	return predicate.Fault(sql.FieldNotIn(FieldSn, vs...))
}

// SnGT applies the GT predicate on the "sn" field.
func SnGT(v string) predicate.Fault {
	return predicate.Fault(sql.FieldGT(FieldSn, v))
}

// SnGTE applies the GTE predicate on the "sn" field.
func SnGTE(v string) predicate.Fault {
	return predicate.Fault(sql.FieldGTE(FieldSn, v))
}

// SnLT applies the LT predicate on the "sn" field.
func SnLT(v string) predicate.Fault {
	return predicate.Fault(sql.FieldLT(FieldSn, v))
}

// SnLTE applies the LTE predicate on the "sn" field.
func SnLTE(v string) predicate.Fault {
	return predicate.Fault(sql.FieldLTE(FieldSn, v))
}

// SnContains applies the Contains predicate on the "sn" field.
func SnContains(v string) predicate.Fault {
	return predicate.Fault(sql.FieldContains(FieldSn, v))
}

// SnHasPrefix applies the HasPrefix predicate on the "sn" field.
func SnHasPrefix(v string) predicate.Fault {
	return predicate.Fault(sql.FieldHasPrefix(FieldSn, v))
}

// SnHasSuffix applies the HasSuffix predicate on the "sn" field.
func SnHasSuffix(v string) predicate.Fault {
	return predicate.Fault(sql.FieldHasSuffix(FieldSn, v))
}

// SnEqualFold applies the EqualFold predicate on the "sn" field.
func SnEqualFold(v string) predicate.Fault {
	return predicate.Fault(sql.FieldEqualFold(FieldSn, v))
}

// SnContainsFold applies the ContainsFold predicate on the "sn" field.
func SnContainsFold(v string) predicate.Fault {
	return predicate.Fault(sql.FieldContainsFold(FieldSn, v))
}

// BatteryIDEQ applies the EQ predicate on the "battery_id" field.
func BatteryIDEQ(v int) predicate.Fault {
	return predicate.Fault(sql.FieldEQ(FieldBatteryID, v))
}

// BatteryIDNEQ applies the NEQ predicate on the "battery_id" field.
func BatteryIDNEQ(v int) predicate.Fault {
	return predicate.Fault(sql.FieldNEQ(FieldBatteryID, v))
}

// BatteryIDIn applies the In predicate on the "battery_id" field.
func BatteryIDIn(vs ...int) predicate.Fault {
	return predicate.Fault(sql.FieldIn(FieldBatteryID, vs...))
}

// BatteryIDNotIn applies the NotIn predicate on the "battery_id" field.
func BatteryIDNotIn(vs ...int) predicate.Fault {
	return predicate.Fault(sql.FieldNotIn(FieldBatteryID, vs...))
}

// FaultEQ applies the EQ predicate on the "fault" field.
func FaultEQ(v xcdef.Fault) predicate.Fault {
	return predicate.Fault(sql.FieldEQ(FieldFault, v))
}

// FaultNEQ applies the NEQ predicate on the "fault" field.
func FaultNEQ(v xcdef.Fault) predicate.Fault {
	return predicate.Fault(sql.FieldNEQ(FieldFault, v))
}

// FaultIn applies the In predicate on the "fault" field.
func FaultIn(vs ...xcdef.Fault) predicate.Fault {
	return predicate.Fault(sql.FieldIn(FieldFault, vs...))
}

// FaultNotIn applies the NotIn predicate on the "fault" field.
func FaultNotIn(vs ...xcdef.Fault) predicate.Fault {
	return predicate.Fault(sql.FieldNotIn(FieldFault, vs...))
}

// FaultGT applies the GT predicate on the "fault" field.
func FaultGT(v xcdef.Fault) predicate.Fault {
	return predicate.Fault(sql.FieldGT(FieldFault, v))
}

// FaultGTE applies the GTE predicate on the "fault" field.
func FaultGTE(v xcdef.Fault) predicate.Fault {
	return predicate.Fault(sql.FieldGTE(FieldFault, v))
}

// FaultLT applies the LT predicate on the "fault" field.
func FaultLT(v xcdef.Fault) predicate.Fault {
	return predicate.Fault(sql.FieldLT(FieldFault, v))
}

// FaultLTE applies the LTE predicate on the "fault" field.
func FaultLTE(v xcdef.Fault) predicate.Fault {
	return predicate.Fault(sql.FieldLTE(FieldFault, v))
}

// BeginAtEQ applies the EQ predicate on the "begin_at" field.
func BeginAtEQ(v time.Time) predicate.Fault {
	return predicate.Fault(sql.FieldEQ(FieldBeginAt, v))
}

// BeginAtNEQ applies the NEQ predicate on the "begin_at" field.
func BeginAtNEQ(v time.Time) predicate.Fault {
	return predicate.Fault(sql.FieldNEQ(FieldBeginAt, v))
}

// BeginAtIn applies the In predicate on the "begin_at" field.
func BeginAtIn(vs ...time.Time) predicate.Fault {
	return predicate.Fault(sql.FieldIn(FieldBeginAt, vs...))
}

// BeginAtNotIn applies the NotIn predicate on the "begin_at" field.
func BeginAtNotIn(vs ...time.Time) predicate.Fault {
	return predicate.Fault(sql.FieldNotIn(FieldBeginAt, vs...))
}

// BeginAtGT applies the GT predicate on the "begin_at" field.
func BeginAtGT(v time.Time) predicate.Fault {
	return predicate.Fault(sql.FieldGT(FieldBeginAt, v))
}

// BeginAtGTE applies the GTE predicate on the "begin_at" field.
func BeginAtGTE(v time.Time) predicate.Fault {
	return predicate.Fault(sql.FieldGTE(FieldBeginAt, v))
}

// BeginAtLT applies the LT predicate on the "begin_at" field.
func BeginAtLT(v time.Time) predicate.Fault {
	return predicate.Fault(sql.FieldLT(FieldBeginAt, v))
}

// BeginAtLTE applies the LTE predicate on the "begin_at" field.
func BeginAtLTE(v time.Time) predicate.Fault {
	return predicate.Fault(sql.FieldLTE(FieldBeginAt, v))
}

// EndAtEQ applies the EQ predicate on the "end_at" field.
func EndAtEQ(v time.Time) predicate.Fault {
	return predicate.Fault(sql.FieldEQ(FieldEndAt, v))
}

// EndAtNEQ applies the NEQ predicate on the "end_at" field.
func EndAtNEQ(v time.Time) predicate.Fault {
	return predicate.Fault(sql.FieldNEQ(FieldEndAt, v))
}

// EndAtIn applies the In predicate on the "end_at" field.
func EndAtIn(vs ...time.Time) predicate.Fault {
	return predicate.Fault(sql.FieldIn(FieldEndAt, vs...))
}

// EndAtNotIn applies the NotIn predicate on the "end_at" field.
func EndAtNotIn(vs ...time.Time) predicate.Fault {
	return predicate.Fault(sql.FieldNotIn(FieldEndAt, vs...))
}

// EndAtGT applies the GT predicate on the "end_at" field.
func EndAtGT(v time.Time) predicate.Fault {
	return predicate.Fault(sql.FieldGT(FieldEndAt, v))
}

// EndAtGTE applies the GTE predicate on the "end_at" field.
func EndAtGTE(v time.Time) predicate.Fault {
	return predicate.Fault(sql.FieldGTE(FieldEndAt, v))
}

// EndAtLT applies the LT predicate on the "end_at" field.
func EndAtLT(v time.Time) predicate.Fault {
	return predicate.Fault(sql.FieldLT(FieldEndAt, v))
}

// EndAtLTE applies the LTE predicate on the "end_at" field.
func EndAtLTE(v time.Time) predicate.Fault {
	return predicate.Fault(sql.FieldLTE(FieldEndAt, v))
}

// EndAtIsNil applies the IsNil predicate on the "end_at" field.
func EndAtIsNil() predicate.Fault {
	return predicate.Fault(sql.FieldIsNull(FieldEndAt))
}

// EndAtNotNil applies the NotNil predicate on the "end_at" field.
func EndAtNotNil() predicate.Fault {
	return predicate.Fault(sql.FieldNotNull(FieldEndAt))
}

// HasBattery applies the HasEdge predicate on the "battery" edge.
func HasBattery() predicate.Fault {
	return predicate.Fault(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BatteryTable, BatteryColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBatteryWith applies the HasEdge predicate on the "battery" edge with a given conditions (other predicates).
func HasBatteryWith(preds ...predicate.Battery) predicate.Fault {
	return predicate.Fault(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BatteryInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BatteryTable, BatteryColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Fault) predicate.Fault {
	return predicate.Fault(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Fault) predicate.Fault {
	return predicate.Fault(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Fault) predicate.Fault {
	return predicate.Fault(func(s *sql.Selector) {
		p(s.Not())
	})
}
