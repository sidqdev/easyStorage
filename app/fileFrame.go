package app

import (
	"bytes"
	"io"
)

type FileFrame struct {
	File         []byte `json:"-"`
	TelegramFile TelegramFileFrame
	FileIndex    int
}

func (f *FileFrame) GetIOReader() io.Reader {
	return bytes.NewReader(f.File)
}

func (f *FileFrame) FromIOReader(r io.Reader) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	f.File = buf.Bytes()
}

func (f *FileFrame) SendFrame(errRes chan error) {
	token, fileID, err := BotsManager.GetRandomBot().SendFile("", f.GetIOReader())
	if err != nil {
		errRes <- err
		return
	}
	f.TelegramFile = TelegramFileFrame{
		Token:  token,
		FileID: fileID,
	}
	errRes <- nil
}

func (f *FileFrame) GetFrame(errRes chan error) {
	fileReader, err := BotsManager.GetCurrentBot(f.TelegramFile.Token).GetFile(f.TelegramFile.FileID)
	if err != nil {
		errRes <- err
		return
	}
	f.FromIOReader(fileReader)
	errRes <- nil
}
