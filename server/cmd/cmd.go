package cmd

import (
	"context"

	"github.com/BryanMwangi/pine"
	"github.com/BryanMwangi/pine/logger"
	"github.com/BryanMwangi/qa/server/db"
	"github.com/BryanMwangi/qa/server/handlers"
)

func Run(port string, ctx context.Context) {
	// start a new Pine server
	//
	// Pine is a web framework for Go that provides a simple and easy-to-use API for building web applications.
	// it was built by me, Bryan, and shares many similarities with Express.js and Fiber.
	//
	// Feel free to star on github if you like it!
	//
	// if you are familiar with either, Pine is the similar with some built in
	// features that make it easier to build web applications.
	app := pine.New()

	// routes for the server
	handlers.Routes(app)

	// start new cache for the server
	db.Init()

	// I used Pine's built-in logger to log messages to the console
	logger.Success("Listening on port: " + port)

	// Start a channel to listen for errors
	ch := make(chan error, 1)

	// start the server in a goroutine and send the error to the channel
	go func() {
		// start the server on the specified port
		ch <- app.Start(":" + port)
	}()
	select {
	case <-ctx.Done():
		logger.Warning("Shutting down gracefully...")

		// Pine offers graceful shutdown, which waits for all active connections to finish before shutting down.
		//
		// Additional shutdown processes can be added during the shutdown process such as closing database connections.
		if err := app.ServeShutDown(ctx); err != nil {
			logger.Error("Error shutting down server: " + err.Error())
			ch <- ctx.Err()
		}
	case err := <-ch:
		if err != nil {
			logger.Error("Error starting server: " + err.Error())
		}
	}
	close(ch)
	logger.Success("Server stopped!")
}
