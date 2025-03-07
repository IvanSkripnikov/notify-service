package helpers

import (
	"database/sql"
	"notify-service/models"

	_ "github.com/go-sql-driver/mysql"

	gormdb "github.com/IvanSkripnikov/go-gormdb"
	logger "github.com/IvanSkripnikov/go-logger"
)

var DB *sql.DB

func InitDatabase(config gormdb.Database) {
	gormDatabase, err := gormdb.AddMysql(models.ServiceDatabase, config)
	if err != nil {
		logger.Fatalf("Cant initialize DB: %v", err)
	}
	db, err := gormDatabase.DB()
	if err != nil {
		logger.Fatalf("Cant get DB: %v", err)
	}
	DB = db
}
