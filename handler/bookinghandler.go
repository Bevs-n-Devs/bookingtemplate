package handler

import (
	"fmt"
	"net/http"

	"github.com/Bevs-n-Devs/bookingtemplate/logs"
)

func BookingHandler(w http.ResponseWriter, r *http.Request) {
	err := Templates.ExecuteTemplate(w, "booking.html", nil)
	if err != nil {
		logs.Logs(logErr, fmt.Sprintf("Unable to load booking page: %s", err.Error()))
		http.Error(w, "Unable to load booking page: "+err.Error(), http.StatusInternalServerError)
	}
}
