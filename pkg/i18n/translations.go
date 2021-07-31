package i18n

var AllLocale = map[string]cmd{
	"en": {
		"English",
		ping{
			DES:  "Return \"Pong!\"",
			PONG: "Pong!",
		},
		ifcfgme{
			DES: "Return IP address of the bot",
		},
		help{
			DES: "Show this help message",
		},
		set{
			DES:     "Set a keyword responser",
			USAGE:   "Usage: set keyword response",
			SUCCESS: "Success!",
			DUP:     "Duplication detected! Overwrite old one",
		},
		locale{
			DES:       "Return current locale if newLocale is not given\nOtherwise set new locale",
			USAGE:     "Usage: locale [newLanguage]",
			SUCCESS:   "Success! The new language is set to ",
			CURRENT:   "The current language is ",
			FAIL:      "Fail! The language you specified is not supported yet",
			SUPPORTED: "Currently supported languages list",
		},
		prefix{
			DES:     "Return current command prefix if newPrefix is not given\nOtherwise set new command prefix",
			USAGE:   "Usage: prefix [newPrefix]",
			SUCCESS: "Success!",
			FAIL:    "Fail! Make sure you are the owner or admin of the guild",
		},
	},
	"zh-tw": {
		"繁體中文(台灣)",
		ping{
			DES:  "回覆 \"乓!\"",
			PONG: "乓!",
		},
		ifcfgme{
			DES: "回覆這個機器人的IP地址",
		},
		help{
			DES: "顯示這則幫助訊息",
		},
		set{
			DES:     "設定關鍵字回應",
			USAGE:   "使用方式: set keyword response",
			SUCCESS: "成功!",
			DUP:     "關鍵字重複! 將直接覆寫舊的回應",
		},
		locale{
			DES:       "如果沒給newLocale就回覆現在所使用的語言\n否則就設定新語言",
			USAGE:     "使用方式: locale [newLocale]",
			SUCCESS:   "成功! 新語言設置為",
			CURRENT:   "目前語言為",
			FAIL:      "失敗! 您所指定的新語言目前並未支援",
			SUPPORTED: "目前有支援的語言清單",
		},
		prefix{
			DES:     "如果沒給newPrefix就回覆現在所使用的命令字首\n否則就設定新命令字首",
			USAGE:   "使用方式: prefix [newPrefix]",
			SUCCESS: "成功!",
			FAIL:    "失敗! 你必須是伺服器擁有者或是管理員",
		},
	},
}
