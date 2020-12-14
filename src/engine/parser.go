package engine

import (
	"reflect"
	"strings"
)

type Parser struct {
	Cmds map[string]reflect.Type
}

func (parser *Parser) AddCmdType(name string, f int) {
	f.
		parser.Cmds[name] = f
}

func (parser *Parser) Parse(str string) Command {
	var s = strings.Fields(str)
	if len(s) > 0 {
		var t, found = parser.Cmds[s[0]]
		if found {
			var _ = reflect.ValueOf(t)
			//			return obj.C
		}
	}
	return nil
}
