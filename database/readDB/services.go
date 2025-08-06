package readDB

import (
	_ "embed"

	_ "github.com/lib/pq"
)

// get services by service service name and service duration from booking_services
// returns the serviceType, durationmins, deposit & cost
// this data will be passed to payments package to calculate remaining balance for each user.
