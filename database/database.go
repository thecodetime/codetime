package database

import (
	"github.com/ekkapob/codetime/models"
	"github.com/jmoiron/sqlx"
)

func InitDB(dataSource string) {
	db := sqlx.MustConnect("postgres", dataSource)
	models.SetDB(db)
}
