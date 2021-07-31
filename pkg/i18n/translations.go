package i18n

var AllLocale = map[string]cmd{
	"en": {
		"English",
		ping{
			PONG: "Pong!",
		},
		set{
			USAGE:   "Usage: set keyword response",
			SUCCESS: "Success!",
			DUP:     "Duplication detected! Overwrite old one",
		},
		locale{
			USAGE:     "Usage: locale [newLanguage]",
			SUCCESS:   "Success! The new language is set to ",
			CURRENT:   "The current language is ",
			FAIL:      "Fail! The language you specified is not supported yet",
			SUPPORTED: "Currently supported languages list",
		},
		prefix{
			USAGE:   "Usage: prefix [newPrefix]",
			SUCCESS: "Success!",
			FAIL:    "Fail! Make sure you are the owner or admin of the guild",
		},
	},
	"zh-tw": {
		"繁體中文(台灣)",
		ping{
			PONG: "乓!",
		},
		set{
			USAGE:   "使用方式: set 關鍵字 回應",
			SUCCESS: "成功!",
			DUP:     "關鍵字重複! 將直接覆寫舊的回應",
		},
		locale{
			USAGE:     "使用方式: locale [新語言]",
			SUCCESS:   "成功! 新語言設置為",
			CURRENT:   "目前語言為",
			FAIL:      "失敗! 您所指定的新語言目前並未支援",
			SUPPORTED: "目前有支援的語言清單",
		},
		prefix{
			USAGE:   "使用方式: prefix [新命令字首]",
			SUCCESS: "成功!",
			FAIL:    "失敗! 你必須是伺服器擁有者或是管理員",
		},
	},
}
