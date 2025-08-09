package readDB

const (
	info  = 1
	warn  = 2
	logDB = 4
	dbErr = 5
)

// var (
// 	db *sql.DB
// )

type GetBookingConfirmation struct {
	UserName               string `json:"username"`
	BookingDate            string `json:"bookingDate"`
	BookingTime            string `json:"bookingTime"`
	ServiceType            string `json:"serviceType"`
	ServiceDuration        string `json:"serviceDuration"`
	DepositAmount          int    `json:"depositAmount"`
	RemainingBalance       int    `json:"remainingBalance"`
	FullPayment            int    `json:"fullPayment"`
	DepositStatus          string `json:"depositStatus"`
	RemainingBalanceStatus string `json:"remainingBalanceStatus"`
}
