package ifconfigme

import (
	"io/ioutil"
	"net/http"

	"github.com/rti56kt/diligent-parrot/pkg/logger"
)

func Dealer() string {
	logger.Logger.WithField("type", "ifcfgme").Info("ifcfgme dealer triggered")
	resp, err := http.Get("https://ifconfig.me")
	if err != nil {
		logger.Logger.WithField("type", "ifcfgme").Error(err)
	}
	defer resp.Body.Close()
	byteCode, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Logger.WithField("type", "ifcfgme").Error(err)
	}

	return string(byteCode)
}
