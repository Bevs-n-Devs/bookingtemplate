package handler

import (
	"fmt"
	"net/http"

	"github.com/Bevs-n-Devs/bookingtemplate/database/readDB"
	"github.com/Bevs-n-Devs/bookingtemplate/database/writeDB"
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
	exists, err := readDB.VerifyUserExistsSQL(username)
	if err != nil {
		logs.Logs(logErr, "Could not verify if user exists in database: "+err.Error())
		http.Error(w, "Could not verify user existence: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// if user exists in database redirect to user dashboard page
	if exists {
		logs.Logs(info, "User exists in database. Redirecting to user dashboard")
		http.Redirect(w, r, "/dashboardlogin", http.StatusSeeOther)
		return
	}

	// if user does not exist then create a new user account
	if !exists {
		logs.Logs(info, "User does not exist in database. Creating new user account")
		err = writeDB.CreateUserSQL(username, userPassword, confirmPassword)
		if err != nil {
			logs.Logs(logErr, "Could not create new user account: "+err.Error())
			http.Error(w, "Could not create new user account: "+err.Error(), http.StatusInternalServerError)
			return
		}

		logs.Logs(info, "New user account created successfully. Redirecting to dashboard login page")
		http.Redirect(w, r, "/dashboardlogin", http.StatusSeeOther)
		return
	}
}
