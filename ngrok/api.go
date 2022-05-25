package ngrokCore

import (
	"context"
	"easyStorage/config"
	"errors"
	"strconv"
	"strings"

	"github.com/ngrok/ngrok-api-go/v4"
	"github.com/ngrok/ngrok-api-go/v4/tunnels"
)

func getPortFromAddr(addr string) int {
	frames := strings.Split(addr, ":")
	port := frames[len(frames)-1]
	iPort, _ := strconv.Atoi(port)
	return iPort
}
func GetNgrokEndpoint() (string, error) {
	clientConfig := ngrok.NewClientConfig(config.NgrokApiKey)
	tunnels := tunnels.NewClient(clientConfig)
	iter := tunnels.List(nil)
	for iter.Next(context.Background()) {
		item := iter.Item()
		if getPortFromAddr(item.ForwardsTo) == getPortFromAddr(config.Host) {
			if item.Proto == "https" {
				if len(item.PublicURL) != 0 {
					return item.PublicURL, nil
				}
			}
		}
	}

	if err := iter.Err(); err != nil {
		return "", err
	}

	return "", errors.New("can not find tunnel")
}
