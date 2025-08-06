package database

import (
	"database/sql"
)

const (
	warn  = 2
	logDB = 4
	dbErr = 5
)

// change this so its NOT global
var (
	Db *sql.DB
)
