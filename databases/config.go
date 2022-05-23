// Refer to https://gorm.io/docs/connecting_to_the_database.html

package databases

import (
	"os"
	"time"

	"github.com/lucidfy/lucid/pkg/errors"
	"github.com/lucidfy/lucid/pkg/facade/path"
	"gorm.io/driver/clickhouse"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
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

func Close(db *gorm.DB) {
	d, _ := db.DB()
	d.Close()
}

func Dialector() *gorm.Dialector {
	var dialect *gorm.Dialector
	switch os.Getenv("DB_CONNECTION") {
	case "sqlite":
		dialect = SQLite()
	case "mysql":
		return MySQL()
	case "postgres":
		return PostgreSQL()
	case "sqlserver":
		return SQLServer()
	case "clickhouse":
		return ClickHouse()
	}
	return dialect
}

func SQLite() *gorm.Dialector {
	filepath := path.Load().DatabasePath(os.Getenv("DB_DATABASE"))
	dialect := sqlite.Open(filepath)
	return &dialect
}

func MySQL() *gorm.Dialector {
	dialect := mysql.New(mysql.Config{
		DriverName: os.Getenv("DB_DRIVER"),
		DSN:        os.Getenv("DB_DATABASE"),
	})
	return &dialect
}

func PostgreSQL() *gorm.Dialector {
	dialect := postgres.New(postgres.Config{
		DriverName: os.Getenv("DB_DRIVER"),
		DSN:        os.Getenv("DB_DATABASE"),
	})
	return &dialect
}

func SQLServer() *gorm.Dialector {
	dialect := sqlserver.Open(os.Getenv("DB_DATABASE"))
	return &dialect
}

func ClickHouse() *gorm.Dialector {
	dialect := clickhouse.Open(os.Getenv("DB_DATABASE"))
	return &dialect
}
