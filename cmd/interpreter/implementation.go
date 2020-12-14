package main

import (
	"fmt"

	engine "github.com/burbokop/simple_interpreter/src/engine"
)

type PrintCommand struct {
	arg string
}

func (p *PrintCommand) Execute(loop engine.Handler) {
	fmt.Println(p.arg)
}
