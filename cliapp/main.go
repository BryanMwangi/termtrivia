package main

import (
	"github.com/BryanMwangi/qa/cliapp/client"
	"github.com/BryanMwangi/qa/cliapp/cmd"
)

func main() {
	// init client
	client.Init()
	// run command
	cmd.Execute()
}
