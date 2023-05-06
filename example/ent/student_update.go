// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	uuid "github.com/gofrs/uuid/v5"
	"github.com/suyuan32/simple-admin-example-api/ent/predicate"
	"github.com/suyuan32/simple-admin-example-api/ent/student"
)

// StudentUpdate is the builder for updating Student entities.
type StudentUpdate struct {
	config
	hooks    []Hook
	mutation *StudentMutation
}

// Where appends a list predicates to the StudentUpdate builder.
func (su *StudentUpdate) Where(ps ...predicate.Student) *StudentUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetUpdatedAt sets the "updated_at" field.
func (su *StudentUpdate) SetUpdatedAt(t time.Time) *StudentUpdate {
	su.mutation.SetUpdatedAt(t)
	return su
}

// SetName sets the "name" field.
func (su *StudentUpdate) SetName(s string) *StudentUpdate {
	su.mutation.SetName(s)
	return su
}

// SetAge sets the "age" field.
func (su *StudentUpdate) SetAge(i int) *StudentUpdate {
	su.mutation.ResetAge()
	su.mutation.SetAge(i)
	return su
}

// AddAge adds i to the "age" field.
func (su *StudentUpdate) AddAge(i int) *StudentUpdate {
	su.mutation.AddAge(i)
	return su
}

// SetAgeInt8 sets the "age_int8" field.
func (su *StudentUpdate) SetAgeInt8(i int8) *StudentUpdate {
	su.mutation.ResetAgeInt8()
	su.mutation.SetAgeInt8(i)
	return su
}

// AddAgeInt8 adds i to the "age_int8" field.
func (su *StudentUpdate) AddAgeInt8(i int8) *StudentUpdate {
	su.mutation.AddAgeInt8(i)
	return su
}

// SetAgeUint8 sets the "age_uint8" field.
func (su *StudentUpdate) SetAgeUint8(u uint8) *StudentUpdate {
	su.mutation.ResetAgeUint8()
	su.mutation.SetAgeUint8(u)
	return su
}

// AddAgeUint8 adds u to the "age_uint8" field.
func (su *StudentUpdate) AddAgeUint8(u int8) *StudentUpdate {
	su.mutation.AddAgeUint8(u)
	return su
}

// SetAgeInt16 sets the "age_int16" field.
func (su *StudentUpdate) SetAgeInt16(i int16) *StudentUpdate {
	su.mutation.ResetAgeInt16()
	su.mutation.SetAgeInt16(i)
	return su
}

// AddAgeInt16 adds i to the "age_int16" field.
func (su *StudentUpdate) AddAgeInt16(i int16) *StudentUpdate {
	su.mutation.AddAgeInt16(i)
	return su
}

// SetAgeUint16 sets the "age_uint16" field.
func (su *StudentUpdate) SetAgeUint16(u uint16) *StudentUpdate {
	su.mutation.ResetAgeUint16()
	su.mutation.SetAgeUint16(u)
	return su
}

// AddAgeUint16 adds u to the "age_uint16" field.
func (su *StudentUpdate) AddAgeUint16(u int16) *StudentUpdate {
	su.mutation.AddAgeUint16(u)
	return su
}

// SetAgeInt32 sets the "age_int32" field.
func (su *StudentUpdate) SetAgeInt32(i int32) *StudentUpdate {
	su.mutation.ResetAgeInt32()
	su.mutation.SetAgeInt32(i)
	return su
}

// AddAgeInt32 adds i to the "age_int32" field.
func (su *StudentUpdate) AddAgeInt32(i int32) *StudentUpdate {
	su.mutation.AddAgeInt32(i)
	return su
}

// SetAgeUint32 sets the "age_uint32" field.
func (su *StudentUpdate) SetAgeUint32(u uint32) *StudentUpdate {
	su.mutation.ResetAgeUint32()
	su.mutation.SetAgeUint32(u)
	return su
}

// AddAgeUint32 adds u to the "age_uint32" field.
func (su *StudentUpdate) AddAgeUint32(u int32) *StudentUpdate {
	su.mutation.AddAgeUint32(u)
	return su
}

// SetAgeInt64 sets the "age_int64" field.
func (su *StudentUpdate) SetAgeInt64(i int64) *StudentUpdate {
	su.mutation.ResetAgeInt64()
	su.mutation.SetAgeInt64(i)
	return su
}

