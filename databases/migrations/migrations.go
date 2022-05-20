package migrations

import (
	"github.com/lucidfy/lucid/app/models/users"
)

var AutoMigrate = []interface{}{
	&users.Model{},
}
