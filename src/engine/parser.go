package engine

import "strings"

type Parser struct {
	Cmds map[string]interface{}
}

func (parser *Parser) AddCmdType(name string, f interface{}) {
	parser.Cmds[name] = f
}

func (parser *Parser) Parse(str string) Command {
	var s = strings.Fields(str)
	if len(s) > 0 {
		parser.Cmds[s[0]]
		//var obj = new()
	}
}
