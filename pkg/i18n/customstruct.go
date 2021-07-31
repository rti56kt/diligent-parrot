package i18n

type ping struct {
	PONG string
}

type set struct {
	USAGE   string
	SUCCESS string
	DUP     string
}

type locale struct {
	USAGE     string
	SUCCESS   string
	CURRENT   string
	FAIL      string
	SUPPORTED string
}

type prefix struct {
	USAGE   string
	SUCCESS string
	FAIL    string
}

type cmd struct {
	NAME   string
	PING   ping
	SET    set
	LOCALE locale
	PREFIX prefix
}
