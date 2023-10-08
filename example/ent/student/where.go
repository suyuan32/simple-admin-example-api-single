// Code generated by ent, DO NOT EDIT.

package student

import (
	"time"

	"entgo.io/ent/dialect/sql"
	uuid "github.com/gofrs/uuid/v5"
	"github.com/suyuan32/simple-admin-example-api/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.Student {
	return predicate.Student(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.Student {
	return predicate.Student(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.Student {
	return predicate.Student(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.Student {
	return predicate.Student(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.Student {
	return predicate.Student(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.Student {
	return predicate.Student(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.Student {
	return predicate.Student(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldUpdatedAt, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldName, v))
}

// Age applies equality check predicate on the "age" field. It's identical to AgeEQ.
func Age(v int) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldAge, v))
}

// AgeInt32 applies equality check predicate on the "age_int32" field. It's identical to AgeInt32EQ.
func AgeInt32(v int32) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldAgeInt32, v))
}

// AgeInt64 applies equality check predicate on the "age_int64" field. It's identical to AgeInt64EQ.
func AgeInt64(v int64) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldAgeInt64, v))
}

// AgeUint applies equality check predicate on the "age_uint" field. It's identical to AgeUintEQ.
func AgeUint(v uint) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldAgeUint, v))
}

// AgeUint32 applies equality check predicate on the "age_uint32" field. It's identical to AgeUint32EQ.
func AgeUint32(v uint32) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldAgeUint32, v))
}

// AgeUint64 applies equality check predicate on the "age_uint64" field. It's identical to AgeUint64EQ.
func AgeUint64(v uint64) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldAgeUint64, v))
}

// WeightFloat applies equality check predicate on the "weight_float" field. It's identical to WeightFloatEQ.
func WeightFloat(v float64) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldWeightFloat, v))
}

// WeightFloat32 applies equality check predicate on the "weight_float32" field. It's identical to WeightFloat32EQ.
func WeightFloat32(v float32) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldWeightFloat32, v))
}

// ClassID applies equality check predicate on the "class_id" field. It's identical to ClassIDEQ.
func ClassID(v uuid.UUID) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldClassID, v))
}

// EnrollAt applies equality check predicate on the "enroll_at" field. It's identical to EnrollAtEQ.
func EnrollAt(v time.Time) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldEnrollAt, v))
}

