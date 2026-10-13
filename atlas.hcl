env "local" {
  src = "ent://internal/ent/schema"
  dev = ""
  migration {
    dir = "file://migrate/migrations"
  }
}