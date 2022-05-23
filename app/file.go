package app

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"easyStorage/config"

	"github.com/google/uuid"
)

type File struct {
	FileName     string      `json:"filename"`
	File         []byte      `json:"-"`
	FileID       string      `json:"fileID"`
	SplittedFile []FileFrame `json:"splittedFile"`
	Hash         string      `json:"hash"`
}

func (f *File) ReadFromIO(r io.Reader) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	f.File = buf.Bytes()
}
func (f *File) generateFileID() {
	u, _ := uuid.NewUUID()
	f.FileID = u.String()
}

func (f *File) split() {
	f.SplittedFile = []FileFrame{}
	frameIndex := 0
	for i := 0; i < len(f.File); i += config.SplitFileLength {
		c := 0
		if (i + config.SplitFileLength) > len(f.File) {
			c = len(f.File)
		} else {
			c = (i + config.SplitFileLength)
		}
		f.SplittedFile = append(f.SplittedFile,
			FileFrame{
				File:      f.File[i:c],
				FileIndex: frameIndex,
			})
		frameIndex += 1
	}
}

func (f *File) join() {
	f.File = []byte{}
	for i := 0; i < len(f.SplittedFile); i += 1 {
		for _, frame := range f.SplittedFile {
			if frame.FileIndex == i {
				f.File = append(f.File, frame.File...)
				break
			}
		}
	}
}

func (f *File) Send() error {
	f.Hash = hex.Dump(md5.New().Sum(f.File))
	f.generateFileID()
	f.split()

	errCh := make(chan error, len(f.SplittedFile))

	for i := 0; i < len(f.SplittedFile); i += 1 {
		go f.SplittedFile[i].SendFrame(errCh)
	}
	for i := 0; i < cap(errCh); i += 1 {
		err := (<-errCh)
		if err != nil {
			return err //TODO: handler error with resend data
		}

	}
	err := f.SaveFileData()
	return err
}

func (f *File) Get() error {
	errCh := make(chan error, len(f.SplittedFile))
	for i := 0; i < len(f.SplittedFile); i += 1 {
		go f.SplittedFile[i].GetFrame(errCh)
	}
	for i := 0; i < cap(errCh); i += 1 {
		err := (<-errCh)
		if err != nil {
			return err //TODO: handler error with resend data
		}

	}

	f.join()
	if f.Hash != hex.Dump(md5.New().Sum(f.File)) {
		return errors.New("differetn file hashs")
	}
	return nil
}

func (f *File) SaveFileData() error {
	j, err := json.Marshal(f)
	if err != nil {
		return err
	}
	err = os.WriteFile(fmt.Sprintf("%s/%s.json", config.StorageDirectory, f.FileID), j, 0644)
	return err
}

func GetFile(fileID string) (File, error) {
	b, err := os.ReadFile(fmt.Sprintf("%s/%s.json", config.StorageDirectory, fileID))
	if err != nil {
		return File{}, err
	}
	file := File{}
	json.Unmarshal(b, &file)
	file.Get()
	return file, nil
}
