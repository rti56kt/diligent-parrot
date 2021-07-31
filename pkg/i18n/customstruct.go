package i18n

type ping struct {
	DES  string
	PONG string
}

type ifcfgme struct {
	DES string
}

type help struct {
	DES string
}

type set struct {
	DES     string
	USAGE   string
	SUCCESS string
	DUP     string
}

type locale struct {
	DES       string
	USAGE     string
	SUCCESS   string
	CURRENT   string
	FAIL      string
	SUPPORTED string
}

type prefix struct {
	DES     string
	USAGE   string
	SUCCESS string
	FAIL    string
}

type cmd struct {
	NAME    string
	PING    ping
	IFCFGME ifcfgme
	HELP    help
	SET     set
	LOCALE  locale
	PREFIX  prefix
}
