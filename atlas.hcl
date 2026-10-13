env "local" {
  src = "ent://internal/ent/schema"
  dev = "postgresql://postres:postgres@localhost:5432/keynote"
  migration {
    dir = "file://migrate/migrations"
  }
}