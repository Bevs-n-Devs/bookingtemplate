package readDB

import (
	"database/sql"
	_ "embed"

	"github.com/Bevs-n-Devs/bookingtemplate/database"
	"github.com/Bevs-n-Devs/bookingtemplate/logs"
	_ "github.com/lib/pq"
)

// checks if the booking date and time is available
//
// returns: true if the booking date and time is available
func VerifyBookingAvailabilitySQL(bookingDate, bookingTime string) (bool, error) {
	if database.Db == nil {
		logs.Logs(dbErr, "Database connection is not initialised. Could not get booking confirmation.")
		return false, nil
	}

	query := `
	SELECT booking_date, booking_time
	FROM  booking_confirmation_table
	WHERE booking_date = $1
		AND booking_time = $2
		AND booking_cancelled = 'N'
	`
	err := database.Db.QueryRow(query, bookingDate, bookingTime).Scan(
		&bookingDate, &bookingTime,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			logs.Logs(info, "No rows found in database. Service is available for booking")
			return true, nil
		}

		logs.Logs(dbErr, "Could not get booking confirmation: "+err.Error())
		return false, nil
	}

	logs.Logs(warn, "Service is not available for booking")
	return false, nil
}

// checks if the service has been cancelled via booking date & time
//
// returns: true if the service has been cancelled
func VerifyCancelledBookingSQL(bookingDate, bookingTime string) (bool, error) {
	if database.Db == nil {
		logs.Logs(dbErr, "Database connection is not initialised. Could not get booking confirmation.")
		return false, nil
	}

	query := `
	SELECT booking_date, booking_time
	FROM  booking_confirmation_table
	WHERE booking_date = $1
		AND booking_time = $2
		AND booking_cancelled = 'Y'
	`

	err := database.Db.QueryRow(query, bookingDate, bookingTime).Scan(
		&bookingDate, &bookingTime,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			logs.Logs(warn, "No rows found in database. Service is available for booking")
			return false, nil
		}

		logs.Logs(dbErr, "Could not get booking confirmation: "+err.Error())
		return false, nil
	}

	logs.Logs(warn, "This service has been cancelled")
	return true, nil
}

// checks if user exists
//
// returns: true if user exists
func VerifyUserExistsSQL(username string) (bool, error) {
	if database.Db == nil {
		logs.Logs(dbErr, "Database connection is not initialised. Could not get booking confirmation.")
		return false, nil
	}

	query := `
	SELECT username
	FROM booking_users
	WHERE username = $1
	`
	err := database.Db.QueryRow(query, username).Scan(&username)
	if err != nil {
		logs.Logs(dbErr, "Could verify user: "+err.Error())
		return false, nil
	}

	logs.Logs(info, "User exists in database")
	return true, nil
}
