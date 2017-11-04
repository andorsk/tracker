package migration

import 
	// Setup the goose configuration
var migrateConf &goose.DBConf{
	MigrationsDir: config.Conf.MigrationsPath,
	Env:           "production",
	Driver: goose.DBDriver{
		Name:    "mysql",
		OpenStr: config.Conf.DBPath,
		Dialect: &goose.MySqlDialect{},
	},
}
