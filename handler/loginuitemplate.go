package handler

import (
	"fmt"
	"net/http"

	"github.com/Bevs-n-Devs/bookingtemplate/logs"
)

func LoginUITemplate(w http.ResponseWriter, r *http.Request) {
	err := Templates.ExecuteTemplate(w, "login.html", nil)
	if err != nil {
		logs.Logs(logErr, fmt.Sprintf("Unable to load login page: %s", err.Error()))
		http.Error(w, "Unable to load login page: "+err.Error(), http.StatusInternalServerError)
	}
}
