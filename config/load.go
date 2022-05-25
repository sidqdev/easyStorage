package config

import (
	"encoding/json"
	"os"
)

func Load() error {
	f, err := os.ReadFile("config.json")
	if err != nil {
		return err
	}
	conf := Config{}

	err = json.Unmarshal(f, &conf)
	if err != nil {
		return err
	}

	Bots = conf.Bots
	SplitFileLength = conf.SplitFileLength
	MaxUploadFileLength = conf.MaxUploadFileLength
	Host = conf.Host
	StorageDirectory = conf.StorageDirectory
	NgrokApiKey = conf.NgrokApiKey
	UseNgrok = conf.UseNgrok != 0
	return nil
}
