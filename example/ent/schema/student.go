package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

type Student struct {
	ent.Schema
}

func (Student) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Annotations(entsql.WithComments(true)).
			Comment("Student name | 学生姓名"),
		field.Int16("age").
			Annotations(entsql.WithComments(true)).
			Comment("Student age | 学生年龄"),
		field.String("address").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("Student's home address | 学生家庭住址 ")}
}

func (Student) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.UUIDMixin{},
	}
}

// Edges of the Student.
func (Student) Edges() []ent.Edge {
	return nil
}
