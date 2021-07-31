package cmdprefix

import (
	"github.com/rti56kt/diligent-parrot/pkg/i18n"
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
	locale := i18n.GetCurrentLocale()
	cmdAndArgs := msgparser.GetCmdAndArgs(m)

	if len(cmdAndArgs) == 1 {
		return GetPrefix()
	} else if len(cmdAndArgs) == 2 {
		if msgparser.IsAuthorAdmin(s, m) || msgparser.IsAuthorOwner(s, m) {
			setPrefix(cmdAndArgs[1])
			return i18n.AllLocale[locale].PREFIX.SUCCESS
		} else {
			return i18n.AllLocale[locale].PREFIX.FAIL
		}
	} else {
		return i18n.AllLocale[locale].PREFIX.USAGE
	}
}
