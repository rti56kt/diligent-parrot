package i18n

import (
	"sort"

	"github.com/rti56kt/diligent-parrot/pkg/logger"

	"github.com/bwmarrin/discordgo"
)

var curLocale string = "en"

func setLocale(input string) bool {
	if _, exist := AllLocale[input]; exist {
		curLocale = input
		return true
	} else {
		return false
	}
}

func getAllLocale() []string {
	allSupportedLocale := make([]string, 0, len(AllLocale))
	for k := range AllLocale {
		allSupportedLocale = append(allSupportedLocale, k)
	}
	sort.Strings(allSupportedLocale)
	return allSupportedLocale
}

func GetCurrentLocale() string {
	return curLocale
}

func Dealer(authorTag string, cmdAndArgs []string) discordgo.MessageSend {
	var respMsg discordgo.MessageSend
	var locale string = GetCurrentLocale()

	if len(cmdAndArgs) == 1 {
		respMsg.Content = authorTag + " " + AllLocale[locale].LOCALE.CURRENT + "\"" + locale + "\""
	} else if len(cmdAndArgs) == 2 {
		if exist := setLocale(cmdAndArgs[1]); exist {
			locale = GetCurrentLocale()
			respMsg.Content = authorTag + " " + AllLocale[locale].LOCALE.SUCCESS + "\"" + locale + "\""
		} else {
			var msgEmbed discordgo.MessageEmbed
			var embedFields []*discordgo.MessageEmbedField

			supportedLocales := getAllLocale()

			msgEmbed.Title = AllLocale[locale].LOCALE.SUPPORTED
			for _, supportedLocale := range supportedLocales {
				var cmdField discordgo.MessageEmbedField
				cmdField.Name = supportedLocale
				cmdField.Value = AllLocale[supportedLocale].NAME
				embedFields = append(embedFields, &cmdField)
			}
			msgEmbed.Fields = embedFields

			respMsg.Content = authorTag + " " + AllLocale[locale].LOCALE.FAIL
			respMsg.Embed = &msgEmbed
		}
	} else {
		logger.Logger.WithField("type", "msg").Debug("Num of args is not correct")
		respMsg.Content = authorTag + " " + AllLocale[locale].LOCALE.USAGE
	}
	return respMsg
}
