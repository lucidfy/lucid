package registrar

import "github.com/lucidfy/lucid/databases/migrations"

var Migrations = []interface{}{
	&migrations.UserMigration{},
}
