package app

import (
	"easyStorage/config"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type Bot struct {
	Token     string
	ChannelID int
	mu        sync.Mutex
}

func (b *Bot) getSendFileEndpoint() string {
	return fmt.Sprintf("https://api.telegram.org/bot%s/sendDocument", b.Token)
}

func (b *Bot) getFileDataEndpoint(fileID string) string {
	return fmt.Sprintf("https://api.telegram.org/bot%s/getFile?file_id=%s", b.Token, fileID)
}

func (b *Bot) getFileEndpoint(filePath string) string {
	return fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", b.Token, filePath)
}

func (b *Bot) SendFile(fileName string, file io.Reader) (string, string, error) { // token, fileid
	b.mu.Lock()
	time.Sleep(time.Millisecond * time.Duration(config.SendFrameDelay))
	b.mu.Unlock()
	fileID, err := SendHttpFormToGetFileID(b.getSendFileEndpoint(),
		"POST",
		[]ValueData{{Key: "chat_id", Value: strconv.Itoa(b.ChannelID)}},
		[]FileData{{FileName: "document", FilePath: fileName, File: file}})

	return b.Token, fileID, err
}

func (b *Bot) GetFile(fileID string) (io.Reader, error) {
	res, err := http.Get(b.getFileDataEndpoint(fileID))
	if err != nil {
		return nil, err
	}

	var p struct {
		Result struct {
			FilePath string `json:"file_path"`
		} `json:"result"`
	}

	if err = json.NewDecoder(res.Body).Decode(&p); err != nil {
		return nil, err
	}
	if p.Result.FilePath == "" {
		return nil, errors.New("file path is null")
	}
	file, err := http.Get(b.getFileEndpoint(p.Result.FilePath))
	return file.Body, err
}

type Bots struct {
	Bots    []Bot
	counter int
}

func (b *Bots) Load(bots []config.Bot) {
	if len(bots) == 0 {
		log.Panic("bots were not loaded")
	}
	for _, bot := range bots {
		b.Bots = append(b.Bots, Bot{
			Token:     bot.Token,
			ChannelID: bot.ChannelID,
		})
	}
}

func (b *Bots) GetRandomBot() *Bot {
	b.counter += 1
	b.counter %= len(b.Bots)
	return &b.Bots[b.counter]
}

func (b *Bots) GetCurrentBot(token string) *Bot {
	for i := 0; i < len(b.Bots); i += 1 {
		if b.Bots[i].Token == token {
			return &b.Bots[i]
		}
	}
	return nil
}

var BotsManager = new(Bots)
