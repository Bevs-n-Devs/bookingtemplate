package logs

import (
	"fmt"
	"log"
)

var (
	logCahnnel = make(chan string)
)

const (
	info     = "INFO: "
	warn     = "WARN: "
	logErr   = "ERROR: "
	database = "DATABASE: "
	dbError  = "DATABASE ERROR: "
)

func ProcessLogs() {
	for logMessage := range logCahnnel {
		log.Println(logMessage)
	}
}

func Logs(logType int, message string) {
	var loggedMessage string

	switch logType {
	case 1:
		loggedMessage = info + message
	case 2:
		loggedMessage = warn + message
	case 3:
		loggedMessage = logErr + message
	case 4:
		loggedMessage = database + message
	case 5:
		loggedMessage = dbError + message
	default:
		loggedMessage = fmt.Sprintf("UNKNOWN LOG TYPE: %d, MESSAGE: %s", logType, message)
	}

	logCahnnel <- loggedMessage
}
