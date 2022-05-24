package main

import (
	"easyStorage/app"
	"easyStorage/config"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	if config.MaxUploadFileLength != -1 {
		r.ParseMultipartForm(int64(config.MaxUploadFileLength))
	}
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		log.Println(err)
		http.Error(w, "cant read file", http.StatusInternalServerError)
		return
	}
	defer file.Close()
	telegramFile := app.File{Name: fileHeader.Filename}
	telegramFile.ReadFromIO(file)
	if err = telegramFile.Send(); err != nil {
		log.Println(err)
		http.Error(w, "cant send file to server", http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, telegramFile.ID)
}

func GetFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fileID := vars["fileID"]
	if len(fileID) == 0 {
		http.Error(w, "no file id in params", http.StatusNotFound)
		return
	}

	file, err := app.GetFile(fileID)
	if err != nil {
		log.Println(err)
		http.Error(w, "cant get file", http.StatusInternalServerError)
		return
	}
	w.Write(file.File)
}

func GetFileInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fileID := vars["fileID"]
	if len(fileID) == 0 {
		http.Error(w, "no file id in params", http.StatusNotFound)
		return
	}

	file, err := app.GetFileInfo(fileID)
	if err != nil {
		log.Println(err)
		http.Error(w, "cant get file info", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(file)

	if err != nil {
		log.Println(err)
		http.Error(w, "cant encode file info", http.StatusInternalServerError)
		return
	}
}
