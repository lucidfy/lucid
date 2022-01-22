package config

import (
	"github.com/daison12006013/gorvel/controllers"
)

var Routes = map[string]interface{}{
	"/": controllers.Home,
}
