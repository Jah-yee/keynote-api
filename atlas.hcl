env "local" {
  src = "ent://internal/ent/schema"
  dev = "postgresql://postgres:postgres@localhost:5432/keynote_atlas?sslmode=disable"
  migration {
    dir = "file://migrate/migrations"
  }
}