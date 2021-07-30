package ifconfigme

import (
	"io/ioutil"
	"net/http"

	"github.com/rti56kt/diligent-parrot/pkg/logger"
)

func Dealer(authorTag string, cmdAndArgs []string) string {
	resp, err := http.Get("https://ifconfig.me")
	if err != nil {
		logger.Logger.WithField("type", "msg").Error(err)
	}
	defer resp.Body.Close()
	byteCode, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Logger.WithField("type", "msg").Error(err)
	}

	return string(byteCode)
}
