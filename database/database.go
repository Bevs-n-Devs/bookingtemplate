package database

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "embed"

	_ "github.com/lib/pq"

	"github.com/Bevs-n-Devs/bookingtemplate/env"
	"github.com/Bevs-n-Devs/bookingtemplate/logs"
)

func ConnectDB() error {
	var err error
	err = env.LoadEnv(".env")
	if err != nil {
		logs.Logs(dbErr, "failed to load database environment variables: "+err.Error())
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		logs.Logs(dbErr, "Database URL is empty!")
		return fmt.Errorf("database URL is empty")
	}

	logs.Logs(logDB, "Connecting to database...")
	Db, err = sql.Open("postgres", databaseURL) // open db connection from global db variable
	if err != nil {
		logs.Logs(dbErr, fmt.Sprintf("Unable to open database connection: %s", err.Error()))
		return err
	}
	// verify connection
	logs.Logs(logDB, "Verifying database connection...")
	if Db == nil {
		logs.Logs(dbErr, "Database connection is nil!")
		return errors.New("database connection not establioshed")
	}
	err = Db.Ping()
	if err != nil {
		logs.Logs(dbErr, fmt.Sprintf("Cannot connect to database: %s", err.Error()))
		return err
	}
	logs.Logs(logDB, "Database connection successful!")
	return nil
}

func CloseDB() error {
	if Db != nil {
		Db.Close()
		logs.Logs(logDB, "Database connection closed")
		return nil
	}
	logs.Logs(dbErr, "Database connection is not initialized. Could not close database.")
	return errors.New("database connection is not initialized")
}
