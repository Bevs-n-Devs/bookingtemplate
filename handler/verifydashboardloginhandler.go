package handler

import (
	"fmt"
	"net/http"

	"github.com/Bevs-n-Devs/bookingtemplate/database/readDB"
	"github.com/Bevs-n-Devs/bookingtemplate/logs"
)

func VerifyDashboardLoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		logs.Logs(warn, fmt.Sprintf("ivalid request method: %s. Redirecting back to home page", r.Method))
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

	// check if the user exits
	exists, err := readDB.VerifyUserExistsByEmailSQL(username, userPassword)
	if err != nil {
		logs.Logs(logErr, "Could not verify user exists: "+err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !exists {
		logs.Logs(logErr, "User does not exist. Redirecting back to home page")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	logs.Logs(info, "User exists in database. Redirecting to user dashboard")
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}
