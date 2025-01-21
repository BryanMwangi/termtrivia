package cmd

import (
	"fmt"
)

// Execute runs the root command.
func Execute() {
	fmt.Println(`
████████╗███████╗██████╗ ███╗   ███╗████████╗██████╗ ██╗██╗   ██╗██╗ █████╗ 
╚══██╔══╝██╔════╝██╔══██╗████╗ ████║╚══██╔══╝██╔══██╗██║██║   ██║██║██╔══██╗
   ██║   █████╗  ██████╔╝██╔████╔██║   ██║   ██████╔╝██║██║   ██║██║███████║
   ██║   ██╔══╝  ██╔══██╗██║╚██╔╝██║   ██║   ██╔══██╗██║╚██╗ ██╔╝██║██╔══██║
   ██║   ███████╗██║  ██║██║ ╚═╝ ██║   ██║   ██║  ██║██║ ╚████╔╝ ██║██║  ██║
   ╚═╝   ╚══════╝╚═╝  ╚═╝╚═╝     ╚═╝   ╚═╝   ╚═╝  ╚═╝╚═╝  ╚═══╝  ╚═╝╚═╝  ╚═╝
	`)
	fmt.Println(`Enter your name to start playing:`)
	// authenticate user
	err := authenticateUser()
	if err != nil {
		panic("Authentication failed due to " + err.Error())
	}
	gameLoop()
}
