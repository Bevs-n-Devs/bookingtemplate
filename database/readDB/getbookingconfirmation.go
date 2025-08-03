package readDB

import (
	"database/sql"
	_ "embed"

	"github.com/Bevs-n-Devs/bookingtemplate/database"
	"github.com/Bevs-n-Devs/bookingtemplate/logs"
	_ "github.com/lib/pq"
)

func CheckBookingConfirmationSQL(email, service, duration, date, serviceTime string, deposit int) bool {
	if database.Db == nil {
		logs.Logs(5, "Database connection is not initialised. Could not get booking confirmation.")
		return false
	}

	query := `
	SELECT email, date, time, service, duration, deposit
	FROM BookingConfirmationTable
	WHERE email = $1
		AND service = $2
		AND duration = $3
		AND date = $4
		AND time = $5
	`

	err := database.Db.QueryRow(query, email, service, duration, date, serviceTime).Scan(&query, &email, &service, &duration, &date, &serviceTime)
	if err != nil {
		if err == sql.ErrNoRows {
			logs.Logs(warn, "No rows found in database. Could not get booking confirmation.")
			return false
		}

		logs.Logs(dbErr, "Could not get booking confirmation: "+err.Error())
		return false
	}

	logs.Logs(logDB, "Successfully got booking confirmation. {service: "+service+", duration: "+duration+", date: "+date+", serviceTime: "+serviceTime+"}")
	return true
}
