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
		field.String("name").Comment("名称"),
		field.Int("age"),
		field.Int32("age_int32"),
		field.Int64("age_int64"),
		field.Uint("age_uint"),
		field.Uint32("age_uint32"),
		field.Uint64("age_uint64"),
		field.Float("weight_float"),
		field.Float32("weight_float32"),
		field.UUID("class_id", uuid.UUID{}),
		field.Time("enroll_at"),
		field.Bool("status_bool"),
	}
}

func (Student) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
	}
}

// Edges of the Student.
func (Student) Edges() []ent.Edge {
	return nil
}
