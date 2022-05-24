package main

import (
	"easyStorage/app"
	"easyStorage/config"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	err := config.Load()
	if err != nil {
		log.Println(err)
	}
	app.BotsManager.Load(config.Bots)

	router := mux.NewRouter()
	router.HandleFunc("/get/{fileID}", GetFile).Methods("GET")
	router.HandleFunc("/getInfo/{fileID}", GetFileInfo).Methods("GET")
	router.HandleFunc("/upload", Upload).Methods("POST")

	http.Handle("/", router)

	log.Println(config.Host)
	http.ListenAndServe(config.Host, nil)
}
