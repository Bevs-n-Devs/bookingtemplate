package handler

import (
	"fmt"
	"net/http"

	"github.com/Bevs-n-Devs/bookingtemplate/database/readDB"
	"github.com/Bevs-n-Devs/bookingtemplate/database/writeDB"
	"github.com/Bevs-n-Devs/bookingtemplate/logs"
)

func ConfiormBookingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		logs.Logs(warn, fmt.Sprintf("ivalid request method: %s. Redirecting back to home page", r.Method))
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	err := r.ParseForm()
	if err != nil {
		logs.Logs(logErr, "Could not extract data from form: "+err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Extract form data username, bookingDate, bookingTime, serviceType, serviceDuration
	var deposit int
	userEmail := r.FormValue("username")
	serviceType := r.FormValue("serviceType")
	serviceDuration := r.FormValue("serviceDuration")
	bookingDate := r.FormValue("bookingDate")
	bookingTime := r.FormValue("bookingTime")

	// calculate the deposit
	switch serviceDuration {
	case "30mins":
		deposit = 10
	case "45mins":
		deposit = 15
	case "60mins":
		deposit = 20
	default:
		deposit = 0
		logs.Logs(logErr, "Invalid duration selected: "+serviceDuration)
		http.Error(w, "Invalid duration selected", http.StatusBadRequest)
		return
	}

	// check if form data [service, duration, date, time] already exsits in the database
	verified, err := readDB.VerifyBookingSQL(userEmail, serviceType, serviceDuration, bookingDate, bookingTime)
	if err != nil {
		logs.Logs(logErr, "Could not verify booking: "+err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// if the booking does not exist in the database
	if !verified {
		// get the service full payment price from the database via serviceType and serviceDuration
		// write insert db using form data and also function above, including pending variables  (depositStatus, remainingBalanceStatus) 
		// 

		err := writeDB.CreateBookingConfirmationSQL(userEmail, serviceType, serviceDuration, bookingDate, bookingTime, deposit)
		if err != nil {
			logs.Logs(logErr, "Could not create booking: "+err.Error())
			http.RedirectHandler("/", http.StatusSeeOther) // change to login page OR direct to payment page
			return
		}
	}

	// if booking already exists in the database
	if verified {
		logs.Logs(warn, fmt.Sprintf("Booking already exists in database {\"email\": \"%s\", \"service\": \"%s\", \"duration\": \"%s\", \"date\": \"%s\", \"time\": \"%s\", \"deposit\": %d}", userEmail, serviceType, serviceDuration, bookingDate, bookingTime, deposit))
		http.RedirectHandler("/", http.StatusSeeOther)
		return
	}
}
