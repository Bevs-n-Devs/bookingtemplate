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
// returns: serviceType, serviceMins, serviceDeposit, serviceDepositLink, serviceBalance, serviceBalanceLink, serviceCost or an error
func getService1_30minsSQL() (serviceType string, serviceMins, serviceDeposit int, serviceDepositLink string, serviceBalance int, serviceBalanceLink string, serviceCost int, err error) {
	if database.Db == nil {
		logs.Logs(dbErr, "Database connection is not initialised. Could not get service info.")
		return "", 0, 0, "", 0, "", 0, errors.New("database connection is not initialised")
	}

	// this query can be storeed in environment variables for later
	query := `
	SELECT service_type, service_mins, service_deposit, service_deposit_link, service_balance, service_balance_link, service_cost
	FROM booking_services
	WHERE service_id = 1
	`

	err = database.Db.QueryRow(query).Scan(
		&serviceType, &serviceMins, &serviceDeposit, &serviceDepositLink, &serviceBalance, &serviceBalanceLink, &serviceCost,
	)

	if err != nil {
		logs.Logs(dbErr, "Could not get service info: "+err.Error())
		return "", 0, 0, "", 0, "", 0, err
	}

	return serviceType, serviceMins, serviceDeposit, serviceDepositLink, serviceBalance, serviceBalanceLink, serviceCost, nil
}

// service 1, 45 minutes
//
// returns: serviceType, serviceMins, serviceDeposit, serviceDepositLink, serviceBalance, serviceBalanceLink, serviceCost or an error
func getService1_45minsSQL() (serviceType string, serviceMins, serviceDeposit int, serviceDepositLink string, serviceBalance int, serviceBalanceLink string, serviceCost int, err error) {
	if database.Db == nil {
		logs.Logs(dbErr, "Database connection is not initialised. Could not get service info.")
		return "", 0, 0, "", 0, "", 0, errors.New("database connection is not initialised")
	}

	query := `
	SELECT service_type, service_mins, service_deposit, service_deposit_link, service_balance, service_balance_link, service_cost
	FROM booking_services
	WHERE service_id = 2
	`

	err = database.Db.QueryRow(query).Scan(
		&serviceType, &serviceMins, &serviceDeposit, &serviceDepositLink, &serviceBalance, &serviceBalanceLink, &serviceCost,
	)

	if err != nil {
		logs.Logs(dbErr, "Could not get service info: "+err.Error())
		return "", 0, 0, "", 0, "", 0, err
	}

	return serviceType, serviceMins, serviceDeposit, serviceDepositLink, serviceBalance, serviceBalanceLink, serviceCost, nil
}

// service 1, 60 minutes
//
// returns: serviceType, serviceMins, serviceDeposit, serviceDepositLink, serviceBalance, serviceBalanceLink, serviceCost or an error
func getService1_60minsSQL() (serviceType string, serviceMins, serviceDeposit int, serviceDepositLink string, serviceBalance int, serviceBalanceLink string, serviceCost int, err error) {
	if database.Db == nil {
		logs.Logs(dbErr, "Database connection is not initialised. Could not get service info.")
		return "", 0, 0, "", 0, "", 0, errors.New("database connection is not initialised")
	}

	query := `
	SELECT service_type, service_mins, service_deposit, service_deposit_link, service_balance, service_balance_link, service_cost
	FROM booking_services
	WHERE service_id = 3
	`

	err = database.Db.QueryRow(query).Scan(
		&serviceType, &serviceMins, &serviceDeposit, &serviceDepositLink, &serviceBalance, &serviceBalanceLink, &serviceCost,
	)

	if err != nil {
		logs.Logs(dbErr, "Could not get service info: "+err.Error())
		return "", 0, 0, "", 0, "", 0, err
	}

	return serviceType, serviceMins, serviceDeposit, serviceDepositLink, serviceBalance, serviceBalanceLink, serviceCost, nil
}

