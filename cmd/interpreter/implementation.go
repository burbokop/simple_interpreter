package main

import (
	"fmt"
	engine "github.com/burbokop/simple_interpreter/src/engine"
)

type printCommand struct {
	arg string
}

func (p *printCommand) Execute(loop engine.Handler) {
	fmt.Println(p.arg)
}
