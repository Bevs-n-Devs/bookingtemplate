package handler

import (
	"fmt"
	"html/template"
	"os"

	"github.com/Bevs-n-Devs/bookingtemplate/logs"
)

// initialise HTML templates
func InitTemplates() {
	var err error
	Templates, err = template.ParseGlob("./handler/userinterface/*.html")
	if err != nil {
		logs.Logs(1, fmt.Sprintf("Failed to parse HTML templates: %s", err.Error()))
		os.Exit(1)
	}
}
