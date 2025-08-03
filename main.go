package main

import (
	"github.com/Bevs-n-Devs/bookingtemplate/handler"
	"github.com/Bevs-n-Devs/bookingtemplate/logs"
)

func main() {
	go logs.ProcessLogs()

	go handler.StartServer()

	select {}
}