// AddAgeInt64 adds i to the "age_int64" field.
func (su *StudentUpdate) AddAgeInt64(i int64) *StudentUpdate {
	su.mutation.AddAgeInt64(i)
	return su
}

// SetAgeUint64 sets the "age_uint64" field.
func (su *StudentUpdate) SetAgeUint64(u uint64) *StudentUpdate {
	su.mutation.ResetAgeUint64()
	su.mutation.SetAgeUint64(u)
	return su
}

// AddAgeUint64 adds u to the "age_uint64" field.
func (su *StudentUpdate) AddAgeUint64(u int64) *StudentUpdate {
	su.mutation.AddAgeUint64(u)
	return su
}

// SetAgeInt sets the "age_int" field.
func (su *StudentUpdate) SetAgeInt(i int) *StudentUpdate {
	su.mutation.ResetAgeInt()
	su.mutation.SetAgeInt(i)
	return su
}

// AddAgeInt adds i to the "age_int" field.
func (su *StudentUpdate) AddAgeInt(i int) *StudentUpdate {
	su.mutation.AddAgeInt(i)
	return su
}

// SetAgeUint sets the "age_uint" field.
func (su *StudentUpdate) SetAgeUint(u uint) *StudentUpdate {
	su.mutation.ResetAgeUint()
	su.mutation.SetAgeUint(u)
	return su
}

// AddAgeUint adds u to the "age_uint" field.
func (su *StudentUpdate) AddAgeUint(u int) *StudentUpdate {
	su.mutation.AddAgeUint(u)
	return su
}

// SetWeightFloat sets the "weight_float" field.
func (su *StudentUpdate) SetWeightFloat(f float64) *StudentUpdate {
	su.mutation.ResetWeightFloat()
	su.mutation.SetWeightFloat(f)
	return su
}

// AddWeightFloat adds f to the "weight_float" field.
func (su *StudentUpdate) AddWeightFloat(f float64) *StudentUpdate {
	su.mutation.AddWeightFloat(f)
	return su
}

// SetWeightFloat32 sets the "weight_float32" field.
func (su *StudentUpdate) SetWeightFloat32(f float32) *StudentUpdate {
	su.mutation.ResetWeightFloat32()
	su.mutation.SetWeightFloat32(f)
	return su
}

// AddWeightFloat32 adds f to the "weight_float32" field.
func (su *StudentUpdate) AddWeightFloat32(f float32) *StudentUpdate {
	su.mutation.AddWeightFloat32(f)
	return su
}

// SetClassID sets the "class_id" field.
func (su *StudentUpdate) SetClassID(u uuid.UUID) *StudentUpdate {
	su.mutation.SetClassID(u)
	return su
}

// SetEnrollAt sets the "enroll_at" field.
func (su *StudentUpdate) SetEnrollAt(t time.Time) *StudentUpdate {
	su.mutation.SetEnrollAt(t)
	return su
}

// SetStatusBool sets the "status_bool" field.
func (su *StudentUpdate) SetStatusBool(b bool) *StudentUpdate {
	su.mutation.SetStatusBool(b)
	return su
}

