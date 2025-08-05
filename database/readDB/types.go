package readDB

const (
	warn  = 2
	logDB = 4
	dbErr = 5
)

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
