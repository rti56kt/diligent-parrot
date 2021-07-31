package cmdlist

type CmdList struct {
	Command    []string
	Argument   string
	Visibility bool
}

type CmdAndDes struct {
	CommandStruct CmdList
	Description   string
}
