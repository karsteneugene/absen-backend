package main

import (
	"github.com/karsteneugene/absen-backend/api"
	"github.com/karsteneugene/absen-backend/settings"
)

func init() {
	settings.LoadEnvVariables()
	settings.ConnectToDB()
}

func main() {
	api.Api()
}
