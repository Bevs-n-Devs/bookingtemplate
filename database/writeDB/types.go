package writeDB

const (
	warn  = 2
	logDB = 4
	dbErr = 5
)

type CreateBookingConfirmation struct {
	Email       string `json:"email"`
	Service     string `json:"service"`
	Duration    string `json:"duration"`
	Date        string `json:"date"`
	ServiceTime string `json:"serviceTime"`
	Deposit     int    `json:"deposit"`
}
