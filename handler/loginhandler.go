package handler

import (
	"fmt"
	"net/http"

	"github.com/Bevs-n-Devs/bookingtemplate/logs"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		logs.Logs(logErr, fmt.Sprintf("ivalid request method: %s. Redirecting back to home page", r.Method))
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	err := r.ParseForm()
	if err != nil {
		logs.Logs(logErr, "Could not extract user details from form: "+err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	userPassword := r.FormValue("userPassword")

	// check if the user exists in database - customers table

	logs.Logs(info, fmt.Sprintf("Username: %s, Password: %s", username, userPassword))
	http.RedirectHandler("/", http.StatusSeeOther) // temporary
	return
}
