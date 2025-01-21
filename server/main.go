package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/BryanMwangi/qa/server/cmd"
)

func main() {
	// create a global context that will be used to signal the server to stop
	GlobalContext, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	cmd.Run("3001", GlobalContext)
}
