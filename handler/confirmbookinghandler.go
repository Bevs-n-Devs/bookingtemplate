package handler

import (
	"fmt"
	"net/http"

	"github.com/Bevs-n-Devs/bookingtemplate/database/readDB"
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
	service := r.FormValue("serviceType")
	serviceDuration := r.FormValue("serviceDuration")
	bookingDate := r.FormValue("bookingDate")
	bookingTime := r.FormValue("bookingTime")

	// check if the booking is available
	bookingAvailable, err := readDB.VerifyBookingAvailabilitySQL(bookingDate, bookingTime)
	if err != nil {
		logs.Logs(logErr, "Could not verify booking availability: "+err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// if the booking is not available
	if !bookingAvailable {
		logs.Logs(warn, "Booking is not available, redirecting user")
		http.RedirectHandler("/", http.StatusSeeOther)
		return
	}

	// get service information by serviceType and serviceDuration
	serviceType, serviceMins, serviceDeposit, serviceCost, err := readDB.GetServiceInfo(service, serviceDuration)
	if err != nil {
		logs.Logs(logErr, "Could not get service information: "+err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// calculate the remaining cost of service after the deposit is taken

	// create a booking confirmation
	// redirect page to login or signup - once user is in account they can view their bookings, make payment etc
}
