package readDB

import (
	_ "embed"
	"errors"

	"github.com/Bevs-n-Devs/bookingtemplate/database"
	"github.com/Bevs-n-Devs/bookingtemplate/logs"
	_ "github.com/lib/pq"
)

/* SQL to get service info by service name and duration */

// service 1, 30 minutes
//
// returns: serviceType, serviceMins, serviceDeposit, serviceCost or an error
func getService1_30minsSQL() (serviceType string, serviceMins, serviceDeposit, serviceCost int, err error) {
	if database.Db == nil {
		logs.Logs(dbErr, "Database connection is not initialised. Could not get service info.")
		return "", 0, 0, 0, errors.New("database connection is not initialised")
	}

	query := `
	SELECT service_type, service_mins, service_deposit, servcie_cost
	FROM booking_srvices
	WHERE serviceid = 10
	`

	err = database.Db.QueryRow(query).Scan(
		&serviceType, &serviceMins, &serviceDeposit, &serviceCost,
	)

	if err != nil {
		logs.Logs(dbErr, "Could not get service info: "+err.Error())
		return "", 0, 0, 0, err
	}

	return serviceType, serviceMins, serviceDeposit, serviceCost, nil
}

// service 1, 45 minutes
//
// returns: serviceType, serviceMins, serviceDeposit, serviceCost or an error
func getService1_45minsSQL() (serviceType string, serviceMins, serviceDeposit, serviceCost int, err error) {
	if database.Db == nil {
		logs.Logs(dbErr, "Database connection is not initialised. Could not get service info.")
		return "", 0, 0, 0, errors.New("database connection is not initialised")
	}

	query := `
	SELECT service_type, service_mins, service_deposit, servcie_cost
	FROM booking_srvices
	WHERE serviceid = 11
	`

	err = database.Db.QueryRow(query).Scan(
		&serviceType, &serviceMins, &serviceDeposit, &serviceCost,
	)

	if err != nil {
		logs.Logs(dbErr, "Could not get service info: "+err.Error())
		return "", 0, 0, 0, err
	}

	return serviceType, serviceMins, serviceDeposit, serviceCost, nil
}

// service 1, 60 minutes
//
// returns: serviceType, serviceMins, serviceDeposit, serviceCost or an error
func getService1_60minsSQL() (serviceType string, serviceMins, serviceDeposit, serviceCost int, err error) {
	if database.Db == nil {
		logs.Logs(dbErr, "Database connection is not initialised. Could not get service info.")
		return "", 0, 0, 0, errors.New("database connection is not initialised")
	}

	query := `
	SELECT service_type, service_mins, service_deposit, servcie_cost
	FROM booking_srvices
	WHERE serviceid = 12
	`

	err = database.Db.QueryRow(query).Scan(
		&serviceType, &serviceMins, &serviceDeposit, &serviceCost,
	)

	if err != nil {
		logs.Logs(dbErr, "Could not get service info: "+err.Error())
		return "", 0, 0, 0, err
	}

	return serviceType, serviceMins, serviceDeposit, serviceCost, nil
}

// service 2, 30 minutes
//
// returns: serviceType, serviceMins, serviceDeposit, serviceCost or an error
func getService2_30minsSQL() (serviceType string, serviceMins, serviceDeposit, serviceCost int, err error) {
	if database.Db == nil {
		logs.Logs(dbErr, "Database connection is not initialised. Could not get service info.")
		return "", 0, 0, 0, errors.New("database connection is not initialised")
	}

	query := `
	SELECT service_type, service_mins, service_deposit, servcie_cost
	FROM booking_srvices
	WHERE serviceid = 13
	`

	err = database.Db.QueryRow(query).Scan(
		&serviceType, &serviceMins, &serviceDeposit, &serviceCost,
	)

	if err != nil {
		logs.Logs(dbErr, "Could not get service info: "+err.Error())
		return "", 0, 0, 0, err
	}

	return serviceType, serviceMins, serviceDeposit, serviceCost, nil
}

