package handler

import (
	"fmt"
	"net/http"

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

	// Extract form data
	var deposit int
	userEmail := r.FormValue("email")
	service := r.FormValue("serviceName")
	duration := r.FormValue("duration")
	date := r.FormValue("date")
	serviceTime := r.FormValue("serviceTime")

	// calculate the deposit
	switch duration {
	case "30mins":
		deposit = 10
	case "45mins":
		deposit = 15
	case "60mins":
		deposit = 20
	default:
		deposit = 0
		logs.Logs(logErr, "Invalid duration selected: "+duration)
		http.Error(w, "Invalid duration selected", http.StatusBadRequest)
		return
	}

	// check if form data [service, duration, date, time] already exsits in the database

	// add the booking to the database if it does not exist

	logs.Logs(1, fmt.Sprintf("{\"email\": \"%s\", \"service\": \"%s\", \"duration\": \"%s\", \"date\": \"%s\", \"time\": \"%s\", \"deposit\": %d}", userEmail, service, duration, date, serviceTime, deposit))
	http.RedirectHandler("/", http.StatusSeeOther)
}
