package global

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func initDatabase() {
	db, err := sqlx.Open("postgres", Config.Database.Connection)
	if err != nil {
		panic("failed to connect to database: " + Config.Database.Connection)
	}
	Db = db
}
