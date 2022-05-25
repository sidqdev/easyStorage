package config

var (
	Bots = []Bot{
		{
			Token:     "5105591629:AAFtTj5py1dcBHcTSumYaXW8TwOXerZO47M",
			ChannelID: -1001500447847,
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
