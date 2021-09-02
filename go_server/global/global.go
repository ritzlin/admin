package global

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var (
	Config config
	Logger *zap.Logger
	Db *sqlx.DB
)

func Init() {
	initConfig()
	initLogger()
	initDatabase()
}
