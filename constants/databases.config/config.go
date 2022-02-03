package databaseconfig

import "os"

// * For SQLite
var DB_DATABASE string = os.Getenv("GOPATH") + "/src/gorvel/databases/sqlite.db"

// * For MySQL,PosgreSQL
// const DB_HOST = "localhost"
// const DB_PORT = "3306"
// const DB_USERNAME = "root"
// const DB_PASSWORD = "password"
// const DB_DATABASE = "gorvel"
