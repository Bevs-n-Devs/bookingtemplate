package readDB

import (
	"database/sql"
	_ "embed"

	"github.com/Bevs-n-Devs/bookingtemplate/logs"
	_ "github.com/lib/pq"
)

func VerifyBookingSQL(username, bookingDate, bookingTime, serviceType, serviceDuration string) (bool, error) {
	if db == nil {
		logs.Logs(dbErr, "Database connection is not initialised. Could not get booking confirmation.")
		return false, nil
	}

	query := `
	SELECT username, bookingDate, bookingTime, serviceType, serviceDuration, depositAmount, remainingBalance, fullPayment, depositStatus, remainingBalanceStatus
	FROM  booking_confirmation_table
	WHERE username = $1
		AND bookingDate = $2
		AND bookingTime = $3
		AND serviceType = $4
		AND serviceDuration = $5
	`
	err := db.QueryRow(query, username, bookingDate, bookingTime, serviceType, serviceDuration).Scan(
		&query, &username, &bookingDate, &bookingTime, &serviceType, &serviceDuration,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			logs.Logs(warn, "No rows found in database. Could not get booking confirmation.")
			return false, nil
		}

		logs.Logs(dbErr, "Could not get booking confirmation: "+err.Error())
		return false, nil
	}

	return true, nil
}

// check if the date and time is available - not booked by someone else
func CheckAvailabilitySQL(bookingDate, bookingTime string) (bool, error) {
	if db == nil {
		logs.Logs(dbErr, "Database connection is not initialised. Could not get booking confirmation.")
		return false, nil
	}

	query := `
	SELECT bookingDate, bookingTime
	FROM  booking_confirmation_table
	WHERE bookingDate = $1
		AND bookingTime = $2
	`
	err := db.QueryRow(query, bookingDate, bookingTime).Scan(
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

// check if the service has been cancelled or not
func CheckCancelledSQL(bookingDate, bookingTime string) (bool, error) {
	if db == nil {
		logs.Logs(dbErr, "Database connection is not initialised. Could not get booking confirmation.")
		return false, nil
	}

	query := `
	SELECT bookingDate, bookingTime
	FROM  booking_confirmation_table
	WHERE bookingDate = $1
		AND bookingTime = $2
		AND cancelled = 'Y'
	`

	err := db.QueryRow(query, bookingDate, bookingTime).Scan(
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

//
