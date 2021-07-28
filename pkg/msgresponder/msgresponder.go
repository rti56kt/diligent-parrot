package msgresponder

import (
	"github.com/rti56kt/diligent-parrot/pkg/logger"
)

var keywords = make(map[string]string)

// Search if the keyword is already exist.
//
// If yes, return the response string and a true boolean represented that keyword is exist.
//
// If no, return an empty string and a false boolean represented that keyword is not exist.
func GetKeywordResp(keyword string) (string, bool) {
	logger.Logger.WithField("type", "msgresper").Debug("keyword: ", keyword)
	if resp, exist := keywords[keyword]; exist {
		return resp, true
	} else {
		return "", false
	}
}

// Set new response for keyword. If there's already a mapping return true, else return false.
func SetKeywordResp(keyword string, resp string) bool {
	logger.Logger.WithField("type", "msgresper").Debug("keyword: ", keyword, "; resp: ", resp)
	if _, exist := keywords[keyword]; exist {
		keywords[keyword] = resp
		return true
	} else {
		keywords[keyword] = resp
		return false
	}
}
