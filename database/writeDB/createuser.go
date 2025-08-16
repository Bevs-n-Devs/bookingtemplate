package writeDB

import (
	_ "embed"
	"errors"

	"github.com/Bevs-n-Devs/bookingtemplate/database"
	"github.com/Bevs-n-Devs/bookingtemplate/logs"
	_ "github.com/lib/pq"
)

func CreateUserSQL(username, userPassword, confirmPassword string) error {
	if database.Db == nil {
		logs.Logs(dbErr, "Database connection is not initialised. Could not create user.")
		return errors.New("database connection is not initialised")
	}

	query := `
	INSERT INTO user_table (
		username,
		user_password,
		confirm_password
	)
	VALUES ($1, $2, $3)
	`

	_, err := database.Db.Exec(query, username, userPassword, confirmPassword)
	if err != nil {
		logs.Logs(dbErr, "Could not create user: "+err.Error())
		return err
	}

	logs.Logs(logDB, "Successfully created user.")
	return nil
}