// Mutation returns the StudentMutation object of the builder.
func (su *StudentUpdate) Mutation() *StudentMutation {
	return su.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *StudentUpdate) Save(ctx context.Context) (int, error) {
	su.defaults()
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *StudentUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *StudentUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *StudentUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (su *StudentUpdate) defaults() {
	if _, ok := su.mutation.UpdatedAt(); !ok {
		v := student.UpdateDefaultUpdatedAt()
		su.mutation.SetUpdatedAt(v)
	}
}

func (su *StudentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(student.Table, student.Columns, sqlgraph.NewFieldSpec(student.FieldID, field.TypeUint64))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.UpdatedAt(); ok {
		_spec.SetField(student.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := su.mutation.Name(); ok {
		_spec.SetField(student.FieldName, field.TypeString, value)
	}
	if value, ok := su.mutation.Age(); ok {
		_spec.SetField(student.FieldAge, field.TypeInt, value)
	}
	if value, ok := su.mutation.AddedAge(); ok {
		_spec.AddField(student.FieldAge, field.TypeInt, value)
	}
	if value, ok := su.mutation.AgeInt8(); ok {
		_spec.SetField(student.FieldAgeInt8, field.TypeInt8, value)
	}
	if value, ok := su.mutation.AddedAgeInt8(); ok {
		_spec.AddField(student.FieldAgeInt8, field.TypeInt8, value)
	}
	if value, ok := su.mutation.AgeUint8(); ok {
		_spec.SetField(student.FieldAgeUint8, field.TypeUint8, value)
	}
	if value, ok := su.mutation.AddedAgeUint8(); ok {
		_spec.AddField(student.FieldAgeUint8, field.TypeUint8, value)
	}
	if value, ok := su.mutation.AgeInt16(); ok {
		_spec.SetField(student.FieldAgeInt16, field.TypeInt16, value)
	}
	if value, ok := su.mutation.AddedAgeInt16(); ok {
		_spec.AddField(student.FieldAgeInt16, field.TypeInt16, value)
	}
	if value, ok := su.mutation.AgeUint16(); ok {
		_spec.SetField(student.FieldAgeUint16, field.TypeUint16, value)
	}
	if value, ok := su.mutation.AddedAgeUint16(); ok {
		_spec.AddField(student.FieldAgeUint16, field.TypeUint16, value)
	}
	if value, ok := su.mutation.AgeInt32(); ok {
		_spec.SetField(student.FieldAgeInt32, field.TypeInt32, value)
	}
	if value, ok := su.mutation.AddedAgeInt32(); ok {
		_spec.AddField(student.FieldAgeInt32, field.TypeInt32, value)
	}
	if value, ok := su.mutation.AgeUint32(); ok {
		_spec.SetField(student.FieldAgeUint32, field.TypeUint32, value)
	}
	if value, ok := su.mutation.AddedAgeUint32(); ok {
		_spec.AddField(student.FieldAgeUint32, field.TypeUint32, value)
	}
	if value, ok := su.mutation.AgeInt64(); ok {
		_spec.SetField(student.FieldAgeInt64, field.TypeInt64, value)
	}
	if value, ok := su.mutation.AddedAgeInt64(); ok {
		_spec.AddField(student.FieldAgeInt64, field.TypeInt64, value)
	}
	if value, ok := su.mutation.AgeUint64(); ok {
		_spec.SetField(student.FieldAgeUint64, field.TypeUint64, value)
	}
	if value, ok := su.mutation.AddedAgeUint64(); ok {
		_spec.AddField(student.FieldAgeUint64, field.TypeUint64, value)
	}
	if value, ok := su.mutation.AgeInt(); ok {
		_spec.SetField(student.FieldAgeInt, field.TypeInt, value)
	}
	if value, ok := su.mutation.AddedAgeInt(); ok {
		_spec.AddField(student.FieldAgeInt, field.TypeInt, value)
	}
	if value, ok := su.mutation.AgeUint(); ok {
		_spec.SetField(student.FieldAgeUint, field.TypeUint, value)
	}
	if value, ok := su.mutation.AddedAgeUint(); ok {
		_spec.AddField(student.FieldAgeUint, field.TypeUint, value)
	}
	if value, ok := su.mutation.WeightFloat(); ok {
		_spec.SetField(student.FieldWeightFloat, field.TypeFloat64, value)
	}
	if value, ok := su.mutation.AddedWeightFloat(); ok {
		_spec.AddField(student.FieldWeightFloat, field.TypeFloat64, value)
	}
	if value, ok := su.mutation.WeightFloat32(); ok {
		_spec.SetField(student.FieldWeightFloat32, field.TypeFloat32, value)
	}
	if value, ok := su.mutation.AddedWeightFloat32(); ok {
		_spec.AddField(student.FieldWeightFloat32, field.TypeFloat32, value)
	}
	if value, ok := su.mutation.ClassID(); ok {
		_spec.SetField(student.FieldClassID, field.TypeUUID, value)
	}
	if value, ok := su.mutation.EnrollAt(); ok {
		_spec.SetField(student.FieldEnrollAt, field.TypeTime, value)
	}
	if value, ok := su.mutation.StatusBool(); ok {
		_spec.SetField(student.FieldStatusBool, field.TypeBool, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{student.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// StudentUpdateOne is the builder for updating a single Student entity.
type StudentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *StudentMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (suo *StudentUpdateOne) SetUpdatedAt(t time.Time) *StudentUpdateOne {
	suo.mutation.SetUpdatedAt(t)
	return suo
}

// SetName sets the "name" field.
func (suo *StudentUpdateOne) SetName(s string) *StudentUpdateOne {
	suo.mutation.SetName(s)
	return suo
}

// SetAge sets the "age" field.
func (suo *StudentUpdateOne) SetAge(i int) *StudentUpdateOne {
	suo.mutation.ResetAge()
	suo.mutation.SetAge(i)
	return suo
}

// AddAge adds i to the "age" field.
func (suo *StudentUpdateOne) AddAge(i int) *StudentUpdateOne {
	suo.mutation.AddAge(i)
	return suo
}

// SetAgeInt8 sets the "age_int8" field.
func (suo *StudentUpdateOne) SetAgeInt8(i int8) *StudentUpdateOne {
	suo.mutation.ResetAgeInt8()
	suo.mutation.SetAgeInt8(i)
	return suo
}

// AddAgeInt8 adds i to the "age_int8" field.
func (suo *StudentUpdateOne) AddAgeInt8(i int8) *StudentUpdateOne {
	suo.mutation.AddAgeInt8(i)
	return suo
}

// SetAgeUint8 sets the "age_uint8" field.
func (suo *StudentUpdateOne) SetAgeUint8(u uint8) *StudentUpdateOne {
	suo.mutation.ResetAgeUint8()
	suo.mutation.SetAgeUint8(u)
	return suo
}

// AddAgeUint8 adds u to the "age_uint8" field.
func (suo *StudentUpdateOne) AddAgeUint8(u int8) *StudentUpdateOne {
	suo.mutation.AddAgeUint8(u)
	return suo
}

// SetAgeInt16 sets the "age_int16" field.
func (suo *StudentUpdateOne) SetAgeInt16(i int16) *StudentUpdateOne {
	suo.mutation.ResetAgeInt16()
	suo.mutation.SetAgeInt16(i)
	return suo
}

// AddAgeInt16 adds i to the "age_int16" field.
func (suo *StudentUpdateOne) AddAgeInt16(i int16) *StudentUpdateOne {
	suo.mutation.AddAgeInt16(i)
	return suo
}

// SetAgeUint16 sets the "age_uint16" field.
func (suo *StudentUpdateOne) SetAgeUint16(u uint16) *StudentUpdateOne {
	suo.mutation.ResetAgeUint16()
	suo.mutation.SetAgeUint16(u)
	return suo
}

// AddAgeUint16 adds u to the "age_uint16" field.
func (suo *StudentUpdateOne) AddAgeUint16(u int16) *StudentUpdateOne {
	suo.mutation.AddAgeUint16(u)
	return suo
}

// SetAgeInt32 sets the "age_int32" field.
func (suo *StudentUpdateOne) SetAgeInt32(i int32) *StudentUpdateOne {
	suo.mutation.ResetAgeInt32()
	suo.mutation.SetAgeInt32(i)
	return suo
}

// AddAgeInt32 adds i to the "age_int32" field.
func (suo *StudentUpdateOne) AddAgeInt32(i int32) *StudentUpdateOne {
	suo.mutation.AddAgeInt32(i)
	return suo
}

// SetAgeUint32 sets the "age_uint32" field.
func (suo *StudentUpdateOne) SetAgeUint32(u uint32) *StudentUpdateOne {
	suo.mutation.ResetAgeUint32()
	suo.mutation.SetAgeUint32(u)
	return suo
}

// AddAgeUint32 adds u to the "age_uint32" field.
func (suo *StudentUpdateOne) AddAgeUint32(u int32) *StudentUpdateOne {
	suo.mutation.AddAgeUint32(u)
	return suo
}

// SetAgeInt64 sets the "age_int64" field.
func (suo *StudentUpdateOne) SetAgeInt64(i int64) *StudentUpdateOne {
	suo.mutation.ResetAgeInt64()
	suo.mutation.SetAgeInt64(i)
	return suo
}

// AddAgeInt64 adds i to the "age_int64" field.
func (suo *StudentUpdateOne) AddAgeInt64(i int64) *StudentUpdateOne {
	suo.mutation.AddAgeInt64(i)
	return suo
}

// SetAgeUint64 sets the "age_uint64" field.
func (suo *StudentUpdateOne) SetAgeUint64(u uint64) *StudentUpdateOne {
	suo.mutation.ResetAgeUint64()
	suo.mutation.SetAgeUint64(u)
	return suo
}

// AddAgeUint64 adds u to the "age_uint64" field.
func (suo *StudentUpdateOne) AddAgeUint64(u int64) *StudentUpdateOne {
	suo.mutation.AddAgeUint64(u)
	return suo
}

// SetAgeInt sets the "age_int" field.
func (suo *StudentUpdateOne) SetAgeInt(i int) *StudentUpdateOne {
	suo.mutation.ResetAgeInt()
	suo.mutation.SetAgeInt(i)
	return suo
}

// AddAgeInt adds i to the "age_int" field.
func (suo *StudentUpdateOne) AddAgeInt(i int) *StudentUpdateOne {
	suo.mutation.AddAgeInt(i)
	return suo
}

// SetAgeUint sets the "age_uint" field.
func (suo *StudentUpdateOne) SetAgeUint(u uint) *StudentUpdateOne {
	suo.mutation.ResetAgeUint()
	suo.mutation.SetAgeUint(u)
	return suo
}

// AddAgeUint adds u to the "age_uint" field.
func (suo *StudentUpdateOne) AddAgeUint(u int) *StudentUpdateOne {
	suo.mutation.AddAgeUint(u)
	return suo
}

// SetWeightFloat sets the "weight_float" field.
func (suo *StudentUpdateOne) SetWeightFloat(f float64) *StudentUpdateOne {
	suo.mutation.ResetWeightFloat()
	suo.mutation.SetWeightFloat(f)
	return suo
}

// AddWeightFloat adds f to the "weight_float" field.
func (suo *StudentUpdateOne) AddWeightFloat(f float64) *StudentUpdateOne {
	suo.mutation.AddWeightFloat(f)
	return suo
}

// SetWeightFloat32 sets the "weight_float32" field.
func (suo *StudentUpdateOne) SetWeightFloat32(f float32) *StudentUpdateOne {
	suo.mutation.ResetWeightFloat32()
	suo.mutation.SetWeightFloat32(f)
	return suo
}

// AddWeightFloat32 adds f to the "weight_float32" field.
func (suo *StudentUpdateOne) AddWeightFloat32(f float32) *StudentUpdateOne {
	suo.mutation.AddWeightFloat32(f)
	return suo
}

// SetClassID sets the "class_id" field.
func (suo *StudentUpdateOne) SetClassID(u uuid.UUID) *StudentUpdateOne {
	suo.mutation.SetClassID(u)
	return suo
}

// SetEnrollAt sets the "enroll_at" field.
func (suo *StudentUpdateOne) SetEnrollAt(t time.Time) *StudentUpdateOne {
	suo.mutation.SetEnrollAt(t)
	return suo
}

// SetStatusBool sets the "status_bool" field.
func (suo *StudentUpdateOne) SetStatusBool(b bool) *StudentUpdateOne {
	suo.mutation.SetStatusBool(b)
	return suo
}

// Mutation returns the StudentMutation object of the builder.
func (suo *StudentUpdateOne) Mutation() *StudentMutation {
	return suo.mutation
}

// Where appends a list predicates to the StudentUpdate builder.
func (suo *StudentUpdateOne) Where(ps ...predicate.Student) *StudentUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *StudentUpdateOne) Select(field string, fields ...string) *StudentUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Student entity.
func (suo *StudentUpdateOne) Save(ctx context.Context) (*Student, error) {
	suo.defaults()
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *StudentUpdateOne) SaveX(ctx context.Context) *Student {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *StudentUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *StudentUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suo *StudentUpdateOne) defaults() {
	if _, ok := suo.mutation.UpdatedAt(); !ok {
		v := student.UpdateDefaultUpdatedAt()
		suo.mutation.SetUpdatedAt(v)
	}
}

func (suo *StudentUpdateOne) sqlSave(ctx context.Context) (_node *Student, err error) {
	_spec := sqlgraph.NewUpdateSpec(student.Table, student.Columns, sqlgraph.NewFieldSpec(student.FieldID, field.TypeUint64))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Student.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, student.FieldID)
		for _, f := range fields {
			if !student.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != student.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.UpdatedAt(); ok {
		_spec.SetField(student.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := suo.mutation.Name(); ok {
		_spec.SetField(student.FieldName, field.TypeString, value)
	}
	if value, ok := suo.mutation.Age(); ok {
		_spec.SetField(student.FieldAge, field.TypeInt, value)
	}
	if value, ok := suo.mutation.AddedAge(); ok {
		_spec.AddField(student.FieldAge, field.TypeInt, value)
	}
	if value, ok := suo.mutation.AgeInt8(); ok {
		_spec.SetField(student.FieldAgeInt8, field.TypeInt8, value)
	}
	if value, ok := suo.mutation.AddedAgeInt8(); ok {
		_spec.AddField(student.FieldAgeInt8, field.TypeInt8, value)
	}
	if value, ok := suo.mutation.AgeUint8(); ok {
		_spec.SetField(student.FieldAgeUint8, field.TypeUint8, value)
	}
	if value, ok := suo.mutation.AddedAgeUint8(); ok {
		_spec.AddField(student.FieldAgeUint8, field.TypeUint8, value)
	}
	if value, ok := suo.mutation.AgeInt16(); ok {
		_spec.SetField(student.FieldAgeInt16, field.TypeInt16, value)
	}
	if value, ok := suo.mutation.AddedAgeInt16(); ok {
		_spec.AddField(student.FieldAgeInt16, field.TypeInt16, value)
	}
	if value, ok := suo.mutation.AgeUint16(); ok {
		_spec.SetField(student.FieldAgeUint16, field.TypeUint16, value)
	}
	if value, ok := suo.mutation.AddedAgeUint16(); ok {
		_spec.AddField(student.FieldAgeUint16, field.TypeUint16, value)
	}
	if value, ok := suo.mutation.AgeInt32(); ok {
		_spec.SetField(student.FieldAgeInt32, field.TypeInt32, value)
	}
	if value, ok := suo.mutation.AddedAgeInt32(); ok {
		_spec.AddField(student.FieldAgeInt32, field.TypeInt32, value)
	}
	if value, ok := suo.mutation.AgeUint32(); ok {
		_spec.SetField(student.FieldAgeUint32, field.TypeUint32, value)
	}
	if value, ok := suo.mutation.AddedAgeUint32(); ok {
		_spec.AddField(student.FieldAgeUint32, field.TypeUint32, value)
	}
	if value, ok := suo.mutation.AgeInt64(); ok {
		_spec.SetField(student.FieldAgeInt64, field.TypeInt64, value)
	}
	if value, ok := suo.mutation.AddedAgeInt64(); ok {
		_spec.AddField(student.FieldAgeInt64, field.TypeInt64, value)
	}
	if value, ok := suo.mutation.AgeUint64(); ok {
		_spec.SetField(student.FieldAgeUint64, field.TypeUint64, value)
	}
	if value, ok := suo.mutation.AddedAgeUint64(); ok {
		_spec.AddField(student.FieldAgeUint64, field.TypeUint64, value)
	}
	if value, ok := suo.mutation.AgeInt(); ok {
		_spec.SetField(student.FieldAgeInt, field.TypeInt, value)
	}
	if value, ok := suo.mutation.AddedAgeInt(); ok {
		_spec.AddField(student.FieldAgeInt, field.TypeInt, value)
	}
	if value, ok := suo.mutation.AgeUint(); ok {
		_spec.SetField(student.FieldAgeUint, field.TypeUint, value)
	}
	if value, ok := suo.mutation.AddedAgeUint(); ok {
		_spec.AddField(student.FieldAgeUint, field.TypeUint, value)
	}
	if value, ok := suo.mutation.WeightFloat(); ok {
		_spec.SetField(student.FieldWeightFloat, field.TypeFloat64, value)
	}
	if value, ok := suo.mutation.AddedWeightFloat(); ok {
		_spec.AddField(student.FieldWeightFloat, field.TypeFloat64, value)
	}
	if value, ok := suo.mutation.WeightFloat32(); ok {
		_spec.SetField(student.FieldWeightFloat32, field.TypeFloat32, value)
	}
	if value, ok := suo.mutation.AddedWeightFloat32(); ok {
		_spec.AddField(student.FieldWeightFloat32, field.TypeFloat32, value)
	}
	if value, ok := suo.mutation.ClassID(); ok {
		_spec.SetField(student.FieldClassID, field.TypeUUID, value)
	}
	if value, ok := suo.mutation.EnrollAt(); ok {
		_spec.SetField(student.FieldEnrollAt, field.TypeTime, value)
	}
	if value, ok := suo.mutation.StatusBool(); ok {
		_spec.SetField(student.FieldStatusBool, field.TypeBool, value)
	}
	_node = &Student{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{student.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
