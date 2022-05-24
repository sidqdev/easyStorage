package main

import (
	"easyStorage/app"
	"easyStorage/config"
	"math/rand"
	"testing"
)

func fileUpload(initialBytes []byte, t *testing.T) {
	file := app.File{Name: "File", File: initialBytes}
	err := file.Send()
	if err != nil {
		t.Error(err)
		return
	}

	fileID := file.ID

	gotFile, err := app.GetFile(fileID)
	if err != nil {
		t.Error(err)
		return
	}
	err = gotFile.Get()
	if err != nil {
		t.Error(err)
		return
	}
	gotBytes := gotFile.File
	if len(initialBytes) != len(gotBytes) {
		t.Fatal("bytes have different lenght")
	}
	for i := 0; i < len(gotBytes); i += 1 {
		if gotBytes[i] != initialBytes[i] {
			t.Fatal("bytes are not equal")
		}
	}
}
func TestLightFile(t *testing.T) {
	err := config.Load()
	if err != nil {
		t.Log(err)
	}
	app.BotsManager.Load(config.Bots)

	initialBytes := make([]byte, 1*1024*1024)
	rand.Read(initialBytes)
	fileUpload(initialBytes, t)
	t.Log("ok")
}
