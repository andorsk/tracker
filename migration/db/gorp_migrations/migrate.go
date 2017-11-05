package gorp_migrations

type MigratorInterface interface {
	CreateTables()
}

type Migrator struct {
}
