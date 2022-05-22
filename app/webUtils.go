package app

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/google/uuid"
)

type ValueData struct {
	Key   string
	Value string
}

type FileData struct {
	FileName string
	FilePath string
	File     io.Reader
}

type result struct {
	Message message `json:"result"`
	Ok      bool    `json:"ok"`
}
type message struct {
	Document document `json:"document"`
}

type document struct {
	FileID string `json:"file_id"`
}

func SendHttpFormToGetFileID(endpoint, method string, data []ValueData, files []FileData) (string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for _, i := range files {
		part, _ := writer.CreateFormFile(i.FileName, filepath.Base(fmt.Sprintf("%s_%s", uuid.New().String(), i.FilePath)))
		io.Copy(part, i.File)
	}

	for _, i := range data {
		writer.WriteField(i.Key, i.Value)
	}
	writer.Close()

	r, err := http.NewRequest(method, endpoint, body)
	if err != nil {
		return "", err
	}
	r.Header.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return "", err
	}

	msg := result{}
	err = json.NewDecoder(res.Body).Decode(&msg)
	// err = json.Unmarshal(b.Bytes(), &msg)
	if err != nil {
		return "", err
	}
	if !msg.Ok {
		return "", errors.New("fucking telegram sorry") // TODO: normal error response
	}
	if msg.Message.Document.FileID == "" {
		return "", errors.New("file id field is null")
	}
	return msg.Message.Document.FileID, nil
}
