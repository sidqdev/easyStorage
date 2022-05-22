package config

type Bot struct {
	Token     string `json:"token"`
	ChannelID int    `json:"channel_id"`
}

type Config struct {
	Bots                []Bot  `json:"bots"`
	StorageDirectory    string `json:"storage_directory"`
	SplitFileLength     int    `json:"split_file_length"`
	MaxUploadFileLength int    `json:"max_upload_file_length"`
	Host                string `json:"host"`
	SendFrameDelay      int    `json:"send_frame_delay"`
}
