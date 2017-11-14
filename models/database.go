package models

import (
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func SetDB(sqlxDb *sqlx.DB) { db = sqlxDb }
