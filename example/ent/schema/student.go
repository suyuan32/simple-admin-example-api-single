package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/gofrs/uuid/v5"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

// Student holds the schema definition for the Student entity.
type Student struct {
	ent.Schema
}

// Fields of the Student.
func (Student) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int("age"),
		field.Int8("age_int8"),
		field.Uint8("age_uint8"),
		field.Int16("age_int16"),
		field.Uint16("age_uint16"),
		field.Int32("age_int32"),
		field.Uint32("age_uint32"),
		field.Int64("age_int64"),
		field.Uint64("age_uint64"),
		field.Int("age_int"),
		field.Uint("age_uint"),
		field.Float("weight_float"),
		field.Float32("weight_float32"),
		field.UUID("class_id", uuid.UUID{}),
		field.Uint64("teacher_id"),
		field.Time("enroll_at"),
		field.Bool("status_bool"),
		// field.JSON("info", Info{}),
	}
}

//type Info struct {
//	BirthDay time.Time
//	Address  string
//}

func (Student) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
	}
}

// Edges of the Student.
func (Student) Edges() []ent.Edge {
	return nil
}
