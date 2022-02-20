// Refer to https://gorm.io/docs/connecting_to_the_database.html

package databases

import (
	"os"
	"time"

	"github.com/daison12006013/gorvel/pkg/errors"
	"github.com/daison12006013/gorvel/pkg/facade/path"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Resolve() *gorm.DB {
	db, err := gorm.Open(*Dialector(), &gorm.Config{
		NowFunc: func() time.Time {
			utc, _ := time.LoadLocation("")
			return time.Now().In(utc)
		},
	})

	if errors.Handler("SQL connection error", err) {
		panic(err)
	}

	return db
}

func Dialector() *gorm.Dialector {
	var dialect *gorm.Dialector
	switch os.Getenv("DB_CONNECTION") {
	case "sqlite":
		dialect = SQLite()
		/*
			case "mysql":
				return MySQL()
			case "postgres":
				return PostgreSQL()
			case "sqlserver":
				return SQLServer()
			case "clickhouse":
				return ClickHouse()
		*/
	}
	return dialect
}

func SQLite() *gorm.Dialector {
	filepath := path.Load().DatabasePath(os.Getenv("DB_DATABASE"))
	dialect := sqlite.Open(filepath)
	return &dialect
}

/*
// install it using `go get gorm.io/driver/mysql`
func MySQL() *gorm.Dialector {
	dialect := mysql.New(mysql.Config{
		DriverName: os.Getenv("DB_DRIVER"),
		DSN:        os.Getenv("DB_DATABASE"),
	})
	return &dialect
}

// install it using `go get gorm.io/driver/postgres`
func PostgreSQL() *gorm.Dialector {
	dialect := postgres.New(postgres.Config{
		DriverName: os.Getenv("DB_DRIVER"),
		DSN:        os.Getenv("DB_DATABASE"),
	})
	return &dialect
}

// install it using `go get gorm.io/driver/sqlserver`
func SQLServer() *gorm.Dialector {
	dialect := sqlserver.Open(os.Getenv("DB_DATABASE"))
	return &dialect
}

// install it using `go get gorm.io/driver/clickhouse`
func ClickHouse() *gorm.Dialector {
	dialect := clickhouse.Open(os.Getenv("DB_DATABASE"))
	return &dialect
}
*/