// StatusBool applies equality check predicate on the "status_bool" field. It's identical to StatusBoolEQ.
func StatusBool(v bool) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldStatusBool, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Student {
	return predicate.Student(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Student {
	return predicate.Student(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Student {
	return predicate.Student(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Student {
	return predicate.Student(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Student {
	return predicate.Student(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Student {
	return predicate.Student(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Student {
	return predicate.Student(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Student {
	return predicate.Student(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Student {
	return predicate.Student(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Student {
	return predicate.Student(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Student {
	return predicate.Student(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Student {
	return predicate.Student(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Student {
	return predicate.Student(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Student {
	return predicate.Student(sql.FieldLTE(FieldUpdatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Student {
	return predicate.Student(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Student {
	return predicate.Student(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Student {
	return predicate.Student(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Student {
	return predicate.Student(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Student {
	return predicate.Student(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Student {
	return predicate.Student(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Student {
	return predicate.Student(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Student {
	return predicate.Student(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Student {
	return predicate.Student(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Student {
	return predicate.Student(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Student {
	return predicate.Student(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Student {
	return predicate.Student(sql.FieldContainsFold(FieldName, v))
}

// AgeEQ applies the EQ predicate on the "age" field.
func AgeEQ(v int) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldAge, v))
}

// AgeNEQ applies the NEQ predicate on the "age" field.
func AgeNEQ(v int) predicate.Student {
	return predicate.Student(sql.FieldNEQ(FieldAge, v))
}

// AgeIn applies the In predicate on the "age" field.
func AgeIn(vs ...int) predicate.Student {
	return predicate.Student(sql.FieldIn(FieldAge, vs...))
}

// AgeNotIn applies the NotIn predicate on the "age" field.
func AgeNotIn(vs ...int) predicate.Student {
	return predicate.Student(sql.FieldNotIn(FieldAge, vs...))
}

// AgeGT applies the GT predicate on the "age" field.
func AgeGT(v int) predicate.Student {
	return predicate.Student(sql.FieldGT(FieldAge, v))
}

// AgeGTE applies the GTE predicate on the "age" field.
func AgeGTE(v int) predicate.Student {
	return predicate.Student(sql.FieldGTE(FieldAge, v))
}

// AgeLT applies the LT predicate on the "age" field.
func AgeLT(v int) predicate.Student {
	return predicate.Student(sql.FieldLT(FieldAge, v))
}

// AgeLTE applies the LTE predicate on the "age" field.
func AgeLTE(v int) predicate.Student {
	return predicate.Student(sql.FieldLTE(FieldAge, v))
}

// AgeInt32EQ applies the EQ predicate on the "age_int32" field.
func AgeInt32EQ(v int32) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldAgeInt32, v))
}

// AgeInt32NEQ applies the NEQ predicate on the "age_int32" field.
func AgeInt32NEQ(v int32) predicate.Student {
	return predicate.Student(sql.FieldNEQ(FieldAgeInt32, v))
}

// AgeInt32In applies the In predicate on the "age_int32" field.
func AgeInt32In(vs ...int32) predicate.Student {
	return predicate.Student(sql.FieldIn(FieldAgeInt32, vs...))
}

// AgeInt32NotIn applies the NotIn predicate on the "age_int32" field.
func AgeInt32NotIn(vs ...int32) predicate.Student {
	return predicate.Student(sql.FieldNotIn(FieldAgeInt32, vs...))
}

// AgeInt32GT applies the GT predicate on the "age_int32" field.
func AgeInt32GT(v int32) predicate.Student {
	return predicate.Student(sql.FieldGT(FieldAgeInt32, v))
}

// AgeInt32GTE applies the GTE predicate on the "age_int32" field.
func AgeInt32GTE(v int32) predicate.Student {
	return predicate.Student(sql.FieldGTE(FieldAgeInt32, v))
}

// AgeInt32LT applies the LT predicate on the "age_int32" field.
func AgeInt32LT(v int32) predicate.Student {
	return predicate.Student(sql.FieldLT(FieldAgeInt32, v))
}

// AgeInt32LTE applies the LTE predicate on the "age_int32" field.
func AgeInt32LTE(v int32) predicate.Student {
	return predicate.Student(sql.FieldLTE(FieldAgeInt32, v))
}

// AgeInt64EQ applies the EQ predicate on the "age_int64" field.
func AgeInt64EQ(v int64) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldAgeInt64, v))
}

// AgeInt64NEQ applies the NEQ predicate on the "age_int64" field.
func AgeInt64NEQ(v int64) predicate.Student {
	return predicate.Student(sql.FieldNEQ(FieldAgeInt64, v))
}

// AgeInt64In applies the In predicate on the "age_int64" field.
func AgeInt64In(vs ...int64) predicate.Student {
	return predicate.Student(sql.FieldIn(FieldAgeInt64, vs...))
}

// AgeInt64NotIn applies the NotIn predicate on the "age_int64" field.
func AgeInt64NotIn(vs ...int64) predicate.Student {
	return predicate.Student(sql.FieldNotIn(FieldAgeInt64, vs...))
}

// AgeInt64GT applies the GT predicate on the "age_int64" field.
func AgeInt64GT(v int64) predicate.Student {
	return predicate.Student(sql.FieldGT(FieldAgeInt64, v))
}

// AgeInt64GTE applies the GTE predicate on the "age_int64" field.
func AgeInt64GTE(v int64) predicate.Student {
	return predicate.Student(sql.FieldGTE(FieldAgeInt64, v))
}

// AgeInt64LT applies the LT predicate on the "age_int64" field.
func AgeInt64LT(v int64) predicate.Student {
	return predicate.Student(sql.FieldLT(FieldAgeInt64, v))
}

// AgeInt64LTE applies the LTE predicate on the "age_int64" field.
func AgeInt64LTE(v int64) predicate.Student {
	return predicate.Student(sql.FieldLTE(FieldAgeInt64, v))
}

// AgeUintEQ applies the EQ predicate on the "age_uint" field.
func AgeUintEQ(v uint) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldAgeUint, v))
}

// AgeUintNEQ applies the NEQ predicate on the "age_uint" field.
func AgeUintNEQ(v uint) predicate.Student {
	return predicate.Student(sql.FieldNEQ(FieldAgeUint, v))
}

// AgeUintIn applies the In predicate on the "age_uint" field.
func AgeUintIn(vs ...uint) predicate.Student {
	return predicate.Student(sql.FieldIn(FieldAgeUint, vs...))
}

// AgeUintNotIn applies the NotIn predicate on the "age_uint" field.
func AgeUintNotIn(vs ...uint) predicate.Student {
	return predicate.Student(sql.FieldNotIn(FieldAgeUint, vs...))
}

// AgeUintGT applies the GT predicate on the "age_uint" field.
func AgeUintGT(v uint) predicate.Student {
	return predicate.Student(sql.FieldGT(FieldAgeUint, v))
}

// AgeUintGTE applies the GTE predicate on the "age_uint" field.
func AgeUintGTE(v uint) predicate.Student {
	return predicate.Student(sql.FieldGTE(FieldAgeUint, v))
}

// AgeUintLT applies the LT predicate on the "age_uint" field.
func AgeUintLT(v uint) predicate.Student {
	return predicate.Student(sql.FieldLT(FieldAgeUint, v))
}

// AgeUintLTE applies the LTE predicate on the "age_uint" field.
func AgeUintLTE(v uint) predicate.Student {
	return predicate.Student(sql.FieldLTE(FieldAgeUint, v))
}

// AgeUint32EQ applies the EQ predicate on the "age_uint32" field.
func AgeUint32EQ(v uint32) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldAgeUint32, v))
}

// AgeUint32NEQ applies the NEQ predicate on the "age_uint32" field.
func AgeUint32NEQ(v uint32) predicate.Student {
	return predicate.Student(sql.FieldNEQ(FieldAgeUint32, v))
}

// AgeUint32In applies the In predicate on the "age_uint32" field.
func AgeUint32In(vs ...uint32) predicate.Student {
	return predicate.Student(sql.FieldIn(FieldAgeUint32, vs...))
}

// AgeUint32NotIn applies the NotIn predicate on the "age_uint32" field.
func AgeUint32NotIn(vs ...uint32) predicate.Student {
	return predicate.Student(sql.FieldNotIn(FieldAgeUint32, vs...))
}

// AgeUint32GT applies the GT predicate on the "age_uint32" field.
func AgeUint32GT(v uint32) predicate.Student {
	return predicate.Student(sql.FieldGT(FieldAgeUint32, v))
}

// AgeUint32GTE applies the GTE predicate on the "age_uint32" field.
func AgeUint32GTE(v uint32) predicate.Student {
	return predicate.Student(sql.FieldGTE(FieldAgeUint32, v))
}

// AgeUint32LT applies the LT predicate on the "age_uint32" field.
func AgeUint32LT(v uint32) predicate.Student {
	return predicate.Student(sql.FieldLT(FieldAgeUint32, v))
}

// AgeUint32LTE applies the LTE predicate on the "age_uint32" field.
func AgeUint32LTE(v uint32) predicate.Student {
	return predicate.Student(sql.FieldLTE(FieldAgeUint32, v))
}

// AgeUint64EQ applies the EQ predicate on the "age_uint64" field.
func AgeUint64EQ(v uint64) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldAgeUint64, v))
}

// AgeUint64NEQ applies the NEQ predicate on the "age_uint64" field.
func AgeUint64NEQ(v uint64) predicate.Student {
	return predicate.Student(sql.FieldNEQ(FieldAgeUint64, v))
}

// AgeUint64In applies the In predicate on the "age_uint64" field.
func AgeUint64In(vs ...uint64) predicate.Student {
	return predicate.Student(sql.FieldIn(FieldAgeUint64, vs...))
}

// AgeUint64NotIn applies the NotIn predicate on the "age_uint64" field.
func AgeUint64NotIn(vs ...uint64) predicate.Student {
	return predicate.Student(sql.FieldNotIn(FieldAgeUint64, vs...))
}

// AgeUint64GT applies the GT predicate on the "age_uint64" field.
func AgeUint64GT(v uint64) predicate.Student {
	return predicate.Student(sql.FieldGT(FieldAgeUint64, v))
}

// AgeUint64GTE applies the GTE predicate on the "age_uint64" field.
func AgeUint64GTE(v uint64) predicate.Student {
	return predicate.Student(sql.FieldGTE(FieldAgeUint64, v))
}

// AgeUint64LT applies the LT predicate on the "age_uint64" field.
func AgeUint64LT(v uint64) predicate.Student {
	return predicate.Student(sql.FieldLT(FieldAgeUint64, v))
}

// AgeUint64LTE applies the LTE predicate on the "age_uint64" field.
func AgeUint64LTE(v uint64) predicate.Student {
	return predicate.Student(sql.FieldLTE(FieldAgeUint64, v))
}

// WeightFloatEQ applies the EQ predicate on the "weight_float" field.
func WeightFloatEQ(v float64) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldWeightFloat, v))
}

// WeightFloatNEQ applies the NEQ predicate on the "weight_float" field.
func WeightFloatNEQ(v float64) predicate.Student {
	return predicate.Student(sql.FieldNEQ(FieldWeightFloat, v))
}

// WeightFloatIn applies the In predicate on the "weight_float" field.
func WeightFloatIn(vs ...float64) predicate.Student {
	return predicate.Student(sql.FieldIn(FieldWeightFloat, vs...))
}

// WeightFloatNotIn applies the NotIn predicate on the "weight_float" field.
func WeightFloatNotIn(vs ...float64) predicate.Student {
	return predicate.Student(sql.FieldNotIn(FieldWeightFloat, vs...))
}

// WeightFloatGT applies the GT predicate on the "weight_float" field.
func WeightFloatGT(v float64) predicate.Student {
	return predicate.Student(sql.FieldGT(FieldWeightFloat, v))
}

// WeightFloatGTE applies the GTE predicate on the "weight_float" field.
func WeightFloatGTE(v float64) predicate.Student {
	return predicate.Student(sql.FieldGTE(FieldWeightFloat, v))
}

// WeightFloatLT applies the LT predicate on the "weight_float" field.
func WeightFloatLT(v float64) predicate.Student {
	return predicate.Student(sql.FieldLT(FieldWeightFloat, v))
}

// WeightFloatLTE applies the LTE predicate on the "weight_float" field.
func WeightFloatLTE(v float64) predicate.Student {
	return predicate.Student(sql.FieldLTE(FieldWeightFloat, v))
}

// WeightFloat32EQ applies the EQ predicate on the "weight_float32" field.
func WeightFloat32EQ(v float32) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldWeightFloat32, v))
}

// WeightFloat32NEQ applies the NEQ predicate on the "weight_float32" field.
func WeightFloat32NEQ(v float32) predicate.Student {
	return predicate.Student(sql.FieldNEQ(FieldWeightFloat32, v))
}

// WeightFloat32In applies the In predicate on the "weight_float32" field.
func WeightFloat32In(vs ...float32) predicate.Student {
	return predicate.Student(sql.FieldIn(FieldWeightFloat32, vs...))
}

// WeightFloat32NotIn applies the NotIn predicate on the "weight_float32" field.
func WeightFloat32NotIn(vs ...float32) predicate.Student {
	return predicate.Student(sql.FieldNotIn(FieldWeightFloat32, vs...))
}

// WeightFloat32GT applies the GT predicate on the "weight_float32" field.
func WeightFloat32GT(v float32) predicate.Student {
	return predicate.Student(sql.FieldGT(FieldWeightFloat32, v))
}

// WeightFloat32GTE applies the GTE predicate on the "weight_float32" field.
func WeightFloat32GTE(v float32) predicate.Student {
	return predicate.Student(sql.FieldGTE(FieldWeightFloat32, v))
}

// WeightFloat32LT applies the LT predicate on the "weight_float32" field.
func WeightFloat32LT(v float32) predicate.Student {
	return predicate.Student(sql.FieldLT(FieldWeightFloat32, v))
}

// WeightFloat32LTE applies the LTE predicate on the "weight_float32" field.
func WeightFloat32LTE(v float32) predicate.Student {
	return predicate.Student(sql.FieldLTE(FieldWeightFloat32, v))
}

// ClassIDEQ applies the EQ predicate on the "class_id" field.
func ClassIDEQ(v uuid.UUID) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldClassID, v))
}

// ClassIDNEQ applies the NEQ predicate on the "class_id" field.
func ClassIDNEQ(v uuid.UUID) predicate.Student {
	return predicate.Student(sql.FieldNEQ(FieldClassID, v))
}

// ClassIDIn applies the In predicate on the "class_id" field.
func ClassIDIn(vs ...uuid.UUID) predicate.Student {
	return predicate.Student(sql.FieldIn(FieldClassID, vs...))
}

// ClassIDNotIn applies the NotIn predicate on the "class_id" field.
func ClassIDNotIn(vs ...uuid.UUID) predicate.Student {
	return predicate.Student(sql.FieldNotIn(FieldClassID, vs...))
}

// ClassIDGT applies the GT predicate on the "class_id" field.
func ClassIDGT(v uuid.UUID) predicate.Student {
	return predicate.Student(sql.FieldGT(FieldClassID, v))
}

// ClassIDGTE applies the GTE predicate on the "class_id" field.
func ClassIDGTE(v uuid.UUID) predicate.Student {
	return predicate.Student(sql.FieldGTE(FieldClassID, v))
}

// ClassIDLT applies the LT predicate on the "class_id" field.
func ClassIDLT(v uuid.UUID) predicate.Student {
	return predicate.Student(sql.FieldLT(FieldClassID, v))
}

// ClassIDLTE applies the LTE predicate on the "class_id" field.
func ClassIDLTE(v uuid.UUID) predicate.Student {
	return predicate.Student(sql.FieldLTE(FieldClassID, v))
}

// EnrollAtEQ applies the EQ predicate on the "enroll_at" field.
func EnrollAtEQ(v time.Time) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldEnrollAt, v))
}

// EnrollAtNEQ applies the NEQ predicate on the "enroll_at" field.
func EnrollAtNEQ(v time.Time) predicate.Student {
	return predicate.Student(sql.FieldNEQ(FieldEnrollAt, v))
}

// EnrollAtIn applies the In predicate on the "enroll_at" field.
func EnrollAtIn(vs ...time.Time) predicate.Student {
	return predicate.Student(sql.FieldIn(FieldEnrollAt, vs...))
}

// EnrollAtNotIn applies the NotIn predicate on the "enroll_at" field.
func EnrollAtNotIn(vs ...time.Time) predicate.Student {
	return predicate.Student(sql.FieldNotIn(FieldEnrollAt, vs...))
}

// EnrollAtGT applies the GT predicate on the "enroll_at" field.
func EnrollAtGT(v time.Time) predicate.Student {
	return predicate.Student(sql.FieldGT(FieldEnrollAt, v))
}

// EnrollAtGTE applies the GTE predicate on the "enroll_at" field.
func EnrollAtGTE(v time.Time) predicate.Student {
	return predicate.Student(sql.FieldGTE(FieldEnrollAt, v))
}

// EnrollAtLT applies the LT predicate on the "enroll_at" field.
func EnrollAtLT(v time.Time) predicate.Student {
	return predicate.Student(sql.FieldLT(FieldEnrollAt, v))
}

// EnrollAtLTE applies the LTE predicate on the "enroll_at" field.
func EnrollAtLTE(v time.Time) predicate.Student {
	return predicate.Student(sql.FieldLTE(FieldEnrollAt, v))
}

// StatusBoolEQ applies the EQ predicate on the "status_bool" field.
func StatusBoolEQ(v bool) predicate.Student {
	return predicate.Student(sql.FieldEQ(FieldStatusBool, v))
}

// StatusBoolNEQ applies the NEQ predicate on the "status_bool" field.
func StatusBoolNEQ(v bool) predicate.Student {
	return predicate.Student(sql.FieldNEQ(FieldStatusBool, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Student) predicate.Student {
	return predicate.Student(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Student) predicate.Student {
	return predicate.Student(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Student) predicate.Student {
	return predicate.Student(sql.NotPredicates(p))
}
