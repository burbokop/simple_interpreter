package main

import (
	"bufio"
	"flag"
	"os"

	engine "github.com/burbokop/simple_interpreter/src/engine"
)

func main() {
	var inputPath = flag.String("i", "", "Input file path")
	if len(*inputPath) == 0 {
		return
	}

	var parser = new(engine.Parser)
	parser.AddCmdType(printCommand)

	eventLoop := new(engine.EventLoop)
	eventLoop.Start()
	if input, err := os.Open(*inputPath); err == nil {
		defer input.Close()
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			commandLine := scanner.Text()
			cmd := engine.Parse(commandLine) // parse the line to get an instance of Command
			eventLoop.Post(cmd)
		}
	}
	eventLoop.AwaitFinish()
}
