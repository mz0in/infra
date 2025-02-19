package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

const (
	DefaultKernelVersion = "vmlinux-5.10.186"
	// The Firecracker version the last tag + the short SHA (so we can build our dev previews)
	DefaultFirecrackerVersion = "v1.7.0-dev_8bb88311"
)

type Env struct {
	ent.Schema
}

func (Env) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().Immutable().SchemaType(map[string]string{dialect.Postgres: "text"}),
		field.Time("created_at").Immutable().Default(time.Now).
			Annotations(
				entsql.Default("CURRENT_TIMESTAMP"),
			),
		field.Time("updated_at").Default(time.Now),
		field.UUID("team_id", uuid.UUID{}),
		field.String("dockerfile").SchemaType(map[string]string{dialect.Postgres: "text"}),
		field.Bool("public").Annotations(entsql.Default("false")),
		field.UUID("build_id", uuid.UUID{}),
		field.Int32("build_count").Default(1),
		field.Int64("spawn_count").Default(0).Comment("Number of times the env was spawned"),
		field.Time("last_spawned_at").Optional().Comment("Timestamp of the last time the env was spawned"),
		field.Int64("vcpu"),
		field.Int64("ram_mb"),
		field.Int64("free_disk_size_mb"),
		field.Int64("total_disk_size_mb"),
		field.String("kernel_version").Default(DefaultKernelVersion),
		field.String("firecracker_version").Default(DefaultFirecrackerVersion),
	}
}

func (Env) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("team", Team.Type).Ref("envs").Unique().Field("team_id").Required(),
		edge.To("env_aliases", EnvAlias.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}

func (Env) Annotations() []schema.Annotation {
	withComments := true

	return []schema.Annotation{
		entsql.Annotation{WithComments: &withComments},
	}
}

func (Env) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Mixin{},
	}
}
