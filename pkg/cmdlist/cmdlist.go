package cmdlist

import (
	"github.com/rti56kt/diligent-parrot/pkg/i18n"
	"github.com/rti56kt/diligent-parrot/pkg/logger"
)

func GetCmdAndDes() []CmdAndDes {
	var cmdAndDesVisible []CmdAndDes

	for idx, cmdObj := range cmdlists {
		if cmdObj.Visibility {
			var cmdAndDes CmdAndDes
			description := GetI18nDes(idx)
			cmdAndDes.CommandStruct = cmdObj
			cmdAndDes.Description = description
			cmdAndDesVisible = append(cmdAndDesVisible, cmdAndDes)
		}
	}
	return cmdAndDesVisible
}

func CheckCmdByFullName(input string, fullName int) bool {
	logger.Logger.WithField("type", "cmdlist").Debug("input: ", input)
	for _, cmd := range cmdlists[fullName].Command {
		logger.Logger.WithField("type", "cmdlist").Debug("cmd: ", cmd)
		if input == cmd {
			return true
		}
	}
	return false
}

func GetI18nDes(cmd int) string {
	locale := i18n.GetCurrentLocale()
	switch cmd {
	case Ping:
		return i18n.AllLocale[locale].PING.DES
	case Ifconfigme:
		return i18n.AllLocale[locale].IFCFGME.DES
	case Help:
		return i18n.AllLocale[locale].HELP.DES
	case Set:
		return i18n.AllLocale[locale].SET.DES
	case Locale:
		return i18n.AllLocale[locale].LOCALE.DES
	case Prefix:
		return i18n.AllLocale[locale].PREFIX.DES
	default:
		return ""
	}
}
