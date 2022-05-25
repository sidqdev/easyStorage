package config

var (
	Bots = []Bot{
		{
			Token:     "",
			ChannelID: 0,
		},
	}
	SplitFileLength     = 20 * 1024 * 1024
	StorageDirectory    = "files"
	MaxUploadFileLength = -1
	Host                = "0.0.0.0:8070"
	SendFrameDelay      = 35
	NgrokApiKey         = ""
	UseNgrok            = false
)
