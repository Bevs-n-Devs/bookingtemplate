package readDB

import (
	"database/sql"
	_ "embed"

	"github.com/Bevs-n-Devs/bookingtemplate/database"
	"github.com/Bevs-n-Devs/bookingtemplate/logs"
	_ "github.com/lib/pq"
)

func VerifyBookingSQL(username, bookingDate, bookingTime, serviceType, serviceDuration string) (bool, error) {
	if database.Db == nil {
		logs.Logs(5, "Database connection is not initialised. Could not get booking confirmation.")
		return false, nil
	}

	query := `
	SELECT username, bookingDate, bookingTime, serviceType, serviceDuration, depositAmount, remainingBalance, fullPayment, depositStatus, remainingBalanceStatus
	WHERE username = $1
		AND bookingDate = $2
		AND bookingTime = $3
		AND serviceType = $4
		AND serviceDuration = $5
	`
	err := database.Db.QueryRow(query, username, bookingDate, bookingTime, serviceType, serviceDuration).Scan(
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
