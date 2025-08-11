package handler

import (
	"fmt"
	"net/http"

	"github.com/Bevs-n-Devs/bookingtemplate/logs"
)

func HomepageUITemplate(w http.ResponseWriter, r *http.Request) {
	err := Templates.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		logs.Logs(logErr, fmt.Sprintf("Unable to load home page: %s", err.Error()))
		http.Error(w, "Unable to load home page: "+err.Error(), http.StatusInternalServerError)
	}
}
