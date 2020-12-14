package main

import (
	"bufio"
	"os"
	"github.com/burbokop/simple_interpreter/src"
)

func main() {
	eventLoop := new(engine.EventLoop)
	eventLoop.Start()
	if input, err := os.Open(inputFile); err == nil {
		defer input.Close()
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			commandLine := scanner.Text()
			cmd := parse(commandLine) // parse the line to get an instance of Command
			eventLoop.Post(cmd)
		}
	}
	eventLoop.AwaitFinish()
}
