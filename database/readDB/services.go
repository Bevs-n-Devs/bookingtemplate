package readDB

import (
	_ "embed"
	"errors"

	"github.com/Bevs-n-Devs/bookingtemplate/logs"
	_ "github.com/lib/pq"
)

/* SQL to get service info by service name and duration */

// service 1 30mins
func getService1_30minsSQL() (serviceType string, durationMins, deposit, cost int, err error) {
	if db == nil {
		logs.Logs(dbErr, "Database connection is not initialised. Could not get service info.")
		return "", 0, 0, 0, errors.New("database connection is not initialised")
	}

	query := `
	SELECT serviceType, durationMins, deposit, cost
	FROM booking_srvices
	WHERE serviceid = 1
	`

	err = db.QueryRow(query).Scan(
		&serviceType, &durationMins, &deposit, &cost,
	)

	if err != nil {
		logs.Logs(dbErr, "Could not get service info: "+err.Error())
		return "", 0, 0, 0, err
	}

	return serviceType, durationMins, deposit, cost, nil
}

// service 1 45mins
// func getService1_45minsSQL(serviceName)

// service 1 60mins
// func getService1_60minsSQL(serviceName)

/* determine which service info to return via service name and duration */
func GetServiceInfo(serviceName, serviceDuration string) (serviceType string, durationMins, deposit, cost int, err error) {
	// switch serviceName {
	// case "service 1":
	// 	switch serviceDuration {
	// 	case "30mins":
	// 		return getService1_30minsSQL()
	// 	case "45mins":
	// 		return getService1_45minsSQL()
	// 	case "60mins":
	// 		return getService1_60minsSQL()
	// 	default:
	// 		return "", 0, 0, 0, errors.New("invalid service duration")
	// 	}
	// default:
	// 	return "", 0, 0, 0, errors.New("invalid service name")
	// }
	switch serviceName {
	case "service 1":
		if serviceDuration == "30mins" {
			return getService1_30minsSQL()
		} //else if serviceDuration == "45mins" {
		// 	return getService1_45minsSQL(serviceName)
		// } else if serviceDuration == "60mins" {
		// 	return getService1_60minsSQL(serviceName)
		// } else {
		// 	return "", 0, 0, 0, errors.New("invalid service duration")
		// }
	}

	return serviceType, durationMins, deposit, cost, nil
}

/* calculate the remaining balance for each service */
