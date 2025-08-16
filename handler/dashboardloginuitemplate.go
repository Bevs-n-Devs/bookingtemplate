package handler

import (
	"fmt"
	"net/http"

	"github.com/Bevs-n-Devs/bookingtemplate/logs"
)

func DashboardLoginUITemplate(w http.ResponseWriter, r *http.Request) {
	err := Templates.ExecuteTemplate(w, "dashboardlogin.html", nil)
	if err != nil {
		logs.Logs(logErr, fmt.Sprintf("Unable to load dashboard login page: %s", err.Error()))
		http.Error(w, "Unable to load dashboard login page: "+err.Error(), http.StatusInternalServerError)
	}
}