// service 2, 45 minutes
//
// returns: serviceType, serviceMins, serviceDeposit, serviceCost or an error
func getService2_45minsSQL() (serviceType string, serviceMins, serviceDeposit, serviceCost int, err error) {
	if database.Db == nil {
		logs.Logs(dbErr, "Database connection is not initialised. Could not get service info.")
		return "", 0, 0, 0, errors.New("database connection is not initialised")
	}

	query := `
	SELECT service_type, service_mins, service_deposit, servcie_cost
	FROM booking_srvices
	WHERE serviceid = 14
	`

	err = database.Db.QueryRow(query).Scan(
		&serviceType, &serviceMins, &serviceDeposit, &serviceCost,
	)

	if err != nil {
		logs.Logs(dbErr, "Could not get service info: "+err.Error())
		return "", 0, 0, 0, err
	}

	return serviceType, serviceMins, serviceDeposit, serviceCost, nil
}

// service 2, 60 minutes
//
// returns: serviceType, serviceMins, serviceDeposit, serviceCost or an error
func getService2_60minsSQL() (serviceType string, serviceMins, serviceDeposit, serviceCost int, err error) {
	if database.Db == nil {
		logs.Logs(dbErr, "Database connection is not initialised. Could not get service info.")
		return "", 0, 0, 0, errors.New("database connection is not initialised")
	}

	query := `
	SELECT service_type, service_mins, service_deposit, servcie_cost
	FROM booking_srvices
	WHERE serviceid = 15
	`

	err = database.Db.QueryRow(query).Scan(
		&serviceType, &serviceMins, &serviceDeposit, &serviceCost,
	)

	if err != nil {
		logs.Logs(dbErr, "Could not get service info: "+err.Error())
		return "", 0, 0, 0, err
	}

	return serviceType, serviceMins, serviceDeposit, serviceCost, nil
}

// service 3, 30 minutes
//
// returns: serviceType, serviceMins, serviceDeposit, serviceCost or an error
func getService3_30minsSQL() (serviceType string, serviceMins, serviceDeposit, serviceCost int, err error) {
	if database.Db == nil {
		logs.Logs(dbErr, "Database connection is not initialised. Could not get service info.")
		return "", 0, 0, 0, errors.New("database connection is not initialised")
	}

	query := `
	SELECT service_type, service_mins, service_deposit, servcie_cost
	FROM booking_srvices
	WHERE serviceid = 16
	`

	err = database.Db.QueryRow(query).Scan(
		&serviceType, &serviceMins, &serviceDeposit, &serviceCost,
	)

	if err != nil {
		logs.Logs(dbErr, "Could not get service info: "+err.Error())
		return "", 0, 0, 0, err
	}

	return serviceType, serviceMins, serviceDeposit, serviceCost, nil
}

// service 3, 45 minutes
//
// returns: serviceType, serviceMins, serviceDeposit, serviceCost or an error
func getService3_45minsSQL() (serviceType string, serviceMins, serviceDeposit, serviceCost int, err error) {
	if database.Db == nil {
		logs.Logs(dbErr, "Database connection is not initialised. Could not get service info.")
		return "", 0, 0, 0, errors.New("database connection is not initialised")
	}

	query := `
	SELECT service_type, service_mins, service_deposit, servcie_cost
	FROM booking_srvices
	WHERE serviceid = 17
	`

	err = database.Db.QueryRow(query).Scan(
		&serviceType, &serviceMins, &serviceDeposit, &serviceCost,
	)

	if err != nil {
		logs.Logs(dbErr, "Could not get service info: "+err.Error())
		return "", 0, 0, 0, err
	}

	return serviceType, serviceMins, serviceDeposit, serviceCost, nil
}

// service 3, 60 minutes
//
// returns: serviceType, serviceMins, serviceDeposit, serviceCost or an error
func getService3_60minsSQL() (serviceType string, serviceMins, serviceDeposit, serviceCost int, err error) {
	if database.Db == nil {
		logs.Logs(dbErr, "Database connection is not initialised. Could not get service info.")
		return "", 0, 0, 0, errors.New("database connection is not initialised")
	}

	query := `
	SELECT service_type, service_mins, service_deposit, servcie_cost
	FROM booking_srvices
	WHERE serviceid = 18
	`

	err = database.Db.QueryRow(query).Scan(
		&serviceType, &serviceMins, &serviceDeposit, &serviceCost,
	)

	if err != nil {
		logs.Logs(dbErr, "Could not get service info: "+err.Error())
		return "", 0, 0, 0, err
	}

	return serviceType, serviceMins, serviceDeposit, serviceCost, nil
}

/* determine which service info to return via service name and duration */
func GetServiceInfo(serviceName, serviceDuration string) (serviceType string, serviceMins, serviceDeposit, serviceCost int, err error) {
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

	return serviceType, serviceMins, serviceDeposit, serviceCost, nil
}
