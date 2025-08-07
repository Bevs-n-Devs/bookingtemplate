package readDB

import (
	_ "embed"
	"errors"

	"github.com/Bevs-n-Devs/bookingtemplate/logs"
	_ "github.com/lib/pq"
)

/* SQL to get service info by service name and duration */

// service 1, 30 minutes
//
// returns: serviceType, durationMins, deposit, cost or an error
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

// service 1, 45 minutes
//
// returns: serviceType, durationMins, deposit, cost or an error
func getService1_45minsSQL() (serviceType string, durationMins, deposit, cost int, err error) {
	if db == nil {
		logs.Logs(dbErr, "Database connection is not initialised. Could not get service info.")
		return "", 0, 0, 0, errors.New("database connection is not initialised")
	}

	query := `
	SELECT serviceType, durationMins, deposit, cost
	FROM booking_srvices
	WHERE serviceid = 2
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

// service 1, 60 minutes
//
// returns: serviceType, durationMins, deposit, cost or an error
func getService1_60minsSQL() (serviceType string, durationMins, deposit, cost int, err error) {
	if db == nil {
		logs.Logs(dbErr, "Database connection is not initialised. Could not get service info.")
		return "", 0, 0, 0, errors.New("database connection is not initialised")
	}

	query := `
	SELECT serviceType, durationMins, deposit, cost
	FROM booking_srvices
	WHERE serviceid = 3
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

// service 2, 30 minutes
//
// returns: serviceType, durationMins, deposit, cost or an error
func getService2_30minsSQL() (serviceType string, durationMins, deposit, cost int, err error) {
	if db == nil {
		logs.Logs(dbErr, "Database connection is not initialised. Could not get service info.")
		return "", 0, 0, 0, errors.New("database connection is not initialised")
	}

	query := `
	SELECT serviceType, durationMins, deposit, cost
	FROM booking_srvices
	WHERE serviceid = 4
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

// service 2, 45 minutes
//
// returns: serviceType, durationMins, deposit, cost or an error
func getService2_45minsSQL() (serviceType string, durationMins, deposit, cost int, err error) {
	if db == nil {
		logs.Logs(dbErr, "Database connection is not initialised. Could not get service info.")
		return "", 0, 0, 0, errors.New("database connection is not initialised")
	}

	query := `
	SELECT serviceType, durationMins, deposit, cost
	FROM booking_srvices
	WHERE serviceid = 5
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

// service 2, 60 minutes
//
// returns: serviceType, durationMins, deposit, cost or an error
func getService2_60minsSQL() (serviceType string, durationMins, deposit, cost int, err error) {
	if db == nil {
		logs.Logs(dbErr, "Database connection is not initialised. Could not get service info.")
		return "", 0, 0, 0, errors.New("database connection is not initialised")
	}

	query := `
	SELECT serviceType, durationMins, deposit, cost
	FROM booking_srvices
	WHERE serviceid = 6
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

// service 3, 30 minutes
//
// returns: serviceType, durationMins, deposit, cost or an error
func getService3_30minsSQL() (serviceType string, durationMins, deposit, cost int, err error) {
	if db == nil {
		logs.Logs(dbErr, "Database connection is not initialised. Could not get service info.")
		return "", 0, 0, 0, errors.New("database connection is not initialised")
	}

	query := `
	SELECT serviceType, durationMins, deposit, cost
	FROM booking_srvices
	WHERE serviceid = 7
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

// service 3, 45 minutes
//
// returns: serviceType, durationMins, deposit, cost or an error
func getService3_45minsSQL() (serviceType string, durationMins, deposit, cost int, err error) {
	if db == nil {
		logs.Logs(dbErr, "Database connection is not initialised. Could not get service info.")
		return "", 0, 0, 0, errors.New("database connection is not initialised")
	}

	query := `
	SELECT serviceType, durationMins, deposit, cost
	FROM booking_srvices
	WHERE serviceid = 8
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

// service 3, 60 minutes
//
// returns: serviceType, durationMins, deposit, cost or an error
func getService3_60minsSQL() (serviceType string, durationMins, deposit, cost int, err error) {
	if db == nil {
		logs.Logs(dbErr, "Database connection is not initialised. Could not get service info.")
		return "", 0, 0, 0, errors.New("database connection is not initialised")
	}

	query := `
	SELECT serviceType, durationMins, deposit, cost
	FROM booking_srvices
	WHERE serviceid = 9
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

/* determine which service info to return via service name and duration */
func GetServiceInfo(serviceName, serviceDuration string) (serviceType string, durationMins, deposit, cost int, err error) {
	switch serviceName {
	case "service 1":
		if serviceDuration == "30mins" {
			return getService1_30minsSQL()
		} else if serviceDuration == "45mins" {
			return getService1_45minsSQL()
		} else if serviceDuration == "60mins" {
			return getService1_60minsSQL()
		} else {
			return "", 0, 0, 0, errors.New("invalid service duration")
		}

	case "service 2":
		if serviceDuration == "30mins" {
			return getService2_30minsSQL()
		} else if serviceDuration == "45mins" {
			return getService2_45minsSQL()
		} else if serviceDuration == "60mins" {
			return getService2_60minsSQL()
		} else {
			return "", 0, 0, 0, errors.New("invalid service duration")
		}

	case "service 3":
		if serviceDuration == "30mins" {
			return getService3_30minsSQL()
		} else if serviceDuration == "45mins" {
			return getService3_45minsSQL()
		} else if serviceDuration == "60mins" {
			return getService3_60minsSQL()
		} else {
			return "", 0, 0, 0, errors.New("invalid service duration")
		}
	}

	return serviceType, durationMins, deposit, cost, nil
}
