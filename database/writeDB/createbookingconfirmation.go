package writeDB

import (
	_ "embed"
	"errors"

	"github.com/Bevs-n-Devs/bookingtemplate/database"
	"github.com/Bevs-n-Devs/bookingtemplate/logs"
	_ "github.com/lib/pq"
)

// creates a new booking confirmation
//
// returns an error if the booking confirmation could not be created
func CreateBookingConfirmationSQL(username, bookingDate, bookingTime, serviceType string, serviceDeposit, remainingBalance int, depositStatus, remainingBalanceStatus string, serviceMins int) error {
	if database.Db == nil {
		logs.Logs(dbErr, "Database connection is not initialised. Could not create booking confirmation.")
		return errors.New("database connection is not initialised")
	}

	query := `
	INSERT INTO booking_confirmation_table (
		username,
		booking_date,
		booking_time,
		service_type,
		service_deposit,
		remaining_balance,
		deposit_status,
		remaining_balance_status,
		booking_cancelled,
		service_mins
	)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, 'N', $9)
	`

	_, err := database.Db.Exec(query, username, bookingDate, bookingTime, serviceType, serviceDeposit, remainingBalance, depositStatus, remainingBalanceStatus, serviceMins)
	if err != nil {
		logs.Logs(dbErr, "Could not create booking confirmation: "+err.Error())
		return err
	}

	logs.Logs(logDB, "Successfully created booking confirmation.")
	return nil
}
