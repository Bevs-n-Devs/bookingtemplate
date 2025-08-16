package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Bevs-n-Devs/bookingtemplate/env"
	"github.com/Bevs-n-Devs/bookingtemplate/logs"
)

func StartServer() {
	logs.Logs(info, "Starting application server...")
	logs.Logs(1, "Loading environment variables...")

	err := env.LoadEnv(".env")
	if err != nil {
		logs.Logs(3, "failed to load environment variables: "+err.Error())
	}

	logs.Logs(info, "Loading HTML templates...")
	InitTemplates()

	http.Handle("/static/", http.StripPrefix("/static/", staticFiles))

	// define backend handlers
	http.HandleFunc("/", HomepageUITemplate)
	http.HandleFunc("/booking", BookingHandler)
	http.HandleFunc("/confirmbooking", ConfiormBookingHandler)
	http.HandleFunc("/customerlogin", CustomerLoginHandler)

	// define UI templates
	http.HandleFunc("/login", LoginUITemplate)
	// http.HandleFunc("/createaccount", CreateAccountHandler)
	// stripe payment page
	// manage booking page
	// admin login
	// add new service

	logs.Logs(1, "Retreiving port from environment variables...")
	httpServerPort := os.Getenv("APP_PORT")

	// start server on local machine
	if httpServerPort == "" {
		logs.Logs(warn, "No port found in environment variables, switching server on default port")
		httpServerPort = defaultPort
		logs.Logs(1, "Starting server on default port http://localhost:"+httpServerPort)
		err := http.ListenAndServe(fmt.Sprintf(":%s", httpServerPort), nil)
		if err != nil {
			logs.Logs(3, fmt.Sprintf("HTTP server failed to start: %s", err.Error()))
		}
	}

	logs.Logs(1, "Starting server on port http://localhost:"+httpServerPort)
	err = http.ListenAndServe(fmt.Sprintf(":%s", httpServerPort), nil)
	if err != nil {
		logs.Logs(3, fmt.Sprintf("HTTP server failed to start: %s", err.Error()))
	}
}