// service 2, 30 minutes
//
// returns: serviceType, serviceMins, serviceDeposit, serviceDepositLink, serviceBalance, serviceBalanceLink, serviceCost or an error
func getService2_30minsSQL() (serviceType string, serviceMins, serviceDeposit int, serviceDepositLink string, serviceBalance int, serviceBalanceLink string, serviceCost int, err error) {
	if database.Db == nil {
		logs.Logs(dbErr, "Database connection is not initialised. Could not get service info.")
		return "", 0, 0, "", 0, "", 0, errors.New("database connection is not initialised")
	}

	query := `
	SELECT service_type, service_mins, service_deposit, service_deposit_link, service_balance, service_balance_link, service_cost
	FROM booking_services
	WHERE service_id = 4
	`

	err = database.Db.QueryRow(query).Scan(
		&serviceType, &serviceMins, &serviceDeposit, &serviceDepositLink, &serviceBalance, &serviceBalanceLink, &serviceCost,
	)

	if err != nil {
		logs.Logs(dbErr, "Could not get service info: "+err.Error())
		return "", 0, 0, "", 0, "", 0, err
	}

	return serviceType, serviceMins, serviceDeposit, serviceDepositLink, serviceBalance, serviceBalanceLink, serviceCost, nil
}

// service 2, 45 minutes
//
// returns: serviceType, serviceMins, serviceDeposit, serviceDepositLink, serviceBalance, serviceBalanceLink, serviceCost or an error
func getService2_45minsSQL() (serviceType string, serviceMins, serviceDeposit int, serviceDepositLink string, serviceBalance int, serviceBalanceLink string, serviceCost int, err error) {
	if database.Db == nil {
		logs.Logs(dbErr, "Database connection is not initialised. Could not get service info.")
		return "", 0, 0, "", 0, "", 0, errors.New("database connection is not initialised")
	}

	query := `
	SELECT service_type, service_mins, service_deposit, service_deposit_link, service_balance, service_balance_link, service_cost
	FROM booking_services
	WHERE service_id = 5
	`

	err = database.Db.QueryRow(query).Scan(
		&serviceType, &serviceMins, &serviceDeposit, &serviceDepositLink, &serviceBalance, &serviceBalanceLink, &serviceCost,
	)

	if err != nil {
		logs.Logs(dbErr, "Could not get service info: "+err.Error())
		return "", 0, 0, "", 0, "", 0, err
	}

	return serviceType, serviceMins, serviceDeposit, serviceDepositLink, serviceBalance, serviceBalanceLink, serviceCost, nil
}

// service 2, 60 minutes
//
// returns: serviceType, serviceMins, serviceDeposit, serviceDepositLink, serviceBalance, serviceBalanceLink, serviceCost or an error
func getService2_60minsSQL() (serviceType string, serviceMins, serviceDeposit int, serviceDepositLink string, serviceBalance int, serviceBalanceLink string, serviceCost int, err error) {
	if database.Db == nil {
		logs.Logs(dbErr, "Database connection is not initialised. Could not get service info.")
		return "", 0, 0, "", 0, "", 0, errors.New("database connection is not initialised")
	}

	query := `
	SELECT service_type, service_mins, service_deposit, service_deposit_link, service_balance, service_balance_link, service_cost
	FROM booking_services
	WHERE service_id = 6
	`

	err = database.Db.QueryRow(query).Scan(
		&serviceType, &serviceMins, &serviceDeposit, &serviceDepositLink, &serviceBalance, &serviceBalanceLink, &serviceCost,
	)

	if err != nil {
		logs.Logs(dbErr, "Could not get service info: "+err.Error())
		return "", 0, 0, "", 0, "", 0, err
	}

	return serviceType, serviceMins, serviceDeposit, serviceDepositLink, serviceBalance, serviceBalanceLink, serviceCost, nil
}

// service 3, 30 minutes
//
// returns: serviceType, serviceMins, serviceDeposit, serviceDepositLink, serviceBalance, serviceBalanceLink, serviceCost or an error
func getService3_30minsSQL() (serviceType string, serviceMins, serviceDeposit int, serviceDepositLink string, serviceBalance int, serviceBalanceLink string, serviceCost int, err error) {
	if database.Db == nil {
		logs.Logs(dbErr, "Database connection is not initialised. Could not get service info.")
		return "", 0, 0, "", 0, "", 0, errors.New("database connection is not initialised")
	}

	query := `
	SELECT service_type, service_mins, service_deposit, service_deposit_link, service_balance, service_balance_link, service_cost
	FROM booking_services
	WHERE service_id = 7
	`

	err = database.Db.QueryRow(query).Scan(
		&serviceType, &serviceMins, &serviceDeposit, &serviceDepositLink, &serviceBalance, &serviceBalanceLink, &serviceCost,
	)

	if err != nil {
		logs.Logs(dbErr, "Could not get service info: "+err.Error())
		return "", 0, 0, "", 0, "", 0, err
	}

	return serviceType, serviceMins, serviceDeposit, serviceDepositLink, serviceBalance, serviceBalanceLink, serviceCost, nil
}

