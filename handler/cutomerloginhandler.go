package handler

import (
	"fmt"
	"net/http"

	"github.com/Bevs-n-Devs/bookingtemplate/logs"
)

func CustomerLoginHandler(w http.ResponseWriter, r *http.Request) {
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
	confirmPassword := r.FormValue("confirmPassword")

	// check if the passwords match from user input
	if userPassword != confirmPassword {
		logs.Logs(logErr, "Passwords do not match. Redirecting back to home page")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// check if the user exists in database - customers table
	
	// if user exists in database redirect to user dashboard page
	// if user does not exist in DB then redirect to create account page

	logs.Logs(info, fmt.Sprintf("Username: %s, Password: %s, Confirm Password: %s", username, userPassword, confirmPassword))
	http.Redirect(w, r, "/", http.StatusSeeOther) // temporary
}
