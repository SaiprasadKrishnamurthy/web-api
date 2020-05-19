package config

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/nats-io/nats.go"
)

var db *sql.DB
var nc *nats.Conn

// InitConfigs initializes all the necessary configs once.
func InitConfigs() {
	initDB()
	initNats()
}

func initDB() {
	dbFile := os.Getenv("DB_FILE")
	if dbFile == "" {
		dbFile = "sai.sqlite"
	}
	db, _ = sql.Open("sqlite3", dbFile)
	_, _ = db.Exec("create table if not exists echo(id text, message text)")
	db.Exec("INSERT into echo(id, message) VALUES ('1', 'Hello, Howdy')")
}

// GetPort returns the Web server port.
func GetPort() string {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8082"
	}
	return port
}

// GetDB get the database.
func GetDB() *sql.DB {
	return db
}

// GetNats gets the nats connection.
func GetNats() *nats.Conn {
	return nc
}

// APIVersion public API version.
func APIVersion() string {
	return "v1"
}

func initNats() {
	nc, _ = nats.Connect(nats.DefaultURL)
}
