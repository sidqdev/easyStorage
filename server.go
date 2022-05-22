package main

import (
	"fmt"
	"main/app"
	"net/http"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	// r.ParseMultipartForm(32 << 20) // limit your max input length!
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "cant read file", http.StatusInternalServerError)
		return
	}
	defer file.Close()
	telegramFile := app.File{FileName: fileHeader.Filename}
	telegramFile.ReadFromIO(file)
	err = telegramFile.Send()
	if err != nil {
		http.Error(w, "cant send file to server", http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, telegramFile.FileID)
}

func GetFile(w http.ResponseWriter, r *http.Request) {
	fileID := r.URL.Query().Get("fileID")
	if fileID == "" {
		http.Error(w, "no file id in params", http.StatusNotFound)
		return
	}

	file, err := app.GetFile(fileID)
	if err != nil {
		http.Error(w, "cant get file", http.StatusInternalServerError)
		return
	}
	w.Write(file.File)
}

func GetFileInfo(w http.ResponseWriter, r *http.Request) {

}
