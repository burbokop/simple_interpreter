package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"reflect"

	engine "github.com/burbokop/simple_interpreter/src/engine"
)

type TestCommand struct {
	arg string
}

func (p *TestCommand) Execute(loop engine.Handler) {
	fmt.Println(p.arg)
}

func main() {
	var inputPath = flag.String("i", "", "Input file path")
	if len(*inputPath) == 0 {
		return
	}

	var parser = new(engine.Parser)
	parser.AddCmdType("test", reflect.TypeOf((*TestCommand)(nil)))

	eventLoop := new(engine.EventLoop)
	eventLoop.Start()
	if input, err := os.Open(*inputPath); err == nil {
		defer input.Close()
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			commandLine := scanner.Text()
			cmd := parser.Parse(commandLine) // parse the line to get an instance of Command
			eventLoop.Post(cmd)
		}
	}
	eventLoop.AwaitFinish()
}
