package cmdprefix

import (
	"github.com/rti56kt/diligent-parrot/pkg/logger"
	"github.com/rti56kt/diligent-parrot/pkg/msgparser"

	"github.com/bwmarrin/discordgo"
)

var (
	prefix string = "$"
)

func GetPrefix() string {
	return prefix
}

func setPrefix(newPrefix string) {
	prefix = newPrefix
}

func Dealer(s *discordgo.Session, m *discordgo.MessageCreate) string {
	logger.Logger.WithField("type", "cmdprefix").Info("cmdprefix dealer triggered")
	cmdAndArgs := msgparser.GetCmdAndArgs(m)

	if len(cmdAndArgs) == 1 {
		return GetPrefix()
	} else if len(cmdAndArgs) == 2 {
		if msgparser.IsAuthorAdmin(s, m) || msgparser.IsAuthorOwner(s, m) {
			setPrefix(cmdAndArgs[1])
			return "Success"
		} else {
			return "401"
		}
	} else {
		return "Usage"
	}
}
