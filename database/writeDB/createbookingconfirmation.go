package writeDB

import (
	_ "embed"
	"errors"

	"github.com/Bevs-n-Devs/bookingtemplate/logs"
	_ "github.com/lib/pq"
)

// creates a new booking confirmation
//
// returns an error if the booking confirmation could not be created
func CreateBookingConfirmationSQL(username, bookingDate, bookingTime, serviceType string, depositAmount, remainingBalance int, depositStatus, remainingBalanceStatus string, serviceMins int) error {
	if db == nil {
		logs.Logs(dbErr, "Database connection is not initialised. Could not create booking confirmation.")
		return errors.New("database connection is not initialised")
	}

	query := `
	INSERT INTO booking_confirmation_table (
		username,
		bookingDate,
		bookingTime,
		serviceType,
		depositAmount,
		remainingBalance,
		depositStatus,
		remainingBalanceStatus,
		bookingCancelled,
		serviceMins
	)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, 'N', $9)
	`

	_, err := db.Exec(query, username, bookingDate, bookingTime, serviceType, depositAmount, remainingBalance, depositStatus, remainingBalanceStatus, serviceMins)
	if err != nil {
		logs.Logs(dbErr, "Could not create booking confirmation: "+err.Error())
		return err
	}

	logs.Logs(logDB, "Successfully created booking confirmation.")
	return nil
}
