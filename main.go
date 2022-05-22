package main

import (
	"main/app"
	"main/config"
	"net/http"
)

func main() {
	app.BotsManager.Load(config.Bots)

	http.HandleFunc("/upload", Upload)
	http.HandleFunc("/get", GetFile)

	http.ListenAndServe(config.Host, nil)
}
