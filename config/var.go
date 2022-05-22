package config

var (
	Bots = []Bot{
		{
			Token:     "5105591629:AAFtTj5py1dcBHcTSumYaXW8TwOXerZO47M",
			ChannelID: -1001500447847,
		},
	}
	MaxFileLength    = 20 * 1024 * 1024
	StorageDirectory = "files"
	Host             = "0.0.0.0:1000"
)
