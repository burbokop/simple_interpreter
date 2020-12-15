package main

import (
	"fmt"

	engine "github.com/burbokop/simple_interpreter/src/engine"
)

type PrintCommand struct {
	arg string
}

func (p *PrintCommand) Init(args []string) {
	if len(args) > 1 {
		p.arg = args[1]
	}
}

func (p *PrintCommand) Execute(loop engine.Handler) {
	fmt.Println(p.arg)
}

type ReverseCommand struct {
	arg   string
	async bool
}

func (p *ReverseCommand) Init(args []string) {
	if len(args) > 1 {
		p.arg = args[1]
		if len(args) > 2 {
			if args[2] == "async" {
				p.async = true
			}
		}
	}
}

func (p *ReverseCommand) Execute(loop engine.Handler) {
	var reverce = func(s string) string {
		rns := []rune(s)
		for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
			rns[i], rns[j] = rns[j], rns[i]
		}
		return string(rns)
	}
	if p.async {
		loop.Post(&PrintCommand{arg: reverce(p.arg)})
	} else {
		fmt.Println(reverce(p.arg))
	}
}
