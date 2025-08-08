package database

import (
	"database/sql"
)

const (
	warn  = 2
	logDB = 4
	dbErr = 5
)

var (
	db *sql.DB
)
