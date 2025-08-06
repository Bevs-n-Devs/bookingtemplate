package handler

import (
	"html/template"
	"net/http"
)

const (
	info        = 1
	warn        = 2
	logErr      = 3
	defaultPort = "8080"
)

var (
	Templates   *template.Template
	staticFiles = http.FileServer(http.Dir("./handler/userinterface/static/"))
	depositStatus, remainingBalanceStatus string = "pending", "pending"
)
