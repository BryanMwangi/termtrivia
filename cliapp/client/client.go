package client

import (
	"github.com/BryanMwangi/pine"
)

var (
	Client *pine.Client
	Url    = "http://localhost:3001"
)

// I used Pine's client for this just to save time but if needed, we could also use
// the basic http client
func Init() {
	Client = pine.NewClient()
	Client.Request().SetRequestURI(Url)
	Client.Request().SetMethod("GET")
}

func GenerateUri(path string) string {
	return Url + path
}
