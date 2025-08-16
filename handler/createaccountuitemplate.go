package handler

import (
	"fmt"
	"net/http"

	"github.com/Bevs-n-Devs/bookingtemplate/logs"
)

func CreateAccountUITemplate(w http.ResponseWriter, r *http.Request) {
	err := Templates.ExecuteTemplate(w, "createAccount.html", nil)
	if err != nil {
		logs.Logs(logErr, fmt.Sprintf("Unable to load create account page: %s", err.Error()))
		http.Error(w, "Unable to load create account page: "+err.Error(), http.StatusInternalServerError)
	}
}
