package writeDB

import (
	_ "embed"
	"errors"

	"github.com/Bevs-n-Devs/bookingtemplate/database"
	"github.com/Bevs-n-Devs/bookingtemplate/logs"
	_ "github.com/lib/pq"
)

func CreateBookingConfirmationSQL(email, service, duration, date, serviceTime string, deposit int) error {
	if database.Db == nil {
		logs.Logs(dbErr, "Database connection is not initialised. Could not create booking confirmation.")
		return errors.New("database connection is not initialised")
	}

	query := `
	INSERT INTO BookingConfirmationTable (
		email,
		date,
		time,
		service,
		duration,
		deposit
	)
	VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := database.Db.Exec(query, email, date, serviceTime, service, duration, deposit)
	if err != nil {
		logs.Logs(dbErr, "Could not create booking confirmation: "+err.Error())
		return err
	}

	logs.Logs(logDB, "Successfully created booking confirmation. {service: "+service+", duration: "+duration+", date: "+date+", serviceTime: "+serviceTime+"}")
	return nil
}