// service 3, 45 minutes
//
// returns: serviceType, serviceMins, serviceDeposit, serviceDepositLink, serviceBalance, serviceBalanceLink, serviceCost or an error
func getService3_45minsSQL() (serviceType string, serviceMins, serviceDeposit int, serviceDepositLink string, serviceBalance int, serviceBalanceLink string, serviceCost int, err error) {
	if database.Db == nil {
		logs.Logs(dbErr, "Database connection is not initialised. Could not get service info.")
		return "", 0, 0, "", 0, "", 0, errors.New("database connection is not initialised")
	}

	query := `
	SELECT service_type, service_mins, service_deposit, service_deposit_link, service_balance, service_balance_link, service_cost
	FROM booking_services
	WHERE service_id = 8
	`

	err = database.Db.QueryRow(query).Scan(
		&serviceType, &serviceMins, &serviceDeposit, &serviceDepositLink, &serviceBalance, &serviceBalanceLink, &serviceCost,
	)

	if err != nil {
		logs.Logs(dbErr, "Could not get service info: "+err.Error())
		return "", 0, 0, "", 0, "", 0, err
	}

	return serviceType, serviceMins, serviceDeposit, serviceDepositLink, serviceBalance, serviceBalanceLink, serviceCost, nil
}

// service 3, 60 minutes
//
// returns: serviceType, serviceMins, serviceDeposit, serviceDepositLink, serviceBalance, serviceBalanceLink, serviceCost or an error
func getService3_60minsSQL() (serviceType string, serviceMins, serviceDeposit int, serviceDepositLink string, serviceBalance int, serviceBalanceLink string, serviceCost int, err error) {
	if database.Db == nil {
		logs.Logs(dbErr, "Database connection is not initialised. Could not get service info.")
		return "", 0, 0, "", 0, "", 0, errors.New("database connection is not initialised")
	}

	query := `
	SELECT service_type, service_mins, service_deposit, service_deposit_link, service_balance, service_balance_link, service_cost
	FROM booking_services
	WHERE service_id = 9
	`

	err = database.Db.QueryRow(query).Scan(
		&serviceType, &serviceMins, &serviceDeposit, &serviceDepositLink, &serviceBalance, &serviceBalanceLink, &serviceCost,
	)

	if err != nil {
		logs.Logs(dbErr, "Could not get service info: "+err.Error())
		return "", 0, 0, "", 0, "", 0, err
	}

	return serviceType, serviceMins, serviceDeposit, serviceDepositLink, serviceBalance, serviceBalanceLink, serviceCost, nil
}

/* determine which service info to return via service name and duration */
func GetServiceInfo(serviceName, serviceDuration string) (serviceType string, serviceMins, serviceDeposit int, serviceDepositLink string, serviceBalance int, serviceBalanceLink string, serviceCost int, err error) {
	switch serviceName {
	case "service 1":
		switch serviceDuration {
		case "30mins":
			return getService1_30minsSQL()
		case "45mins":
			return getService1_45minsSQL()
		case "60mins":
			return getService1_60minsSQL()
		default:
			return "", 0, 0, "", 0, "", 0, errors.New("invalid service duration")
		}

	case "service 2":
		switch serviceDuration {
		case "30mins":
			return getService2_30minsSQL()
		case "45mins":
			return getService2_45minsSQL()
		case "60mins":
			return getService2_60minsSQL()
		default:
			return "", 0, 0, "", 0, "", 0, errors.New("invalid service duration")
		}

	case "service 3":
		switch serviceDuration {
		case "30mins":
			return getService3_30minsSQL()
		case "45mins":
			return getService3_45minsSQL()
		case "60mins":
			return getService3_60minsSQL()
		default:
			return "", 0, 0, "", 0, "", 0, errors.New("invalid service duration")
		}
	}

	return serviceType, serviceMins, serviceDeposit, serviceDepositLink, serviceBalance, serviceBalanceLink, serviceCost, nil
}
