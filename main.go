package main

import (
	"github.com/Bevs-n-Devs/bookingtemplate/database"
	"github.com/Bevs-n-Devs/bookingtemplate/handler"
	"github.com/Bevs-n-Devs/bookingtemplate/logs"
)

func main() {
	go logs.ProcessLogs()
	err := database.ConnectDB()
	if err != nil {
		logs.Logs(3, "Could not connect to database: "+err.Error())
	}

	go handler.StartServer()

	select {}
}
