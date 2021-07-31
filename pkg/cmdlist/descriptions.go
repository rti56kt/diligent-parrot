package cmdlist

const (
	Ping int = iota
	Ifconfigme
	Help
	Set
	Locale
	Prefix
)

var cmdlists = []CmdList{
	{[]string{"ping"}, "", false},
	{[]string{"ifconfigme", "ip"}, "", false},
	{[]string{"help", "h"}, "", true},
	{[]string{"set", "s"}, "keyword response", true},
	{[]string{"locale", "l"}, "[newLocale]", true},
	{[]string{"prefix"}, "[newPrefix]", true},
}
