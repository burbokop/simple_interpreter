package engine

import (
	"fmt"
	"reflect"
	"strings"
)

type Parser struct {
	Cmds map[string]reflect.Type
}

func (parser *Parser) AddCmdType(name string, t reflect.Type) {
	parser.Cmds[name] = t
}

func (parser *Parser) Parse(str string) Command {
	var s = strings.Fields(str)
	if len(s) > 0 {
		var t, found = parser.Cmds[s[0]]
		if found {
			var ptr = reflect.New(t)
			if ptr.IsNil() {
				return nil
			} else {
				return ptr.Elem().Interface().(Command)
			}
		}
	}
	return nil
}

func Print(cmds []Command) {
	for i, cmd := range cmds {
		fmt.Println(i, reflect.TypeOf(cmd).Name())
	}
}
